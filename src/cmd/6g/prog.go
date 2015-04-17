// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cmd/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/x86"
)

const (
	LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
	RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
//
// The table is formatted for 8-space tabs.
var progtable = [x86.ALAST]obj.ProgInfo{
	obj.ATYPE:     {gc.Pseudo | gc.Skip, 0, 0, 0},
	obj.ATEXT:     {gc.Pseudo, 0, 0, 0},
	obj.AFUNCDATA: {gc.Pseudo, 0, 0, 0},
	obj.APCDATA:   {gc.Pseudo, 0, 0, 0},
	obj.AUNDEF:    {gc.Break, 0, 0, 0},
	obj.AUSEFIELD: {gc.OK, 0, 0, 0},
	obj.ACHECKNIL: {gc.LeftRead, 0, 0, 0},
	obj.AVARDEF:   {gc.Pseudo | gc.RightWrite, 0, 0, 0},
	obj.AVARKILL:  {gc.Pseudo | gc.RightWrite, 0, 0, 0},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the Intel opcode.
	obj.ANOP:       {gc.LeftRead | gc.RightWrite, 0, 0, 0},
	x86.AADCL:      {gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.AADCQ:      {gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.AADCW:      {gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.AADDB:      {gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AADDL:      {gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AADDW:      {gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AADDQ:      {gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AADDSD:     {gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.AADDSS:     {gc.SizeF | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.AANDB:      {gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AANDL:      {gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AANDQ:      {gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AANDW:      {gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	obj.ACALL:      {gc.RightAddr | gc.Call | gc.KillCarry, 0, 0, 0},
	x86.ACDQ:       {gc.OK, AX, AX | DX, 0},
	x86.ACQO:       {gc.OK, AX, AX | DX, 0},
	x86.ACWD:       {gc.OK, AX, AX | DX, 0},
	x86.ACLD:       {gc.OK, 0, 0, 0},
	x86.ASTD:       {gc.OK, 0, 0, 0},
	x86.ACMPB:      {gc.SizeB | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ACMPL:      {gc.SizeL | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ACMPQ:      {gc.SizeQ | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ACMPW:      {gc.SizeW | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ACOMISD:    {gc.SizeD | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ACOMISS:    {gc.SizeF | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ACVTSD2SL:  {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSD2SQ:  {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSD2SS:  {gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSL2SD:  {gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSL2SS:  {gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSQ2SD:  {gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSQ2SS:  {gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSS2SD:  {gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSS2SL:  {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTSS2SQ:  {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTTSD2SL: {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTTSD2SQ: {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTTSS2SL: {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ACVTTSS2SQ: {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.ADECB:      {gc.SizeB | RightRdwr, 0, 0, 0},
	x86.ADECL:      {gc.SizeL | RightRdwr, 0, 0, 0},
	x86.ADECQ:      {gc.SizeQ | RightRdwr, 0, 0, 0},
	x86.ADECW:      {gc.SizeW | RightRdwr, 0, 0, 0},
	x86.ADIVB:      {gc.SizeB | gc.LeftRead | gc.SetCarry, AX, AX, 0},
	x86.ADIVL:      {gc.SizeL | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	x86.ADIVQ:      {gc.SizeQ | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	x86.ADIVW:      {gc.SizeW | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	x86.ADIVSD:     {gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.ADIVSS:     {gc.SizeF | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.AIDIVB:     {gc.SizeB | gc.LeftRead | gc.SetCarry, AX, AX, 0},
	x86.AIDIVL:     {gc.SizeL | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	x86.AIDIVQ:     {gc.SizeQ | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	x86.AIDIVW:     {gc.SizeW | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	x86.AIMULB:     {gc.SizeB | gc.LeftRead | gc.SetCarry, AX, AX, 0},
	x86.AIMULL:     {gc.SizeL | gc.LeftRead | gc.ImulAXDX | gc.SetCarry, 0, 0, 0},
	x86.AIMULQ:     {gc.SizeQ | gc.LeftRead | gc.ImulAXDX | gc.SetCarry, 0, 0, 0},
	x86.AIMULW:     {gc.SizeW | gc.LeftRead | gc.ImulAXDX | gc.SetCarry, 0, 0, 0},
	x86.AINCB:      {gc.SizeB | RightRdwr, 0, 0, 0},
	x86.AINCL:      {gc.SizeL | RightRdwr, 0, 0, 0},
	x86.AINCQ:      {gc.SizeQ | RightRdwr, 0, 0, 0},
	x86.AINCW:      {gc.SizeW | RightRdwr, 0, 0, 0},
	x86.AJCC:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJCS:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJEQ:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJGE:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJGT:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJHI:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJLE:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJLS:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJLT:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJMI:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJNE:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJOC:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJOS:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJPC:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJPL:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	x86.AJPS:       {gc.Cjmp | gc.UseCarry, 0, 0, 0},
	obj.AJMP:       {gc.Jump | gc.Break | gc.KillCarry, 0, 0, 0},
	x86.ALEAL:      {gc.LeftAddr | gc.RightWrite, 0, 0, 0},
	x86.ALEAQ:      {gc.LeftAddr | gc.RightWrite, 0, 0, 0},
	x86.AMOVBLSX:   {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVBLZX:   {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVBQSX:   {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVBQZX:   {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVBWSX:   {gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVBWZX:   {gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVLQSX:   {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVLQZX:   {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVWLSX:   {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVWLZX:   {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVWQSX:   {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVWQZX:   {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVQL:     {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	x86.AMOVB:      {gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	x86.AMOVL:      {gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	x86.AMOVQ:      {gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	x86.AMOVW:      {gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	x86.AMOVSB:     {gc.OK, DI | SI, DI | SI, 0},
	x86.AMOVSL:     {gc.OK, DI | SI, DI | SI, 0},
	x86.AMOVSQ:     {gc.OK, DI | SI, DI | SI, 0},
	x86.AMOVSW:     {gc.OK, DI | SI, DI | SI, 0},
	obj.ADUFFCOPY:  {gc.OK, DI | SI, DI | SI | CX, 0},
	x86.AMOVSD:     {gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	x86.AMOVSS:     {gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},

	// We use MOVAPD as a faster synonym for MOVSD.
	x86.AMOVAPD:   {gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	x86.AMULB:     {gc.SizeB | gc.LeftRead | gc.SetCarry, AX, AX, 0},
	x86.AMULL:     {gc.SizeL | gc.LeftRead | gc.SetCarry, AX, AX | DX, 0},
	x86.AMULQ:     {gc.SizeQ | gc.LeftRead | gc.SetCarry, AX, AX | DX, 0},
	x86.AMULW:     {gc.SizeW | gc.LeftRead | gc.SetCarry, AX, AX | DX, 0},
	x86.AMULSD:    {gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.AMULSS:    {gc.SizeF | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.ANEGB:     {gc.SizeB | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.ANEGL:     {gc.SizeL | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.ANEGQ:     {gc.SizeQ | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.ANEGW:     {gc.SizeW | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.ANOTB:     {gc.SizeB | RightRdwr, 0, 0, 0},
	x86.ANOTL:     {gc.SizeL | RightRdwr, 0, 0, 0},
	x86.ANOTQ:     {gc.SizeQ | RightRdwr, 0, 0, 0},
	x86.ANOTW:     {gc.SizeW | RightRdwr, 0, 0, 0},
	x86.AORB:      {gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AORL:      {gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AORQ:      {gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AORW:      {gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.APOPQ:     {gc.SizeQ | gc.RightWrite, 0, 0, 0},
	x86.APUSHQ:    {gc.SizeQ | gc.LeftRead, 0, 0, 0},
	x86.ARCLB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ARCLL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ARCLQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ARCLW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ARCRB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ARCRL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ARCRQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ARCRW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.AREP:      {gc.OK, CX, CX, 0},
	x86.AREPN:     {gc.OK, CX, CX, 0},
	obj.ARET:      {gc.Break | gc.KillCarry, 0, 0, 0},
	x86.AROLB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.AROLL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.AROLQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.AROLW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ARORB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ARORL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ARORQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ARORW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASALB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASALL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASALQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASALW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASARB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASARL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASARQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASARW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASBBB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ASBBL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ASBBQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ASBBW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	x86.ASETCC:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETCS:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETEQ:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETGE:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETGT:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETHI:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETLE:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETLS:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETLT:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETMI:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETNE:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETOC:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETOS:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETPC:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETPL:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASETPS:    {gc.SizeB | gc.RightWrite | gc.UseCarry, 0, 0, 0},
	x86.ASHLB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASHLL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASHLQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASHLW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASHRB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASHRL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASHRQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASHRW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	x86.ASQRTSD:   {gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.ASTOSB:    {gc.OK, AX | DI, DI, 0},
	x86.ASTOSL:    {gc.OK, AX | DI, DI, 0},
	x86.ASTOSQ:    {gc.OK, AX | DI, DI, 0},
	x86.ASTOSW:    {gc.OK, AX | DI, DI, 0},
	obj.ADUFFZERO: {gc.OK, AX | DI, DI, 0},
	x86.ASUBB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.ASUBL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.ASUBQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.ASUBW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.ASUBSD:    {gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.ASUBSS:    {gc.SizeF | gc.LeftRead | RightRdwr, 0, 0, 0},
	x86.ATESTB:    {gc.SizeB | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ATESTL:    {gc.SizeL | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ATESTQ:    {gc.SizeQ | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.ATESTW:    {gc.SizeW | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	x86.AUCOMISD:  {gc.SizeD | gc.LeftRead | gc.RightRead, 0, 0, 0},
	x86.AUCOMISS:  {gc.SizeF | gc.LeftRead | gc.RightRead, 0, 0, 0},
	x86.AXCHGB:    {gc.SizeB | LeftRdwr | RightRdwr, 0, 0, 0},
	x86.AXCHGL:    {gc.SizeL | LeftRdwr | RightRdwr, 0, 0, 0},
	x86.AXCHGQ:    {gc.SizeQ | LeftRdwr | RightRdwr, 0, 0, 0},
	x86.AXCHGW:    {gc.SizeW | LeftRdwr | RightRdwr, 0, 0, 0},
	x86.AXORB:     {gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AXORL:     {gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AXORQ:     {gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	x86.AXORW:     {gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
}

func progflags(p *obj.Prog) uint32 {
	flags := progtable[p.As].Flags
	if flags&gc.ImulAXDX != 0 && p.To.Type != obj.TYPE_NONE {
		flags |= RightRdwr
	}
	return flags
}

func progcarryflags(p *obj.Prog) uint32 {
	return progtable[p.As].Flags
}

func proginfo(p *obj.Prog) {
	info := &p.Info
	*info = progtable[p.As]
	if info.Flags == 0 {
		gc.Fatal("unknown instruction %v", p)
	}

	if (info.Flags&gc.ShiftCX != 0) && p.From.Type != obj.TYPE_CONST {
		info.Reguse |= CX
	}

	if info.Flags&gc.ImulAXDX != 0 {
		if p.To.Type == obj.TYPE_NONE {
			info.Reguse |= AX
			info.Regset |= AX | DX
		} else {
			info.Flags |= RightRdwr
		}
	}

	// Addressing makes some registers used.
	if p.From.Type == obj.TYPE_MEM && p.From.Name == obj.NAME_NONE {
		info.Regindex |= RtoB(int(p.From.Reg))
	}
	if p.From.Index != x86.REG_NONE {
		info.Regindex |= RtoB(int(p.From.Index))
	}
	if p.To.Type == obj.TYPE_MEM && p.To.Name == obj.NAME_NONE {
		info.Regindex |= RtoB(int(p.To.Reg))
	}
	if p.To.Index != x86.REG_NONE {
		info.Regindex |= RtoB(int(p.To.Index))
	}
	if gc.Ctxt.Flag_dynlink {
		// When -dynlink is passed, many operations on external names (and
		// also calling duffzero/duffcopy) use R15 as a scratch register.
		if p.As == x86.ALEAQ || info.Flags == gc.Pseudo || p.As == obj.ACALL || p.As == obj.ARET || p.As == obj.AJMP {
			return
		}
		if p.As == obj.ADUFFZERO || p.As == obj.ADUFFCOPY || p.From.Name == obj.NAME_EXTERN || p.To.Name == obj.NAME_EXTERN {
			info.Reguse |= R15
			info.Regset |= R15
			return
		}
	}
}
