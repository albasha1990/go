// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <u.h>
#include <libc.h>
#include "gg.h"
#include "opt.h"

enum
{
	RightRdwr = RightRead | RightWrite,
};

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
//
// The table is formatted for 8-space tabs.
static ProgInfo progtable[ALAST] = {
	[ATYPE]=	{Pseudo | Skip},
	[ATEXT]=	{Pseudo},
	[AFUNCDATA]=	{Pseudo},
	[APCDATA]=	{Pseudo},
	[AUNDEF]=	{OK},
	[AUSEFIELD]=	{OK},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the Intel opcode.
	[ANOP]=		{LeftRead | RightWrite},
	
	// Integer.
	[AADC]=		{SizeL | LeftRead | RegRead | RightWrite},
	[AADD]=		{SizeL | LeftRead | RegRead | RightWrite},
	[AAND]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ABIC]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ACMN]=		{SizeL | LeftRead | RightRead},
	[ACMP]=		{SizeL | LeftRead | RightRead},
	[ADIVU]=	{SizeL | LeftRead | RegRead | RightWrite},
	[ADIV]=		{SizeL | LeftRead | RegRead | RightWrite},
	[AEOR]=		{SizeL | LeftRead | RegRead | RightWrite},
	[AMODU]=	{SizeL | LeftRead | RegRead | RightWrite},
	[AMOD]=		{SizeL | LeftRead | RegRead | RightWrite},
	[AMULALU]=	{SizeL | LeftRead | RegRead | RightRdwr},
	[AMULAL]=	{SizeL | LeftRead | RegRead | RightRdwr},
	[AMULA]=	{SizeL | LeftRead | RegRead | RightRdwr},
	[AMULU]=	{SizeL | LeftRead | RegRead | RightWrite},
	[AMUL]=		{SizeL | LeftRead | RegRead | RightWrite},
	[AMULL]=	{SizeL | LeftRead | RegRead | RightWrite},
	[AMULLU]=	{SizeL | LeftRead | RegRead | RightWrite},
	[AMVN]=		{SizeL | LeftRead | RightWrite},
	[AORR]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ARSB]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ARSC]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ASBC]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ASLL]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ASRA]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ASRL]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ASUB]=		{SizeL | LeftRead | RegRead | RightWrite},
	[ATEQ]=		{SizeL | LeftRead | RightRead},
	[ATST]=		{SizeL | LeftRead | RightRead},

	// Floating point.
	[AADDD]=	{SizeD | LeftRead | RightRdwr},
	[AADDF]=	{SizeF | LeftRead | RightRdwr},
	[ACMPD]=	{SizeD | LeftRead | RightRead},
	[ACMPF]=	{SizeF | LeftRead | RightRead},
	[ADIVD]=	{SizeD | LeftRead | RightRdwr},
	[ADIVF]=	{SizeF | LeftRead | RightRdwr},
	[AMULD]=	{SizeD | LeftRead | RightRdwr},
	[AMULF]=	{SizeF | LeftRead | RightRdwr},
	[ASUBD]=	{SizeD | LeftRead | RightRdwr},
	[ASUBF]=	{SizeF | LeftRead | RightRdwr},

	// Conversions.
	[AMOVWD]=		{SizeD | LeftRead | RightWrite | Conv},
	[AMOVWF]=		{SizeF | LeftRead | RightWrite | Conv},
	[AMOVDF]=		{SizeF | LeftRead | RightWrite | Conv},
	[AMOVDW]=		{SizeL | LeftRead | RightWrite | Conv},
	[AMOVFD]=		{SizeD | LeftRead | RightWrite | Conv},
	[AMOVFW]=		{SizeL | LeftRead | RightWrite | Conv},

	// Moves.
	[AMOVB]=		{SizeB | LeftRead | RightWrite | Move},
	[AMOVD]=		{SizeD | LeftRead | RightWrite | Move},
	[AMOVF]=		{SizeF | LeftRead | RightWrite | Move},
	[AMOVH]=		{SizeW | LeftRead | RightWrite | Move},
	[AMOVW]=		{SizeL | LeftRead | RightWrite | Move},

	// These should be split into the two different conversions instead
	// of overloading the one.
	[AMOVBS]=		{SizeB | LeftRead | RightWrite | Conv},
	[AMOVBU]=		{SizeB | LeftRead | RightWrite | Conv},
	[AMOVHS]=		{SizeW | LeftRead | RightWrite | Conv},
	[AMOVHU]=		{SizeW | LeftRead | RightWrite | Conv},
	
	// Jumps.
	[AB]=		{Jump},
	[ABL]=		{Call},
	[ABEQ]=		{Cjmp},
	[ABNE]=		{Cjmp},
	[ABCS]=		{Cjmp},
	[ABHS]=		{Cjmp},
	[ABCC]=		{Cjmp},
	[ABLO]=		{Cjmp},
	[ABMI]=		{Cjmp},
	[ABPL]=		{Cjmp},
	[ABVS]=		{Cjmp},
	[ABVC]=		{Cjmp},
	[ABHI]=		{Cjmp},
	[ABLS]=		{Cjmp},
	[ABGE]=		{Cjmp},
	[ABLT]=		{Cjmp},
	[ABGT]=		{Cjmp},
	[ABLE]=		{Cjmp},
	[ARET]=		{Break},
};

void
proginfo(ProgInfo *info, Prog *p)
{
	*info = progtable[p->as];
	if(info->flags == 0)
		fatal("unknown instruction %P", p);

	if((info->flags & RegRead) && p->reg == NREG) {
		info->flags &= ~RegRead;
		info->flags |= CanRegRead | RightRead;
	}
	
	if(((p->scond & C_SCOND) != C_SCOND_NONE) && (info->flags & RightWrite))
		info->flags |= RightRead;
}
