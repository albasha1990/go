// Code generated by "stringer -type=Op -trimprefix=O"; DO NOT EDIT.

package ir

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OXXX-0]
	_ = x[ONAME-1]
	_ = x[ONONAME-2]
	_ = x[OTYPE-3]
	_ = x[OPACK-4]
	_ = x[OLITERAL-5]
	_ = x[ONIL-6]
	_ = x[OADD-7]
	_ = x[OSUB-8]
	_ = x[OOR-9]
	_ = x[OXOR-10]
	_ = x[OADDSTR-11]
	_ = x[OADDR-12]
	_ = x[OANDAND-13]
	_ = x[OAPPEND-14]
	_ = x[OBYTES2STR-15]
	_ = x[OBYTES2STRTMP-16]
	_ = x[ORUNES2STR-17]
	_ = x[OSTR2BYTES-18]
	_ = x[OSTR2BYTESTMP-19]
	_ = x[OSTR2RUNES-20]
	_ = x[OAS-21]
	_ = x[OAS2-22]
	_ = x[OAS2DOTTYPE-23]
	_ = x[OAS2FUNC-24]
	_ = x[OAS2MAPR-25]
	_ = x[OAS2RECV-26]
	_ = x[OASOP-27]
	_ = x[OCALL-28]
	_ = x[OCALLFUNC-29]
	_ = x[OCALLMETH-30]
	_ = x[OCALLINTER-31]
	_ = x[OCALLPART-32]
	_ = x[OCAP-33]
	_ = x[OCLOSE-34]
	_ = x[OCLOSURE-35]
	_ = x[OCOMPLIT-36]
	_ = x[OMAPLIT-37]
	_ = x[OSTRUCTLIT-38]
	_ = x[OARRAYLIT-39]
	_ = x[OSLICELIT-40]
	_ = x[OPTRLIT-41]
	_ = x[OCONV-42]
	_ = x[OCONVIFACE-43]
	_ = x[OCONVNOP-44]
	_ = x[OCOPY-45]
	_ = x[ODCL-46]
	_ = x[ODCLFUNC-47]
	_ = x[ODCLCONST-48]
	_ = x[ODCLTYPE-49]
	_ = x[ODELETE-50]
	_ = x[ODOT-51]
	_ = x[ODOTPTR-52]
	_ = x[ODOTMETH-53]
	_ = x[ODOTINTER-54]
	_ = x[OXDOT-55]
	_ = x[ODOTTYPE-56]
	_ = x[ODOTTYPE2-57]
	_ = x[OEQ-58]
	_ = x[ONE-59]
	_ = x[OLT-60]
	_ = x[OLE-61]
	_ = x[OGE-62]
	_ = x[OGT-63]
	_ = x[ODEREF-64]
	_ = x[OINDEX-65]
	_ = x[OINDEXMAP-66]
	_ = x[OKEY-67]
	_ = x[OSTRUCTKEY-68]
	_ = x[OLEN-69]
	_ = x[OMAKE-70]
	_ = x[OMAKECHAN-71]
	_ = x[OMAKEMAP-72]
	_ = x[OMAKESLICE-73]
	_ = x[OMAKESLICECOPY-74]
	_ = x[OMUL-75]
	_ = x[ODIV-76]
	_ = x[OMOD-77]
	_ = x[OLSH-78]
	_ = x[ORSH-79]
	_ = x[OAND-80]
	_ = x[OANDNOT-81]
	_ = x[ONEW-82]
	_ = x[ONEWOBJ-83]
	_ = x[ONOT-84]
	_ = x[OBITNOT-85]
	_ = x[OPLUS-86]
	_ = x[ONEG-87]
	_ = x[OOROR-88]
	_ = x[OPANIC-89]
	_ = x[OPRINT-90]
	_ = x[OPRINTN-91]
	_ = x[OPAREN-92]
	_ = x[OSEND-93]
	_ = x[OSLICE-94]
	_ = x[OSLICEARR-95]
	_ = x[OSLICESTR-96]
	_ = x[OSLICE3-97]
	_ = x[OSLICE3ARR-98]
	_ = x[OSLICEHEADER-99]
	_ = x[ORECOVER-100]
	_ = x[ORECV-101]
	_ = x[ORUNESTR-102]
	_ = x[OSELRECV-103]
	_ = x[OSELRECV2-104]
	_ = x[OIOTA-105]
	_ = x[OREAL-106]
	_ = x[OIMAG-107]
	_ = x[OCOMPLEX-108]
	_ = x[OALIGNOF-109]
	_ = x[OOFFSETOF-110]
	_ = x[OSIZEOF-111]
	_ = x[OMETHEXPR-112]
	_ = x[OSTMTEXPR-113]
	_ = x[OBLOCK-114]
	_ = x[OBREAK-115]
	_ = x[OCASE-116]
	_ = x[OCONTINUE-117]
	_ = x[ODEFER-118]
	_ = x[OEMPTY-119]
	_ = x[OFALL-120]
	_ = x[OFOR-121]
	_ = x[OFORUNTIL-122]
	_ = x[OGOTO-123]
	_ = x[OIF-124]
	_ = x[OLABEL-125]
	_ = x[OGO-126]
	_ = x[ORANGE-127]
	_ = x[ORETURN-128]
	_ = x[OSELECT-129]
	_ = x[OSWITCH-130]
	_ = x[OTYPESW-131]
	_ = x[OTCHAN-132]
	_ = x[OTMAP-133]
	_ = x[OTSTRUCT-134]
	_ = x[OTINTER-135]
	_ = x[OTFUNC-136]
	_ = x[OTARRAY-137]
	_ = x[OTSLICE-138]
	_ = x[OINLCALL-139]
	_ = x[OEFACE-140]
	_ = x[OITAB-141]
	_ = x[OIDATA-142]
	_ = x[OSPTR-143]
	_ = x[OCLOSUREREAD-144]
	_ = x[OCFUNC-145]
	_ = x[OCHECKNIL-146]
	_ = x[OVARDEF-147]
	_ = x[OVARKILL-148]
	_ = x[OVARLIVE-149]
	_ = x[ORESULT-150]
	_ = x[OINLMARK-151]
	_ = x[ORETJMP-152]
	_ = x[OGETG-153]
	_ = x[OEND-154]
}

const _Op_name = "XXXNAMENONAMETYPEPACKLITERALNILADDSUBORXORADDSTRADDRANDANDAPPENDBYTES2STRBYTES2STRTMPRUNES2STRSTR2BYTESSTR2BYTESTMPSTR2RUNESASAS2AS2DOTTYPEAS2FUNCAS2MAPRAS2RECVASOPCALLCALLFUNCCALLMETHCALLINTERCALLPARTCAPCLOSECLOSURECOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVNOPCOPYDCLDCLFUNCDCLCONSTDCLTYPEDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTDEREFINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMAKESLICECOPYMULDIVMODLSHRSHANDANDNOTNEWNEWOBJNOTBITNOTPLUSNEGORORPANICPRINTPRINTNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRSLICEHEADERRECOVERRECVRUNESTRSELRECVSELRECV2IOTAREALIMAGCOMPLEXALIGNOFOFFSETOFSIZEOFMETHEXPRSTMTEXPRBLOCKBREAKCASECONTINUEDEFEREMPTYFALLFORFORUNTILGOTOIFLABELGORANGERETURNSELECTSWITCHTYPESWTCHANTMAPTSTRUCTTINTERTFUNCTARRAYTSLICEINLCALLEFACEITABIDATASPTRCLOSUREREADCFUNCCHECKNILVARDEFVARKILLVARLIVERESULTINLMARKRETJMPGETGEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 21, 28, 31, 34, 37, 39, 42, 48, 52, 58, 64, 73, 85, 94, 103, 115, 124, 126, 129, 139, 146, 153, 160, 164, 168, 176, 184, 193, 201, 204, 209, 216, 223, 229, 238, 246, 254, 260, 264, 273, 280, 284, 287, 294, 302, 309, 315, 318, 324, 331, 339, 343, 350, 358, 360, 362, 364, 366, 368, 370, 375, 380, 388, 391, 400, 403, 407, 415, 422, 431, 444, 447, 450, 453, 456, 459, 462, 468, 471, 477, 480, 486, 490, 493, 497, 502, 507, 513, 518, 522, 527, 535, 543, 549, 558, 569, 576, 580, 587, 594, 602, 606, 610, 614, 621, 628, 636, 642, 650, 658, 663, 668, 672, 680, 685, 690, 694, 697, 705, 709, 711, 716, 718, 723, 729, 735, 741, 747, 752, 756, 763, 769, 774, 780, 786, 793, 798, 802, 807, 811, 822, 827, 835, 841, 848, 855, 861, 868, 874, 878, 881}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
