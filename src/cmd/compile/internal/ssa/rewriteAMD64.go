// Code generated from gen/AMD64.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "math"
import "cmd/compile/internal/types"

func rewriteValueAMD64(v *Value) bool {
	switch v.Op {
	case OpAMD64ADCQ:
		return rewriteValueAMD64_OpAMD64ADCQ_0(v)
	case OpAMD64ADCQconst:
		return rewriteValueAMD64_OpAMD64ADCQconst_0(v)
	case OpAMD64ADDL:
		return rewriteValueAMD64_OpAMD64ADDL_0(v) || rewriteValueAMD64_OpAMD64ADDL_10(v)
	case OpAMD64ADDLconst:
		return rewriteValueAMD64_OpAMD64ADDLconst_0(v) || rewriteValueAMD64_OpAMD64ADDLconst_10(v)
	case OpAMD64ADDLconstmodify:
		return rewriteValueAMD64_OpAMD64ADDLconstmodify_0(v)
	case OpAMD64ADDLload:
		return rewriteValueAMD64_OpAMD64ADDLload_0(v)
	case OpAMD64ADDLmodify:
		return rewriteValueAMD64_OpAMD64ADDLmodify_0(v)
	case OpAMD64ADDQ:
		return rewriteValueAMD64_OpAMD64ADDQ_0(v) || rewriteValueAMD64_OpAMD64ADDQ_10(v)
	case OpAMD64ADDQcarry:
		return rewriteValueAMD64_OpAMD64ADDQcarry_0(v)
	case OpAMD64ADDQconst:
		return rewriteValueAMD64_OpAMD64ADDQconst_0(v) || rewriteValueAMD64_OpAMD64ADDQconst_10(v)
	case OpAMD64ADDQconstmodify:
		return rewriteValueAMD64_OpAMD64ADDQconstmodify_0(v)
	case OpAMD64ADDQload:
		return rewriteValueAMD64_OpAMD64ADDQload_0(v)
	case OpAMD64ADDQmodify:
		return rewriteValueAMD64_OpAMD64ADDQmodify_0(v)
	case OpAMD64ADDSD:
		return rewriteValueAMD64_OpAMD64ADDSD_0(v)
	case OpAMD64ADDSDload:
		return rewriteValueAMD64_OpAMD64ADDSDload_0(v)
	case OpAMD64ADDSS:
		return rewriteValueAMD64_OpAMD64ADDSS_0(v)
	case OpAMD64ADDSSload:
		return rewriteValueAMD64_OpAMD64ADDSSload_0(v)
	case OpAMD64ANDL:
		return rewriteValueAMD64_OpAMD64ANDL_0(v)
	case OpAMD64ANDLconst:
		return rewriteValueAMD64_OpAMD64ANDLconst_0(v)
	case OpAMD64ANDLconstmodify:
		return rewriteValueAMD64_OpAMD64ANDLconstmodify_0(v)
	case OpAMD64ANDLload:
		return rewriteValueAMD64_OpAMD64ANDLload_0(v)
	case OpAMD64ANDLmodify:
		return rewriteValueAMD64_OpAMD64ANDLmodify_0(v)
	case OpAMD64ANDQ:
		return rewriteValueAMD64_OpAMD64ANDQ_0(v)
	case OpAMD64ANDQconst:
		return rewriteValueAMD64_OpAMD64ANDQconst_0(v)
	case OpAMD64ANDQconstmodify:
		return rewriteValueAMD64_OpAMD64ANDQconstmodify_0(v)
	case OpAMD64ANDQload:
		return rewriteValueAMD64_OpAMD64ANDQload_0(v)
	case OpAMD64ANDQmodify:
		return rewriteValueAMD64_OpAMD64ANDQmodify_0(v)
	case OpAMD64BSFQ:
		return rewriteValueAMD64_OpAMD64BSFQ_0(v)
	case OpAMD64BTCLconst:
		return rewriteValueAMD64_OpAMD64BTCLconst_0(v)
	case OpAMD64BTCLconstmodify:
		return rewriteValueAMD64_OpAMD64BTCLconstmodify_0(v)
	case OpAMD64BTCLmodify:
		return rewriteValueAMD64_OpAMD64BTCLmodify_0(v)
	case OpAMD64BTCQconst:
		return rewriteValueAMD64_OpAMD64BTCQconst_0(v)
	case OpAMD64BTCQconstmodify:
		return rewriteValueAMD64_OpAMD64BTCQconstmodify_0(v)
	case OpAMD64BTCQmodify:
		return rewriteValueAMD64_OpAMD64BTCQmodify_0(v)
	case OpAMD64BTLconst:
		return rewriteValueAMD64_OpAMD64BTLconst_0(v)
	case OpAMD64BTQconst:
		return rewriteValueAMD64_OpAMD64BTQconst_0(v)
	case OpAMD64BTRLconst:
		return rewriteValueAMD64_OpAMD64BTRLconst_0(v)
	case OpAMD64BTRLconstmodify:
		return rewriteValueAMD64_OpAMD64BTRLconstmodify_0(v)
	case OpAMD64BTRLmodify:
		return rewriteValueAMD64_OpAMD64BTRLmodify_0(v)
	case OpAMD64BTRQconst:
		return rewriteValueAMD64_OpAMD64BTRQconst_0(v)
	case OpAMD64BTRQconstmodify:
		return rewriteValueAMD64_OpAMD64BTRQconstmodify_0(v)
	case OpAMD64BTRQmodify:
		return rewriteValueAMD64_OpAMD64BTRQmodify_0(v)
	case OpAMD64BTSLconst:
		return rewriteValueAMD64_OpAMD64BTSLconst_0(v)
	case OpAMD64BTSLconstmodify:
		return rewriteValueAMD64_OpAMD64BTSLconstmodify_0(v)
	case OpAMD64BTSLmodify:
		return rewriteValueAMD64_OpAMD64BTSLmodify_0(v)
	case OpAMD64BTSQconst:
		return rewriteValueAMD64_OpAMD64BTSQconst_0(v)
	case OpAMD64BTSQconstmodify:
		return rewriteValueAMD64_OpAMD64BTSQconstmodify_0(v)
	case OpAMD64BTSQmodify:
		return rewriteValueAMD64_OpAMD64BTSQmodify_0(v)
	case OpAMD64CMOVLCC:
		return rewriteValueAMD64_OpAMD64CMOVLCC_0(v)
	case OpAMD64CMOVLCS:
		return rewriteValueAMD64_OpAMD64CMOVLCS_0(v)
	case OpAMD64CMOVLEQ:
		return rewriteValueAMD64_OpAMD64CMOVLEQ_0(v)
	case OpAMD64CMOVLGE:
		return rewriteValueAMD64_OpAMD64CMOVLGE_0(v)
	case OpAMD64CMOVLGT:
		return rewriteValueAMD64_OpAMD64CMOVLGT_0(v)
	case OpAMD64CMOVLHI:
		return rewriteValueAMD64_OpAMD64CMOVLHI_0(v)
	case OpAMD64CMOVLLE:
		return rewriteValueAMD64_OpAMD64CMOVLLE_0(v)
	case OpAMD64CMOVLLS:
		return rewriteValueAMD64_OpAMD64CMOVLLS_0(v)
	case OpAMD64CMOVLLT:
		return rewriteValueAMD64_OpAMD64CMOVLLT_0(v)
	case OpAMD64CMOVLNE:
		return rewriteValueAMD64_OpAMD64CMOVLNE_0(v)
	case OpAMD64CMOVQCC:
		return rewriteValueAMD64_OpAMD64CMOVQCC_0(v)
	case OpAMD64CMOVQCS:
		return rewriteValueAMD64_OpAMD64CMOVQCS_0(v)
	case OpAMD64CMOVQEQ:
		return rewriteValueAMD64_OpAMD64CMOVQEQ_0(v)
	case OpAMD64CMOVQGE:
		return rewriteValueAMD64_OpAMD64CMOVQGE_0(v)
	case OpAMD64CMOVQGT:
		return rewriteValueAMD64_OpAMD64CMOVQGT_0(v)
	case OpAMD64CMOVQHI:
		return rewriteValueAMD64_OpAMD64CMOVQHI_0(v)
	case OpAMD64CMOVQLE:
		return rewriteValueAMD64_OpAMD64CMOVQLE_0(v)
	case OpAMD64CMOVQLS:
		return rewriteValueAMD64_OpAMD64CMOVQLS_0(v)
	case OpAMD64CMOVQLT:
		return rewriteValueAMD64_OpAMD64CMOVQLT_0(v)
	case OpAMD64CMOVQNE:
		return rewriteValueAMD64_OpAMD64CMOVQNE_0(v)
	case OpAMD64CMOVWCC:
		return rewriteValueAMD64_OpAMD64CMOVWCC_0(v)
	case OpAMD64CMOVWCS:
		return rewriteValueAMD64_OpAMD64CMOVWCS_0(v)
	case OpAMD64CMOVWEQ:
		return rewriteValueAMD64_OpAMD64CMOVWEQ_0(v)
	case OpAMD64CMOVWGE:
		return rewriteValueAMD64_OpAMD64CMOVWGE_0(v)
	case OpAMD64CMOVWGT:
		return rewriteValueAMD64_OpAMD64CMOVWGT_0(v)
	case OpAMD64CMOVWHI:
		return rewriteValueAMD64_OpAMD64CMOVWHI_0(v)
	case OpAMD64CMOVWLE:
		return rewriteValueAMD64_OpAMD64CMOVWLE_0(v)
	case OpAMD64CMOVWLS:
		return rewriteValueAMD64_OpAMD64CMOVWLS_0(v)
	case OpAMD64CMOVWLT:
		return rewriteValueAMD64_OpAMD64CMOVWLT_0(v)
	case OpAMD64CMOVWNE:
		return rewriteValueAMD64_OpAMD64CMOVWNE_0(v)
	case OpAMD64CMPB:
		return rewriteValueAMD64_OpAMD64CMPB_0(v)
	case OpAMD64CMPBconst:
		return rewriteValueAMD64_OpAMD64CMPBconst_0(v)
	case OpAMD64CMPBconstload:
		return rewriteValueAMD64_OpAMD64CMPBconstload_0(v)
	case OpAMD64CMPBload:
		return rewriteValueAMD64_OpAMD64CMPBload_0(v)
	case OpAMD64CMPL:
		return rewriteValueAMD64_OpAMD64CMPL_0(v)
	case OpAMD64CMPLconst:
		return rewriteValueAMD64_OpAMD64CMPLconst_0(v) || rewriteValueAMD64_OpAMD64CMPLconst_10(v)
	case OpAMD64CMPLconstload:
		return rewriteValueAMD64_OpAMD64CMPLconstload_0(v)
	case OpAMD64CMPLload:
		return rewriteValueAMD64_OpAMD64CMPLload_0(v)
	case OpAMD64CMPQ:
		return rewriteValueAMD64_OpAMD64CMPQ_0(v)
	case OpAMD64CMPQconst:
		return rewriteValueAMD64_OpAMD64CMPQconst_0(v) || rewriteValueAMD64_OpAMD64CMPQconst_10(v)
	case OpAMD64CMPQconstload:
		return rewriteValueAMD64_OpAMD64CMPQconstload_0(v)
	case OpAMD64CMPQload:
		return rewriteValueAMD64_OpAMD64CMPQload_0(v)
	case OpAMD64CMPW:
		return rewriteValueAMD64_OpAMD64CMPW_0(v)
	case OpAMD64CMPWconst:
		return rewriteValueAMD64_OpAMD64CMPWconst_0(v)
	case OpAMD64CMPWconstload:
		return rewriteValueAMD64_OpAMD64CMPWconstload_0(v)
	case OpAMD64CMPWload:
		return rewriteValueAMD64_OpAMD64CMPWload_0(v)
	case OpAMD64CMPXCHGLlock:
		return rewriteValueAMD64_OpAMD64CMPXCHGLlock_0(v)
	case OpAMD64CMPXCHGQlock:
		return rewriteValueAMD64_OpAMD64CMPXCHGQlock_0(v)
	case OpAMD64DIVSD:
		return rewriteValueAMD64_OpAMD64DIVSD_0(v)
	case OpAMD64DIVSDload:
		return rewriteValueAMD64_OpAMD64DIVSDload_0(v)
	case OpAMD64DIVSS:
		return rewriteValueAMD64_OpAMD64DIVSS_0(v)
	case OpAMD64DIVSSload:
		return rewriteValueAMD64_OpAMD64DIVSSload_0(v)
	case OpAMD64HMULL:
		return rewriteValueAMD64_OpAMD64HMULL_0(v)
	case OpAMD64HMULLU:
		return rewriteValueAMD64_OpAMD64HMULLU_0(v)
	case OpAMD64HMULQ:
		return rewriteValueAMD64_OpAMD64HMULQ_0(v)
	case OpAMD64HMULQU:
		return rewriteValueAMD64_OpAMD64HMULQU_0(v)
	case OpAMD64LEAL:
		return rewriteValueAMD64_OpAMD64LEAL_0(v)
	case OpAMD64LEAL1:
		return rewriteValueAMD64_OpAMD64LEAL1_0(v)
	case OpAMD64LEAL2:
		return rewriteValueAMD64_OpAMD64LEAL2_0(v)
	case OpAMD64LEAL4:
		return rewriteValueAMD64_OpAMD64LEAL4_0(v)
	case OpAMD64LEAL8:
		return rewriteValueAMD64_OpAMD64LEAL8_0(v)
	case OpAMD64LEAQ:
		return rewriteValueAMD64_OpAMD64LEAQ_0(v)
	case OpAMD64LEAQ1:
		return rewriteValueAMD64_OpAMD64LEAQ1_0(v)
	case OpAMD64LEAQ2:
		return rewriteValueAMD64_OpAMD64LEAQ2_0(v)
	case OpAMD64LEAQ4:
		return rewriteValueAMD64_OpAMD64LEAQ4_0(v)
	case OpAMD64LEAQ8:
		return rewriteValueAMD64_OpAMD64LEAQ8_0(v)
	case OpAMD64MOVBQSX:
		return rewriteValueAMD64_OpAMD64MOVBQSX_0(v)
	case OpAMD64MOVBQSXload:
		return rewriteValueAMD64_OpAMD64MOVBQSXload_0(v)
	case OpAMD64MOVBQZX:
		return rewriteValueAMD64_OpAMD64MOVBQZX_0(v)
	case OpAMD64MOVBatomicload:
		return rewriteValueAMD64_OpAMD64MOVBatomicload_0(v)
	case OpAMD64MOVBload:
		return rewriteValueAMD64_OpAMD64MOVBload_0(v)
	case OpAMD64MOVBloadidx1:
		return rewriteValueAMD64_OpAMD64MOVBloadidx1_0(v)
	case OpAMD64MOVBstore:
		return rewriteValueAMD64_OpAMD64MOVBstore_0(v) || rewriteValueAMD64_OpAMD64MOVBstore_10(v) || rewriteValueAMD64_OpAMD64MOVBstore_20(v) || rewriteValueAMD64_OpAMD64MOVBstore_30(v)
	case OpAMD64MOVBstoreconst:
		return rewriteValueAMD64_OpAMD64MOVBstoreconst_0(v)
	case OpAMD64MOVBstoreconstidx1:
		return rewriteValueAMD64_OpAMD64MOVBstoreconstidx1_0(v)
	case OpAMD64MOVBstoreidx1:
		return rewriteValueAMD64_OpAMD64MOVBstoreidx1_0(v) || rewriteValueAMD64_OpAMD64MOVBstoreidx1_10(v)
	case OpAMD64MOVLQSX:
		return rewriteValueAMD64_OpAMD64MOVLQSX_0(v)
	case OpAMD64MOVLQSXload:
		return rewriteValueAMD64_OpAMD64MOVLQSXload_0(v)
	case OpAMD64MOVLQZX:
		return rewriteValueAMD64_OpAMD64MOVLQZX_0(v)
	case OpAMD64MOVLatomicload:
		return rewriteValueAMD64_OpAMD64MOVLatomicload_0(v)
	case OpAMD64MOVLf2i:
		return rewriteValueAMD64_OpAMD64MOVLf2i_0(v)
	case OpAMD64MOVLi2f:
		return rewriteValueAMD64_OpAMD64MOVLi2f_0(v)
	case OpAMD64MOVLload:
		return rewriteValueAMD64_OpAMD64MOVLload_0(v) || rewriteValueAMD64_OpAMD64MOVLload_10(v)
	case OpAMD64MOVLloadidx1:
		return rewriteValueAMD64_OpAMD64MOVLloadidx1_0(v)
	case OpAMD64MOVLloadidx4:
		return rewriteValueAMD64_OpAMD64MOVLloadidx4_0(v)
	case OpAMD64MOVLloadidx8:
		return rewriteValueAMD64_OpAMD64MOVLloadidx8_0(v)
	case OpAMD64MOVLstore:
		return rewriteValueAMD64_OpAMD64MOVLstore_0(v) || rewriteValueAMD64_OpAMD64MOVLstore_10(v) || rewriteValueAMD64_OpAMD64MOVLstore_20(v) || rewriteValueAMD64_OpAMD64MOVLstore_30(v)
	case OpAMD64MOVLstoreconst:
		return rewriteValueAMD64_OpAMD64MOVLstoreconst_0(v)
	case OpAMD64MOVLstoreconstidx1:
		return rewriteValueAMD64_OpAMD64MOVLstoreconstidx1_0(v)
	case OpAMD64MOVLstoreconstidx4:
		return rewriteValueAMD64_OpAMD64MOVLstoreconstidx4_0(v)
	case OpAMD64MOVLstoreidx1:
		return rewriteValueAMD64_OpAMD64MOVLstoreidx1_0(v)
	case OpAMD64MOVLstoreidx4:
		return rewriteValueAMD64_OpAMD64MOVLstoreidx4_0(v)
	case OpAMD64MOVLstoreidx8:
		return rewriteValueAMD64_OpAMD64MOVLstoreidx8_0(v)
	case OpAMD64MOVOload:
		return rewriteValueAMD64_OpAMD64MOVOload_0(v)
	case OpAMD64MOVOstore:
		return rewriteValueAMD64_OpAMD64MOVOstore_0(v)
	case OpAMD64MOVQatomicload:
		return rewriteValueAMD64_OpAMD64MOVQatomicload_0(v)
	case OpAMD64MOVQf2i:
		return rewriteValueAMD64_OpAMD64MOVQf2i_0(v)
	case OpAMD64MOVQi2f:
		return rewriteValueAMD64_OpAMD64MOVQi2f_0(v)
	case OpAMD64MOVQload:
		return rewriteValueAMD64_OpAMD64MOVQload_0(v)
	case OpAMD64MOVQloadidx1:
		return rewriteValueAMD64_OpAMD64MOVQloadidx1_0(v)
	case OpAMD64MOVQloadidx8:
		return rewriteValueAMD64_OpAMD64MOVQloadidx8_0(v)
	case OpAMD64MOVQstore:
		return rewriteValueAMD64_OpAMD64MOVQstore_0(v) || rewriteValueAMD64_OpAMD64MOVQstore_10(v) || rewriteValueAMD64_OpAMD64MOVQstore_20(v)
	case OpAMD64MOVQstoreconst:
		return rewriteValueAMD64_OpAMD64MOVQstoreconst_0(v)
	case OpAMD64MOVQstoreconstidx1:
		return rewriteValueAMD64_OpAMD64MOVQstoreconstidx1_0(v)
	case OpAMD64MOVQstoreconstidx8:
		return rewriteValueAMD64_OpAMD64MOVQstoreconstidx8_0(v)
	case OpAMD64MOVQstoreidx1:
		return rewriteValueAMD64_OpAMD64MOVQstoreidx1_0(v)
	case OpAMD64MOVQstoreidx8:
		return rewriteValueAMD64_OpAMD64MOVQstoreidx8_0(v)
	case OpAMD64MOVSDload:
		return rewriteValueAMD64_OpAMD64MOVSDload_0(v)
	case OpAMD64MOVSDloadidx1:
		return rewriteValueAMD64_OpAMD64MOVSDloadidx1_0(v)
	case OpAMD64MOVSDloadidx8:
		return rewriteValueAMD64_OpAMD64MOVSDloadidx8_0(v)
	case OpAMD64MOVSDstore:
		return rewriteValueAMD64_OpAMD64MOVSDstore_0(v)
	case OpAMD64MOVSDstoreidx1:
		return rewriteValueAMD64_OpAMD64MOVSDstoreidx1_0(v)
	case OpAMD64MOVSDstoreidx8:
		return rewriteValueAMD64_OpAMD64MOVSDstoreidx8_0(v)
	case OpAMD64MOVSSload:
		return rewriteValueAMD64_OpAMD64MOVSSload_0(v)
	case OpAMD64MOVSSloadidx1:
		return rewriteValueAMD64_OpAMD64MOVSSloadidx1_0(v)
	case OpAMD64MOVSSloadidx4:
		return rewriteValueAMD64_OpAMD64MOVSSloadidx4_0(v)
	case OpAMD64MOVSSstore:
		return rewriteValueAMD64_OpAMD64MOVSSstore_0(v)
	case OpAMD64MOVSSstoreidx1:
		return rewriteValueAMD64_OpAMD64MOVSSstoreidx1_0(v)
	case OpAMD64MOVSSstoreidx4:
		return rewriteValueAMD64_OpAMD64MOVSSstoreidx4_0(v)
	case OpAMD64MOVWQSX:
		return rewriteValueAMD64_OpAMD64MOVWQSX_0(v)
	case OpAMD64MOVWQSXload:
		return rewriteValueAMD64_OpAMD64MOVWQSXload_0(v)
	case OpAMD64MOVWQZX:
		return rewriteValueAMD64_OpAMD64MOVWQZX_0(v)
	case OpAMD64MOVWload:
		return rewriteValueAMD64_OpAMD64MOVWload_0(v)
	case OpAMD64MOVWloadidx1:
		return rewriteValueAMD64_OpAMD64MOVWloadidx1_0(v)
	case OpAMD64MOVWloadidx2:
		return rewriteValueAMD64_OpAMD64MOVWloadidx2_0(v)
	case OpAMD64MOVWstore:
		return rewriteValueAMD64_OpAMD64MOVWstore_0(v) || rewriteValueAMD64_OpAMD64MOVWstore_10(v)
	case OpAMD64MOVWstoreconst:
		return rewriteValueAMD64_OpAMD64MOVWstoreconst_0(v)
	case OpAMD64MOVWstoreconstidx1:
		return rewriteValueAMD64_OpAMD64MOVWstoreconstidx1_0(v)
	case OpAMD64MOVWstoreconstidx2:
		return rewriteValueAMD64_OpAMD64MOVWstoreconstidx2_0(v)
	case OpAMD64MOVWstoreidx1:
		return rewriteValueAMD64_OpAMD64MOVWstoreidx1_0(v)
	case OpAMD64MOVWstoreidx2:
		return rewriteValueAMD64_OpAMD64MOVWstoreidx2_0(v)
	case OpAMD64MULL:
		return rewriteValueAMD64_OpAMD64MULL_0(v)
	case OpAMD64MULLconst:
		return rewriteValueAMD64_OpAMD64MULLconst_0(v) || rewriteValueAMD64_OpAMD64MULLconst_10(v) || rewriteValueAMD64_OpAMD64MULLconst_20(v) || rewriteValueAMD64_OpAMD64MULLconst_30(v)
	case OpAMD64MULQ:
		return rewriteValueAMD64_OpAMD64MULQ_0(v)
	case OpAMD64MULQconst:
		return rewriteValueAMD64_OpAMD64MULQconst_0(v) || rewriteValueAMD64_OpAMD64MULQconst_10(v) || rewriteValueAMD64_OpAMD64MULQconst_20(v) || rewriteValueAMD64_OpAMD64MULQconst_30(v)
	case OpAMD64MULSD:
		return rewriteValueAMD64_OpAMD64MULSD_0(v)
	case OpAMD64MULSDload:
		return rewriteValueAMD64_OpAMD64MULSDload_0(v)
	case OpAMD64MULSS:
		return rewriteValueAMD64_OpAMD64MULSS_0(v)
	case OpAMD64MULSSload:
		return rewriteValueAMD64_OpAMD64MULSSload_0(v)
	case OpAMD64NEGL:
		return rewriteValueAMD64_OpAMD64NEGL_0(v)
	case OpAMD64NEGQ:
		return rewriteValueAMD64_OpAMD64NEGQ_0(v)
	case OpAMD64NOTL:
		return rewriteValueAMD64_OpAMD64NOTL_0(v)
	case OpAMD64NOTQ:
		return rewriteValueAMD64_OpAMD64NOTQ_0(v)
	case OpAMD64ORL:
		return rewriteValueAMD64_OpAMD64ORL_0(v) || rewriteValueAMD64_OpAMD64ORL_10(v) || rewriteValueAMD64_OpAMD64ORL_20(v) || rewriteValueAMD64_OpAMD64ORL_30(v)
	case OpAMD64ORLconst:
		return rewriteValueAMD64_OpAMD64ORLconst_0(v)
	case OpAMD64ORLconstmodify:
		return rewriteValueAMD64_OpAMD64ORLconstmodify_0(v)
	case OpAMD64ORLload:
		return rewriteValueAMD64_OpAMD64ORLload_0(v)
	case OpAMD64ORLmodify:
		return rewriteValueAMD64_OpAMD64ORLmodify_0(v)
	case OpAMD64ORQ:
		return rewriteValueAMD64_OpAMD64ORQ_0(v) || rewriteValueAMD64_OpAMD64ORQ_10(v) || rewriteValueAMD64_OpAMD64ORQ_20(v)
	case OpAMD64ORQconst:
		return rewriteValueAMD64_OpAMD64ORQconst_0(v)
	case OpAMD64ORQconstmodify:
		return rewriteValueAMD64_OpAMD64ORQconstmodify_0(v)
	case OpAMD64ORQload:
		return rewriteValueAMD64_OpAMD64ORQload_0(v)
	case OpAMD64ORQmodify:
		return rewriteValueAMD64_OpAMD64ORQmodify_0(v)
	case OpAMD64ROLB:
		return rewriteValueAMD64_OpAMD64ROLB_0(v)
	case OpAMD64ROLBconst:
		return rewriteValueAMD64_OpAMD64ROLBconst_0(v)
	case OpAMD64ROLL:
		return rewriteValueAMD64_OpAMD64ROLL_0(v)
	case OpAMD64ROLLconst:
		return rewriteValueAMD64_OpAMD64ROLLconst_0(v)
	case OpAMD64ROLQ:
		return rewriteValueAMD64_OpAMD64ROLQ_0(v)
	case OpAMD64ROLQconst:
		return rewriteValueAMD64_OpAMD64ROLQconst_0(v)
	case OpAMD64ROLW:
		return rewriteValueAMD64_OpAMD64ROLW_0(v)
	case OpAMD64ROLWconst:
		return rewriteValueAMD64_OpAMD64ROLWconst_0(v)
	case OpAMD64RORB:
		return rewriteValueAMD64_OpAMD64RORB_0(v)
	case OpAMD64RORL:
		return rewriteValueAMD64_OpAMD64RORL_0(v)
	case OpAMD64RORQ:
		return rewriteValueAMD64_OpAMD64RORQ_0(v)
	case OpAMD64RORW:
		return rewriteValueAMD64_OpAMD64RORW_0(v)
	case OpAMD64SARB:
		return rewriteValueAMD64_OpAMD64SARB_0(v)
	case OpAMD64SARBconst:
		return rewriteValueAMD64_OpAMD64SARBconst_0(v)
	case OpAMD64SARL:
		return rewriteValueAMD64_OpAMD64SARL_0(v)
	case OpAMD64SARLconst:
		return rewriteValueAMD64_OpAMD64SARLconst_0(v)
	case OpAMD64SARQ:
		return rewriteValueAMD64_OpAMD64SARQ_0(v)
	case OpAMD64SARQconst:
		return rewriteValueAMD64_OpAMD64SARQconst_0(v)
	case OpAMD64SARW:
		return rewriteValueAMD64_OpAMD64SARW_0(v)
	case OpAMD64SARWconst:
		return rewriteValueAMD64_OpAMD64SARWconst_0(v)
	case OpAMD64SBBLcarrymask:
		return rewriteValueAMD64_OpAMD64SBBLcarrymask_0(v)
	case OpAMD64SBBQ:
		return rewriteValueAMD64_OpAMD64SBBQ_0(v)
	case OpAMD64SBBQcarrymask:
		return rewriteValueAMD64_OpAMD64SBBQcarrymask_0(v)
	case OpAMD64SBBQconst:
		return rewriteValueAMD64_OpAMD64SBBQconst_0(v)
	case OpAMD64SETA:
		return rewriteValueAMD64_OpAMD64SETA_0(v)
	case OpAMD64SETAE:
		return rewriteValueAMD64_OpAMD64SETAE_0(v)
	case OpAMD64SETAEstore:
		return rewriteValueAMD64_OpAMD64SETAEstore_0(v)
	case OpAMD64SETAstore:
		return rewriteValueAMD64_OpAMD64SETAstore_0(v)
	case OpAMD64SETB:
		return rewriteValueAMD64_OpAMD64SETB_0(v)
	case OpAMD64SETBE:
		return rewriteValueAMD64_OpAMD64SETBE_0(v)
	case OpAMD64SETBEstore:
		return rewriteValueAMD64_OpAMD64SETBEstore_0(v)
	case OpAMD64SETBstore:
		return rewriteValueAMD64_OpAMD64SETBstore_0(v)
	case OpAMD64SETEQ:
		return rewriteValueAMD64_OpAMD64SETEQ_0(v) || rewriteValueAMD64_OpAMD64SETEQ_10(v)
	case OpAMD64SETEQstore:
		return rewriteValueAMD64_OpAMD64SETEQstore_0(v) || rewriteValueAMD64_OpAMD64SETEQstore_10(v) || rewriteValueAMD64_OpAMD64SETEQstore_20(v)
	case OpAMD64SETG:
		return rewriteValueAMD64_OpAMD64SETG_0(v)
	case OpAMD64SETGE:
		return rewriteValueAMD64_OpAMD64SETGE_0(v)
	case OpAMD64SETGEstore:
		return rewriteValueAMD64_OpAMD64SETGEstore_0(v)
	case OpAMD64SETGstore:
		return rewriteValueAMD64_OpAMD64SETGstore_0(v)
	case OpAMD64SETL:
		return rewriteValueAMD64_OpAMD64SETL_0(v)
	case OpAMD64SETLE:
		return rewriteValueAMD64_OpAMD64SETLE_0(v)
	case OpAMD64SETLEstore:
		return rewriteValueAMD64_OpAMD64SETLEstore_0(v)
	case OpAMD64SETLstore:
		return rewriteValueAMD64_OpAMD64SETLstore_0(v)
	case OpAMD64SETNE:
		return rewriteValueAMD64_OpAMD64SETNE_0(v) || rewriteValueAMD64_OpAMD64SETNE_10(v)
	case OpAMD64SETNEstore:
		return rewriteValueAMD64_OpAMD64SETNEstore_0(v) || rewriteValueAMD64_OpAMD64SETNEstore_10(v) || rewriteValueAMD64_OpAMD64SETNEstore_20(v)
	case OpAMD64SHLL:
		return rewriteValueAMD64_OpAMD64SHLL_0(v)
	case OpAMD64SHLLconst:
		return rewriteValueAMD64_OpAMD64SHLLconst_0(v)
	case OpAMD64SHLQ:
		return rewriteValueAMD64_OpAMD64SHLQ_0(v)
	case OpAMD64SHLQconst:
		return rewriteValueAMD64_OpAMD64SHLQconst_0(v)
	case OpAMD64SHRB:
		return rewriteValueAMD64_OpAMD64SHRB_0(v)
	case OpAMD64SHRBconst:
		return rewriteValueAMD64_OpAMD64SHRBconst_0(v)
	case OpAMD64SHRL:
		return rewriteValueAMD64_OpAMD64SHRL_0(v)
	case OpAMD64SHRLconst:
		return rewriteValueAMD64_OpAMD64SHRLconst_0(v)
	case OpAMD64SHRQ:
		return rewriteValueAMD64_OpAMD64SHRQ_0(v)
	case OpAMD64SHRQconst:
		return rewriteValueAMD64_OpAMD64SHRQconst_0(v)
	case OpAMD64SHRW:
		return rewriteValueAMD64_OpAMD64SHRW_0(v)
	case OpAMD64SHRWconst:
		return rewriteValueAMD64_OpAMD64SHRWconst_0(v)
	case OpAMD64SUBL:
		return rewriteValueAMD64_OpAMD64SUBL_0(v)
	case OpAMD64SUBLconst:
		return rewriteValueAMD64_OpAMD64SUBLconst_0(v)
	case OpAMD64SUBLload:
		return rewriteValueAMD64_OpAMD64SUBLload_0(v)
	case OpAMD64SUBLmodify:
		return rewriteValueAMD64_OpAMD64SUBLmodify_0(v)
	case OpAMD64SUBQ:
		return rewriteValueAMD64_OpAMD64SUBQ_0(v)
	case OpAMD64SUBQborrow:
		return rewriteValueAMD64_OpAMD64SUBQborrow_0(v)
	case OpAMD64SUBQconst:
		return rewriteValueAMD64_OpAMD64SUBQconst_0(v)
	case OpAMD64SUBQload:
		return rewriteValueAMD64_OpAMD64SUBQload_0(v)
	case OpAMD64SUBQmodify:
		return rewriteValueAMD64_OpAMD64SUBQmodify_0(v)
	case OpAMD64SUBSD:
		return rewriteValueAMD64_OpAMD64SUBSD_0(v)
	case OpAMD64SUBSDload:
		return rewriteValueAMD64_OpAMD64SUBSDload_0(v)
	case OpAMD64SUBSS:
		return rewriteValueAMD64_OpAMD64SUBSS_0(v)
	case OpAMD64SUBSSload:
		return rewriteValueAMD64_OpAMD64SUBSSload_0(v)
	case OpAMD64TESTB:
		return rewriteValueAMD64_OpAMD64TESTB_0(v)
	case OpAMD64TESTBconst:
		return rewriteValueAMD64_OpAMD64TESTBconst_0(v)
	case OpAMD64TESTL:
		return rewriteValueAMD64_OpAMD64TESTL_0(v)
	case OpAMD64TESTLconst:
		return rewriteValueAMD64_OpAMD64TESTLconst_0(v)
	case OpAMD64TESTQ:
		return rewriteValueAMD64_OpAMD64TESTQ_0(v)
	case OpAMD64TESTQconst:
		return rewriteValueAMD64_OpAMD64TESTQconst_0(v)
	case OpAMD64TESTW:
		return rewriteValueAMD64_OpAMD64TESTW_0(v)
	case OpAMD64TESTWconst:
		return rewriteValueAMD64_OpAMD64TESTWconst_0(v)
	case OpAMD64XADDLlock:
		return rewriteValueAMD64_OpAMD64XADDLlock_0(v)
	case OpAMD64XADDQlock:
		return rewriteValueAMD64_OpAMD64XADDQlock_0(v)
	case OpAMD64XCHGL:
		return rewriteValueAMD64_OpAMD64XCHGL_0(v)
	case OpAMD64XCHGQ:
		return rewriteValueAMD64_OpAMD64XCHGQ_0(v)
	case OpAMD64XORL:
		return rewriteValueAMD64_OpAMD64XORL_0(v)
	case OpAMD64XORLconst:
		return rewriteValueAMD64_OpAMD64XORLconst_0(v) || rewriteValueAMD64_OpAMD64XORLconst_10(v)
	case OpAMD64XORLconstmodify:
		return rewriteValueAMD64_OpAMD64XORLconstmodify_0(v)
	case OpAMD64XORLload:
		return rewriteValueAMD64_OpAMD64XORLload_0(v)
	case OpAMD64XORLmodify:
		return rewriteValueAMD64_OpAMD64XORLmodify_0(v)
	case OpAMD64XORQ:
		return rewriteValueAMD64_OpAMD64XORQ_0(v)
	case OpAMD64XORQconst:
		return rewriteValueAMD64_OpAMD64XORQconst_0(v)
	case OpAMD64XORQconstmodify:
		return rewriteValueAMD64_OpAMD64XORQconstmodify_0(v)
	case OpAMD64XORQload:
		return rewriteValueAMD64_OpAMD64XORQload_0(v)
	case OpAMD64XORQmodify:
		return rewriteValueAMD64_OpAMD64XORQmodify_0(v)
	case OpAdd16:
		return rewriteValueAMD64_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueAMD64_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueAMD64_OpAdd32F_0(v)
	case OpAdd64:
		return rewriteValueAMD64_OpAdd64_0(v)
	case OpAdd64F:
		return rewriteValueAMD64_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueAMD64_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueAMD64_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueAMD64_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueAMD64_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueAMD64_OpAnd32_0(v)
	case OpAnd64:
		return rewriteValueAMD64_OpAnd64_0(v)
	case OpAnd8:
		return rewriteValueAMD64_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueAMD64_OpAndB_0(v)
	case OpAtomicAdd32:
		return rewriteValueAMD64_OpAtomicAdd32_0(v)
	case OpAtomicAdd64:
		return rewriteValueAMD64_OpAtomicAdd64_0(v)
	case OpAtomicAnd8:
		return rewriteValueAMD64_OpAtomicAnd8_0(v)
	case OpAtomicCompareAndSwap32:
		return rewriteValueAMD64_OpAtomicCompareAndSwap32_0(v)
	case OpAtomicCompareAndSwap64:
		return rewriteValueAMD64_OpAtomicCompareAndSwap64_0(v)
	case OpAtomicExchange32:
		return rewriteValueAMD64_OpAtomicExchange32_0(v)
	case OpAtomicExchange64:
		return rewriteValueAMD64_OpAtomicExchange64_0(v)
	case OpAtomicLoad32:
		return rewriteValueAMD64_OpAtomicLoad32_0(v)
	case OpAtomicLoad64:
		return rewriteValueAMD64_OpAtomicLoad64_0(v)
	case OpAtomicLoad8:
		return rewriteValueAMD64_OpAtomicLoad8_0(v)
	case OpAtomicLoadPtr:
		return rewriteValueAMD64_OpAtomicLoadPtr_0(v)
	case OpAtomicOr8:
		return rewriteValueAMD64_OpAtomicOr8_0(v)
	case OpAtomicStore32:
		return rewriteValueAMD64_OpAtomicStore32_0(v)
	case OpAtomicStore64:
		return rewriteValueAMD64_OpAtomicStore64_0(v)
	case OpAtomicStore8:
		return rewriteValueAMD64_OpAtomicStore8_0(v)
	case OpAtomicStorePtrNoWB:
		return rewriteValueAMD64_OpAtomicStorePtrNoWB_0(v)
	case OpAvg64u:
		return rewriteValueAMD64_OpAvg64u_0(v)
	case OpBitLen16:
		return rewriteValueAMD64_OpBitLen16_0(v)
	case OpBitLen32:
		return rewriteValueAMD64_OpBitLen32_0(v)
	case OpBitLen64:
		return rewriteValueAMD64_OpBitLen64_0(v)
	case OpBitLen8:
		return rewriteValueAMD64_OpBitLen8_0(v)
	case OpBswap32:
		return rewriteValueAMD64_OpBswap32_0(v)
	case OpBswap64:
		return rewriteValueAMD64_OpBswap64_0(v)
	case OpCeil:
		return rewriteValueAMD64_OpCeil_0(v)
	case OpClosureCall:
		return rewriteValueAMD64_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueAMD64_OpCom16_0(v)
	case OpCom32:
		return rewriteValueAMD64_OpCom32_0(v)
	case OpCom64:
		return rewriteValueAMD64_OpCom64_0(v)
	case OpCom8:
		return rewriteValueAMD64_OpCom8_0(v)
	case OpCondSelect:
		return rewriteValueAMD64_OpCondSelect_0(v) || rewriteValueAMD64_OpCondSelect_10(v) || rewriteValueAMD64_OpCondSelect_20(v) || rewriteValueAMD64_OpCondSelect_30(v) || rewriteValueAMD64_OpCondSelect_40(v)
	case OpConst16:
		return rewriteValueAMD64_OpConst16_0(v)
	case OpConst32:
		return rewriteValueAMD64_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueAMD64_OpConst32F_0(v)
	case OpConst64:
		return rewriteValueAMD64_OpConst64_0(v)
	case OpConst64F:
		return rewriteValueAMD64_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueAMD64_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueAMD64_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueAMD64_OpConstNil_0(v)
	case OpCtz16:
		return rewriteValueAMD64_OpCtz16_0(v)
	case OpCtz16NonZero:
		return rewriteValueAMD64_OpCtz16NonZero_0(v)
	case OpCtz32:
		return rewriteValueAMD64_OpCtz32_0(v)
	case OpCtz32NonZero:
		return rewriteValueAMD64_OpCtz32NonZero_0(v)
	case OpCtz64:
		return rewriteValueAMD64_OpCtz64_0(v)
	case OpCtz64NonZero:
		return rewriteValueAMD64_OpCtz64NonZero_0(v)
	case OpCtz8:
		return rewriteValueAMD64_OpCtz8_0(v)
	case OpCtz8NonZero:
		return rewriteValueAMD64_OpCtz8NonZero_0(v)
	case OpCvt32Fto32:
		return rewriteValueAMD64_OpCvt32Fto32_0(v)
	case OpCvt32Fto64:
		return rewriteValueAMD64_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValueAMD64_OpCvt32Fto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueAMD64_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueAMD64_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueAMD64_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueAMD64_OpCvt64Fto32F_0(v)
	case OpCvt64Fto64:
		return rewriteValueAMD64_OpCvt64Fto64_0(v)
	case OpCvt64to32F:
		return rewriteValueAMD64_OpCvt64to32F_0(v)
	case OpCvt64to64F:
		return rewriteValueAMD64_OpCvt64to64F_0(v)
	case OpDiv128u:
		return rewriteValueAMD64_OpDiv128u_0(v)
	case OpDiv16:
		return rewriteValueAMD64_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueAMD64_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueAMD64_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueAMD64_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueAMD64_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValueAMD64_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValueAMD64_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValueAMD64_OpDiv64u_0(v)
	case OpDiv8:
		return rewriteValueAMD64_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueAMD64_OpDiv8u_0(v)
	case OpEq16:
		return rewriteValueAMD64_OpEq16_0(v)
	case OpEq32:
		return rewriteValueAMD64_OpEq32_0(v)
	case OpEq32F:
		return rewriteValueAMD64_OpEq32F_0(v)
	case OpEq64:
		return rewriteValueAMD64_OpEq64_0(v)
	case OpEq64F:
		return rewriteValueAMD64_OpEq64F_0(v)
	case OpEq8:
		return rewriteValueAMD64_OpEq8_0(v)
	case OpEqB:
		return rewriteValueAMD64_OpEqB_0(v)
	case OpEqPtr:
		return rewriteValueAMD64_OpEqPtr_0(v)
	case OpFMA:
		return rewriteValueAMD64_OpFMA_0(v)
	case OpFloor:
		return rewriteValueAMD64_OpFloor_0(v)
	case OpGeq16:
		return rewriteValueAMD64_OpGeq16_0(v)
	case OpGeq16U:
		return rewriteValueAMD64_OpGeq16U_0(v)
	case OpGeq32:
		return rewriteValueAMD64_OpGeq32_0(v)
	case OpGeq32F:
		return rewriteValueAMD64_OpGeq32F_0(v)
	case OpGeq32U:
		return rewriteValueAMD64_OpGeq32U_0(v)
	case OpGeq64:
		return rewriteValueAMD64_OpGeq64_0(v)
	case OpGeq64F:
		return rewriteValueAMD64_OpGeq64F_0(v)
	case OpGeq64U:
		return rewriteValueAMD64_OpGeq64U_0(v)
	case OpGeq8:
		return rewriteValueAMD64_OpGeq8_0(v)
	case OpGeq8U:
		return rewriteValueAMD64_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueAMD64_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueAMD64_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueAMD64_OpGetClosurePtr_0(v)
	case OpGetG:
		return rewriteValueAMD64_OpGetG_0(v)
	case OpGreater16:
		return rewriteValueAMD64_OpGreater16_0(v)
	case OpGreater16U:
		return rewriteValueAMD64_OpGreater16U_0(v)
	case OpGreater32:
		return rewriteValueAMD64_OpGreater32_0(v)
	case OpGreater32F:
		return rewriteValueAMD64_OpGreater32F_0(v)
	case OpGreater32U:
		return rewriteValueAMD64_OpGreater32U_0(v)
	case OpGreater64:
		return rewriteValueAMD64_OpGreater64_0(v)
	case OpGreater64F:
		return rewriteValueAMD64_OpGreater64F_0(v)
	case OpGreater64U:
		return rewriteValueAMD64_OpGreater64U_0(v)
	case OpGreater8:
		return rewriteValueAMD64_OpGreater8_0(v)
	case OpGreater8U:
		return rewriteValueAMD64_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValueAMD64_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValueAMD64_OpHmul32u_0(v)
	case OpHmul64:
		return rewriteValueAMD64_OpHmul64_0(v)
	case OpHmul64u:
		return rewriteValueAMD64_OpHmul64u_0(v)
	case OpInterCall:
		return rewriteValueAMD64_OpInterCall_0(v)
	case OpIsInBounds:
		return rewriteValueAMD64_OpIsInBounds_0(v)
	case OpIsNonNil:
		return rewriteValueAMD64_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return rewriteValueAMD64_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return rewriteValueAMD64_OpLeq16_0(v)
	case OpLeq16U:
		return rewriteValueAMD64_OpLeq16U_0(v)
	case OpLeq32:
		return rewriteValueAMD64_OpLeq32_0(v)
	case OpLeq32F:
		return rewriteValueAMD64_OpLeq32F_0(v)
	case OpLeq32U:
		return rewriteValueAMD64_OpLeq32U_0(v)
	case OpLeq64:
		return rewriteValueAMD64_OpLeq64_0(v)
	case OpLeq64F:
		return rewriteValueAMD64_OpLeq64F_0(v)
	case OpLeq64U:
		return rewriteValueAMD64_OpLeq64U_0(v)
	case OpLeq8:
		return rewriteValueAMD64_OpLeq8_0(v)
	case OpLeq8U:
		return rewriteValueAMD64_OpLeq8U_0(v)
	case OpLess16:
		return rewriteValueAMD64_OpLess16_0(v)
	case OpLess16U:
		return rewriteValueAMD64_OpLess16U_0(v)
	case OpLess32:
		return rewriteValueAMD64_OpLess32_0(v)
	case OpLess32F:
		return rewriteValueAMD64_OpLess32F_0(v)
	case OpLess32U:
		return rewriteValueAMD64_OpLess32U_0(v)
	case OpLess64:
		return rewriteValueAMD64_OpLess64_0(v)
	case OpLess64F:
		return rewriteValueAMD64_OpLess64F_0(v)
	case OpLess64U:
		return rewriteValueAMD64_OpLess64U_0(v)
	case OpLess8:
		return rewriteValueAMD64_OpLess8_0(v)
	case OpLess8U:
		return rewriteValueAMD64_OpLess8U_0(v)
	case OpLoad:
		return rewriteValueAMD64_OpLoad_0(v)
	case OpLocalAddr:
		return rewriteValueAMD64_OpLocalAddr_0(v)
	case OpLsh16x16:
		return rewriteValueAMD64_OpLsh16x16_0(v)
	case OpLsh16x32:
		return rewriteValueAMD64_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValueAMD64_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValueAMD64_OpLsh16x8_0(v)
	case OpLsh32x16:
		return rewriteValueAMD64_OpLsh32x16_0(v)
	case OpLsh32x32:
		return rewriteValueAMD64_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValueAMD64_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValueAMD64_OpLsh32x8_0(v)
	case OpLsh64x16:
		return rewriteValueAMD64_OpLsh64x16_0(v)
	case OpLsh64x32:
		return rewriteValueAMD64_OpLsh64x32_0(v)
	case OpLsh64x64:
		return rewriteValueAMD64_OpLsh64x64_0(v)
	case OpLsh64x8:
		return rewriteValueAMD64_OpLsh64x8_0(v)
	case OpLsh8x16:
		return rewriteValueAMD64_OpLsh8x16_0(v)
	case OpLsh8x32:
		return rewriteValueAMD64_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValueAMD64_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValueAMD64_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValueAMD64_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueAMD64_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueAMD64_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueAMD64_OpMod32u_0(v)
	case OpMod64:
		return rewriteValueAMD64_OpMod64_0(v)
	case OpMod64u:
		return rewriteValueAMD64_OpMod64u_0(v)
	case OpMod8:
		return rewriteValueAMD64_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueAMD64_OpMod8u_0(v)
	case OpMove:
		return rewriteValueAMD64_OpMove_0(v) || rewriteValueAMD64_OpMove_10(v) || rewriteValueAMD64_OpMove_20(v)
	case OpMul16:
		return rewriteValueAMD64_OpMul16_0(v)
	case OpMul32:
		return rewriteValueAMD64_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueAMD64_OpMul32F_0(v)
	case OpMul64:
		return rewriteValueAMD64_OpMul64_0(v)
	case OpMul64F:
		return rewriteValueAMD64_OpMul64F_0(v)
	case OpMul64uhilo:
		return rewriteValueAMD64_OpMul64uhilo_0(v)
	case OpMul8:
		return rewriteValueAMD64_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueAMD64_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueAMD64_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueAMD64_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValueAMD64_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValueAMD64_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueAMD64_OpNeg8_0(v)
	case OpNeq16:
		return rewriteValueAMD64_OpNeq16_0(v)
	case OpNeq32:
		return rewriteValueAMD64_OpNeq32_0(v)
	case OpNeq32F:
		return rewriteValueAMD64_OpNeq32F_0(v)
	case OpNeq64:
		return rewriteValueAMD64_OpNeq64_0(v)
	case OpNeq64F:
		return rewriteValueAMD64_OpNeq64F_0(v)
	case OpNeq8:
		return rewriteValueAMD64_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueAMD64_OpNeqB_0(v)
	case OpNeqPtr:
		return rewriteValueAMD64_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueAMD64_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueAMD64_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueAMD64_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueAMD64_OpOr16_0(v)
	case OpOr32:
		return rewriteValueAMD64_OpOr32_0(v)
	case OpOr64:
		return rewriteValueAMD64_OpOr64_0(v)
	case OpOr8:
		return rewriteValueAMD64_OpOr8_0(v)
	case OpOrB:
		return rewriteValueAMD64_OpOrB_0(v)
	case OpPanicBounds:
		return rewriteValueAMD64_OpPanicBounds_0(v)
	case OpPopCount16:
		return rewriteValueAMD64_OpPopCount16_0(v)
	case OpPopCount32:
		return rewriteValueAMD64_OpPopCount32_0(v)
	case OpPopCount64:
		return rewriteValueAMD64_OpPopCount64_0(v)
	case OpPopCount8:
		return rewriteValueAMD64_OpPopCount8_0(v)
	case OpRotateLeft16:
		return rewriteValueAMD64_OpRotateLeft16_0(v)
	case OpRotateLeft32:
		return rewriteValueAMD64_OpRotateLeft32_0(v)
	case OpRotateLeft64:
		return rewriteValueAMD64_OpRotateLeft64_0(v)
	case OpRotateLeft8:
		return rewriteValueAMD64_OpRotateLeft8_0(v)
	case OpRound32F:
		return rewriteValueAMD64_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueAMD64_OpRound64F_0(v)
	case OpRoundToEven:
		return rewriteValueAMD64_OpRoundToEven_0(v)
	case OpRsh16Ux16:
		return rewriteValueAMD64_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return rewriteValueAMD64_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValueAMD64_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return rewriteValueAMD64_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return rewriteValueAMD64_OpRsh16x16_0(v)
	case OpRsh16x32:
		return rewriteValueAMD64_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValueAMD64_OpRsh16x64_0(v)
	case OpRsh16x8:
		return rewriteValueAMD64_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return rewriteValueAMD64_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return rewriteValueAMD64_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValueAMD64_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValueAMD64_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return rewriteValueAMD64_OpRsh32x16_0(v)
	case OpRsh32x32:
		return rewriteValueAMD64_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValueAMD64_OpRsh32x64_0(v)
	case OpRsh32x8:
		return rewriteValueAMD64_OpRsh32x8_0(v)
	case OpRsh64Ux16:
		return rewriteValueAMD64_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return rewriteValueAMD64_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return rewriteValueAMD64_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return rewriteValueAMD64_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return rewriteValueAMD64_OpRsh64x16_0(v)
	case OpRsh64x32:
		return rewriteValueAMD64_OpRsh64x32_0(v)
	case OpRsh64x64:
		return rewriteValueAMD64_OpRsh64x64_0(v)
	case OpRsh64x8:
		return rewriteValueAMD64_OpRsh64x8_0(v)
	case OpRsh8Ux16:
		return rewriteValueAMD64_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return rewriteValueAMD64_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValueAMD64_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValueAMD64_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return rewriteValueAMD64_OpRsh8x16_0(v)
	case OpRsh8x32:
		return rewriteValueAMD64_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValueAMD64_OpRsh8x64_0(v)
	case OpRsh8x8:
		return rewriteValueAMD64_OpRsh8x8_0(v)
	case OpSelect0:
		return rewriteValueAMD64_OpSelect0_0(v)
	case OpSelect1:
		return rewriteValueAMD64_OpSelect1_0(v)
	case OpSignExt16to32:
		return rewriteValueAMD64_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValueAMD64_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValueAMD64_OpSignExt32to64_0(v)
	case OpSignExt8to16:
		return rewriteValueAMD64_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueAMD64_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValueAMD64_OpSignExt8to64_0(v)
	case OpSlicemask:
		return rewriteValueAMD64_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValueAMD64_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValueAMD64_OpStaticCall_0(v)
	case OpStore:
		return rewriteValueAMD64_OpStore_0(v)
	case OpSub16:
		return rewriteValueAMD64_OpSub16_0(v)
	case OpSub32:
		return rewriteValueAMD64_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueAMD64_OpSub32F_0(v)
	case OpSub64:
		return rewriteValueAMD64_OpSub64_0(v)
	case OpSub64F:
		return rewriteValueAMD64_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueAMD64_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueAMD64_OpSubPtr_0(v)
	case OpTrunc:
		return rewriteValueAMD64_OpTrunc_0(v)
	case OpTrunc16to8:
		return rewriteValueAMD64_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueAMD64_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueAMD64_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValueAMD64_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValueAMD64_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValueAMD64_OpTrunc64to8_0(v)
	case OpWB:
		return rewriteValueAMD64_OpWB_0(v)
	case OpXor16:
		return rewriteValueAMD64_OpXor16_0(v)
	case OpXor32:
		return rewriteValueAMD64_OpXor32_0(v)
	case OpXor64:
		return rewriteValueAMD64_OpXor64_0(v)
	case OpXor8:
		return rewriteValueAMD64_OpXor8_0(v)
	case OpZero:
		return rewriteValueAMD64_OpZero_0(v) || rewriteValueAMD64_OpZero_10(v) || rewriteValueAMD64_OpZero_20(v)
	case OpZeroExt16to32:
		return rewriteValueAMD64_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValueAMD64_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValueAMD64_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValueAMD64_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueAMD64_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValueAMD64_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADCQ_0(v *Value) bool {
	// match: (ADCQ x (MOVQconst [c]) carry)
	// cond: is32Bit(c)
	// result: (ADCQconst x [c] carry)
	for {
		carry := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpAMD64ADCQconst)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(carry)
			return true
		}
		break
	}
	// match: (ADCQ x y (FlagEQ))
	// result: (ADDQcarry x y)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64ADDQcarry)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADCQconst_0(v *Value) bool {
	// match: (ADCQconst x [c] (FlagEQ))
	// result: (ADDQconstcarry x [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64ADDQconstcarry)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDL_0(v *Value) bool {
	// match: (ADDL x (MOVLconst [c]))
	// result: (ADDLconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpAMD64ADDLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDL (SHLLconst x [c]) (SHRLconst x [d]))
	// cond: d==32-c
	// result: (ROLLconst x [c])
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRLconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpAMD64ROLLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDL <t> (SHLLconst x [c]) (SHRWconst x [d]))
	// cond: d==16-c && c < 16 && t.Size() == 2
	// result: (ROLWconst x [c])
	for {
		t := v.Type
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 16-c && c < 16 && t.Size() == 2) {
				continue
			}
			v.reset(OpAMD64ROLWconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDL <t> (SHLLconst x [c]) (SHRBconst x [d]))
	// cond: d==8-c && c < 8 && t.Size() == 1
	// result: (ROLBconst x [c])
	for {
		t := v.Type
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRBconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 8-c && c < 8 && t.Size() == 1) {
				continue
			}
			v.reset(OpAMD64ROLBconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDL x (SHLLconst [3] y))
	// result: (LEAL8 x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 3 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAL8)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDL x (SHLLconst [2] y))
	// result: (LEAL4 x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 2 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAL4)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDL x (SHLLconst [1] y))
	// result: (LEAL2 x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 1 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAL2)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDL x (ADDL y y))
	// result: (LEAL2 x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ADDL {
				continue
			}
			y := v_1.Args[1]
			if y != v_1.Args[0] {
				continue
			}
			v.reset(OpAMD64LEAL2)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDL x (ADDL x y))
	// result: (LEAL2 y x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ADDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				if x != v_1.Args[_i1] {
					continue
				}
				y := v_1.Args[1^_i1]
				v.reset(OpAMD64LEAL2)
				v.AddArg(y)
				v.AddArg(x)
				return true
			}
		}
		break
	}
	// match: (ADDL (ADDLconst [c] x) y)
	// result: (LEAL1 [c] x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64ADDLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			y := v.Args[1^_i0]
			v.reset(OpAMD64LEAL1)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDL_10(v *Value) bool {
	// match: (ADDL x (LEAL [c] {s} y))
	// cond: x.Op != OpSB && y.Op != OpSB
	// result: (LEAL1 [c] {s} x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64LEAL {
				continue
			}
			c := v_1.AuxInt
			s := v_1.Aux
			y := v_1.Args[0]
			if !(x.Op != OpSB && y.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64LEAL1)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDL x (NEGL y))
	// result: (SUBL x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64NEGL {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64SUBL)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDL x l:(MOVLload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (ADDLload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVLload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ADDLload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDLconst_0(v *Value) bool {
	// match: (ADDLconst [c] (ADDL x y))
	// result: (LEAL1 [c] x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpAMD64LEAL1)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDLconst [c] (SHLLconst [1] x))
	// result: (LEAL1 [c] x x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHLLconst || v_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64LEAL1)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (ADDLconst [c] (LEAL [d] {s} x))
	// cond: is32Bit(c+d)
	// result: (LEAL [c+d] {s} x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAL)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}
	// match: (ADDLconst [c] (LEAL1 [d] {s} x y))
	// cond: is32Bit(c+d)
	// result: (LEAL1 [c+d] {s} x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL1 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAL1)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDLconst [c] (LEAL2 [d] {s} x y))
	// cond: is32Bit(c+d)
	// result: (LEAL2 [c+d] {s} x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL2 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAL2)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDLconst [c] (LEAL4 [d] {s} x y))
	// cond: is32Bit(c+d)
	// result: (LEAL4 [c+d] {s} x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL4 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAL4)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDLconst [c] (LEAL8 [d] {s} x y))
	// cond: is32Bit(c+d)
	// result: (LEAL8 [c+d] {s} x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL8 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAL8)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDLconst [c] x)
	// cond: int32(c)==0
	// result: x
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ADDLconst [c] (MOVLconst [d]))
	// result: (MOVLconst [int64(int32(c+d))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}
	// match: (ADDLconst [c] (ADDLconst [d] x))
	// result: (ADDLconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ADDLconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDLconst_10(v *Value) bool {
	// match: (ADDLconst [off] x:(SP))
	// result: (LEAL [off] x)
	for {
		off := v.AuxInt
		x := v.Args[0]
		if x.Op != OpSP {
			break
		}
		v.reset(OpAMD64LEAL)
		v.AuxInt = off
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDLconstmodify_0(v *Value) bool {
	// match: (ADDLconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (ADDLconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64ADDLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDLconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (ADDLconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ADDLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDLload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ADDLload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (ADDLload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ADDLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDLload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ADDLload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ADDLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDLload x [off] {sym} ptr (MOVSSstore [off] {sym} ptr y _))
	// result: (ADDL x (MOVLf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSSstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64ADDL)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVLf2i, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDLmodify_0(v *Value) bool {
	// match: (ADDLmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (ADDLmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ADDLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (ADDLmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ADDLmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ADDLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDQ_0(v *Value) bool {
	// match: (ADDQ x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (ADDQconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpAMD64ADDQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDQ (SHLQconst x [c]) (SHRQconst x [d]))
	// cond: d==64-c
	// result: (ROLQconst x [c])
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLQconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRQconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpAMD64ROLQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDQ x (SHLQconst [3] y))
	// result: (LEAQ8 x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAQ8)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDQ x (SHLQconst [2] y))
	// result: (LEAQ4 x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 2 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAQ4)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDQ x (SHLQconst [1] y))
	// result: (LEAQ2 x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 1 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAQ2)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDQ x (ADDQ y y))
	// result: (LEAQ2 x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ADDQ {
				continue
			}
			y := v_1.Args[1]
			if y != v_1.Args[0] {
				continue
			}
			v.reset(OpAMD64LEAQ2)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDQ x (ADDQ x y))
	// result: (LEAQ2 y x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ADDQ {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				if x != v_1.Args[_i1] {
					continue
				}
				y := v_1.Args[1^_i1]
				v.reset(OpAMD64LEAQ2)
				v.AddArg(y)
				v.AddArg(x)
				return true
			}
		}
		break
	}
	// match: (ADDQ (ADDQconst [c] x) y)
	// result: (LEAQ1 [c] x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64ADDQconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			y := v.Args[1^_i0]
			v.reset(OpAMD64LEAQ1)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDQ x (LEAQ [c] {s} y))
	// cond: x.Op != OpSB && y.Op != OpSB
	// result: (LEAQ1 [c] {s} x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64LEAQ {
				continue
			}
			c := v_1.AuxInt
			s := v_1.Aux
			y := v_1.Args[0]
			if !(x.Op != OpSB && y.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64LEAQ1)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDQ x (NEGQ y))
	// result: (SUBQ x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64NEGQ {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64SUBQ)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDQ_10(v *Value) bool {
	// match: (ADDQ x l:(MOVQload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (ADDQload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVQload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ADDQload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDQcarry_0(v *Value) bool {
	// match: (ADDQcarry x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (ADDQconstcarry x [c])
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpAMD64ADDQconstcarry)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDQconst_0(v *Value) bool {
	// match: (ADDQconst [c] (ADDQ x y))
	// result: (LEAQ1 [c] x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpAMD64LEAQ1)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDQconst [c] (SHLQconst [1] x))
	// result: (LEAQ1 [c] x x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHLQconst || v_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64LEAQ1)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (ADDQconst [c] (LEAQ [d] {s} x))
	// cond: is32Bit(c+d)
	// result: (LEAQ [c+d] {s} x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAQ)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}
	// match: (ADDQconst [c] (LEAQ1 [d] {s} x y))
	// cond: is32Bit(c+d)
	// result: (LEAQ1 [c+d] {s} x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAQ1)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDQconst [c] (LEAQ2 [d] {s} x y))
	// cond: is32Bit(c+d)
	// result: (LEAQ2 [c+d] {s} x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ2 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAQ2)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDQconst [c] (LEAQ4 [d] {s} x y))
	// cond: is32Bit(c+d)
	// result: (LEAQ4 [c+d] {s} x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ4 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAQ4)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDQconst [c] (LEAQ8 [d] {s} x y))
	// cond: is32Bit(c+d)
	// result: (LEAQ8 [c+d] {s} x y)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAQ8)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDQconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ADDQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [c+d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = c + d
		return true
	}
	// match: (ADDQconst [c] (ADDQconst [d] x))
	// cond: is32Bit(c+d)
	// result: (ADDQconst [c+d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64ADDQconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDQconst_10(v *Value) bool {
	// match: (ADDQconst [off] x:(SP))
	// result: (LEAQ [off] x)
	for {
		off := v.AuxInt
		x := v.Args[0]
		if x.Op != OpSP {
			break
		}
		v.reset(OpAMD64LEAQ)
		v.AuxInt = off
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDQconstmodify_0(v *Value) bool {
	// match: (ADDQconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (ADDQconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64ADDQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDQconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (ADDQconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ADDQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDQload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ADDQload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (ADDQload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ADDQload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDQload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ADDQload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ADDQload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDQload x [off] {sym} ptr (MOVSDstore [off] {sym} ptr y _))
	// result: (ADDQ x (MOVQf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64ADDQ)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVQf2i, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDQmodify_0(v *Value) bool {
	// match: (ADDQmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (ADDQmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ADDQmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (ADDQmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ADDQmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ADDQmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDSD_0(v *Value) bool {
	// match: (ADDSD x l:(MOVSDload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (ADDSDload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVSDload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ADDSDload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDSDload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ADDSDload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (ADDSDload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ADDSDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDSDload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ADDSDload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ADDSDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDSDload x [off] {sym} ptr (MOVQstore [off] {sym} ptr y _))
	// result: (ADDSD x (MOVQi2f y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVQstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64ADDSD)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVQi2f, typ.Float64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDSS_0(v *Value) bool {
	// match: (ADDSS x l:(MOVSSload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (ADDSSload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVSSload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ADDSSload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ADDSSload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ADDSSload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (ADDSSload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ADDSSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDSSload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ADDSSload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ADDSSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ADDSSload x [off] {sym} ptr (MOVLstore [off] {sym} ptr y _))
	// result: (ADDSS x (MOVLi2f y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVLstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64ADDSS)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVLi2f, typ.Float32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDL_0(v *Value) bool {
	// match: (ANDL (NOTL (SHLL (MOVLconst [1]) y)) x)
	// result: (BTRL x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64NOTL {
				continue
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SHLL {
				continue
			}
			y := v_0_0.Args[1]
			v_0_0_0 := v_0_0.Args[0]
			if v_0_0_0.Op != OpAMD64MOVLconst || v_0_0_0.AuxInt != 1 {
				continue
			}
			x := v.Args[1^_i0]
			v.reset(OpAMD64BTRL)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ANDL (MOVLconst [c]) x)
	// cond: isUint32PowerOfTwo(^c) && uint64(^c) >= 128
	// result: (BTRLconst [log2uint32(^c)] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			if !(isUint32PowerOfTwo(^c) && uint64(^c) >= 128) {
				continue
			}
			v.reset(OpAMD64BTRLconst)
			v.AuxInt = log2uint32(^c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ANDL x (MOVLconst [c]))
	// result: (ANDLconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpAMD64ANDLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ANDL x x)
	// result: x
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDL x l:(MOVLload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (ANDLload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVLload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ANDLload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDLconst_0(v *Value) bool {
	// match: (ANDLconst [c] x)
	// cond: isUint32PowerOfTwo(^c) && uint64(^c) >= 128
	// result: (BTRLconst [log2uint32(^c)] x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isUint32PowerOfTwo(^c) && uint64(^c) >= 128) {
			break
		}
		v.reset(OpAMD64BTRLconst)
		v.AuxInt = log2uint32(^c)
		v.AddArg(x)
		return true
	}
	// match: (ANDLconst [c] (ANDLconst [d] x))
	// result: (ANDLconst [c & d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	// match: (ANDLconst [c] (BTRLconst [d] x))
	// result: (ANDLconst [c &^ (1<<uint32(d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTRLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = c &^ (1 << uint32(d))
		v.AddArg(x)
		return true
	}
	// match: (ANDLconst [ 0xFF] x)
	// result: (MOVBQZX x)
	for {
		if v.AuxInt != 0xFF {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
	// match: (ANDLconst [0xFFFF] x)
	// result: (MOVWQZX x)
	for {
		if v.AuxInt != 0xFFFF {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64MOVWQZX)
		v.AddArg(x)
		return true
	}
	// match: (ANDLconst [c] _)
	// cond: int32(c)==0
	// result: (MOVLconst [0])
	for {
		c := v.AuxInt
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDLconst [c] x)
	// cond: int32(c)==-1
	// result: x
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDLconst [c] (MOVLconst [d]))
	// result: (MOVLconst [c&d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = c & d
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDLconstmodify_0(v *Value) bool {
	// match: (ANDLconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (ANDLconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64ANDLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ANDLconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (ANDLconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ANDLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDLload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ANDLload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (ANDLload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ANDLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ANDLload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ANDLload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ANDLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ANDLload x [off] {sym} ptr (MOVSSstore [off] {sym} ptr y _))
	// result: (ANDL x (MOVLf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSSstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64ANDL)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVLf2i, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDLmodify_0(v *Value) bool {
	// match: (ANDLmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (ANDLmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ANDLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (ANDLmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ANDLmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ANDLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDQ_0(v *Value) bool {
	// match: (ANDQ (NOTQ (SHLQ (MOVQconst [1]) y)) x)
	// result: (BTRQ x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64NOTQ {
				continue
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SHLQ {
				continue
			}
			y := v_0_0.Args[1]
			v_0_0_0 := v_0_0.Args[0]
			if v_0_0_0.Op != OpAMD64MOVQconst || v_0_0_0.AuxInt != 1 {
				continue
			}
			x := v.Args[1^_i0]
			v.reset(OpAMD64BTRQ)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ANDQ (MOVQconst [c]) x)
	// cond: isUint64PowerOfTwo(^c) && uint64(^c) >= 128
	// result: (BTRQconst [log2(^c)] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			if !(isUint64PowerOfTwo(^c) && uint64(^c) >= 128) {
				continue
			}
			v.reset(OpAMD64BTRQconst)
			v.AuxInt = log2(^c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ANDQ x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (ANDQconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpAMD64ANDQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ANDQ x x)
	// result: x
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDQ x l:(MOVQload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (ANDQload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVQload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ANDQload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDQconst_0(v *Value) bool {
	// match: (ANDQconst [c] x)
	// cond: isUint64PowerOfTwo(^c) && uint64(^c) >= 128
	// result: (BTRQconst [log2(^c)] x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isUint64PowerOfTwo(^c) && uint64(^c) >= 128) {
			break
		}
		v.reset(OpAMD64BTRQconst)
		v.AuxInt = log2(^c)
		v.AddArg(x)
		return true
	}
	// match: (ANDQconst [c] (ANDQconst [d] x))
	// result: (ANDQconst [c & d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDQconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	// match: (ANDQconst [c] (BTRQconst [d] x))
	// result: (ANDQconst [c &^ (1<<uint32(d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTRQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDQconst)
		v.AuxInt = c &^ (1 << uint32(d))
		v.AddArg(x)
		return true
	}
	// match: (ANDQconst [ 0xFF] x)
	// result: (MOVBQZX x)
	for {
		if v.AuxInt != 0xFF {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
	// match: (ANDQconst [0xFFFF] x)
	// result: (MOVWQZX x)
	for {
		if v.AuxInt != 0xFFFF {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64MOVWQZX)
		v.AddArg(x)
		return true
	}
	// match: (ANDQconst [0xFFFFFFFF] x)
	// result: (MOVLQZX x)
	for {
		if v.AuxInt != 0xFFFFFFFF {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64MOVLQZX)
		v.AddArg(x)
		return true
	}
	// match: (ANDQconst [0] _)
	// result: (MOVQconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDQconst [-1] x)
	// result: x
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [c&d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = c & d
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDQconstmodify_0(v *Value) bool {
	// match: (ANDQconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (ANDQconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64ANDQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ANDQconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (ANDQconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ANDQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDQload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ANDQload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (ANDQload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ANDQload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ANDQload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ANDQload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ANDQload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ANDQload x [off] {sym} ptr (MOVSDstore [off] {sym} ptr y _))
	// result: (ANDQ x (MOVQf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64ANDQ)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVQf2i, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ANDQmodify_0(v *Value) bool {
	// match: (ANDQmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (ANDQmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ANDQmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (ANDQmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ANDQmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ANDQmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BSFQ_0(v *Value) bool {
	b := v.Block
	// match: (BSFQ (ORQconst <t> [1<<8] (MOVBQZX x)))
	// result: (BSFQ (ORQconst <t> [1<<8] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ORQconst {
			break
		}
		t := v_0.Type
		if v_0.AuxInt != 1<<8 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAMD64MOVBQZX {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpAMD64BSFQ)
		v0 := b.NewValue0(v.Pos, OpAMD64ORQconst, t)
		v0.AuxInt = 1 << 8
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (BSFQ (ORQconst <t> [1<<16] (MOVWQZX x)))
	// result: (BSFQ (ORQconst <t> [1<<16] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ORQconst {
			break
		}
		t := v_0.Type
		if v_0.AuxInt != 1<<16 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAMD64MOVWQZX {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpAMD64BSFQ)
		v0 := b.NewValue0(v.Pos, OpAMD64ORQconst, t)
		v0.AuxInt = 1 << 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTCLconst_0(v *Value) bool {
	// match: (BTCLconst [c] (XORLconst [d] x))
	// result: (XORLconst [d ^ 1<<uint32(c)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64XORLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64XORLconst)
		v.AuxInt = d ^ 1<<uint32(c)
		v.AddArg(x)
		return true
	}
	// match: (BTCLconst [c] (BTCLconst [d] x))
	// result: (XORLconst [1<<uint32(c) ^ 1<<uint32(d)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTCLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64XORLconst)
		v.AuxInt = 1<<uint32(c) ^ 1<<uint32(d)
		v.AddArg(x)
		return true
	}
	// match: (BTCLconst [c] (MOVLconst [d]))
	// result: (MOVLconst [d^(1<<uint32(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = d ^ (1 << uint32(c))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTCLconstmodify_0(v *Value) bool {
	// match: (BTCLconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (BTCLconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64BTCLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (BTCLconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (BTCLconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTCLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTCLmodify_0(v *Value) bool {
	// match: (BTCLmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (BTCLmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64BTCLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (BTCLmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (BTCLmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTCLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTCQconst_0(v *Value) bool {
	// match: (BTCQconst [c] (XORQconst [d] x))
	// result: (XORQconst [d ^ 1<<uint32(c)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64XORQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64XORQconst)
		v.AuxInt = d ^ 1<<uint32(c)
		v.AddArg(x)
		return true
	}
	// match: (BTCQconst [c] (BTCQconst [d] x))
	// result: (XORQconst [1<<uint32(c) ^ 1<<uint32(d)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTCQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64XORQconst)
		v.AuxInt = 1<<uint32(c) ^ 1<<uint32(d)
		v.AddArg(x)
		return true
	}
	// match: (BTCQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [d^(1<<uint32(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = d ^ (1 << uint32(c))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTCQconstmodify_0(v *Value) bool {
	// match: (BTCQconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (BTCQconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64BTCQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (BTCQconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (BTCQconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTCQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTCQmodify_0(v *Value) bool {
	// match: (BTCQmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (BTCQmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64BTCQmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (BTCQmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (BTCQmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTCQmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTLconst_0(v *Value) bool {
	// match: (BTLconst [c] (SHRQconst [d] x))
	// cond: (c+d)<64
	// result: (BTQconst [c+d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHRQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !((c + d) < 64) {
			break
		}
		v.reset(OpAMD64BTQconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	// match: (BTLconst [c] (SHLQconst [d] x))
	// cond: c>d
	// result: (BTLconst [c-d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHLQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(c > d) {
			break
		}
		v.reset(OpAMD64BTLconst)
		v.AuxInt = c - d
		v.AddArg(x)
		return true
	}
	// match: (BTLconst [0] s:(SHRQ x y))
	// result: (BTQ y x)
	for {
		if v.AuxInt != 0 {
			break
		}
		s := v.Args[0]
		if s.Op != OpAMD64SHRQ {
			break
		}
		y := s.Args[1]
		x := s.Args[0]
		v.reset(OpAMD64BTQ)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	// match: (BTLconst [c] (SHRLconst [d] x))
	// cond: (c+d)<32
	// result: (BTLconst [c+d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHRLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !((c + d) < 32) {
			break
		}
		v.reset(OpAMD64BTLconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	// match: (BTLconst [c] (SHLLconst [d] x))
	// cond: c>d
	// result: (BTLconst [c-d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHLLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(c > d) {
			break
		}
		v.reset(OpAMD64BTLconst)
		v.AuxInt = c - d
		v.AddArg(x)
		return true
	}
	// match: (BTLconst [0] s:(SHRL x y))
	// result: (BTL y x)
	for {
		if v.AuxInt != 0 {
			break
		}
		s := v.Args[0]
		if s.Op != OpAMD64SHRL {
			break
		}
		y := s.Args[1]
		x := s.Args[0]
		v.reset(OpAMD64BTL)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTQconst_0(v *Value) bool {
	// match: (BTQconst [c] (SHRQconst [d] x))
	// cond: (c+d)<64
	// result: (BTQconst [c+d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHRQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !((c + d) < 64) {
			break
		}
		v.reset(OpAMD64BTQconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	// match: (BTQconst [c] (SHLQconst [d] x))
	// cond: c>d
	// result: (BTQconst [c-d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHLQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(c > d) {
			break
		}
		v.reset(OpAMD64BTQconst)
		v.AuxInt = c - d
		v.AddArg(x)
		return true
	}
	// match: (BTQconst [0] s:(SHRQ x y))
	// result: (BTQ y x)
	for {
		if v.AuxInt != 0 {
			break
		}
		s := v.Args[0]
		if s.Op != OpAMD64SHRQ {
			break
		}
		y := s.Args[1]
		x := s.Args[0]
		v.reset(OpAMD64BTQ)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTRLconst_0(v *Value) bool {
	// match: (BTRLconst [c] (BTSLconst [c] x))
	// result: (BTRLconst [c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTSLconst || v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BTRLconst [c] (BTCLconst [c] x))
	// result: (BTRLconst [c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTCLconst || v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BTRLconst [c] (ANDLconst [d] x))
	// result: (ANDLconst [d &^ (1<<uint32(c))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = d &^ (1 << uint32(c))
		v.AddArg(x)
		return true
	}
	// match: (BTRLconst [c] (BTRLconst [d] x))
	// result: (ANDLconst [^(1<<uint32(c) | 1<<uint32(d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTRLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = ^(1<<uint32(c) | 1<<uint32(d))
		v.AddArg(x)
		return true
	}
	// match: (BTRLconst [c] (MOVLconst [d]))
	// result: (MOVLconst [d&^(1<<uint32(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = d &^ (1 << uint32(c))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTRLconstmodify_0(v *Value) bool {
	// match: (BTRLconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (BTRLconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64BTRLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (BTRLconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (BTRLconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTRLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTRLmodify_0(v *Value) bool {
	// match: (BTRLmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (BTRLmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64BTRLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (BTRLmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (BTRLmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTRLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTRQconst_0(v *Value) bool {
	// match: (BTRQconst [c] (BTSQconst [c] x))
	// result: (BTRQconst [c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTSQconst || v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTRQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BTRQconst [c] (BTCQconst [c] x))
	// result: (BTRQconst [c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTCQconst || v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTRQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BTRQconst [c] (ANDQconst [d] x))
	// result: (ANDQconst [d &^ (1<<uint32(c))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDQconst)
		v.AuxInt = d &^ (1 << uint32(c))
		v.AddArg(x)
		return true
	}
	// match: (BTRQconst [c] (BTRQconst [d] x))
	// result: (ANDQconst [^(1<<uint32(c) | 1<<uint32(d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTRQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDQconst)
		v.AuxInt = ^(1<<uint32(c) | 1<<uint32(d))
		v.AddArg(x)
		return true
	}
	// match: (BTRQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [d&^(1<<uint32(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = d &^ (1 << uint32(c))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTRQconstmodify_0(v *Value) bool {
	// match: (BTRQconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (BTRQconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64BTRQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (BTRQconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (BTRQconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTRQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTRQmodify_0(v *Value) bool {
	// match: (BTRQmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (BTRQmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64BTRQmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (BTRQmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (BTRQmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTRQmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTSLconst_0(v *Value) bool {
	// match: (BTSLconst [c] (BTRLconst [c] x))
	// result: (BTSLconst [c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTRLconst || v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTSLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BTSLconst [c] (BTCLconst [c] x))
	// result: (BTSLconst [c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTCLconst || v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTSLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BTSLconst [c] (ORLconst [d] x))
	// result: (ORLconst [d | 1<<uint32(c)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ORLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ORLconst)
		v.AuxInt = d | 1<<uint32(c)
		v.AddArg(x)
		return true
	}
	// match: (BTSLconst [c] (BTSLconst [d] x))
	// result: (ORLconst [1<<uint32(d) | 1<<uint32(c)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTSLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ORLconst)
		v.AuxInt = 1<<uint32(d) | 1<<uint32(c)
		v.AddArg(x)
		return true
	}
	// match: (BTSLconst [c] (MOVLconst [d]))
	// result: (MOVLconst [d|(1<<uint32(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = d | (1 << uint32(c))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTSLconstmodify_0(v *Value) bool {
	// match: (BTSLconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (BTSLconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64BTSLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (BTSLconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (BTSLconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTSLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTSLmodify_0(v *Value) bool {
	// match: (BTSLmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (BTSLmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64BTSLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (BTSLmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (BTSLmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTSLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTSQconst_0(v *Value) bool {
	// match: (BTSQconst [c] (BTRQconst [c] x))
	// result: (BTSQconst [c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTRQconst || v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTSQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BTSQconst [c] (BTCQconst [c] x))
	// result: (BTSQconst [c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTCQconst || v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTSQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BTSQconst [c] (ORQconst [d] x))
	// result: (ORQconst [d | 1<<uint32(c)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ORQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ORQconst)
		v.AuxInt = d | 1<<uint32(c)
		v.AddArg(x)
		return true
	}
	// match: (BTSQconst [c] (BTSQconst [d] x))
	// result: (ORQconst [1<<uint32(d) | 1<<uint32(c)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTSQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ORQconst)
		v.AuxInt = 1<<uint32(d) | 1<<uint32(c)
		v.AddArg(x)
		return true
	}
	// match: (BTSQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [d|(1<<uint32(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = d | (1 << uint32(c))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTSQconstmodify_0(v *Value) bool {
	// match: (BTSQconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (BTSQconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64BTSQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (BTSQconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (BTSQconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTSQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64BTSQmodify_0(v *Value) bool {
	// match: (BTSQmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (BTSQmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64BTSQmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (BTSQmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (BTSQmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64BTSQmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLCC_0(v *Value) bool {
	// match: (CMOVLCC x y (InvertFlags cond))
	// result: (CMOVLLS x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLLS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLCC _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLCC _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLCC y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLCC y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLCC _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLCS_0(v *Value) bool {
	// match: (CMOVLCS x y (InvertFlags cond))
	// result: (CMOVLHI x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLHI)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLCS y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLCS y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLCS _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLCS _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLCS y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLEQ_0(v *Value) bool {
	// match: (CMOVLEQ x y (InvertFlags cond))
	// result: (CMOVLEQ x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLEQ)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLEQ _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLEQ y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLEQ y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLEQ y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLEQ y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLGE_0(v *Value) bool {
	// match: (CMOVLGE x y (InvertFlags cond))
	// result: (CMOVLLE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLLE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLGE _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLGE _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLGE _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLGE y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLGE y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLGT_0(v *Value) bool {
	// match: (CMOVLGT x y (InvertFlags cond))
	// result: (CMOVLLT x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLLT)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLGT y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLGT _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLGT _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLGT y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLGT y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLHI_0(v *Value) bool {
	// match: (CMOVLHI x y (InvertFlags cond))
	// result: (CMOVLCS x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLCS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLHI y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLHI _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLHI y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLHI y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLHI _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLLE_0(v *Value) bool {
	// match: (CMOVLLE x y (InvertFlags cond))
	// result: (CMOVLGE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLGE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLLE _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLLE y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLLE y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLLE _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLLE _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLLS_0(v *Value) bool {
	// match: (CMOVLLS x y (InvertFlags cond))
	// result: (CMOVLCC x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLCC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLLS _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLLS y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLLS _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLLS _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLLS y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLLT_0(v *Value) bool {
	// match: (CMOVLLT x y (InvertFlags cond))
	// result: (CMOVLGT x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLGT)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLLT y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLLT y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLLT y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLLT _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLLT _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVLNE_0(v *Value) bool {
	// match: (CMOVLNE x y (InvertFlags cond))
	// result: (CMOVLNE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVLNE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVLNE y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVLNE _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLNE _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLNE _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVLNE _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQCC_0(v *Value) bool {
	// match: (CMOVQCC x y (InvertFlags cond))
	// result: (CMOVQLS x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQLS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQCC _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQCC _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQCC y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQCC y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQCC _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQCS_0(v *Value) bool {
	// match: (CMOVQCS x y (InvertFlags cond))
	// result: (CMOVQHI x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQHI)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQCS y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQCS y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQCS _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQCS _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQCS y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQEQ_0(v *Value) bool {
	// match: (CMOVQEQ x y (InvertFlags cond))
	// result: (CMOVQEQ x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQEQ)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQEQ _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQEQ y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQEQ y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQEQ y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQEQ y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQEQ x _ (Select1 (BSFQ (ORQconst [c] _))))
	// cond: c != 0
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpSelect1 {
			break
		}
		v_2_0 := v_2.Args[0]
		if v_2_0.Op != OpAMD64BSFQ {
			break
		}
		v_2_0_0 := v_2_0.Args[0]
		if v_2_0_0.Op != OpAMD64ORQconst {
			break
		}
		c := v_2_0_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQGE_0(v *Value) bool {
	// match: (CMOVQGE x y (InvertFlags cond))
	// result: (CMOVQLE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQLE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQGE _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQGE _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQGE _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQGE y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQGE y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQGT_0(v *Value) bool {
	// match: (CMOVQGT x y (InvertFlags cond))
	// result: (CMOVQLT x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQLT)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQGT y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQGT _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQGT _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQGT y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQGT y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQHI_0(v *Value) bool {
	// match: (CMOVQHI x y (InvertFlags cond))
	// result: (CMOVQCS x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQCS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQHI y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQHI _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQHI y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQHI y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQHI _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQLE_0(v *Value) bool {
	// match: (CMOVQLE x y (InvertFlags cond))
	// result: (CMOVQGE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQGE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQLE _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQLE y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQLE y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQLE _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQLE _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQLS_0(v *Value) bool {
	// match: (CMOVQLS x y (InvertFlags cond))
	// result: (CMOVQCC x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQCC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQLS _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQLS y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQLS _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQLS _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQLS y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQLT_0(v *Value) bool {
	// match: (CMOVQLT x y (InvertFlags cond))
	// result: (CMOVQGT x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQGT)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQLT y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQLT y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQLT y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQLT _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQLT _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVQNE_0(v *Value) bool {
	// match: (CMOVQNE x y (InvertFlags cond))
	// result: (CMOVQNE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVQNE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVQNE y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVQNE _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQNE _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQNE _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVQNE _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWCC_0(v *Value) bool {
	// match: (CMOVWCC x y (InvertFlags cond))
	// result: (CMOVWLS x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWLS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWCC _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWCC _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWCC y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWCC y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWCC _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWCS_0(v *Value) bool {
	// match: (CMOVWCS x y (InvertFlags cond))
	// result: (CMOVWHI x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWHI)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWCS y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWCS y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWCS _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWCS _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWCS y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWEQ_0(v *Value) bool {
	// match: (CMOVWEQ x y (InvertFlags cond))
	// result: (CMOVWEQ x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWEQ)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWEQ _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWEQ y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWEQ y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWEQ y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWEQ y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWGE_0(v *Value) bool {
	// match: (CMOVWGE x y (InvertFlags cond))
	// result: (CMOVWLE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWLE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWGE _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWGE _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWGE _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWGE y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWGE y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWGT_0(v *Value) bool {
	// match: (CMOVWGT x y (InvertFlags cond))
	// result: (CMOVWLT x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWLT)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWGT y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWGT _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWGT _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWGT y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWGT y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWHI_0(v *Value) bool {
	// match: (CMOVWHI x y (InvertFlags cond))
	// result: (CMOVWCS x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWCS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWHI y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWHI _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWHI y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWHI y _ (FlagLT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWHI _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWLE_0(v *Value) bool {
	// match: (CMOVWLE x y (InvertFlags cond))
	// result: (CMOVWGE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWGE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWLE _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLE y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWLE y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWLE _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLE _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWLS_0(v *Value) bool {
	// match: (CMOVWLS x y (InvertFlags cond))
	// result: (CMOVWCC x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWCC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWLS _ x (FlagEQ))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLS y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWLS _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLS _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLS y _ (FlagLT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWLT_0(v *Value) bool {
	// match: (CMOVWLT x y (InvertFlags cond))
	// result: (CMOVWGT x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWGT)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWLT y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWLT y _ (FlagGT_UGT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWLT y _ (FlagGT_ULT))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWLT _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLT _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMOVWNE_0(v *Value) bool {
	// match: (CMOVWNE x y (InvertFlags cond))
	// result: (CMOVWNE x y cond)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64InvertFlags {
			break
		}
		cond := v_2.Args[0]
		v.reset(OpAMD64CMOVWNE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cond)
		return true
	}
	// match: (CMOVWNE y _ (FlagEQ))
	// result: y
	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (CMOVWNE _ x (FlagGT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWNE _ x (FlagGT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWNE _ x (FlagLT_ULT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWNE _ x (FlagLT_UGT))
	// result: x
	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPB_0(v *Value) bool {
	b := v.Block
	// match: (CMPB x (MOVLconst [c]))
	// result: (CMPBconst x [int64(int8(c))])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64CMPBconst)
		v.AuxInt = int64(int8(c))
		v.AddArg(x)
		return true
	}
	// match: (CMPB (MOVLconst [c]) x)
	// result: (InvertFlags (CMPBconst x [int64(int8(c))]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpAMD64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v0.AuxInt = int64(int8(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPB l:(MOVBload {sym} [off] ptr mem) x)
	// cond: canMergeLoad(v, l) && clobber(l)
	// result: (CMPBload {sym} [off] ptr x mem)
	for {
		x := v.Args[1]
		l := v.Args[0]
		if l.Op != OpAMD64MOVBload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoad(v, l) && clobber(l)) {
			break
		}
		v.reset(OpAMD64CMPBload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPB x l:(MOVBload {sym} [off] ptr mem))
	// cond: canMergeLoad(v, l) && clobber(l)
	// result: (InvertFlags (CMPBload {sym} [off] ptr x mem))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVBload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoad(v, l) && clobber(l)) {
			break
		}
		v.reset(OpAMD64InvertFlags)
		v0 := b.NewValue0(l.Pos, OpAMD64CMPBload, types.TypeFlags)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(x)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPBconst_0(v *Value) bool {
	b := v.Block
	// match: (CMPBconst (MOVLconst [x]) [y])
	// cond: int8(x)==int8(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) == int8(y)) {
			break
		}
		v.reset(OpAMD64FlagEQ)
		return true
	}
	// match: (CMPBconst (MOVLconst [x]) [y])
	// cond: int8(x)<int8(y) && uint8(x)<uint8(y)
	// result: (FlagLT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) < int8(y) && uint8(x) < uint8(y)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPBconst (MOVLconst [x]) [y])
	// cond: int8(x)<int8(y) && uint8(x)>uint8(y)
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) < int8(y) && uint8(x) > uint8(y)) {
			break
		}
		v.reset(OpAMD64FlagLT_UGT)
		return true
	}
	// match: (CMPBconst (MOVLconst [x]) [y])
	// cond: int8(x)>int8(y) && uint8(x)<uint8(y)
	// result: (FlagGT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) > int8(y) && uint8(x) < uint8(y)) {
			break
		}
		v.reset(OpAMD64FlagGT_ULT)
		return true
	}
	// match: (CMPBconst (MOVLconst [x]) [y])
	// cond: int8(x)>int8(y) && uint8(x)>uint8(y)
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) > int8(y) && uint8(x) > uint8(y)) {
			break
		}
		v.reset(OpAMD64FlagGT_UGT)
		return true
	}
	// match: (CMPBconst (ANDLconst _ [m]) [n])
	// cond: 0 <= int8(m) && int8(m) < int8(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int8(m) && int8(m) < int8(n)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPBconst (ANDL x y) [0])
	// result: (TESTB x y)
	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpAMD64TESTB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPBconst (ANDLconst [c] x) [0])
	// result: (TESTBconst [int64(int8(c))] x)
	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64TESTBconst)
		v.AuxInt = int64(int8(c))
		v.AddArg(x)
		return true
	}
	// match: (CMPBconst x [0])
	// result: (TESTB x x)
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64TESTB)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (CMPBconst l:(MOVBload {sym} [off] ptr mem) [c])
	// cond: l.Uses == 1 && validValAndOff(c, off) && clobber(l)
	// result: @l.Block (CMPBconstload {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		c := v.AuxInt
		l := v.Args[0]
		if l.Op != OpAMD64MOVBload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(l.Uses == 1 && validValAndOff(c, off) && clobber(l)) {
			break
		}
		b = l.Block
		v0 := b.NewValue0(l.Pos, OpAMD64CMPBconstload, types.TypeFlags)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = makeValAndOff(c, off)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPBconstload_0(v *Value) bool {
	// match: (CMPBconstload [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (CMPBconstload [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64CMPBconstload)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (CMPBconstload [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (CMPBconstload [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64CMPBconstload)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPBload_0(v *Value) bool {
	// match: (CMPBload [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (CMPBload [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64CMPBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (CMPBload [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (CMPBload [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64CMPBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (CMPBload {sym} [off] ptr (MOVLconst [c]) mem)
	// cond: validValAndOff(int64(int8(c)),off)
	// result: (CMPBconstload {sym} [makeValAndOff(int64(int8(c)),off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(validValAndOff(int64(int8(c)), off)) {
			break
		}
		v.reset(OpAMD64CMPBconstload)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPL_0(v *Value) bool {
	b := v.Block
	// match: (CMPL x (MOVLconst [c]))
	// result: (CMPLconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64CMPLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPL (MOVLconst [c]) x)
	// result: (InvertFlags (CMPLconst x [c]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpAMD64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPL l:(MOVLload {sym} [off] ptr mem) x)
	// cond: canMergeLoad(v, l) && clobber(l)
	// result: (CMPLload {sym} [off] ptr x mem)
	for {
		x := v.Args[1]
		l := v.Args[0]
		if l.Op != OpAMD64MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoad(v, l) && clobber(l)) {
			break
		}
		v.reset(OpAMD64CMPLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPL x l:(MOVLload {sym} [off] ptr mem))
	// cond: canMergeLoad(v, l) && clobber(l)
	// result: (InvertFlags (CMPLload {sym} [off] ptr x mem))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoad(v, l) && clobber(l)) {
			break
		}
		v.reset(OpAMD64InvertFlags)
		v0 := b.NewValue0(l.Pos, OpAMD64CMPLload, types.TypeFlags)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(x)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPLconst_0(v *Value) bool {
	// match: (CMPLconst (MOVLconst [x]) [y])
	// cond: int32(x)==int32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpAMD64FlagEQ)
		return true
	}
	// match: (CMPLconst (MOVLconst [x]) [y])
	// cond: int32(x)<int32(y) && uint32(x)<uint32(y)
	// result: (FlagLT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPLconst (MOVLconst [x]) [y])
	// cond: int32(x)<int32(y) && uint32(x)>uint32(y)
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpAMD64FlagLT_UGT)
		return true
	}
	// match: (CMPLconst (MOVLconst [x]) [y])
	// cond: int32(x)>int32(y) && uint32(x)<uint32(y)
	// result: (FlagGT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpAMD64FlagGT_ULT)
		return true
	}
	// match: (CMPLconst (MOVLconst [x]) [y])
	// cond: int32(x)>int32(y) && uint32(x)>uint32(y)
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpAMD64FlagGT_UGT)
		return true
	}
	// match: (CMPLconst (SHRLconst _ [c]) [n])
	// cond: 0 <= n && 0 < c && c <= 32 && (1<<uint64(32-c)) <= uint64(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHRLconst {
			break
		}
		c := v_0.AuxInt
		if !(0 <= n && 0 < c && c <= 32 && (1<<uint64(32-c)) <= uint64(n)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPLconst (ANDLconst _ [m]) [n])
	// cond: 0 <= int32(m) && int32(m) < int32(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int32(m) && int32(m) < int32(n)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPLconst (ANDL x y) [0])
	// result: (TESTL x y)
	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpAMD64TESTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPLconst (ANDLconst [c] x) [0])
	// result: (TESTLconst [c] x)
	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64TESTLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPLconst x [0])
	// result: (TESTL x x)
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64TESTL)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPLconst_10(v *Value) bool {
	b := v.Block
	// match: (CMPLconst l:(MOVLload {sym} [off] ptr mem) [c])
	// cond: l.Uses == 1 && validValAndOff(c, off) && clobber(l)
	// result: @l.Block (CMPLconstload {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		c := v.AuxInt
		l := v.Args[0]
		if l.Op != OpAMD64MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(l.Uses == 1 && validValAndOff(c, off) && clobber(l)) {
			break
		}
		b = l.Block
		v0 := b.NewValue0(l.Pos, OpAMD64CMPLconstload, types.TypeFlags)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = makeValAndOff(c, off)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPLconstload_0(v *Value) bool {
	// match: (CMPLconstload [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (CMPLconstload [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64CMPLconstload)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (CMPLconstload [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (CMPLconstload [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64CMPLconstload)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPLload_0(v *Value) bool {
	// match: (CMPLload [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (CMPLload [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64CMPLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (CMPLload [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (CMPLload [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64CMPLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (CMPLload {sym} [off] ptr (MOVLconst [c]) mem)
	// cond: validValAndOff(c,off)
	// result: (CMPLconstload {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(validValAndOff(c, off)) {
			break
		}
		v.reset(OpAMD64CMPLconstload)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPQ_0(v *Value) bool {
	b := v.Block
	// match: (CMPQ x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (CMPQconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpAMD64CMPQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPQ (MOVQconst [c]) x)
	// cond: is32Bit(c)
	// result: (InvertFlags (CMPQconst x [c]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		c := v_0.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpAMD64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPQ l:(MOVQload {sym} [off] ptr mem) x)
	// cond: canMergeLoad(v, l) && clobber(l)
	// result: (CMPQload {sym} [off] ptr x mem)
	for {
		x := v.Args[1]
		l := v.Args[0]
		if l.Op != OpAMD64MOVQload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoad(v, l) && clobber(l)) {
			break
		}
		v.reset(OpAMD64CMPQload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQ x l:(MOVQload {sym} [off] ptr mem))
	// cond: canMergeLoad(v, l) && clobber(l)
	// result: (InvertFlags (CMPQload {sym} [off] ptr x mem))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVQload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoad(v, l) && clobber(l)) {
			break
		}
		v.reset(OpAMD64InvertFlags)
		v0 := b.NewValue0(l.Pos, OpAMD64CMPQload, types.TypeFlags)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(x)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPQconst_0(v *Value) bool {
	// match: (CMPQconst (NEGQ (ADDQconst [-16] (ANDQconst [15] _))) [32])
	// result: (FlagLT_ULT)
	for {
		if v.AuxInt != 32 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64NEGQ {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAMD64ADDQconst || v_0_0.AuxInt != -16 {
			break
		}
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpAMD64ANDQconst || v_0_0_0.AuxInt != 15 {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPQconst (NEGQ (ADDQconst [ -8] (ANDQconst [7] _))) [32])
	// result: (FlagLT_ULT)
	for {
		if v.AuxInt != 32 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64NEGQ {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAMD64ADDQconst || v_0_0.AuxInt != -8 {
			break
		}
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpAMD64ANDQconst || v_0_0_0.AuxInt != 7 {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPQconst (MOVQconst [x]) [y])
	// cond: x==y
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		x := v_0.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpAMD64FlagEQ)
		return true
	}
	// match: (CMPQconst (MOVQconst [x]) [y])
	// cond: x<y && uint64(x)<uint64(y)
	// result: (FlagLT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		x := v_0.AuxInt
		if !(x < y && uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPQconst (MOVQconst [x]) [y])
	// cond: x<y && uint64(x)>uint64(y)
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		x := v_0.AuxInt
		if !(x < y && uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpAMD64FlagLT_UGT)
		return true
	}
	// match: (CMPQconst (MOVQconst [x]) [y])
	// cond: x>y && uint64(x)<uint64(y)
	// result: (FlagGT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		x := v_0.AuxInt
		if !(x > y && uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpAMD64FlagGT_ULT)
		return true
	}
	// match: (CMPQconst (MOVQconst [x]) [y])
	// cond: x>y && uint64(x)>uint64(y)
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		x := v_0.AuxInt
		if !(x > y && uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpAMD64FlagGT_UGT)
		return true
	}
	// match: (CMPQconst (MOVBQZX _) [c])
	// cond: 0xFF < c
	// result: (FlagLT_ULT)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVBQZX || !(0xFF < c) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPQconst (MOVWQZX _) [c])
	// cond: 0xFFFF < c
	// result: (FlagLT_ULT)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVWQZX || !(0xFFFF < c) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPQconst (MOVLQZX _) [c])
	// cond: 0xFFFFFFFF < c
	// result: (FlagLT_ULT)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLQZX || !(0xFFFFFFFF < c) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPQconst_10(v *Value) bool {
	b := v.Block
	// match: (CMPQconst (SHRQconst _ [c]) [n])
	// cond: 0 <= n && 0 < c && c <= 64 && (1<<uint64(64-c)) <= uint64(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHRQconst {
			break
		}
		c := v_0.AuxInt
		if !(0 <= n && 0 < c && c <= 64 && (1<<uint64(64-c)) <= uint64(n)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPQconst (ANDQconst _ [m]) [n])
	// cond: 0 <= m && m < n
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDQconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= m && m < n) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPQconst (ANDLconst _ [m]) [n])
	// cond: 0 <= m && m < n
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= m && m < n) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPQconst (ANDQ x y) [0])
	// result: (TESTQ x y)
	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDQ {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpAMD64TESTQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPQconst (ANDQconst [c] x) [0])
	// result: (TESTQconst [c] x)
	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDQconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64TESTQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPQconst x [0])
	// result: (TESTQ x x)
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64TESTQ)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (CMPQconst l:(MOVQload {sym} [off] ptr mem) [c])
	// cond: l.Uses == 1 && validValAndOff(c, off) && clobber(l)
	// result: @l.Block (CMPQconstload {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		c := v.AuxInt
		l := v.Args[0]
		if l.Op != OpAMD64MOVQload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(l.Uses == 1 && validValAndOff(c, off) && clobber(l)) {
			break
		}
		b = l.Block
		v0 := b.NewValue0(l.Pos, OpAMD64CMPQconstload, types.TypeFlags)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = makeValAndOff(c, off)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPQconstload_0(v *Value) bool {
	// match: (CMPQconstload [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (CMPQconstload [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64CMPQconstload)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQconstload [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (CMPQconstload [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64CMPQconstload)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPQload_0(v *Value) bool {
	// match: (CMPQload [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (CMPQload [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64CMPQload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQload [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (CMPQload [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64CMPQload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQload {sym} [off] ptr (MOVQconst [c]) mem)
	// cond: validValAndOff(c,off)
	// result: (CMPQconstload {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(validValAndOff(c, off)) {
			break
		}
		v.reset(OpAMD64CMPQconstload)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPW_0(v *Value) bool {
	b := v.Block
	// match: (CMPW x (MOVLconst [c]))
	// result: (CMPWconst x [int64(int16(c))])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64CMPWconst)
		v.AuxInt = int64(int16(c))
		v.AddArg(x)
		return true
	}
	// match: (CMPW (MOVLconst [c]) x)
	// result: (InvertFlags (CMPWconst x [int64(int16(c))]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpAMD64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v0.AuxInt = int64(int16(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMPW l:(MOVWload {sym} [off] ptr mem) x)
	// cond: canMergeLoad(v, l) && clobber(l)
	// result: (CMPWload {sym} [off] ptr x mem)
	for {
		x := v.Args[1]
		l := v.Args[0]
		if l.Op != OpAMD64MOVWload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoad(v, l) && clobber(l)) {
			break
		}
		v.reset(OpAMD64CMPWload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPW x l:(MOVWload {sym} [off] ptr mem))
	// cond: canMergeLoad(v, l) && clobber(l)
	// result: (InvertFlags (CMPWload {sym} [off] ptr x mem))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVWload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoad(v, l) && clobber(l)) {
			break
		}
		v.reset(OpAMD64InvertFlags)
		v0 := b.NewValue0(l.Pos, OpAMD64CMPWload, types.TypeFlags)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(x)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPWconst_0(v *Value) bool {
	b := v.Block
	// match: (CMPWconst (MOVLconst [x]) [y])
	// cond: int16(x)==int16(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) == int16(y)) {
			break
		}
		v.reset(OpAMD64FlagEQ)
		return true
	}
	// match: (CMPWconst (MOVLconst [x]) [y])
	// cond: int16(x)<int16(y) && uint16(x)<uint16(y)
	// result: (FlagLT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) < int16(y) && uint16(x) < uint16(y)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPWconst (MOVLconst [x]) [y])
	// cond: int16(x)<int16(y) && uint16(x)>uint16(y)
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) < int16(y) && uint16(x) > uint16(y)) {
			break
		}
		v.reset(OpAMD64FlagLT_UGT)
		return true
	}
	// match: (CMPWconst (MOVLconst [x]) [y])
	// cond: int16(x)>int16(y) && uint16(x)<uint16(y)
	// result: (FlagGT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) > int16(y) && uint16(x) < uint16(y)) {
			break
		}
		v.reset(OpAMD64FlagGT_ULT)
		return true
	}
	// match: (CMPWconst (MOVLconst [x]) [y])
	// cond: int16(x)>int16(y) && uint16(x)>uint16(y)
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) > int16(y) && uint16(x) > uint16(y)) {
			break
		}
		v.reset(OpAMD64FlagGT_UGT)
		return true
	}
	// match: (CMPWconst (ANDLconst _ [m]) [n])
	// cond: 0 <= int16(m) && int16(m) < int16(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int16(m) && int16(m) < int16(n)) {
			break
		}
		v.reset(OpAMD64FlagLT_ULT)
		return true
	}
	// match: (CMPWconst (ANDL x y) [0])
	// result: (TESTW x y)
	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpAMD64TESTW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPWconst (ANDLconst [c] x) [0])
	// result: (TESTWconst [int64(int16(c))] x)
	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64TESTWconst)
		v.AuxInt = int64(int16(c))
		v.AddArg(x)
		return true
	}
	// match: (CMPWconst x [0])
	// result: (TESTW x x)
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64TESTW)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (CMPWconst l:(MOVWload {sym} [off] ptr mem) [c])
	// cond: l.Uses == 1 && validValAndOff(c, off) && clobber(l)
	// result: @l.Block (CMPWconstload {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		c := v.AuxInt
		l := v.Args[0]
		if l.Op != OpAMD64MOVWload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(l.Uses == 1 && validValAndOff(c, off) && clobber(l)) {
			break
		}
		b = l.Block
		v0 := b.NewValue0(l.Pos, OpAMD64CMPWconstload, types.TypeFlags)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = makeValAndOff(c, off)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPWconstload_0(v *Value) bool {
	// match: (CMPWconstload [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (CMPWconstload [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64CMPWconstload)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (CMPWconstload [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (CMPWconstload [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64CMPWconstload)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPWload_0(v *Value) bool {
	// match: (CMPWload [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (CMPWload [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64CMPWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (CMPWload [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (CMPWload [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64CMPWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (CMPWload {sym} [off] ptr (MOVLconst [c]) mem)
	// cond: validValAndOff(int64(int16(c)),off)
	// result: (CMPWconstload {sym} [makeValAndOff(int64(int16(c)),off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(validValAndOff(int64(int16(c)), off)) {
			break
		}
		v.reset(OpAMD64CMPWconstload)
		v.AuxInt = makeValAndOff(int64(int16(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPXCHGLlock_0(v *Value) bool {
	// match: (CMPXCHGLlock [off1] {sym} (ADDQconst [off2] ptr) old new_ mem)
	// cond: is32Bit(off1+off2)
	// result: (CMPXCHGLlock [off1+off2] {sym} ptr old new_ mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64CMPXCHGLlock)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64CMPXCHGQlock_0(v *Value) bool {
	// match: (CMPXCHGQlock [off1] {sym} (ADDQconst [off2] ptr) old new_ mem)
	// cond: is32Bit(off1+off2)
	// result: (CMPXCHGQlock [off1+off2] {sym} ptr old new_ mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64CMPXCHGQlock)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64DIVSD_0(v *Value) bool {
	// match: (DIVSD x l:(MOVSDload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (DIVSDload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVSDload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
			break
		}
		v.reset(OpAMD64DIVSDload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64DIVSDload_0(v *Value) bool {
	// match: (DIVSDload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (DIVSDload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64DIVSDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (DIVSDload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (DIVSDload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64DIVSDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64DIVSS_0(v *Value) bool {
	// match: (DIVSS x l:(MOVSSload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (DIVSSload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVSSload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
			break
		}
		v.reset(OpAMD64DIVSSload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64DIVSSload_0(v *Value) bool {
	// match: (DIVSSload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (DIVSSload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64DIVSSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (DIVSSload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (DIVSSload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64DIVSSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64HMULL_0(v *Value) bool {
	// match: (HMULL x y)
	// cond: !x.rematerializeable() && y.rematerializeable()
	// result: (HMULL y x)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(!x.rematerializeable() && y.rematerializeable()) {
			break
		}
		v.reset(OpAMD64HMULL)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64HMULLU_0(v *Value) bool {
	// match: (HMULLU x y)
	// cond: !x.rematerializeable() && y.rematerializeable()
	// result: (HMULLU y x)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(!x.rematerializeable() && y.rematerializeable()) {
			break
		}
		v.reset(OpAMD64HMULLU)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64HMULQ_0(v *Value) bool {
	// match: (HMULQ x y)
	// cond: !x.rematerializeable() && y.rematerializeable()
	// result: (HMULQ y x)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(!x.rematerializeable() && y.rematerializeable()) {
			break
		}
		v.reset(OpAMD64HMULQ)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64HMULQU_0(v *Value) bool {
	// match: (HMULQU x y)
	// cond: !x.rematerializeable() && y.rematerializeable()
	// result: (HMULQU y x)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(!x.rematerializeable() && y.rematerializeable()) {
			break
		}
		v.reset(OpAMD64HMULQU)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAL_0(v *Value) bool {
	// match: (LEAL [c] {s} (ADDLconst [d] x))
	// cond: is32Bit(c+d)
	// result: (LEAL [c+d] {s} x)
	for {
		c := v.AuxInt
		s := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAL)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}
	// match: (LEAL [c] {s} (ADDL x y))
	// cond: x.Op != OpSB && y.Op != OpSB
	// result: (LEAL1 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v_0.Args[_i0]
			y := v_0.Args[1^_i0]
			if !(x.Op != OpSB && y.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64LEAL1)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAL1_0(v *Value) bool {
	// match: (LEAL1 [c] {s} (ADDLconst [d] x) y)
	// cond: is32Bit(c+d) && x.Op != OpSB
	// result: (LEAL1 [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64ADDLconst {
				continue
			}
			d := v_0.AuxInt
			x := v_0.Args[0]
			y := v.Args[1^_i0]
			if !(is32Bit(c+d) && x.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64LEAL1)
			v.AuxInt = c + d
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (LEAL1 [c] {s} x (SHLLconst [1] y))
	// result: (LEAL2 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 1 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAL2)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (LEAL1 [c] {s} x (SHLLconst [2] y))
	// result: (LEAL4 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 2 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAL4)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (LEAL1 [c] {s} x (SHLLconst [3] y))
	// result: (LEAL8 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 3 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAL8)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAL2_0(v *Value) bool {
	// match: (LEAL2 [c] {s} (ADDLconst [d] x) y)
	// cond: is32Bit(c+d) && x.Op != OpSB
	// result: (LEAL2 [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAL2)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAL2 [c] {s} x (ADDLconst [d] y))
	// cond: is32Bit(c+2*d) && y.Op != OpSB
	// result: (LEAL2 [c+2*d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+2*d) && y.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAL2)
		v.AuxInt = c + 2*d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAL2 [c] {s} x (SHLLconst [1] y))
	// result: (LEAL4 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64LEAL4)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAL2 [c] {s} x (SHLLconst [2] y))
	// result: (LEAL8 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 2 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64LEAL8)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAL4_0(v *Value) bool {
	// match: (LEAL4 [c] {s} (ADDLconst [d] x) y)
	// cond: is32Bit(c+d) && x.Op != OpSB
	// result: (LEAL4 [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAL4)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAL4 [c] {s} x (ADDLconst [d] y))
	// cond: is32Bit(c+4*d) && y.Op != OpSB
	// result: (LEAL4 [c+4*d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+4*d) && y.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAL4)
		v.AuxInt = c + 4*d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAL4 [c] {s} x (SHLLconst [1] y))
	// result: (LEAL8 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLLconst || v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64LEAL8)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAL8_0(v *Value) bool {
	// match: (LEAL8 [c] {s} (ADDLconst [d] x) y)
	// cond: is32Bit(c+d) && x.Op != OpSB
	// result: (LEAL8 [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAL8)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAL8 [c] {s} x (ADDLconst [d] y))
	// cond: is32Bit(c+8*d) && y.Op != OpSB
	// result: (LEAL8 [c+8*d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+8*d) && y.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAL8)
		v.AuxInt = c + 8*d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAQ_0(v *Value) bool {
	// match: (LEAQ [c] {s} (ADDQconst [d] x))
	// cond: is32Bit(c+d)
	// result: (LEAQ [c+d] {s} x)
	for {
		c := v.AuxInt
		s := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64LEAQ)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}
	// match: (LEAQ [c] {s} (ADDQ x y))
	// cond: x.Op != OpSB && y.Op != OpSB
	// result: (LEAQ1 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v_0.Args[_i0]
			y := v_0.Args[1^_i0]
			if !(x.Op != OpSB && y.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64LEAQ1)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (LEAQ [off1] {sym1} (LEAQ [off2] {sym2} x))
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (LEAQ [off1+off2] {mergeSym(sym1,sym2)} x)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64LEAQ)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		return true
	}
	// match: (LEAQ [off1] {sym1} (LEAQ1 [off2] {sym2} x y))
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (LEAQ1 [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64LEAQ1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ [off1] {sym1} (LEAQ2 [off2] {sym2} x y))
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (LEAQ2 [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ2 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64LEAQ2)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ [off1] {sym1} (LEAQ4 [off2] {sym2} x y))
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (LEAQ4 [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ4 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64LEAQ4)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ [off1] {sym1} (LEAQ8 [off2] {sym2} x y))
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (LEAQ8 [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64LEAQ8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAQ1_0(v *Value) bool {
	// match: (LEAQ1 [c] {s} (ADDQconst [d] x) y)
	// cond: is32Bit(c+d) && x.Op != OpSB
	// result: (LEAQ1 [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_0.AuxInt
			x := v_0.Args[0]
			y := v.Args[1^_i0]
			if !(is32Bit(c+d) && x.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64LEAQ1)
			v.AuxInt = c + d
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (LEAQ1 [c] {s} x (SHLQconst [1] y))
	// result: (LEAQ2 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 1 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAQ2)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (LEAQ1 [c] {s} x (SHLQconst [2] y))
	// result: (LEAQ4 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 2 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAQ4)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (LEAQ1 [c] {s} x (SHLQconst [3] y))
	// result: (LEAQ8 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpAMD64LEAQ8)
			v.AuxInt = c
			v.Aux = s
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (LEAQ1 [off1] {sym1} (LEAQ [off2] {sym2} x) y)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB
	// result: (LEAQ1 [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64LEAQ {
				continue
			}
			off2 := v_0.AuxInt
			sym2 := v_0.Aux
			x := v_0.Args[0]
			y := v.Args[1^_i0]
			if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64LEAQ1)
			v.AuxInt = off1 + off2
			v.Aux = mergeSym(sym1, sym2)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAQ2_0(v *Value) bool {
	// match: (LEAQ2 [c] {s} (ADDQconst [d] x) y)
	// cond: is32Bit(c+d) && x.Op != OpSB
	// result: (LEAQ2 [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ2)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ2 [c] {s} x (ADDQconst [d] y))
	// cond: is32Bit(c+2*d) && y.Op != OpSB
	// result: (LEAQ2 [c+2*d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+2*d) && y.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ2)
		v.AuxInt = c + 2*d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ2 [c] {s} x (SHLQconst [1] y))
	// result: (LEAQ4 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64LEAQ4)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ2 [c] {s} x (SHLQconst [2] y))
	// result: (LEAQ8 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 2 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64LEAQ8)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ2 [off1] {sym1} (LEAQ [off2] {sym2} x) y)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB
	// result: (LEAQ2 [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ2)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAQ4_0(v *Value) bool {
	// match: (LEAQ4 [c] {s} (ADDQconst [d] x) y)
	// cond: is32Bit(c+d) && x.Op != OpSB
	// result: (LEAQ4 [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ4)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ4 [c] {s} x (ADDQconst [d] y))
	// cond: is32Bit(c+4*d) && y.Op != OpSB
	// result: (LEAQ4 [c+4*d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+4*d) && y.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ4)
		v.AuxInt = c + 4*d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ4 [c] {s} x (SHLQconst [1] y))
	// result: (LEAQ8 [c] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64LEAQ8)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ4 [off1] {sym1} (LEAQ [off2] {sym2} x) y)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB
	// result: (LEAQ4 [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ4)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64LEAQ8_0(v *Value) bool {
	// match: (LEAQ8 [c] {s} (ADDQconst [d] x) y)
	// cond: is32Bit(c+d) && x.Op != OpSB
	// result: (LEAQ8 [c+d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ8)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ8 [c] {s} x (ADDQconst [d] y))
	// cond: is32Bit(c+8*d) && y.Op != OpSB
	// result: (LEAQ8 [c+8*d] {s} x y)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+8*d) && y.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ8)
		v.AuxInt = c + 8*d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (LEAQ8 [off1] {sym1} (LEAQ [off2] {sym2} x) y)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB
	// result: (LEAQ8 [off1+off2] {mergeSym(sym1,sym2)} x y)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(OpAMD64LEAQ8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBQSX_0(v *Value) bool {
	b := v.Block
	// match: (MOVBQSX x:(MOVBload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVBload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVBQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQSX x:(MOVWload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVWload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVBQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQSX x:(MOVLload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVLload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVBQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQSX x:(MOVQload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVQload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVBQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQSX (ANDLconst [c] x))
	// cond: c & 0x80 == 0
	// result: (ANDLconst [c & 0x7f] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = c & 0x7f
		v.AddArg(x)
		return true
	}
	// match: (MOVBQSX (MOVBQSX x))
	// result: (MOVBQSX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVBQSX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVBQSX)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBQSXload_0(v *Value) bool {
	// match: (MOVBQSXload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBQSX x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpAMD64MOVBQSX)
		v.AddArg(x)
		return true
	}
	// match: (MOVBQSXload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBQSXload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVBQSXload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBQZX_0(v *Value) bool {
	b := v.Block
	// match: (MOVBQZX x:(MOVBload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVBload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVBload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQZX x:(MOVWload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVWload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVBload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQZX x:(MOVLload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVLload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVBload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQZX x:(MOVQload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVQload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVBload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQZX x)
	// cond: zeroUpper56Bits(x,3)
	// result: x
	for {
		x := v.Args[0]
		if !(zeroUpper56Bits(x, 3)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBQZX x:(MOVBloadidx1 [off] {sym} ptr idx mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBloadidx1 <v.Type> [off] {sym} ptr idx mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVBloadidx1 {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBloadidx1, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBQZX (ANDLconst [c] x))
	// result: (ANDLconst [c & 0xff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = c & 0xff
		v.AddArg(x)
		return true
	}
	// match: (MOVBQZX (MOVBQZX x))
	// result: (MOVBQZX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVBQZX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBatomicload_0(v *Value) bool {
	// match: (MOVBatomicload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBatomicload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVBatomicload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBatomicload [off1] {sym1} (LEAQ [off2] {sym2} ptr) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBatomicload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVBatomicload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBload_0(v *Value) bool {
	// match: (MOVBload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBQZX x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
	// match: (MOVBload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBloadidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVBloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off] {sym} (ADDQ ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVBloadidx1 [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVBloadidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBload [off1] {sym1} (LEAL [off2] {sym2} base) mem)
	// cond: canMergeSym(sym1, sym2) && is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpAMD64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym} (ADDLconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVLconst [int64(read8(sym, off))])
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = int64(read8(sym, off))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBloadidx1_0(v *Value) bool {
	// match: (MOVBloadidx1 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVBloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v.Args[1^_i0]
			if !(is32Bit(c + d)) {
				continue
			}
			v.reset(OpAMD64MOVBloadidx1)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBloadidx1 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+d)
	// result: (MOVBloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			if !(is32Bit(c + d)) {
				continue
			}
			v.reset(OpAMD64MOVBloadidx1)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBloadidx1 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+c)
	// result: (MOVBload [i+c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			p := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(i + c)) {
				continue
			}
			v.reset(OpAMD64MOVBload)
			v.AuxInt = i + c
			v.Aux = s
			v.AddArg(p)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBstore_0(v *Value) bool {
	// match: (MOVBstore [off] {sym} ptr y:(SETL x) mem)
	// cond: y.Uses == 1
	// result: (SETLstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETL {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETLstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETLE x) mem)
	// cond: y.Uses == 1
	// result: (SETLEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETLE {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETLEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETG x) mem)
	// cond: y.Uses == 1
	// result: (SETGstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETG {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETGstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETGE x) mem)
	// cond: y.Uses == 1
	// result: (SETGEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETGE {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETGEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETEQ x) mem)
	// cond: y.Uses == 1
	// result: (SETEQstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETEQ {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETEQstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETNE x) mem)
	// cond: y.Uses == 1
	// result: (SETNEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETNE {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETNEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETB x) mem)
	// cond: y.Uses == 1
	// result: (SETBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETB {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETBE x) mem)
	// cond: y.Uses == 1
	// result: (SETBEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETBE {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETBEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETA x) mem)
	// cond: y.Uses == 1
	// result: (SETAstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETA {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETAstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr y:(SETAE x) mem)
	// cond: y.Uses == 1
	// result: (SETAEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SETAE {
			break
		}
		x := y.Args[0]
		if !(y.Uses == 1) {
			break
		}
		v.reset(OpAMD64SETAEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBstore_10(v *Value) bool {
	b := v.Block
	// match: (MOVBstore [off] {sym} ptr (MOVBQSX x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVBQSX {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBQZX x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVBQZX {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (ADDQconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVLconst [c]) mem)
	// cond: validOff(off)
	// result: (MOVBstoreconst [makeValAndOff(int64(int8(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(validOff(off)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVQconst [c]) mem)
	// cond: validOff(off)
	// result: (MOVBstoreconst [makeValAndOff(int64(int8(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(validOff(off)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBstoreidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVBstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} (ADDQ ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVBstoreidx1 [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			val := v.Args[1]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVBstoreidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVBstore [i] {s} p w x0:(MOVBstore [i-1] {s} p (SHRWconst [8] w) mem))
	// cond: x0.Uses == 1 && clobber(x0)
	// result: (MOVWstore [i-1] {s} p (ROLWconst <w.Type> [8] w) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpAMD64MOVBstore || x0.AuxInt != i-1 || x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpAMD64SHRWconst || x0_1.AuxInt != 8 || w != x0_1.Args[0] || !(x0.Uses == 1 && clobber(x0)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x0.Pos, OpAMD64ROLWconst, w.Type)
		v0.AuxInt = 8
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p w x2:(MOVBstore [i-1] {s} p (SHRLconst [8] w) x1:(MOVBstore [i-2] {s} p (SHRLconst [16] w) x0:(MOVBstore [i-3] {s} p (SHRLconst [24] w) mem))))
	// cond: x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)
	// result: (MOVLstore [i-3] {s} p (BSWAPL <w.Type> w) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x2 := v.Args[2]
		if x2.Op != OpAMD64MOVBstore || x2.AuxInt != i-1 || x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpAMD64SHRLconst || x2_1.AuxInt != 8 || w != x2_1.Args[0] {
			break
		}
		x1 := x2.Args[2]
		if x1.Op != OpAMD64MOVBstore || x1.AuxInt != i-2 || x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpAMD64SHRLconst || x1_1.AuxInt != 16 || w != x1_1.Args[0] {
			break
		}
		x0 := x1.Args[2]
		if x0.Op != OpAMD64MOVBstore || x0.AuxInt != i-3 || x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpAMD64SHRLconst || x0_1.AuxInt != 24 || w != x0_1.Args[0] || !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i - 3
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x0.Pos, OpAMD64BSWAPL, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBstore_20(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVBstore [i] {s} p w x6:(MOVBstore [i-1] {s} p (SHRQconst [8] w) x5:(MOVBstore [i-2] {s} p (SHRQconst [16] w) x4:(MOVBstore [i-3] {s} p (SHRQconst [24] w) x3:(MOVBstore [i-4] {s} p (SHRQconst [32] w) x2:(MOVBstore [i-5] {s} p (SHRQconst [40] w) x1:(MOVBstore [i-6] {s} p (SHRQconst [48] w) x0:(MOVBstore [i-7] {s} p (SHRQconst [56] w) mem))))))))
	// cond: x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)
	// result: (MOVQstore [i-7] {s} p (BSWAPQ <w.Type> w) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x6 := v.Args[2]
		if x6.Op != OpAMD64MOVBstore || x6.AuxInt != i-1 || x6.Aux != s {
			break
		}
		_ = x6.Args[2]
		if p != x6.Args[0] {
			break
		}
		x6_1 := x6.Args[1]
		if x6_1.Op != OpAMD64SHRQconst || x6_1.AuxInt != 8 || w != x6_1.Args[0] {
			break
		}
		x5 := x6.Args[2]
		if x5.Op != OpAMD64MOVBstore || x5.AuxInt != i-2 || x5.Aux != s {
			break
		}
		_ = x5.Args[2]
		if p != x5.Args[0] {
			break
		}
		x5_1 := x5.Args[1]
		if x5_1.Op != OpAMD64SHRQconst || x5_1.AuxInt != 16 || w != x5_1.Args[0] {
			break
		}
		x4 := x5.Args[2]
		if x4.Op != OpAMD64MOVBstore || x4.AuxInt != i-3 || x4.Aux != s {
			break
		}
		_ = x4.Args[2]
		if p != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpAMD64SHRQconst || x4_1.AuxInt != 24 || w != x4_1.Args[0] {
			break
		}
		x3 := x4.Args[2]
		if x3.Op != OpAMD64MOVBstore || x3.AuxInt != i-4 || x3.Aux != s {
			break
		}
		_ = x3.Args[2]
		if p != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpAMD64SHRQconst || x3_1.AuxInt != 32 || w != x3_1.Args[0] {
			break
		}
		x2 := x3.Args[2]
		if x2.Op != OpAMD64MOVBstore || x2.AuxInt != i-5 || x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpAMD64SHRQconst || x2_1.AuxInt != 40 || w != x2_1.Args[0] {
			break
		}
		x1 := x2.Args[2]
		if x1.Op != OpAMD64MOVBstore || x1.AuxInt != i-6 || x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpAMD64SHRQconst || x1_1.AuxInt != 48 || w != x1_1.Args[0] {
			break
		}
		x0 := x1.Args[2]
		if x0.Op != OpAMD64MOVBstore || x0.AuxInt != i-7 || x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpAMD64SHRQconst || x0_1.AuxInt != 56 || w != x0_1.Args[0] || !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = i - 7
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x0.Pos, OpAMD64BSWAPQ, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SHRWconst [8] w) x:(MOVBstore [i-1] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-1] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRWconst || v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SHRLconst [8] w) x:(MOVBstore [i-1] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-1] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRLconst || v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SHRQconst [8] w) x:(MOVBstore [i-1] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-1] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRQconst || v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p w x:(MOVBstore [i+1] {s} p (SHRWconst [8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstore || x.AuxInt != i+1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpAMD64SHRWconst || x_1.AuxInt != 8 || w != x_1.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p w x:(MOVBstore [i+1] {s} p (SHRLconst [8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstore || x.AuxInt != i+1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpAMD64SHRLconst || x_1.AuxInt != 8 || w != x_1.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p w x:(MOVBstore [i+1] {s} p (SHRQconst [8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstore || x.AuxInt != i+1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpAMD64SHRQconst || x_1.AuxInt != 8 || w != x_1.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SHRLconst [j] w) x:(MOVBstore [i-1] {s} p w0:(SHRLconst [j-8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-1] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpAMD64SHRLconst || w0.AuxInt != j-8 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p (SHRQconst [j] w) x:(MOVBstore [i-1] {s} p w0:(SHRQconst [j-8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstore [i-1] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRQconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstore || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpAMD64SHRQconst || w0.AuxInt != j-8 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i] {s} p x1:(MOVBload [j] {s2} p2 mem) mem2:(MOVBstore [i-1] {s} p x2:(MOVBload [j-1] {s2} p2 mem) mem))
	// cond: x1.Uses == 1 && x2.Uses == 1 && mem2.Uses == 1 && clobber(x1) && clobber(x2) && clobber(mem2)
	// result: (MOVWstore [i-1] {s} p (MOVWload [j-1] {s2} p2 mem) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpAMD64MOVBload {
			break
		}
		j := x1.AuxInt
		s2 := x1.Aux
		mem := x1.Args[1]
		p2 := x1.Args[0]
		mem2 := v.Args[2]
		if mem2.Op != OpAMD64MOVBstore || mem2.AuxInt != i-1 || mem2.Aux != s {
			break
		}
		_ = mem2.Args[2]
		if p != mem2.Args[0] {
			break
		}
		x2 := mem2.Args[1]
		if x2.Op != OpAMD64MOVBload || x2.AuxInt != j-1 || x2.Aux != s2 {
			break
		}
		_ = x2.Args[1]
		if p2 != x2.Args[0] || mem != x2.Args[1] || mem != mem2.Args[2] || !(x1.Uses == 1 && x2.Uses == 1 && mem2.Uses == 1 && clobber(x1) && clobber(x2) && clobber(mem2)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x2.Pos, OpAMD64MOVWload, typ.UInt16)
		v0.AuxInt = j - 1
		v0.Aux = s2
		v0.AddArg(p2)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBstore_30(v *Value) bool {
	// match: (MOVBstore [off1] {sym1} (LEAL [off2] {sym2} base) val mem)
	// cond: canMergeSym(sym1, sym2) && is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (ADDLconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBstoreconst_0(v *Value) bool {
	// match: (MOVBstoreconst [sc] {s} (ADDQconst [off] ptr) mem)
	// cond: ValAndOff(sc).canAdd(off)
	// result: (MOVBstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [sc] {sym1} (LEAQ [off] {sym2} ptr) mem)
	// cond: canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVBstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [x] {sym1} (LEAQ1 [off] {sym2} ptr idx) mem)
	// cond: canMergeSym(sym1, sym2)
	// result: (MOVBstoreconstidx1 [ValAndOff(x).add(off)] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		x := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [x] {sym} (ADDQ ptr idx) mem)
	// result: (MOVBstoreconstidx1 [x] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		v.reset(OpAMD64MOVBstoreconstidx1)
		v.AuxInt = x
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [c] {s} p x:(MOVBstoreconst [a] {s} p mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 1 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVWstoreconst [makeValAndOff(ValAndOff(a).Val()&0xff | ValAndOff(c).Val()<<8, ValAndOff(a).Off())] {s} p mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpAMD64MOVBstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(x.Uses == 1 && ValAndOff(a).Off()+1 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xff|ValAndOff(c).Val()<<8, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [a] {s} p x:(MOVBstoreconst [c] {s} p mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 1 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVWstoreconst [makeValAndOff(ValAndOff(a).Val()&0xff | ValAndOff(c).Val()<<8, ValAndOff(a).Off())] {s} p mem)
	for {
		a := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpAMD64MOVBstoreconst {
			break
		}
		c := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(x.Uses == 1 && ValAndOff(a).Off()+1 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xff|ValAndOff(c).Val()<<8, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [sc] {sym1} (LEAL [off] {sym2} ptr) mem)
	// cond: canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVBstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconst [sc] {s} (ADDLconst [off] ptr) mem)
	// cond: ValAndOff(sc).canAdd(off)
	// result: (MOVBstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBstoreconstidx1_0(v *Value) bool {
	// match: (MOVBstoreconstidx1 [x] {sym} (ADDQconst [c] ptr) idx mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVBstoreconstidx1 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconstidx1 [x] {sym} ptr (ADDQconst [c] idx) mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVBstoreconstidx1 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVBstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreconstidx1 [c] {s} p i x:(MOVBstoreconstidx1 [a] {s} p i mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 1 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVWstoreconstidx1 [makeValAndOff(ValAndOff(a).Val()&0xff | ValAndOff(c).Val()<<8, ValAndOff(a).Off())] {s} p i mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		i := v.Args[1]
		x := v.Args[2]
		if x.Op != OpAMD64MOVBstoreconstidx1 {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || i != x.Args[1] || !(x.Uses == 1 && ValAndOff(a).Off()+1 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconstidx1)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xff|ValAndOff(c).Val()<<8, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(i)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBstoreidx1_0(v *Value) bool {
	b := v.Block
	// match: (MOVBstoreidx1 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVBstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVBstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+d)
	// result: (MOVBstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVBstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [i] {s} p idx w x0:(MOVBstoreidx1 [i-1] {s} p idx (SHRWconst [8] w) mem))
	// cond: x0.Uses == 1 && clobber(x0)
	// result: (MOVWstoreidx1 [i-1] {s} p idx (ROLWconst <w.Type> [8] w) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x0 := v.Args[3]
		if x0.Op != OpAMD64MOVBstoreidx1 || x0.AuxInt != i-1 || x0.Aux != s {
			break
		}
		mem := x0.Args[3]
		if p != x0.Args[0] || idx != x0.Args[1] {
			break
		}
		x0_2 := x0.Args[2]
		if x0_2.Op != OpAMD64SHRWconst || x0_2.AuxInt != 8 || w != x0_2.Args[0] || !(x0.Uses == 1 && clobber(x0)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpAMD64ROLWconst, w.Type)
		v0.AuxInt = 8
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [i] {s} p idx w x2:(MOVBstoreidx1 [i-1] {s} p idx (SHRLconst [8] w) x1:(MOVBstoreidx1 [i-2] {s} p idx (SHRLconst [16] w) x0:(MOVBstoreidx1 [i-3] {s} p idx (SHRLconst [24] w) mem))))
	// cond: x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)
	// result: (MOVLstoreidx1 [i-3] {s} p idx (BSWAPL <w.Type> w) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x2 := v.Args[3]
		if x2.Op != OpAMD64MOVBstoreidx1 || x2.AuxInt != i-1 || x2.Aux != s {
			break
		}
		_ = x2.Args[3]
		if p != x2.Args[0] || idx != x2.Args[1] {
			break
		}
		x2_2 := x2.Args[2]
		if x2_2.Op != OpAMD64SHRLconst || x2_2.AuxInt != 8 || w != x2_2.Args[0] {
			break
		}
		x1 := x2.Args[3]
		if x1.Op != OpAMD64MOVBstoreidx1 || x1.AuxInt != i-2 || x1.Aux != s {
			break
		}
		_ = x1.Args[3]
		if p != x1.Args[0] || idx != x1.Args[1] {
			break
		}
		x1_2 := x1.Args[2]
		if x1_2.Op != OpAMD64SHRLconst || x1_2.AuxInt != 16 || w != x1_2.Args[0] {
			break
		}
		x0 := x1.Args[3]
		if x0.Op != OpAMD64MOVBstoreidx1 || x0.AuxInt != i-3 || x0.Aux != s {
			break
		}
		mem := x0.Args[3]
		if p != x0.Args[0] || idx != x0.Args[1] {
			break
		}
		x0_2 := x0.Args[2]
		if x0_2.Op != OpAMD64SHRLconst || x0_2.AuxInt != 24 || w != x0_2.Args[0] || !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = i - 3
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpAMD64BSWAPL, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [i] {s} p idx w x6:(MOVBstoreidx1 [i-1] {s} p idx (SHRQconst [8] w) x5:(MOVBstoreidx1 [i-2] {s} p idx (SHRQconst [16] w) x4:(MOVBstoreidx1 [i-3] {s} p idx (SHRQconst [24] w) x3:(MOVBstoreidx1 [i-4] {s} p idx (SHRQconst [32] w) x2:(MOVBstoreidx1 [i-5] {s} p idx (SHRQconst [40] w) x1:(MOVBstoreidx1 [i-6] {s} p idx (SHRQconst [48] w) x0:(MOVBstoreidx1 [i-7] {s} p idx (SHRQconst [56] w) mem))))))))
	// cond: x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)
	// result: (MOVQstoreidx1 [i-7] {s} p idx (BSWAPQ <w.Type> w) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x6 := v.Args[3]
		if x6.Op != OpAMD64MOVBstoreidx1 || x6.AuxInt != i-1 || x6.Aux != s {
			break
		}
		_ = x6.Args[3]
		if p != x6.Args[0] || idx != x6.Args[1] {
			break
		}
		x6_2 := x6.Args[2]
		if x6_2.Op != OpAMD64SHRQconst || x6_2.AuxInt != 8 || w != x6_2.Args[0] {
			break
		}
		x5 := x6.Args[3]
		if x5.Op != OpAMD64MOVBstoreidx1 || x5.AuxInt != i-2 || x5.Aux != s {
			break
		}
		_ = x5.Args[3]
		if p != x5.Args[0] || idx != x5.Args[1] {
			break
		}
		x5_2 := x5.Args[2]
		if x5_2.Op != OpAMD64SHRQconst || x5_2.AuxInt != 16 || w != x5_2.Args[0] {
			break
		}
		x4 := x5.Args[3]
		if x4.Op != OpAMD64MOVBstoreidx1 || x4.AuxInt != i-3 || x4.Aux != s {
			break
		}
		_ = x4.Args[3]
		if p != x4.Args[0] || idx != x4.Args[1] {
			break
		}
		x4_2 := x4.Args[2]
		if x4_2.Op != OpAMD64SHRQconst || x4_2.AuxInt != 24 || w != x4_2.Args[0] {
			break
		}
		x3 := x4.Args[3]
		if x3.Op != OpAMD64MOVBstoreidx1 || x3.AuxInt != i-4 || x3.Aux != s {
			break
		}
		_ = x3.Args[3]
		if p != x3.Args[0] || idx != x3.Args[1] {
			break
		}
		x3_2 := x3.Args[2]
		if x3_2.Op != OpAMD64SHRQconst || x3_2.AuxInt != 32 || w != x3_2.Args[0] {
			break
		}
		x2 := x3.Args[3]
		if x2.Op != OpAMD64MOVBstoreidx1 || x2.AuxInt != i-5 || x2.Aux != s {
			break
		}
		_ = x2.Args[3]
		if p != x2.Args[0] || idx != x2.Args[1] {
			break
		}
		x2_2 := x2.Args[2]
		if x2_2.Op != OpAMD64SHRQconst || x2_2.AuxInt != 40 || w != x2_2.Args[0] {
			break
		}
		x1 := x2.Args[3]
		if x1.Op != OpAMD64MOVBstoreidx1 || x1.AuxInt != i-6 || x1.Aux != s {
			break
		}
		_ = x1.Args[3]
		if p != x1.Args[0] || idx != x1.Args[1] {
			break
		}
		x1_2 := x1.Args[2]
		if x1_2.Op != OpAMD64SHRQconst || x1_2.AuxInt != 48 || w != x1_2.Args[0] {
			break
		}
		x0 := x1.Args[3]
		if x0.Op != OpAMD64MOVBstoreidx1 || x0.AuxInt != i-7 || x0.Aux != s {
			break
		}
		mem := x0.Args[3]
		if p != x0.Args[0] || idx != x0.Args[1] {
			break
		}
		x0_2 := x0.Args[2]
		if x0_2.Op != OpAMD64SHRQconst || x0_2.AuxInt != 56 || w != x0_2.Args[0] || !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = i - 7
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpAMD64BSWAPQ, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [i] {s} p idx (SHRWconst [8] w) x:(MOVBstoreidx1 [i-1] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx1 [i-1] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRWconst || v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVBstoreidx1 || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [i] {s} p idx (SHRLconst [8] w) x:(MOVBstoreidx1 [i-1] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx1 [i-1] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRLconst || v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVBstoreidx1 || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [i] {s} p idx (SHRQconst [8] w) x:(MOVBstoreidx1 [i-1] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx1 [i-1] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst || v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVBstoreidx1 || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [i] {s} p idx (SHRLconst [j] w) x:(MOVBstoreidx1 [i-1] {s} p idx w0:(SHRLconst [j-8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx1 [i-1] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVBstoreidx1 || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpAMD64SHRLconst || w0.AuxInt != j-8 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx1 [i] {s} p idx (SHRQconst [j] w) x:(MOVBstoreidx1 [i-1] {s} p idx w0:(SHRQconst [j-8] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVWstoreidx1 [i-1] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVBstoreidx1 || x.AuxInt != i-1 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpAMD64SHRQconst || w0.AuxInt != j-8 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVBstoreidx1_10(v *Value) bool {
	// match: (MOVBstoreidx1 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+c)
	// result: (MOVBstore [i+c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + c)) {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = i + c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLQSX_0(v *Value) bool {
	b := v.Block
	// match: (MOVLQSX x:(MOVLload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVLQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVLload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVLQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVLQSX x:(MOVQload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVLQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVQload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVLQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVLQSX (ANDLconst [c] x))
	// cond: c & 0x80000000 == 0
	// result: (ANDLconst [c & 0x7fffffff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80000000 == 0) {
			break
		}
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = c & 0x7fffffff
		v.AddArg(x)
		return true
	}
	// match: (MOVLQSX (MOVLQSX x))
	// result: (MOVLQSX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLQSX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVLQSX)
		v.AddArg(x)
		return true
	}
	// match: (MOVLQSX (MOVWQSX x))
	// result: (MOVWQSX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVWQSX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVWQSX)
		v.AddArg(x)
		return true
	}
	// match: (MOVLQSX (MOVBQSX x))
	// result: (MOVBQSX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVBQSX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVBQSX)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLQSXload_0(v *Value) bool {
	// match: (MOVLQSXload [off] {sym} ptr (MOVLstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVLQSX x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpAMD64MOVLQSX)
		v.AddArg(x)
		return true
	}
	// match: (MOVLQSXload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLQSXload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLQSXload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLQZX_0(v *Value) bool {
	b := v.Block
	// match: (MOVLQZX x:(MOVLload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVLload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVLload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVLload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVLQZX x:(MOVQload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVLload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVQload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVLload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVLQZX x)
	// cond: zeroUpper32Bits(x,3)
	// result: x
	for {
		x := v.Args[0]
		if !(zeroUpper32Bits(x, 3)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVLQZX x:(MOVLloadidx1 [off] {sym} ptr idx mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVLloadidx1 <v.Type> [off] {sym} ptr idx mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVLloadidx1 {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVLQZX x:(MOVLloadidx4 [off] {sym} ptr idx mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVLloadidx4 <v.Type> [off] {sym} ptr idx mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVLloadidx4 {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx4, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVLQZX (ANDLconst [c] x))
	// result: (ANDLconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVLQZX (MOVLQZX x))
	// result: (MOVLQZX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLQZX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVLQZX)
		v.AddArg(x)
		return true
	}
	// match: (MOVLQZX (MOVWQZX x))
	// result: (MOVWQZX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVWQZX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVWQZX)
		v.AddArg(x)
		return true
	}
	// match: (MOVLQZX (MOVBQZX x))
	// result: (MOVBQZX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVBQZX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLatomicload_0(v *Value) bool {
	// match: (MOVLatomicload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVLatomicload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVLatomicload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLatomicload [off1] {sym1} (LEAQ [off2] {sym2} ptr) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLatomicload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLatomicload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLf2i_0(v *Value) bool {
	b := v.Block
	// match: (MOVLf2i <t> (Arg <u> [off] {sym}))
	// cond: t.Size() == u.Size()
	// result: @b.Func.Entry (Arg <t> [off] {sym})
	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpArg {
			break
		}
		u := v_0.Type
		off := v_0.AuxInt
		sym := v_0.Aux
		if !(t.Size() == u.Size()) {
			break
		}
		b = b.Func.Entry
		v0 := b.NewValue0(v.Pos, OpArg, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLi2f_0(v *Value) bool {
	b := v.Block
	// match: (MOVLi2f <t> (Arg <u> [off] {sym}))
	// cond: t.Size() == u.Size()
	// result: @b.Func.Entry (Arg <t> [off] {sym})
	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpArg {
			break
		}
		u := v_0.Type
		off := v_0.AuxInt
		sym := v_0.Aux
		if !(t.Size() == u.Size()) {
			break
		}
		b = b.Func.Entry
		v0 := b.NewValue0(v.Pos, OpArg, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLload_0(v *Value) bool {
	// match: (MOVLload [off] {sym} ptr (MOVLstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVLQZX x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpAMD64MOVLQZX)
		v.AddArg(x)
		return true
	}
	// match: (MOVLload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVLload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLload [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLloadidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLload [off1] {sym1} (LEAQ4 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLloadidx4 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ4 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLloadidx4)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLload [off1] {sym1} (LEAQ8 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLloadidx8 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLloadidx8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLload [off] {sym} (ADDQ ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVLloadidx1 [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVLloadidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVLload [off1] {sym1} (LEAL [off2] {sym2} base) mem)
	// cond: canMergeSym(sym1, sym2) && is32Bit(off1+off2)
	// result: (MOVLload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpAMD64MOVLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLload [off1] {sym} (ADDLconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVLload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLload [off] {sym} ptr (MOVSSstore [off] {sym} ptr val _))
	// result: (MOVLf2i val)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVSSstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		if ptr != v_1.Args[0] {
			break
		}
		val := v_1.Args[1]
		v.reset(OpAMD64MOVLf2i)
		v.AddArg(val)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLload_10(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	// match: (MOVLload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVQconst [int64(read32(sym, off, config.BigEndian))])
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = int64(read32(sym, off, config.BigEndian))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLloadidx1_0(v *Value) bool {
	// match: (MOVLloadidx1 [c] {sym} ptr (SHLQconst [2] idx) mem)
	// result: (MOVLloadidx4 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 2 {
				continue
			}
			idx := v_1.Args[0]
			v.reset(OpAMD64MOVLloadidx4)
			v.AuxInt = c
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVLloadidx1 [c] {sym} ptr (SHLQconst [3] idx) mem)
	// result: (MOVLloadidx8 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
				continue
			}
			idx := v_1.Args[0]
			v.reset(OpAMD64MOVLloadidx8)
			v.AuxInt = c
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVLloadidx1 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVLloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v.Args[1^_i0]
			if !(is32Bit(c + d)) {
				continue
			}
			v.reset(OpAMD64MOVLloadidx1)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVLloadidx1 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+d)
	// result: (MOVLloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			if !(is32Bit(c + d)) {
				continue
			}
			v.reset(OpAMD64MOVLloadidx1)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVLloadidx1 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+c)
	// result: (MOVLload [i+c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			p := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(i + c)) {
				continue
			}
			v.reset(OpAMD64MOVLload)
			v.AuxInt = i + c
			v.Aux = s
			v.AddArg(p)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLloadidx4_0(v *Value) bool {
	// match: (MOVLloadidx4 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVLloadidx4 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVLloadidx4)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLloadidx4 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+4*d)
	// result: (MOVLloadidx4 [c+4*d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		if !(is32Bit(c + 4*d)) {
			break
		}
		v.reset(OpAMD64MOVLloadidx4)
		v.AuxInt = c + 4*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLloadidx4 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+4*c)
	// result: (MOVLload [i+4*c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(i + 4*c)) {
			break
		}
		v.reset(OpAMD64MOVLload)
		v.AuxInt = i + 4*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLloadidx8_0(v *Value) bool {
	// match: (MOVLloadidx8 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVLloadidx8 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVLloadidx8)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLloadidx8 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+8*d)
	// result: (MOVLloadidx8 [c+8*d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		if !(is32Bit(c + 8*d)) {
			break
		}
		v.reset(OpAMD64MOVLloadidx8)
		v.AuxInt = c + 8*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLloadidx8 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+8*c)
	// result: (MOVLload [i+8*c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(i + 8*c)) {
			break
		}
		v.reset(OpAMD64MOVLload)
		v.AuxInt = i + 8*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstore_0(v *Value) bool {
	// match: (MOVLstore [off] {sym} ptr (MOVLQSX x) mem)
	// result: (MOVLstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLQSX {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr (MOVLQZX x) mem)
	// result: (MOVLstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLQZX {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off1] {sym} (ADDQconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVLstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr (MOVLconst [c]) mem)
	// cond: validOff(off)
	// result: (MOVLstoreconst [makeValAndOff(int64(int32(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(validOff(off)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = makeValAndOff(int64(int32(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr (MOVQconst [c]) mem)
	// cond: validOff(off)
	// result: (MOVLstoreconst [makeValAndOff(int64(int32(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(validOff(off)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = makeValAndOff(int64(int32(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLstoreidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off1] {sym1} (LEAQ4 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLstoreidx4 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ4 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx4)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off1] {sym1} (LEAQ8 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVLstoreidx8 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} (ADDQ ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVLstoreidx1 [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			val := v.Args[1]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVLstoreidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstore_10(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVLstore [i] {s} p (SHRQconst [32] w) x:(MOVLstore [i-4] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVQstore [i-4] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRQconst || v_1.AuxInt != 32 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVLstore || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [i] {s} p (SHRQconst [j] w) x:(MOVLstore [i-4] {s} p w0:(SHRQconst [j-32] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVQstore [i-4] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRQconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVLstore || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpAMD64SHRQconst || w0.AuxInt != j-32 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [i] {s} p x1:(MOVLload [j] {s2} p2 mem) mem2:(MOVLstore [i-4] {s} p x2:(MOVLload [j-4] {s2} p2 mem) mem))
	// cond: x1.Uses == 1 && x2.Uses == 1 && mem2.Uses == 1 && clobber(x1) && clobber(x2) && clobber(mem2)
	// result: (MOVQstore [i-4] {s} p (MOVQload [j-4] {s2} p2 mem) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpAMD64MOVLload {
			break
		}
		j := x1.AuxInt
		s2 := x1.Aux
		mem := x1.Args[1]
		p2 := x1.Args[0]
		mem2 := v.Args[2]
		if mem2.Op != OpAMD64MOVLstore || mem2.AuxInt != i-4 || mem2.Aux != s {
			break
		}
		_ = mem2.Args[2]
		if p != mem2.Args[0] {
			break
		}
		x2 := mem2.Args[1]
		if x2.Op != OpAMD64MOVLload || x2.AuxInt != j-4 || x2.Aux != s2 {
			break
		}
		_ = x2.Args[1]
		if p2 != x2.Args[0] || mem != x2.Args[1] || mem != mem2.Args[2] || !(x1.Uses == 1 && x2.Uses == 1 && mem2.Uses == 1 && clobber(x1) && clobber(x2) && clobber(mem2)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x2.Pos, OpAMD64MOVQload, typ.UInt64)
		v0.AuxInt = j - 4
		v0.Aux = s2
		v0.AddArg(p2)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off1] {sym1} (LEAL [off2] {sym2} base) val mem)
	// cond: canMergeSym(sym1, sym2) && is32Bit(off1+off2)
	// result: (MOVLstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off1] {sym} (ADDLconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVLstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore {sym} [off] ptr y:(ADDLload x [off] {sym} ptr mem) mem)
	// cond: y.Uses==1 && clobber(y)
	// result: (ADDLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ADDLload || y.AuxInt != off || y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] || mem != y.Args[2] || !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(OpAMD64ADDLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore {sym} [off] ptr y:(ANDLload x [off] {sym} ptr mem) mem)
	// cond: y.Uses==1 && clobber(y)
	// result: (ANDLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ANDLload || y.AuxInt != off || y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] || mem != y.Args[2] || !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(OpAMD64ANDLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore {sym} [off] ptr y:(ORLload x [off] {sym} ptr mem) mem)
	// cond: y.Uses==1 && clobber(y)
	// result: (ORLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ORLload || y.AuxInt != off || y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] || mem != y.Args[2] || !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(OpAMD64ORLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore {sym} [off] ptr y:(XORLload x [off] {sym} ptr mem) mem)
	// cond: y.Uses==1 && clobber(y)
	// result: (XORLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64XORLload || y.AuxInt != off || y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] || mem != y.Args[2] || !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(OpAMD64XORLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore {sym} [off] ptr y:(ADDL l:(MOVLload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (ADDLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ADDL {
			break
		}
		_ = y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := y.Args[_i0]
			if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
				continue
			}
			_ = l.Args[1]
			if ptr != l.Args[0] || mem != l.Args[1] {
				continue
			}
			x := y.Args[1^_i0]
			if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ADDLmodify)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(x)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstore_20(v *Value) bool {
	// match: (MOVLstore {sym} [off] ptr y:(SUBL l:(MOVLload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (SUBLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SUBL {
			break
		}
		x := y.Args[1]
		l := y.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] || mem != l.Args[1] || !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(OpAMD64SUBLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore {sym} [off] ptr y:(ANDL l:(MOVLload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (ANDLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ANDL {
			break
		}
		_ = y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := y.Args[_i0]
			if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
				continue
			}
			_ = l.Args[1]
			if ptr != l.Args[0] || mem != l.Args[1] {
				continue
			}
			x := y.Args[1^_i0]
			if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ANDLmodify)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(x)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVLstore {sym} [off] ptr y:(ORL l:(MOVLload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (ORLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ORL {
			break
		}
		_ = y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := y.Args[_i0]
			if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
				continue
			}
			_ = l.Args[1]
			if ptr != l.Args[0] || mem != l.Args[1] {
				continue
			}
			x := y.Args[1^_i0]
			if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ORLmodify)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(x)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVLstore {sym} [off] ptr y:(XORL l:(MOVLload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (XORLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64XORL {
			break
		}
		_ = y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := y.Args[_i0]
			if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
				continue
			}
			_ = l.Args[1]
			if ptr != l.Args[0] || mem != l.Args[1] {
				continue
			}
			x := y.Args[1^_i0]
			if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64XORLmodify)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(x)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVLstore {sym} [off] ptr y:(BTCL l:(MOVLload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (BTCLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64BTCL {
			break
		}
		x := y.Args[1]
		l := y.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] || mem != l.Args[1] || !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(OpAMD64BTCLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore {sym} [off] ptr y:(BTRL l:(MOVLload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (BTRLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64BTRL {
			break
		}
		x := y.Args[1]
		l := y.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] || mem != l.Args[1] || !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(OpAMD64BTRLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore {sym} [off] ptr y:(BTSL l:(MOVLload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (BTSLmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64BTSL {
			break
		}
		x := y.Args[1]
		l := y.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] || mem != l.Args[1] || !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(OpAMD64BTSLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr a:(ADDLconst [c] l:(MOVLload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (ADDLconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64ADDLconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64ADDLconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr a:(ANDLconst [c] l:(MOVLload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (ANDLconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64ANDLconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64ANDLconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr a:(ORLconst [c] l:(MOVLload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (ORLconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64ORLconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64ORLconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstore_30(v *Value) bool {
	// match: (MOVLstore [off] {sym} ptr a:(XORLconst [c] l:(MOVLload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (XORLconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64XORLconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64XORLconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr a:(BTCLconst [c] l:(MOVLload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (BTCLconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64BTCLconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64BTCLconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr a:(BTRLconst [c] l:(MOVLload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (BTRLconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64BTRLconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64BTRLconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr a:(BTSLconst [c] l:(MOVLload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (BTSLconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64BTSLconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVLload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64BTSLconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstore [off] {sym} ptr (MOVLf2i val) mem)
	// result: (MOVSSstore [off] {sym} ptr val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLf2i {
			break
		}
		val := v_1.Args[0]
		v.reset(OpAMD64MOVSSstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstoreconst_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVLstoreconst [sc] {s} (ADDQconst [off] ptr) mem)
	// cond: ValAndOff(sc).canAdd(off)
	// result: (MOVLstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconst [sc] {sym1} (LEAQ [off] {sym2} ptr) mem)
	// cond: canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVLstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconst [x] {sym1} (LEAQ1 [off] {sym2} ptr idx) mem)
	// cond: canMergeSym(sym1, sym2)
	// result: (MOVLstoreconstidx1 [ValAndOff(x).add(off)] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		x := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconst [x] {sym1} (LEAQ4 [off] {sym2} ptr idx) mem)
	// cond: canMergeSym(sym1, sym2)
	// result: (MOVLstoreconstidx4 [ValAndOff(x).add(off)] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		x := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ4 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconstidx4)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconst [x] {sym} (ADDQ ptr idx) mem)
	// result: (MOVLstoreconstidx1 [x] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		v.reset(OpAMD64MOVLstoreconstidx1)
		v.AuxInt = x
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconst [c] {s} p x:(MOVLstoreconst [a] {s} p mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 4 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVQstore [ValAndOff(a).Off()] {s} p (MOVQconst [ValAndOff(a).Val()&0xffffffff | ValAndOff(c).Val()<<32]) mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpAMD64MOVLstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(x.Uses == 1 && ValAndOff(a).Off()+4 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = ValAndOff(a).Off()
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x.Pos, OpAMD64MOVQconst, typ.UInt64)
		v0.AuxInt = ValAndOff(a).Val()&0xffffffff | ValAndOff(c).Val()<<32
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconst [a] {s} p x:(MOVLstoreconst [c] {s} p mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 4 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVQstore [ValAndOff(a).Off()] {s} p (MOVQconst [ValAndOff(a).Val()&0xffffffff | ValAndOff(c).Val()<<32]) mem)
	for {
		a := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpAMD64MOVLstoreconst {
			break
		}
		c := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(x.Uses == 1 && ValAndOff(a).Off()+4 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = ValAndOff(a).Off()
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x.Pos, OpAMD64MOVQconst, typ.UInt64)
		v0.AuxInt = ValAndOff(a).Val()&0xffffffff | ValAndOff(c).Val()<<32
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconst [sc] {sym1} (LEAL [off] {sym2} ptr) mem)
	// cond: canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVLstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconst [sc] {s} (ADDLconst [off] ptr) mem)
	// cond: ValAndOff(sc).canAdd(off)
	// result: (MOVLstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstoreconstidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVLstoreconstidx1 [c] {sym} ptr (SHLQconst [2] idx) mem)
	// result: (MOVLstoreconstidx4 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		v.reset(OpAMD64MOVLstoreconstidx4)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconstidx1 [x] {sym} (ADDQconst [c] ptr) idx mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVLstoreconstidx1 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconstidx1 [x] {sym} ptr (ADDQconst [c] idx) mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVLstoreconstidx1 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconstidx1 [c] {s} p i x:(MOVLstoreconstidx1 [a] {s} p i mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 4 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVQstoreidx1 [ValAndOff(a).Off()] {s} p i (MOVQconst [ValAndOff(a).Val()&0xffffffff | ValAndOff(c).Val()<<32]) mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		i := v.Args[1]
		x := v.Args[2]
		if x.Op != OpAMD64MOVLstoreconstidx1 {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || i != x.Args[1] || !(x.Uses == 1 && ValAndOff(a).Off()+4 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = ValAndOff(a).Off()
		v.Aux = s
		v.AddArg(p)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQconst, typ.UInt64)
		v0.AuxInt = ValAndOff(a).Val()&0xffffffff | ValAndOff(c).Val()<<32
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstoreconstidx4_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVLstoreconstidx4 [x] {sym} (ADDQconst [c] ptr) idx mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVLstoreconstidx4 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconstidx4)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconstidx4 [x] {sym} ptr (ADDQconst [c] idx) mem)
	// cond: ValAndOff(x).canAdd(4*c)
	// result: (MOVLstoreconstidx4 [ValAndOff(x).add(4*c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(ValAndOff(x).canAdd(4 * c)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconstidx4)
		v.AuxInt = ValAndOff(x).add(4 * c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreconstidx4 [c] {s} p i x:(MOVLstoreconstidx4 [a] {s} p i mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 4 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVQstoreidx1 [ValAndOff(a).Off()] {s} p (SHLQconst <i.Type> [2] i) (MOVQconst [ValAndOff(a).Val()&0xffffffff | ValAndOff(c).Val()<<32]) mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		i := v.Args[1]
		x := v.Args[2]
		if x.Op != OpAMD64MOVLstoreconstidx4 {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || i != x.Args[1] || !(x.Uses == 1 && ValAndOff(a).Off()+4 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = ValAndOff(a).Off()
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, i.Type)
		v0.AuxInt = 2
		v0.AddArg(i)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQconst, typ.UInt64)
		v1.AuxInt = ValAndOff(a).Val()&0xffffffff | ValAndOff(c).Val()<<32
		v.AddArg(v1)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstoreidx1_0(v *Value) bool {
	// match: (MOVLstoreidx1 [c] {sym} ptr (SHLQconst [2] idx) val mem)
	// result: (MOVLstoreidx4 [c] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		v.reset(OpAMD64MOVLstoreidx4)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx1 [c] {sym} ptr (SHLQconst [3] idx) val mem)
	// result: (MOVLstoreidx8 [c] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		v.reset(OpAMD64MOVLstoreidx8)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx1 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVLstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx1 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+d)
	// result: (MOVLstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx1 [i] {s} p idx (SHRQconst [32] w) x:(MOVLstoreidx1 [i-4] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVQstoreidx1 [i-4] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst || v_2.AuxInt != 32 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVLstoreidx1 || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx1 [i] {s} p idx (SHRQconst [j] w) x:(MOVLstoreidx1 [i-4] {s} p idx w0:(SHRQconst [j-32] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVQstoreidx1 [i-4] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVLstoreidx1 || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpAMD64SHRQconst || w0.AuxInt != j-32 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx1 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+c)
	// result: (MOVLstore [i+c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + c)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i + c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstoreidx4_0(v *Value) bool {
	b := v.Block
	// match: (MOVLstoreidx4 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVLstoreidx4 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx4)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx4 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+4*d)
	// result: (MOVLstoreidx4 [c+4*d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + 4*d)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx4)
		v.AuxInt = c + 4*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx4 [i] {s} p idx (SHRQconst [32] w) x:(MOVLstoreidx4 [i-4] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVQstoreidx1 [i-4] {s} p (SHLQconst <idx.Type> [2] idx) w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst || v_2.AuxInt != 32 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVLstoreidx4 || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, idx.Type)
		v0.AuxInt = 2
		v0.AddArg(idx)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx4 [i] {s} p idx (SHRQconst [j] w) x:(MOVLstoreidx4 [i-4] {s} p idx w0:(SHRQconst [j-32] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVQstoreidx1 [i-4] {s} p (SHLQconst <idx.Type> [2] idx) w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVLstoreidx4 || x.AuxInt != i-4 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpAMD64SHRQconst || w0.AuxInt != j-32 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, idx.Type)
		v0.AuxInt = 2
		v0.AddArg(idx)
		v.AddArg(v0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx4 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+4*c)
	// result: (MOVLstore [i+4*c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + 4*c)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i + 4*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVLstoreidx8_0(v *Value) bool {
	// match: (MOVLstoreidx8 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVLstoreidx8 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx8)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx8 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+8*d)
	// result: (MOVLstoreidx8 [c+8*d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + 8*d)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx8)
		v.AuxInt = c + 8*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVLstoreidx8 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+8*c)
	// result: (MOVLstore [i+8*c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + 8*c)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i + 8*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVOload_0(v *Value) bool {
	// match: (MOVOload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVOload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVOload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVOload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVOload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVOload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVOstore_0(v *Value) bool {
	// match: (MOVOstore [off1] {sym} (ADDQconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVOstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVOstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVOstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVOstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVOstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQatomicload_0(v *Value) bool {
	// match: (MOVQatomicload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVQatomicload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVQatomicload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQatomicload [off1] {sym1} (LEAQ [off2] {sym2} ptr) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVQatomicload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQatomicload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQf2i_0(v *Value) bool {
	b := v.Block
	// match: (MOVQf2i <t> (Arg <u> [off] {sym}))
	// cond: t.Size() == u.Size()
	// result: @b.Func.Entry (Arg <t> [off] {sym})
	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpArg {
			break
		}
		u := v_0.Type
		off := v_0.AuxInt
		sym := v_0.Aux
		if !(t.Size() == u.Size()) {
			break
		}
		b = b.Func.Entry
		v0 := b.NewValue0(v.Pos, OpArg, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQi2f_0(v *Value) bool {
	b := v.Block
	// match: (MOVQi2f <t> (Arg <u> [off] {sym}))
	// cond: t.Size() == u.Size()
	// result: @b.Func.Entry (Arg <t> [off] {sym})
	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpArg {
			break
		}
		u := v_0.Type
		off := v_0.AuxInt
		sym := v_0.Aux
		if !(t.Size() == u.Size()) {
			break
		}
		b = b.Func.Entry
		v0 := b.NewValue0(v.Pos, OpArg, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQload_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	// match: (MOVQload [off] {sym} ptr (MOVQstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVQload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVQload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVQload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVQload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQload [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVQloadidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQload [off1] {sym1} (LEAQ8 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVQloadidx8 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQloadidx8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQload [off] {sym} (ADDQ ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVQloadidx1 [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVQloadidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQload [off1] {sym1} (LEAL [off2] {sym2} base) mem)
	// cond: canMergeSym(sym1, sym2) && is32Bit(off1+off2)
	// result: (MOVQload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpAMD64MOVQload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQload [off1] {sym} (ADDLconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVQload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVQload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQload [off] {sym} ptr (MOVSDstore [off] {sym} ptr val _))
	// result: (MOVQf2i val)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVSDstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		if ptr != v_1.Args[0] {
			break
		}
		val := v_1.Args[1]
		v.reset(OpAMD64MOVQf2i)
		v.AddArg(val)
		return true
	}
	// match: (MOVQload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVQconst [int64(read64(sym, off, config.BigEndian))])
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = int64(read64(sym, off, config.BigEndian))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQloadidx1_0(v *Value) bool {
	// match: (MOVQloadidx1 [c] {sym} ptr (SHLQconst [3] idx) mem)
	// result: (MOVQloadidx8 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
				continue
			}
			idx := v_1.Args[0]
			v.reset(OpAMD64MOVQloadidx8)
			v.AuxInt = c
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQloadidx1 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVQloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v.Args[1^_i0]
			if !(is32Bit(c + d)) {
				continue
			}
			v.reset(OpAMD64MOVQloadidx1)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQloadidx1 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+d)
	// result: (MOVQloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			if !(is32Bit(c + d)) {
				continue
			}
			v.reset(OpAMD64MOVQloadidx1)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQloadidx1 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+c)
	// result: (MOVQload [i+c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			p := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(i + c)) {
				continue
			}
			v.reset(OpAMD64MOVQload)
			v.AuxInt = i + c
			v.Aux = s
			v.AddArg(p)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQloadidx8_0(v *Value) bool {
	// match: (MOVQloadidx8 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVQloadidx8 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVQloadidx8)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQloadidx8 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+8*d)
	// result: (MOVQloadidx8 [c+8*d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		if !(is32Bit(c + 8*d)) {
			break
		}
		v.reset(OpAMD64MOVQloadidx8)
		v.AuxInt = c + 8*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQloadidx8 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+8*c)
	// result: (MOVQload [i+8*c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(i + 8*c)) {
			break
		}
		v.reset(OpAMD64MOVQload)
		v.AuxInt = i + 8*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQstore_0(v *Value) bool {
	// match: (MOVQstore [off1] {sym} (ADDQconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVQstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} ptr (MOVQconst [c]) mem)
	// cond: validValAndOff(c,off)
	// result: (MOVQstoreconst [makeValAndOff(c,off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(validValAndOff(c, off)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVQstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVQstoreidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off1] {sym1} (LEAQ8 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVQstoreidx8 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} (ADDQ ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVQstoreidx1 [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			val := v.Args[1]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVQstoreidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQstore [off1] {sym1} (LEAL [off2] {sym2} base) val mem)
	// cond: canMergeSym(sym1, sym2) && is32Bit(off1+off2)
	// result: (MOVQstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off1] {sym} (ADDLconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVQstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore {sym} [off] ptr y:(ADDQload x [off] {sym} ptr mem) mem)
	// cond: y.Uses==1 && clobber(y)
	// result: (ADDQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ADDQload || y.AuxInt != off || y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] || mem != y.Args[2] || !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(OpAMD64ADDQmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore {sym} [off] ptr y:(ANDQload x [off] {sym} ptr mem) mem)
	// cond: y.Uses==1 && clobber(y)
	// result: (ANDQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ANDQload || y.AuxInt != off || y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] || mem != y.Args[2] || !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(OpAMD64ANDQmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQstore_10(v *Value) bool {
	// match: (MOVQstore {sym} [off] ptr y:(ORQload x [off] {sym} ptr mem) mem)
	// cond: y.Uses==1 && clobber(y)
	// result: (ORQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ORQload || y.AuxInt != off || y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] || mem != y.Args[2] || !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(OpAMD64ORQmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore {sym} [off] ptr y:(XORQload x [off] {sym} ptr mem) mem)
	// cond: y.Uses==1 && clobber(y)
	// result: (XORQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64XORQload || y.AuxInt != off || y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] || mem != y.Args[2] || !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(OpAMD64XORQmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore {sym} [off] ptr y:(ADDQ l:(MOVQload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (ADDQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ADDQ {
			break
		}
		_ = y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := y.Args[_i0]
			if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
				continue
			}
			_ = l.Args[1]
			if ptr != l.Args[0] || mem != l.Args[1] {
				continue
			}
			x := y.Args[1^_i0]
			if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ADDQmodify)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(x)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQstore {sym} [off] ptr y:(SUBQ l:(MOVQload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (SUBQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64SUBQ {
			break
		}
		x := y.Args[1]
		l := y.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] || mem != l.Args[1] || !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(OpAMD64SUBQmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore {sym} [off] ptr y:(ANDQ l:(MOVQload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (ANDQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ANDQ {
			break
		}
		_ = y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := y.Args[_i0]
			if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
				continue
			}
			_ = l.Args[1]
			if ptr != l.Args[0] || mem != l.Args[1] {
				continue
			}
			x := y.Args[1^_i0]
			if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ANDQmodify)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(x)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQstore {sym} [off] ptr y:(ORQ l:(MOVQload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (ORQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64ORQ {
			break
		}
		_ = y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := y.Args[_i0]
			if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
				continue
			}
			_ = l.Args[1]
			if ptr != l.Args[0] || mem != l.Args[1] {
				continue
			}
			x := y.Args[1^_i0]
			if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ORQmodify)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(x)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQstore {sym} [off] ptr y:(XORQ l:(MOVQload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (XORQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64XORQ {
			break
		}
		_ = y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := y.Args[_i0]
			if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
				continue
			}
			_ = l.Args[1]
			if ptr != l.Args[0] || mem != l.Args[1] {
				continue
			}
			x := y.Args[1^_i0]
			if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64XORQmodify)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(x)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVQstore {sym} [off] ptr y:(BTCQ l:(MOVQload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (BTCQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64BTCQ {
			break
		}
		x := y.Args[1]
		l := y.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] || mem != l.Args[1] || !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(OpAMD64BTCQmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore {sym} [off] ptr y:(BTRQ l:(MOVQload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (BTRQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64BTRQ {
			break
		}
		x := y.Args[1]
		l := y.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] || mem != l.Args[1] || !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(OpAMD64BTRQmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore {sym} [off] ptr y:(BTSQ l:(MOVQload [off] {sym} ptr mem) x) mem)
	// cond: y.Uses==1 && l.Uses==1 && clobber(y) && clobber(l)
	// result: (BTSQmodify [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != OpAMD64BTSQ {
			break
		}
		x := y.Args[1]
		l := y.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] || mem != l.Args[1] || !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(OpAMD64BTSQmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQstore_20(v *Value) bool {
	// match: (MOVQstore [off] {sym} ptr a:(ADDQconst [c] l:(MOVQload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (ADDQconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64ADDQconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64ADDQconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} ptr a:(ANDQconst [c] l:(MOVQload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (ANDQconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64ANDQconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64ANDQconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} ptr a:(ORQconst [c] l:(MOVQload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (ORQconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64ORQconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64ORQconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} ptr a:(XORQconst [c] l:(MOVQload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (XORQconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64XORQconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64XORQconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} ptr a:(BTCQconst [c] l:(MOVQload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (BTCQconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64BTCQconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64BTCQconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} ptr a:(BTRQconst [c] l:(MOVQload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (BTRQconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64BTRQconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64BTRQconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} ptr a:(BTSQconst [c] l:(MOVQload [off] {sym} ptr2 mem)) mem)
	// cond: isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c,off) && clobber(l) && clobber(a)
	// result: (BTSQconstmodify {sym} [makeValAndOff(c,off)] ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		a := v.Args[1]
		if a.Op != OpAMD64BTSQconst {
			break
		}
		c := a.AuxInt
		l := a.Args[0]
		if l.Op != OpAMD64MOVQload || l.AuxInt != off || l.Aux != sym {
			break
		}
		_ = l.Args[1]
		ptr2 := l.Args[0]
		if mem != l.Args[1] || !(isSamePtr(ptr, ptr2) && a.Uses == 1 && l.Uses == 1 && validValAndOff(c, off) && clobber(l) && clobber(a)) {
			break
		}
		v.reset(OpAMD64BTSQconstmodify)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstore [off] {sym} ptr (MOVQf2i val) mem)
	// result: (MOVSDstore [off] {sym} ptr val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQf2i {
			break
		}
		val := v_1.Args[0]
		v.reset(OpAMD64MOVSDstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQstoreconst_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	// match: (MOVQstoreconst [sc] {s} (ADDQconst [off] ptr) mem)
	// cond: ValAndOff(sc).canAdd(off)
	// result: (MOVQstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconst [sc] {sym1} (LEAQ [off] {sym2} ptr) mem)
	// cond: canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVQstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconst [x] {sym1} (LEAQ1 [off] {sym2} ptr idx) mem)
	// cond: canMergeSym(sym1, sym2)
	// result: (MOVQstoreconstidx1 [ValAndOff(x).add(off)] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		x := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconst [x] {sym1} (LEAQ8 [off] {sym2} ptr idx) mem)
	// cond: canMergeSym(sym1, sym2)
	// result: (MOVQstoreconstidx8 [ValAndOff(x).add(off)] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		x := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconstidx8)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconst [x] {sym} (ADDQ ptr idx) mem)
	// result: (MOVQstoreconstidx1 [x] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		v.reset(OpAMD64MOVQstoreconstidx1)
		v.AuxInt = x
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconst [c] {s} p x:(MOVQstoreconst [c2] {s} p mem))
	// cond: config.useSSE && x.Uses == 1 && ValAndOff(c2).Off() + 8 == ValAndOff(c).Off() && ValAndOff(c).Val() == 0 && ValAndOff(c2).Val() == 0 && clobber(x)
	// result: (MOVOstore [ValAndOff(c2).Off()] {s} p (MOVOconst [0]) mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpAMD64MOVQstoreconst {
			break
		}
		c2 := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(config.useSSE && x.Uses == 1 && ValAndOff(c2).Off()+8 == ValAndOff(c).Off() && ValAndOff(c).Val() == 0 && ValAndOff(c2).Val() == 0 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVOstore)
		v.AuxInt = ValAndOff(c2).Off()
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconst [sc] {sym1} (LEAL [off] {sym2} ptr) mem)
	// cond: canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVQstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconst [sc] {s} (ADDLconst [off] ptr) mem)
	// cond: ValAndOff(sc).canAdd(off)
	// result: (MOVQstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQstoreconstidx1_0(v *Value) bool {
	// match: (MOVQstoreconstidx1 [c] {sym} ptr (SHLQconst [3] idx) mem)
	// result: (MOVQstoreconstidx8 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		v.reset(OpAMD64MOVQstoreconstidx8)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconstidx1 [x] {sym} (ADDQconst [c] ptr) idx mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVQstoreconstidx1 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconstidx1 [x] {sym} ptr (ADDQconst [c] idx) mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVQstoreconstidx1 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQstoreconstidx8_0(v *Value) bool {
	// match: (MOVQstoreconstidx8 [x] {sym} (ADDQconst [c] ptr) idx mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVQstoreconstidx8 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconstidx8)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreconstidx8 [x] {sym} ptr (ADDQconst [c] idx) mem)
	// cond: ValAndOff(x).canAdd(8*c)
	// result: (MOVQstoreconstidx8 [ValAndOff(x).add(8*c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(ValAndOff(x).canAdd(8 * c)) {
			break
		}
		v.reset(OpAMD64MOVQstoreconstidx8)
		v.AuxInt = ValAndOff(x).add(8 * c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQstoreidx1_0(v *Value) bool {
	// match: (MOVQstoreidx1 [c] {sym} ptr (SHLQconst [3] idx) val mem)
	// result: (MOVQstoreidx8 [c] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		v.reset(OpAMD64MOVQstoreidx8)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreidx1 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVQstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreidx1 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+d)
	// result: (MOVQstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreidx1 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+c)
	// result: (MOVQstore [i+c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + c)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = i + c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVQstoreidx8_0(v *Value) bool {
	// match: (MOVQstoreidx8 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVQstoreidx8 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx8)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreidx8 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+8*d)
	// result: (MOVQstoreidx8 [c+8*d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + 8*d)) {
			break
		}
		v.reset(OpAMD64MOVQstoreidx8)
		v.AuxInt = c + 8*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVQstoreidx8 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+8*c)
	// result: (MOVQstore [i+8*c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + 8*c)) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = i + 8*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSDload_0(v *Value) bool {
	// match: (MOVSDload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVSDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVSDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSDload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDload [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSDloadidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSDloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDload [off1] {sym1} (LEAQ8 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSDloadidx8 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSDloadidx8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDload [off] {sym} (ADDQ ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVSDloadidx1 [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVSDloadidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVSDload [off] {sym} ptr (MOVQstore [off] {sym} ptr val _))
	// result: (MOVQi2f val)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		if ptr != v_1.Args[0] {
			break
		}
		val := v_1.Args[1]
		v.reset(OpAMD64MOVQi2f)
		v.AddArg(val)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSDloadidx1_0(v *Value) bool {
	// match: (MOVSDloadidx1 [c] {sym} ptr (SHLQconst [3] idx) mem)
	// result: (MOVSDloadidx8 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		v.reset(OpAMD64MOVSDloadidx8)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDloadidx1 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVSDloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSDloadidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDloadidx1 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+d)
	// result: (MOVSDloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSDloadidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDloadidx1 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+c)
	// result: (MOVSDload [i+c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(i + c)) {
			break
		}
		v.reset(OpAMD64MOVSDload)
		v.AuxInt = i + c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSDloadidx8_0(v *Value) bool {
	// match: (MOVSDloadidx8 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVSDloadidx8 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSDloadidx8)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDloadidx8 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+8*d)
	// result: (MOVSDloadidx8 [c+8*d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		if !(is32Bit(c + 8*d)) {
			break
		}
		v.reset(OpAMD64MOVSDloadidx8)
		v.AuxInt = c + 8*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDloadidx8 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+8*c)
	// result: (MOVSDload [i+8*c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(i + 8*c)) {
			break
		}
		v.reset(OpAMD64MOVSDload)
		v.AuxInt = i + 8*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSDstore_0(v *Value) bool {
	// match: (MOVSDstore [off1] {sym} (ADDQconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVSDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVSDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSDstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstore [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSDstoreidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSDstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstore [off1] {sym1} (LEAQ8 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSDstoreidx8 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ8 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSDstoreidx8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstore [off] {sym} (ADDQ ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVSDstoreidx1 [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			val := v.Args[1]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVSDstoreidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVSDstore [off] {sym} ptr (MOVQi2f val) mem)
	// result: (MOVQstore [off] {sym} ptr val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQi2f {
			break
		}
		val := v_1.Args[0]
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSDstoreidx1_0(v *Value) bool {
	// match: (MOVSDstoreidx1 [c] {sym} ptr (SHLQconst [3] idx) val mem)
	// result: (MOVSDstoreidx8 [c] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		v.reset(OpAMD64MOVSDstoreidx8)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstoreidx1 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVSDstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSDstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstoreidx1 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+d)
	// result: (MOVSDstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSDstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstoreidx1 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+c)
	// result: (MOVSDstore [i+c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + c)) {
			break
		}
		v.reset(OpAMD64MOVSDstore)
		v.AuxInt = i + c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSDstoreidx8_0(v *Value) bool {
	// match: (MOVSDstoreidx8 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVSDstoreidx8 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSDstoreidx8)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstoreidx8 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+8*d)
	// result: (MOVSDstoreidx8 [c+8*d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + 8*d)) {
			break
		}
		v.reset(OpAMD64MOVSDstoreidx8)
		v.AuxInt = c + 8*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSDstoreidx8 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+8*c)
	// result: (MOVSDstore [i+8*c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + 8*c)) {
			break
		}
		v.reset(OpAMD64MOVSDstore)
		v.AuxInt = i + 8*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSSload_0(v *Value) bool {
	// match: (MOVSSload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVSSload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVSSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSSload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSload [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSSloadidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSSloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSload [off1] {sym1} (LEAQ4 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSSloadidx4 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ4 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSSloadidx4)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSload [off] {sym} (ADDQ ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVSSloadidx1 [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVSSloadidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVSSload [off] {sym} ptr (MOVLstore [off] {sym} ptr val _))
	// result: (MOVLi2f val)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		if ptr != v_1.Args[0] {
			break
		}
		val := v_1.Args[1]
		v.reset(OpAMD64MOVLi2f)
		v.AddArg(val)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSSloadidx1_0(v *Value) bool {
	// match: (MOVSSloadidx1 [c] {sym} ptr (SHLQconst [2] idx) mem)
	// result: (MOVSSloadidx4 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		v.reset(OpAMD64MOVSSloadidx4)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSloadidx1 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVSSloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSSloadidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSloadidx1 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+d)
	// result: (MOVSSloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSSloadidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSloadidx1 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+c)
	// result: (MOVSSload [i+c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(i + c)) {
			break
		}
		v.reset(OpAMD64MOVSSload)
		v.AuxInt = i + c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSSloadidx4_0(v *Value) bool {
	// match: (MOVSSloadidx4 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVSSloadidx4 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSSloadidx4)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSloadidx4 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+4*d)
	// result: (MOVSSloadidx4 [c+4*d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		if !(is32Bit(c + 4*d)) {
			break
		}
		v.reset(OpAMD64MOVSSloadidx4)
		v.AuxInt = c + 4*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSloadidx4 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+4*c)
	// result: (MOVSSload [i+4*c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(i + 4*c)) {
			break
		}
		v.reset(OpAMD64MOVSSload)
		v.AuxInt = i + 4*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSSstore_0(v *Value) bool {
	// match: (MOVSSstore [off1] {sym} (ADDQconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVSSstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVSSstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSSstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSSstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstore [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSSstoreidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSSstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstore [off1] {sym1} (LEAQ4 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVSSstoreidx4 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ4 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVSSstoreidx4)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstore [off] {sym} (ADDQ ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVSSstoreidx1 [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			val := v.Args[1]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVSSstoreidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVSSstore [off] {sym} ptr (MOVLi2f val) mem)
	// result: (MOVLstore [off] {sym} ptr val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLi2f {
			break
		}
		val := v_1.Args[0]
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSSstoreidx1_0(v *Value) bool {
	// match: (MOVSSstoreidx1 [c] {sym} ptr (SHLQconst [2] idx) val mem)
	// result: (MOVSSstoreidx4 [c] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		v.reset(OpAMD64MOVSSstoreidx4)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstoreidx1 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVSSstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSSstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstoreidx1 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+d)
	// result: (MOVSSstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSSstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstoreidx1 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+c)
	// result: (MOVSSstore [i+c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + c)) {
			break
		}
		v.reset(OpAMD64MOVSSstore)
		v.AuxInt = i + c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVSSstoreidx4_0(v *Value) bool {
	// match: (MOVSSstoreidx4 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVSSstoreidx4 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVSSstoreidx4)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstoreidx4 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+4*d)
	// result: (MOVSSstoreidx4 [c+4*d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + 4*d)) {
			break
		}
		v.reset(OpAMD64MOVSSstoreidx4)
		v.AuxInt = c + 4*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVSSstoreidx4 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+4*c)
	// result: (MOVSSstore [i+4*c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + 4*c)) {
			break
		}
		v.reset(OpAMD64MOVSSstore)
		v.AuxInt = i + 4*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWQSX_0(v *Value) bool {
	b := v.Block
	// match: (MOVWQSX x:(MOVWload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVWload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVWQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWQSX x:(MOVLload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVLload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVWQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWQSX x:(MOVQload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWQSXload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVQload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVWQSXload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWQSX (ANDLconst [c] x))
	// cond: c & 0x8000 == 0
	// result: (ANDLconst [c & 0x7fff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = c & 0x7fff
		v.AddArg(x)
		return true
	}
	// match: (MOVWQSX (MOVWQSX x))
	// result: (MOVWQSX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVWQSX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVWQSX)
		v.AddArg(x)
		return true
	}
	// match: (MOVWQSX (MOVBQSX x))
	// result: (MOVBQSX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVBQSX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVBQSX)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWQSXload_0(v *Value) bool {
	// match: (MOVWQSXload [off] {sym} ptr (MOVWstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVWQSX x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVWstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpAMD64MOVWQSX)
		v.AddArg(x)
		return true
	}
	// match: (MOVWQSXload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWQSXload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWQSXload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWQZX_0(v *Value) bool {
	b := v.Block
	// match: (MOVWQZX x:(MOVWload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVWload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVWload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWQZX x:(MOVLload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVLload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVWload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWQZX x:(MOVQload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWload <v.Type> [off] {sym} ptr mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVQload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpAMD64MOVWload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWQZX x)
	// cond: zeroUpper48Bits(x,3)
	// result: x
	for {
		x := v.Args[0]
		if !(zeroUpper48Bits(x, 3)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWQZX x:(MOVWloadidx1 [off] {sym} ptr idx mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWloadidx1 <v.Type> [off] {sym} ptr idx mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVWloadidx1 {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWQZX x:(MOVWloadidx2 [off] {sym} ptr idx mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVWloadidx2 <v.Type> [off] {sym} ptr idx mem)
	for {
		x := v.Args[0]
		if x.Op != OpAMD64MOVWloadidx2 {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx2, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVWQZX (ANDLconst [c] x))
	// result: (ANDLconst [c & 0xffff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ANDLconst)
		v.AuxInt = c & 0xffff
		v.AddArg(x)
		return true
	}
	// match: (MOVWQZX (MOVWQZX x))
	// result: (MOVWQZX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVWQZX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVWQZX)
		v.AddArg(x)
		return true
	}
	// match: (MOVWQZX (MOVBQZX x))
	// result: (MOVBQZX x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVBQZX {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWload_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	// match: (MOVWload [off] {sym} ptr (MOVWstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVWQZX x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVWstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpAMD64MOVWQZX)
		v.AddArg(x)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWloadidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (LEAQ2 [off2] {sym2} ptr idx) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWloadidx2 [off1+off2] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ2 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWloadidx2)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off] {sym} (ADDQ ptr idx) mem)
	// cond: ptr.Op != OpSB
	// result: (MOVWloadidx1 [off] {sym} ptr idx mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVWloadidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWload [off1] {sym1} (LEAL [off2] {sym2} base) mem)
	// cond: canMergeSym(sym1, sym2) && is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpAMD64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDLconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVLconst [int64(read16(sym, off, config.BigEndian))])
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = int64(read16(sym, off, config.BigEndian))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWloadidx1_0(v *Value) bool {
	// match: (MOVWloadidx1 [c] {sym} ptr (SHLQconst [1] idx) mem)
	// result: (MOVWloadidx2 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 1 {
				continue
			}
			idx := v_1.Args[0]
			v.reset(OpAMD64MOVWloadidx2)
			v.AuxInt = c
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWloadidx1 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVWloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_0.AuxInt
			ptr := v_0.Args[0]
			idx := v.Args[1^_i0]
			if !(is32Bit(c + d)) {
				continue
			}
			v.reset(OpAMD64MOVWloadidx1)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWloadidx1 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+d)
	// result: (MOVWloadidx1 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ADDQconst {
				continue
			}
			d := v_1.AuxInt
			idx := v_1.Args[0]
			if !(is32Bit(c + d)) {
				continue
			}
			v.reset(OpAMD64MOVWloadidx1)
			v.AuxInt = c + d
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWloadidx1 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+c)
	// result: (MOVWload [i+c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		for _i0 := 0; _i0 <= 1; _i0++ {
			p := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(i + c)) {
				continue
			}
			v.reset(OpAMD64MOVWload)
			v.AuxInt = i + c
			v.Aux = s
			v.AddArg(p)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWloadidx2_0(v *Value) bool {
	// match: (MOVWloadidx2 [c] {sym} (ADDQconst [d] ptr) idx mem)
	// cond: is32Bit(c+d)
	// result: (MOVWloadidx2 [c+d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVWloadidx2)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx2 [c] {sym} ptr (ADDQconst [d] idx) mem)
	// cond: is32Bit(c+2*d)
	// result: (MOVWloadidx2 [c+2*d] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		if !(is32Bit(c + 2*d)) {
			break
		}
		v.reset(OpAMD64MOVWloadidx2)
		v.AuxInt = c + 2*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx2 [i] {s} p (MOVQconst [c]) mem)
	// cond: is32Bit(i+2*c)
	// result: (MOVWload [i+2*c] {s} p mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(i + 2*c)) {
			break
		}
		v.reset(OpAMD64MOVWload)
		v.AuxInt = i + 2*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWstore_0(v *Value) bool {
	// match: (MOVWstore [off] {sym} ptr (MOVWQSX x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVWQSX {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWQZX x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVWQZX {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (ADDQconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVLconst [c]) mem)
	// cond: validOff(off)
	// result: (MOVWstoreconst [makeValAndOff(int64(int16(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(validOff(off)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = makeValAndOff(int64(int16(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVQconst [c]) mem)
	// cond: validOff(off)
	// result: (MOVWstoreconst [makeValAndOff(int64(int16(c)),off)] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(validOff(off)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = makeValAndOff(int64(int16(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (LEAQ1 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWstoreidx1 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (LEAQ2 [off2] {sym2} ptr idx) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWstoreidx2 [off1+off2] {mergeSym(sym1,sym2)} ptr idx val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ2 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx2)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} (ADDQ ptr idx) val mem)
	// cond: ptr.Op != OpSB
	// result: (MOVWstoreidx1 [off] {sym} ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			ptr := v_0.Args[_i0]
			idx := v_0.Args[1^_i0]
			val := v.Args[1]
			if !(ptr.Op != OpSB) {
				continue
			}
			v.reset(OpAMD64MOVWstoreidx1)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (MOVWstore [i] {s} p (SHRLconst [16] w) x:(MOVWstore [i-2] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstore [i-2] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRLconst || v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVWstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWstore_10(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVWstore [i] {s} p (SHRQconst [16] w) x:(MOVWstore [i-2] {s} p w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstore [i-2] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRQconst || v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVWstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || w != x.Args[1] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [i] {s} p (SHRLconst [j] w) x:(MOVWstore [i-2] {s} p w0:(SHRLconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstore [i-2] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVWstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpAMD64SHRLconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [i] {s} p (SHRQconst [j] w) x:(MOVWstore [i-2] {s} p w0:(SHRQconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstore [i-2] {s} p w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHRQconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpAMD64MOVWstore || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpAMD64SHRQconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [i] {s} p x1:(MOVWload [j] {s2} p2 mem) mem2:(MOVWstore [i-2] {s} p x2:(MOVWload [j-2] {s2} p2 mem) mem))
	// cond: x1.Uses == 1 && x2.Uses == 1 && mem2.Uses == 1 && clobber(x1) && clobber(x2) && clobber(mem2)
	// result: (MOVLstore [i-2] {s} p (MOVLload [j-2] {s2} p2 mem) mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpAMD64MOVWload {
			break
		}
		j := x1.AuxInt
		s2 := x1.Aux
		mem := x1.Args[1]
		p2 := x1.Args[0]
		mem2 := v.Args[2]
		if mem2.Op != OpAMD64MOVWstore || mem2.AuxInt != i-2 || mem2.Aux != s {
			break
		}
		_ = mem2.Args[2]
		if p != mem2.Args[0] {
			break
		}
		x2 := mem2.Args[1]
		if x2.Op != OpAMD64MOVWload || x2.AuxInt != j-2 || x2.Aux != s2 {
			break
		}
		_ = x2.Args[1]
		if p2 != x2.Args[0] || mem != x2.Args[1] || mem != mem2.Args[2] || !(x1.Uses == 1 && x2.Uses == 1 && mem2.Uses == 1 && clobber(x1) && clobber(x2) && clobber(mem2)) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x2.Pos, OpAMD64MOVLload, typ.UInt32)
		v0.AuxInt = j - 2
		v0.Aux = s2
		v0.AddArg(p2)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (LEAL [off2] {sym2} base) val mem)
	// cond: canMergeSym(sym1, sym2) && is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (ADDLconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWstoreconst_0(v *Value) bool {
	// match: (MOVWstoreconst [sc] {s} (ADDQconst [off] ptr) mem)
	// cond: ValAndOff(sc).canAdd(off)
	// result: (MOVWstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [sc] {sym1} (LEAQ [off] {sym2} ptr) mem)
	// cond: canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVWstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [x] {sym1} (LEAQ1 [off] {sym2} ptr idx) mem)
	// cond: canMergeSym(sym1, sym2)
	// result: (MOVWstoreconstidx1 [ValAndOff(x).add(off)] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		x := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ1 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [x] {sym1} (LEAQ2 [off] {sym2} ptr idx) mem)
	// cond: canMergeSym(sym1, sym2)
	// result: (MOVWstoreconstidx2 [ValAndOff(x).add(off)] {mergeSym(sym1,sym2)} ptr idx mem)
	for {
		x := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ2 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconstidx2)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [x] {sym} (ADDQ ptr idx) mem)
	// result: (MOVWstoreconstidx1 [x] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQ {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		v.reset(OpAMD64MOVWstoreconstidx1)
		v.AuxInt = x
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [c] {s} p x:(MOVWstoreconst [a] {s} p mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 2 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVLstoreconst [makeValAndOff(ValAndOff(a).Val()&0xffff | ValAndOff(c).Val()<<16, ValAndOff(a).Off())] {s} p mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpAMD64MOVWstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xffff|ValAndOff(c).Val()<<16, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [a] {s} p x:(MOVWstoreconst [c] {s} p mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 2 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVLstoreconst [makeValAndOff(ValAndOff(a).Val()&0xffff | ValAndOff(c).Val()<<16, ValAndOff(a).Off())] {s} p mem)
	for {
		a := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpAMD64MOVWstoreconst {
			break
		}
		c := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[1]
		if p != x.Args[0] || !(x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xffff|ValAndOff(c).Val()<<16, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [sc] {sym1} (LEAL [off] {sym2} ptr) mem)
	// cond: canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)
	// result: (MOVWstoreconst [ValAndOff(sc).add(off)] {mergeSym(sym1, sym2)} ptr mem)
	for {
		sc := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAL {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconst [sc] {s} (ADDLconst [off] ptr) mem)
	// cond: ValAndOff(sc).canAdd(off)
	// result: (MOVWstoreconst [ValAndOff(sc).add(off)] {s} ptr mem)
	for {
		sc := v.AuxInt
		s := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDLconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWstoreconstidx1_0(v *Value) bool {
	// match: (MOVWstoreconstidx1 [c] {sym} ptr (SHLQconst [1] idx) mem)
	// result: (MOVWstoreconstidx2 [c] {sym} ptr idx mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		v.reset(OpAMD64MOVWstoreconstidx2)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconstidx1 [x] {sym} (ADDQconst [c] ptr) idx mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVWstoreconstidx1 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconstidx1 [x] {sym} ptr (ADDQconst [c] idx) mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVWstoreconstidx1 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconstidx1 [c] {s} p i x:(MOVWstoreconstidx1 [a] {s} p i mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 2 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVLstoreconstidx1 [makeValAndOff(ValAndOff(a).Val()&0xffff | ValAndOff(c).Val()<<16, ValAndOff(a).Off())] {s} p i mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		i := v.Args[1]
		x := v.Args[2]
		if x.Op != OpAMD64MOVWstoreconstidx1 {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || i != x.Args[1] || !(x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconstidx1)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xffff|ValAndOff(c).Val()<<16, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(i)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWstoreconstidx2_0(v *Value) bool {
	b := v.Block
	// match: (MOVWstoreconstidx2 [x] {sym} (ADDQconst [c] ptr) idx mem)
	// cond: ValAndOff(x).canAdd(c)
	// result: (MOVWstoreconstidx2 [ValAndOff(x).add(c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(ValAndOff(x).canAdd(c)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconstidx2)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconstidx2 [x] {sym} ptr (ADDQconst [c] idx) mem)
	// cond: ValAndOff(x).canAdd(2*c)
	// result: (MOVWstoreconstidx2 [ValAndOff(x).add(2*c)] {sym} ptr idx mem)
	for {
		x := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(ValAndOff(x).canAdd(2 * c)) {
			break
		}
		v.reset(OpAMD64MOVWstoreconstidx2)
		v.AuxInt = ValAndOff(x).add(2 * c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreconstidx2 [c] {s} p i x:(MOVWstoreconstidx2 [a] {s} p i mem))
	// cond: x.Uses == 1 && ValAndOff(a).Off() + 2 == ValAndOff(c).Off() && clobber(x)
	// result: (MOVLstoreconstidx1 [makeValAndOff(ValAndOff(a).Val()&0xffff | ValAndOff(c).Val()<<16, ValAndOff(a).Off())] {s} p (SHLQconst <i.Type> [1] i) mem)
	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		i := v.Args[1]
		x := v.Args[2]
		if x.Op != OpAMD64MOVWstoreconstidx2 {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		mem := x.Args[2]
		if p != x.Args[0] || i != x.Args[1] || !(x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreconstidx1)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xffff|ValAndOff(c).Val()<<16, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, i.Type)
		v0.AuxInt = 1
		v0.AddArg(i)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWstoreidx1_0(v *Value) bool {
	// match: (MOVWstoreidx1 [c] {sym} ptr (SHLQconst [1] idx) val mem)
	// result: (MOVWstoreidx2 [c] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64SHLQconst || v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		v.reset(OpAMD64MOVWstoreidx2)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx1 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVWstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx1 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+d)
	// result: (MOVWstoreidx1 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx1)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx1 [i] {s} p idx (SHRLconst [16] w) x:(MOVWstoreidx1 [i-2] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstoreidx1 [i-2] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRLconst || v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVWstoreidx1 || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx1 [i] {s} p idx (SHRQconst [16] w) x:(MOVWstoreidx1 [i-2] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstoreidx1 [i-2] {s} p idx w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst || v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVWstoreidx1 || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx1 [i] {s} p idx (SHRLconst [j] w) x:(MOVWstoreidx1 [i-2] {s} p idx w0:(SHRLconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstoreidx1 [i-2] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVWstoreidx1 || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpAMD64SHRLconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx1 [i] {s} p idx (SHRQconst [j] w) x:(MOVWstoreidx1 [i-2] {s} p idx w0:(SHRQconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstoreidx1 [i-2] {s} p idx w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVWstoreidx1 || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpAMD64SHRQconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx1 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+c)
	// result: (MOVWstore [i+c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + c)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i + c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MOVWstoreidx2_0(v *Value) bool {
	b := v.Block
	// match: (MOVWstoreidx2 [c] {sym} (ADDQconst [d] ptr) idx val mem)
	// cond: is32Bit(c+d)
	// result: (MOVWstoreidx2 [c+d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx2)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx2 [c] {sym} ptr (ADDQconst [d] idx) val mem)
	// cond: is32Bit(c+2*d)
	// result: (MOVWstoreidx2 [c+2*d] {sym} ptr idx val mem)
	for {
		c := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(is32Bit(c + 2*d)) {
			break
		}
		v.reset(OpAMD64MOVWstoreidx2)
		v.AuxInt = c + 2*d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx2 [i] {s} p idx (SHRLconst [16] w) x:(MOVWstoreidx2 [i-2] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstoreidx1 [i-2] {s} p (SHLQconst <idx.Type> [1] idx) w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRLconst || v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVWstoreidx2 || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, idx.Type)
		v0.AuxInt = 1
		v0.AddArg(idx)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx2 [i] {s} p idx (SHRQconst [16] w) x:(MOVWstoreidx2 [i-2] {s} p idx w mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstoreidx1 [i-2] {s} p (SHLQconst <idx.Type> [1] idx) w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst || v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVWstoreidx2 || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] || w != x.Args[2] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, idx.Type)
		v0.AuxInt = 1
		v0.AddArg(idx)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx2 [i] {s} p idx (SHRQconst [j] w) x:(MOVWstoreidx2 [i-2] {s} p idx w0:(SHRQconst [j-16] w) mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: (MOVLstoreidx1 [i-2] {s} p (SHLQconst <idx.Type> [1] idx) w0 mem)
	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SHRQconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpAMD64MOVWstoreidx2 || x.AuxInt != i-2 || x.Aux != s {
			break
		}
		mem := x.Args[3]
		if p != x.Args[0] || idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpAMD64SHRQconst || w0.AuxInt != j-16 || w != w0.Args[0] || !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpAMD64MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, idx.Type)
		v0.AuxInt = 1
		v0.AddArg(idx)
		v.AddArg(v0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx2 [i] {s} p (MOVQconst [c]) w mem)
	// cond: is32Bit(i+2*c)
	// result: (MOVWstore [i+2*c] {s} p w mem)
	for {
		i := v.AuxInt
		s := v.Aux
		mem := v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		w := v.Args[2]
		if !(is32Bit(i + 2*c)) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = i + 2*c
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULL_0(v *Value) bool {
	// match: (MULL x (MOVLconst [c]))
	// result: (MULLconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpAMD64MULLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULLconst_0(v *Value) bool {
	b := v.Block
	// match: (MULLconst [c] (MULLconst [d] x))
	// result: (MULLconst [int64(int32(c * d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MULLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64MULLconst)
		v.AuxInt = int64(int32(c * d))
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [-9] x)
	// result: (NEGL (LEAL8 <v.Type> x x))
	for {
		if v.AuxInt != -9 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64NEGL)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [-5] x)
	// result: (NEGL (LEAL4 <v.Type> x x))
	for {
		if v.AuxInt != -5 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64NEGL)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [-3] x)
	// result: (NEGL (LEAL2 <v.Type> x x))
	for {
		if v.AuxInt != -3 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64NEGL)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [-1] x)
	// result: (NEGL x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64NEGL)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [ 0] _)
	// result: (MOVLconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (MULLconst [ 1] x)
	// result: x
	for {
		if v.AuxInt != 1 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [ 3] x)
	// result: (LEAL2 x x)
	for {
		if v.AuxInt != 3 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL2)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [ 5] x)
	// result: (LEAL4 x x)
	for {
		if v.AuxInt != 5 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL4)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [ 7] x)
	// result: (LEAL2 x (LEAL2 <v.Type> x x))
	for {
		if v.AuxInt != 7 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULLconst_10(v *Value) bool {
	b := v.Block
	// match: (MULLconst [ 9] x)
	// result: (LEAL8 x x)
	for {
		if v.AuxInt != 9 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL8)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [11] x)
	// result: (LEAL2 x (LEAL4 <v.Type> x x))
	for {
		if v.AuxInt != 11 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [13] x)
	// result: (LEAL4 x (LEAL2 <v.Type> x x))
	for {
		if v.AuxInt != 13 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [19] x)
	// result: (LEAL2 x (LEAL8 <v.Type> x x))
	for {
		if v.AuxInt != 19 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [21] x)
	// result: (LEAL4 x (LEAL4 <v.Type> x x))
	for {
		if v.AuxInt != 21 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [25] x)
	// result: (LEAL8 x (LEAL2 <v.Type> x x))
	for {
		if v.AuxInt != 25 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [27] x)
	// result: (LEAL8 (LEAL2 <v.Type> x x) (LEAL2 <v.Type> x x))
	for {
		if v.AuxInt != 27 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL8)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64LEAL2, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
	// match: (MULLconst [37] x)
	// result: (LEAL4 x (LEAL8 <v.Type> x x))
	for {
		if v.AuxInt != 37 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [41] x)
	// result: (LEAL8 x (LEAL4 <v.Type> x x))
	for {
		if v.AuxInt != 41 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [45] x)
	// result: (LEAL8 (LEAL4 <v.Type> x x) (LEAL4 <v.Type> x x))
	for {
		if v.AuxInt != 45 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL8)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64LEAL4, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULLconst_20(v *Value) bool {
	b := v.Block
	// match: (MULLconst [73] x)
	// result: (LEAL8 x (LEAL8 <v.Type> x x))
	for {
		if v.AuxInt != 73 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [81] x)
	// result: (LEAL8 (LEAL8 <v.Type> x x) (LEAL8 <v.Type> x x))
	for {
		if v.AuxInt != 81 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAL8)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64LEAL8, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
	// match: (MULLconst [c] x)
	// cond: isPowerOfTwo(c+1) && c >= 15
	// result: (SUBL (SHLLconst <v.Type> [log2(c+1)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c+1) && c >= 15) {
			break
		}
		v.reset(OpAMD64SUBL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLLconst, v.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [c] x)
	// cond: isPowerOfTwo(c-1) && c >= 17
	// result: (LEAL1 (SHLLconst <v.Type> [log2(c-1)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-1) && c >= 17) {
			break
		}
		v.reset(OpAMD64LEAL1)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLLconst, v.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [c] x)
	// cond: isPowerOfTwo(c-2) && c >= 34
	// result: (LEAL2 (SHLLconst <v.Type> [log2(c-2)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-2) && c >= 34) {
			break
		}
		v.reset(OpAMD64LEAL2)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLLconst, v.Type)
		v0.AuxInt = log2(c - 2)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [c] x)
	// cond: isPowerOfTwo(c-4) && c >= 68
	// result: (LEAL4 (SHLLconst <v.Type> [log2(c-4)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-4) && c >= 68) {
			break
		}
		v.reset(OpAMD64LEAL4)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLLconst, v.Type)
		v0.AuxInt = log2(c - 4)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [c] x)
	// cond: isPowerOfTwo(c-8) && c >= 136
	// result: (LEAL8 (SHLLconst <v.Type> [log2(c-8)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-8) && c >= 136) {
			break
		}
		v.reset(OpAMD64LEAL8)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLLconst, v.Type)
		v0.AuxInt = log2(c - 8)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULLconst [c] x)
	// cond: c%3 == 0 && isPowerOfTwo(c/3)
	// result: (SHLLconst [log2(c/3)] (LEAL2 <v.Type> x x))
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%3 == 0 && isPowerOfTwo(c/3)) {
			break
		}
		v.reset(OpAMD64SHLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [c] x)
	// cond: c%5 == 0 && isPowerOfTwo(c/5)
	// result: (SHLLconst [log2(c/5)] (LEAL4 <v.Type> x x))
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%5 == 0 && isPowerOfTwo(c/5)) {
			break
		}
		v.reset(OpAMD64SHLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULLconst [c] x)
	// cond: c%9 == 0 && isPowerOfTwo(c/9)
	// result: (SHLLconst [log2(c/9)] (LEAL8 <v.Type> x x))
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%9 == 0 && isPowerOfTwo(c/9)) {
			break
		}
		v.reset(OpAMD64SHLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULLconst_30(v *Value) bool {
	// match: (MULLconst [c] (MOVLconst [d]))
	// result: (MOVLconst [int64(int32(c*d))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = int64(int32(c * d))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULQ_0(v *Value) bool {
	// match: (MULQ x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (MULQconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpAMD64MULQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULQconst_0(v *Value) bool {
	b := v.Block
	// match: (MULQconst [c] (MULQconst [d] x))
	// cond: is32Bit(c*d)
	// result: (MULQconst [c * d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MULQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c * d)) {
			break
		}
		v.reset(OpAMD64MULQconst)
		v.AuxInt = c * d
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [-9] x)
	// result: (NEGQ (LEAQ8 <v.Type> x x))
	for {
		if v.AuxInt != -9 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64NEGQ)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [-5] x)
	// result: (NEGQ (LEAQ4 <v.Type> x x))
	for {
		if v.AuxInt != -5 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64NEGQ)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [-3] x)
	// result: (NEGQ (LEAQ2 <v.Type> x x))
	for {
		if v.AuxInt != -3 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64NEGQ)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [-1] x)
	// result: (NEGQ x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64NEGQ)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [ 0] _)
	// result: (MOVQconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = 0
		return true
	}
	// match: (MULQconst [ 1] x)
	// result: x
	for {
		if v.AuxInt != 1 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [ 3] x)
	// result: (LEAQ2 x x)
	for {
		if v.AuxInt != 3 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ2)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [ 5] x)
	// result: (LEAQ4 x x)
	for {
		if v.AuxInt != 5 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ4)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [ 7] x)
	// result: (LEAQ2 x (LEAQ2 <v.Type> x x))
	for {
		if v.AuxInt != 7 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULQconst_10(v *Value) bool {
	b := v.Block
	// match: (MULQconst [ 9] x)
	// result: (LEAQ8 x x)
	for {
		if v.AuxInt != 9 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ8)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [11] x)
	// result: (LEAQ2 x (LEAQ4 <v.Type> x x))
	for {
		if v.AuxInt != 11 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [13] x)
	// result: (LEAQ4 x (LEAQ2 <v.Type> x x))
	for {
		if v.AuxInt != 13 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [19] x)
	// result: (LEAQ2 x (LEAQ8 <v.Type> x x))
	for {
		if v.AuxInt != 19 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [21] x)
	// result: (LEAQ4 x (LEAQ4 <v.Type> x x))
	for {
		if v.AuxInt != 21 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [25] x)
	// result: (LEAQ8 x (LEAQ2 <v.Type> x x))
	for {
		if v.AuxInt != 25 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [27] x)
	// result: (LEAQ8 (LEAQ2 <v.Type> x x) (LEAQ2 <v.Type> x x))
	for {
		if v.AuxInt != 27 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ8)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64LEAQ2, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
	// match: (MULQconst [37] x)
	// result: (LEAQ4 x (LEAQ8 <v.Type> x x))
	for {
		if v.AuxInt != 37 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [41] x)
	// result: (LEAQ8 x (LEAQ4 <v.Type> x x))
	for {
		if v.AuxInt != 41 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [45] x)
	// result: (LEAQ8 (LEAQ4 <v.Type> x x) (LEAQ4 <v.Type> x x))
	for {
		if v.AuxInt != 45 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ8)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64LEAQ4, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULQconst_20(v *Value) bool {
	b := v.Block
	// match: (MULQconst [73] x)
	// result: (LEAQ8 x (LEAQ8 <v.Type> x x))
	for {
		if v.AuxInt != 73 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [81] x)
	// result: (LEAQ8 (LEAQ8 <v.Type> x x) (LEAQ8 <v.Type> x x))
	for {
		if v.AuxInt != 81 {
			break
		}
		x := v.Args[0]
		v.reset(OpAMD64LEAQ8)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64LEAQ8, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
	// match: (MULQconst [c] x)
	// cond: isPowerOfTwo(c+1) && c >= 15
	// result: (SUBQ (SHLQconst <v.Type> [log2(c+1)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c+1) && c >= 15) {
			break
		}
		v.reset(OpAMD64SUBQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [c] x)
	// cond: isPowerOfTwo(c-1) && c >= 17
	// result: (LEAQ1 (SHLQconst <v.Type> [log2(c-1)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-1) && c >= 17) {
			break
		}
		v.reset(OpAMD64LEAQ1)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [c] x)
	// cond: isPowerOfTwo(c-2) && c >= 34
	// result: (LEAQ2 (SHLQconst <v.Type> [log2(c-2)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-2) && c >= 34) {
			break
		}
		v.reset(OpAMD64LEAQ2)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
		v0.AuxInt = log2(c - 2)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [c] x)
	// cond: isPowerOfTwo(c-4) && c >= 68
	// result: (LEAQ4 (SHLQconst <v.Type> [log2(c-4)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-4) && c >= 68) {
			break
		}
		v.reset(OpAMD64LEAQ4)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
		v0.AuxInt = log2(c - 4)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [c] x)
	// cond: isPowerOfTwo(c-8) && c >= 136
	// result: (LEAQ8 (SHLQconst <v.Type> [log2(c-8)] x) x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-8) && c >= 136) {
			break
		}
		v.reset(OpAMD64LEAQ8)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
		v0.AuxInt = log2(c - 8)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (MULQconst [c] x)
	// cond: c%3 == 0 && isPowerOfTwo(c/3)
	// result: (SHLQconst [log2(c/3)] (LEAQ2 <v.Type> x x))
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%3 == 0 && isPowerOfTwo(c/3)) {
			break
		}
		v.reset(OpAMD64SHLQconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [c] x)
	// cond: c%5 == 0 && isPowerOfTwo(c/5)
	// result: (SHLQconst [log2(c/5)] (LEAQ4 <v.Type> x x))
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%5 == 0 && isPowerOfTwo(c/5)) {
			break
		}
		v.reset(OpAMD64SHLQconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MULQconst [c] x)
	// cond: c%9 == 0 && isPowerOfTwo(c/9)
	// result: (SHLQconst [log2(c/9)] (LEAQ8 <v.Type> x x))
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%9 == 0 && isPowerOfTwo(c/9)) {
			break
		}
		v.reset(OpAMD64SHLQconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAQ8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULQconst_30(v *Value) bool {
	// match: (MULQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [c*d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = c * d
		return true
	}
	// match: (MULQconst [c] (NEGQ x))
	// cond: c != -(1<<31)
	// result: (MULQconst [-c] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64NEGQ {
			break
		}
		x := v_0.Args[0]
		if !(c != -(1 << 31)) {
			break
		}
		v.reset(OpAMD64MULQconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULSD_0(v *Value) bool {
	// match: (MULSD x l:(MOVSDload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (MULSDload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVSDload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64MULSDload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULSDload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MULSDload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MULSDload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MULSDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MULSDload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MULSDload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MULSDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MULSDload x [off] {sym} ptr (MOVQstore [off] {sym} ptr y _))
	// result: (MULSD x (MOVQi2f y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVQstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64MULSD)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVQi2f, typ.Float64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULSS_0(v *Value) bool {
	// match: (MULSS x l:(MOVSSload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (MULSSload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVSSload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64MULSSload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64MULSSload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MULSSload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MULSSload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64MULSSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MULSSload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MULSSload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64MULSSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MULSSload x [off] {sym} ptr (MOVLstore [off] {sym} ptr y _))
	// result: (MULSS x (MOVLi2f y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVLstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64MULSS)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVLi2f, typ.Float32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64NEGL_0(v *Value) bool {
	// match: (NEGL (NEGL x))
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64NEGL {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (NEGL s:(SUBL x y))
	// cond: s.Uses == 1
	// result: (SUBL y x)
	for {
		s := v.Args[0]
		if s.Op != OpAMD64SUBL {
			break
		}
		y := s.Args[1]
		x := s.Args[0]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpAMD64SUBL)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	// match: (NEGL (MOVLconst [c]))
	// result: (MOVLconst [int64(int32(-c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = int64(int32(-c))
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64NEGQ_0(v *Value) bool {
	// match: (NEGQ (NEGQ x))
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64NEGQ {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (NEGQ s:(SUBQ x y))
	// cond: s.Uses == 1
	// result: (SUBQ y x)
	for {
		s := v.Args[0]
		if s.Op != OpAMD64SUBQ {
			break
		}
		y := s.Args[1]
		x := s.Args[0]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpAMD64SUBQ)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	// match: (NEGQ (MOVQconst [c]))
	// result: (MOVQconst [-c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = -c
		return true
	}
	// match: (NEGQ (ADDQconst [c] (NEGQ x)))
	// cond: c != -(1<<31)
	// result: (ADDQconst [-c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAMD64NEGQ {
			break
		}
		x := v_0_0.Args[0]
		if !(c != -(1 << 31)) {
			break
		}
		v.reset(OpAMD64ADDQconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64NOTL_0(v *Value) bool {
	// match: (NOTL (MOVLconst [c]))
	// result: (MOVLconst [^c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = ^c
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64NOTQ_0(v *Value) bool {
	// match: (NOTQ (MOVQconst [c]))
	// result: (MOVQconst [^c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = ^c
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORL_0(v *Value) bool {
	// match: (ORL (SHLL (MOVLconst [1]) y) x)
	// result: (BTSL x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLL {
				continue
			}
			y := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64MOVLconst || v_0_0.AuxInt != 1 {
				continue
			}
			x := v.Args[1^_i0]
			v.reset(OpAMD64BTSL)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ORL (MOVLconst [c]) x)
	// cond: isUint32PowerOfTwo(c) && uint64(c) >= 128
	// result: (BTSLconst [log2uint32(c)] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			if !(isUint32PowerOfTwo(c) && uint64(c) >= 128) {
				continue
			}
			v.reset(OpAMD64BTSLconst)
			v.AuxInt = log2uint32(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORL x (MOVLconst [c]))
	// result: (ORLconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpAMD64ORLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORL (SHLLconst x [c]) (SHRLconst x [d]))
	// cond: d==32-c
	// result: (ROLLconst x [c])
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRLconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpAMD64ROLLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORL <t> (SHLLconst x [c]) (SHRWconst x [d]))
	// cond: d==16-c && c < 16 && t.Size() == 2
	// result: (ROLWconst x [c])
	for {
		t := v.Type
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 16-c && c < 16 && t.Size() == 2) {
				continue
			}
			v.reset(OpAMD64ROLWconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORL <t> (SHLLconst x [c]) (SHRBconst x [d]))
	// cond: d==8-c && c < 8 && t.Size() == 1
	// result: (ROLBconst x [c])
	for {
		t := v.Type
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRBconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 8-c && c < 8 && t.Size() == 1) {
				continue
			}
			v.reset(OpAMD64ROLBconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORL (SHLL x y) (ANDL (SHRL x (NEGQ y)) (SBBLcarrymask (CMPQconst (NEGQ (ADDQconst (ANDQconst y [31]) [-32])) [32]))))
	// result: (ROLL x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLL {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHRL {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGQ || y != v_1_0_1.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBLcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPQconst || v_1_1_0.AuxInt != 32 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGQ {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDQconst || v_1_1_0_0_0.AuxInt != -32 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDQconst || v_1_1_0_0_0_0.AuxInt != 31 || y != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v.reset(OpAMD64ROLL)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL (SHLL x y) (ANDL (SHRL x (NEGL y)) (SBBLcarrymask (CMPLconst (NEGL (ADDLconst (ANDLconst y [31]) [-32])) [32]))))
	// result: (ROLL x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLL {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHRL {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGL || y != v_1_0_1.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBLcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPLconst || v_1_1_0.AuxInt != 32 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGL {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDLconst || v_1_1_0_0_0.AuxInt != -32 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDLconst || v_1_1_0_0_0_0.AuxInt != 31 || y != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v.reset(OpAMD64ROLL)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL (SHRL x y) (ANDL (SHLL x (NEGQ y)) (SBBLcarrymask (CMPQconst (NEGQ (ADDQconst (ANDQconst y [31]) [-32])) [32]))))
	// result: (RORL x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHRL {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHLL {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGQ || y != v_1_0_1.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBLcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPQconst || v_1_1_0.AuxInt != 32 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGQ {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDQconst || v_1_1_0_0_0.AuxInt != -32 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDQconst || v_1_1_0_0_0_0.AuxInt != 31 || y != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v.reset(OpAMD64RORL)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL (SHRL x y) (ANDL (SHLL x (NEGL y)) (SBBLcarrymask (CMPLconst (NEGL (ADDLconst (ANDLconst y [31]) [-32])) [32]))))
	// result: (RORL x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHRL {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHLL {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGL || y != v_1_0_1.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBLcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPLconst || v_1_1_0.AuxInt != 32 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGL {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDLconst || v_1_1_0_0_0.AuxInt != -32 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDLconst || v_1_1_0_0_0_0.AuxInt != 31 || y != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v.reset(OpAMD64RORL)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORL_10(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORL (SHLL x (ANDQconst y [15])) (ANDL (SHRW x (NEGQ (ADDQconst (ANDQconst y [15]) [-16]))) (SBBLcarrymask (CMPQconst (NEGQ (ADDQconst (ANDQconst y [15]) [-16])) [16]))))
	// cond: v.Type.Size() == 2
	// result: (ROLW x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLL {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64ANDQconst || v_0_1.AuxInt != 15 {
				continue
			}
			y := v_0_1.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHRW {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGQ {
					continue
				}
				v_1_0_1_0 := v_1_0_1.Args[0]
				if v_1_0_1_0.Op != OpAMD64ADDQconst || v_1_0_1_0.AuxInt != -16 {
					continue
				}
				v_1_0_1_0_0 := v_1_0_1_0.Args[0]
				if v_1_0_1_0_0.Op != OpAMD64ANDQconst || v_1_0_1_0_0.AuxInt != 15 || y != v_1_0_1_0_0.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBLcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPQconst || v_1_1_0.AuxInt != 16 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGQ {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDQconst || v_1_1_0_0_0.AuxInt != -16 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDQconst || v_1_1_0_0_0_0.AuxInt != 15 || y != v_1_1_0_0_0_0.Args[0] || !(v.Type.Size() == 2) {
					continue
				}
				v.reset(OpAMD64ROLW)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL (SHLL x (ANDLconst y [15])) (ANDL (SHRW x (NEGL (ADDLconst (ANDLconst y [15]) [-16]))) (SBBLcarrymask (CMPLconst (NEGL (ADDLconst (ANDLconst y [15]) [-16])) [16]))))
	// cond: v.Type.Size() == 2
	// result: (ROLW x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLL {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64ANDLconst || v_0_1.AuxInt != 15 {
				continue
			}
			y := v_0_1.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHRW {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGL {
					continue
				}
				v_1_0_1_0 := v_1_0_1.Args[0]
				if v_1_0_1_0.Op != OpAMD64ADDLconst || v_1_0_1_0.AuxInt != -16 {
					continue
				}
				v_1_0_1_0_0 := v_1_0_1_0.Args[0]
				if v_1_0_1_0_0.Op != OpAMD64ANDLconst || v_1_0_1_0_0.AuxInt != 15 || y != v_1_0_1_0_0.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBLcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPLconst || v_1_1_0.AuxInt != 16 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGL {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDLconst || v_1_1_0_0_0.AuxInt != -16 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDLconst || v_1_1_0_0_0_0.AuxInt != 15 || y != v_1_1_0_0_0_0.Args[0] || !(v.Type.Size() == 2) {
					continue
				}
				v.reset(OpAMD64ROLW)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL (SHRW x (ANDQconst y [15])) (SHLL x (NEGQ (ADDQconst (ANDQconst y [15]) [-16]))))
	// cond: v.Type.Size() == 2
	// result: (RORW x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHRW {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64ANDQconst || v_0_1.AuxInt != 15 {
				continue
			}
			y := v_0_1.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLL {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpAMD64NEGQ {
				continue
			}
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpAMD64ADDQconst || v_1_1_0.AuxInt != -16 {
				continue
			}
			v_1_1_0_0 := v_1_1_0.Args[0]
			if v_1_1_0_0.Op != OpAMD64ANDQconst || v_1_1_0_0.AuxInt != 15 || y != v_1_1_0_0.Args[0] || !(v.Type.Size() == 2) {
				continue
			}
			v.reset(OpAMD64RORW)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ORL (SHRW x (ANDLconst y [15])) (SHLL x (NEGL (ADDLconst (ANDLconst y [15]) [-16]))))
	// cond: v.Type.Size() == 2
	// result: (RORW x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHRW {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64ANDLconst || v_0_1.AuxInt != 15 {
				continue
			}
			y := v_0_1.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLL {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpAMD64NEGL {
				continue
			}
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpAMD64ADDLconst || v_1_1_0.AuxInt != -16 {
				continue
			}
			v_1_1_0_0 := v_1_1_0.Args[0]
			if v_1_1_0_0.Op != OpAMD64ANDLconst || v_1_1_0_0.AuxInt != 15 || y != v_1_1_0_0.Args[0] || !(v.Type.Size() == 2) {
				continue
			}
			v.reset(OpAMD64RORW)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ORL (SHLL x (ANDQconst y [ 7])) (ANDL (SHRB x (NEGQ (ADDQconst (ANDQconst y [ 7]) [ -8]))) (SBBLcarrymask (CMPQconst (NEGQ (ADDQconst (ANDQconst y [ 7]) [ -8])) [ 8]))))
	// cond: v.Type.Size() == 1
	// result: (ROLB x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLL {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64ANDQconst || v_0_1.AuxInt != 7 {
				continue
			}
			y := v_0_1.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHRB {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGQ {
					continue
				}
				v_1_0_1_0 := v_1_0_1.Args[0]
				if v_1_0_1_0.Op != OpAMD64ADDQconst || v_1_0_1_0.AuxInt != -8 {
					continue
				}
				v_1_0_1_0_0 := v_1_0_1_0.Args[0]
				if v_1_0_1_0_0.Op != OpAMD64ANDQconst || v_1_0_1_0_0.AuxInt != 7 || y != v_1_0_1_0_0.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBLcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPQconst || v_1_1_0.AuxInt != 8 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGQ {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDQconst || v_1_1_0_0_0.AuxInt != -8 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDQconst || v_1_1_0_0_0_0.AuxInt != 7 || y != v_1_1_0_0_0_0.Args[0] || !(v.Type.Size() == 1) {
					continue
				}
				v.reset(OpAMD64ROLB)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL (SHLL x (ANDLconst y [ 7])) (ANDL (SHRB x (NEGL (ADDLconst (ANDLconst y [ 7]) [ -8]))) (SBBLcarrymask (CMPLconst (NEGL (ADDLconst (ANDLconst y [ 7]) [ -8])) [ 8]))))
	// cond: v.Type.Size() == 1
	// result: (ROLB x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLL {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64ANDLconst || v_0_1.AuxInt != 7 {
				continue
			}
			y := v_0_1.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDL {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHRB {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGL {
					continue
				}
				v_1_0_1_0 := v_1_0_1.Args[0]
				if v_1_0_1_0.Op != OpAMD64ADDLconst || v_1_0_1_0.AuxInt != -8 {
					continue
				}
				v_1_0_1_0_0 := v_1_0_1_0.Args[0]
				if v_1_0_1_0_0.Op != OpAMD64ANDLconst || v_1_0_1_0_0.AuxInt != 7 || y != v_1_0_1_0_0.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBLcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPLconst || v_1_1_0.AuxInt != 8 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGL {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDLconst || v_1_1_0_0_0.AuxInt != -8 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDLconst || v_1_1_0_0_0_0.AuxInt != 7 || y != v_1_1_0_0_0_0.Args[0] || !(v.Type.Size() == 1) {
					continue
				}
				v.reset(OpAMD64ROLB)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL (SHRB x (ANDQconst y [ 7])) (SHLL x (NEGQ (ADDQconst (ANDQconst y [ 7]) [ -8]))))
	// cond: v.Type.Size() == 1
	// result: (RORB x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHRB {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64ANDQconst || v_0_1.AuxInt != 7 {
				continue
			}
			y := v_0_1.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLL {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpAMD64NEGQ {
				continue
			}
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpAMD64ADDQconst || v_1_1_0.AuxInt != -8 {
				continue
			}
			v_1_1_0_0 := v_1_1_0.Args[0]
			if v_1_1_0_0.Op != OpAMD64ANDQconst || v_1_1_0_0.AuxInt != 7 || y != v_1_1_0_0.Args[0] || !(v.Type.Size() == 1) {
				continue
			}
			v.reset(OpAMD64RORB)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ORL (SHRB x (ANDLconst y [ 7])) (SHLL x (NEGL (ADDLconst (ANDLconst y [ 7]) [ -8]))))
	// cond: v.Type.Size() == 1
	// result: (RORB x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHRB {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64ANDLconst || v_0_1.AuxInt != 7 {
				continue
			}
			y := v_0_1.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHLL {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpAMD64NEGL {
				continue
			}
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpAMD64ADDLconst || v_1_1_0.AuxInt != -8 {
				continue
			}
			v_1_1_0_0 := v_1_1_0.Args[0]
			if v_1_1_0_0.Op != OpAMD64ANDLconst || v_1_1_0_0.AuxInt != 7 || y != v_1_1_0_0.Args[0] || !(v.Type.Size() == 1) {
				continue
			}
			v.reset(OpAMD64RORB)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ORL x x)
	// result: x
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORL x0:(MOVBload [i0] {s} p mem) sh:(SHLLconst [8] x1:(MOVBload [i1] {s} p mem)))
	// cond: i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVWload [i0] {s} p mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVBload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLLconst || sh.AuxInt != 8 {
				continue
			}
			x1 := sh.Args[0]
			if x1.Op != OpAMD64MOVBload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpAMD64MOVWload, typ.UInt16)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORL_20(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORL x0:(MOVWload [i0] {s} p mem) sh:(SHLLconst [16] x1:(MOVWload [i1] {s} p mem)))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVLload [i0] {s} p mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVWload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLLconst || sh.AuxInt != 16 {
				continue
			}
			x1 := sh.Args[0]
			if x1.Op != OpAMD64MOVWload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpAMD64MOVLload, typ.UInt32)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (ORL s1:(SHLLconst [j1] x1:(MOVBload [i1] {s} p mem)) or:(ORL s0:(SHLLconst [j0] x0:(MOVBload [i0] {s} p mem)) y))
	// cond: i1 == i0+1 && j1 == j0+8 && j0 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORL <v.Type> (SHLLconst <v.Type> [j0] (MOVWload [i0] {s} p mem)) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s1 := v.Args[_i0]
			if s1.Op != OpAMD64SHLLconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpAMD64MOVBload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			or := v.Args[1^_i0]
			if or.Op != OpAMD64ORL {
				continue
			}
			_ = or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				s0 := or.Args[_i1]
				if s0.Op != OpAMD64SHLLconst {
					continue
				}
				j0 := s0.AuxInt
				x0 := s0.Args[0]
				if x0.Op != OpAMD64MOVBload {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[1]
				if p != x0.Args[0] || mem != x0.Args[1] {
					continue
				}
				y := or.Args[1^_i1]
				if !(i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x0.Pos, OpAMD64ORL, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x0.Pos, OpAMD64SHLLconst, v.Type)
				v1.AuxInt = j0
				v2 := b.NewValue0(x0.Pos, OpAMD64MOVWload, typ.UInt16)
				v2.AuxInt = i0
				v2.Aux = s
				v2.AddArg(p)
				v2.AddArg(mem)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL x0:(MOVBloadidx1 [i0] {s} p idx mem) sh:(SHLLconst [8] x1:(MOVBloadidx1 [i1] {s} p idx mem)))
	// cond: i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVWloadidx1 <v.Type> [i0] {s} p idx mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVBloadidx1 {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x0.Args[_i1]
				idx := x0.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLLconst || sh.AuxInt != 8 {
					continue
				}
				x1 := sh.Args[0]
				if x1.Op != OpAMD64MOVBloadidx1 {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x1.Args[_i2] || idx != x1.Args[1^_i2] || mem != x1.Args[2] || !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, v.Type)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (ORL x0:(MOVWloadidx1 [i0] {s} p idx mem) sh:(SHLLconst [16] x1:(MOVWloadidx1 [i1] {s} p idx mem)))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVLloadidx1 [i0] {s} p idx mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVWloadidx1 {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x0.Args[_i1]
				idx := x0.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLLconst || sh.AuxInt != 16 {
					continue
				}
				x1 := sh.Args[0]
				if x1.Op != OpAMD64MOVWloadidx1 {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x1.Args[_i2] || idx != x1.Args[1^_i2] || mem != x1.Args[2] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (ORL s1:(SHLLconst [j1] x1:(MOVBloadidx1 [i1] {s} p idx mem)) or:(ORL s0:(SHLLconst [j0] x0:(MOVBloadidx1 [i0] {s} p idx mem)) y))
	// cond: i1 == i0+1 && j1 == j0+8 && j0 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORL <v.Type> (SHLLconst <v.Type> [j0] (MOVWloadidx1 [i0] {s} p idx mem)) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s1 := v.Args[_i0]
			if s1.Op != OpAMD64SHLLconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpAMD64MOVBloadidx1 {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x1.Args[_i1]
				idx := x1.Args[1^_i1]
				or := v.Args[1^_i0]
				if or.Op != OpAMD64ORL {
					continue
				}
				_ = or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2++ {
					s0 := or.Args[_i2]
					if s0.Op != OpAMD64SHLLconst {
						continue
					}
					j0 := s0.AuxInt
					x0 := s0.Args[0]
					if x0.Op != OpAMD64MOVBloadidx1 {
						continue
					}
					i0 := x0.AuxInt
					if x0.Aux != s {
						continue
					}
					_ = x0.Args[2]
					for _i3 := 0; _i3 <= 1; _i3++ {
						if p != x0.Args[_i3] || idx != x0.Args[1^_i3] || mem != x0.Args[2] {
							continue
						}
						y := or.Args[1^_i2]
						if !(i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpAMD64ORL, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpAMD64SHLLconst, v.Type)
						v1.AuxInt = j0
						v2 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, typ.UInt16)
						v2.AuxInt = i0
						v2.Aux = s
						v2.AddArg(p)
						v2.AddArg(idx)
						v2.AddArg(mem)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (ORL x1:(MOVBload [i1] {s} p mem) sh:(SHLLconst [8] x0:(MOVBload [i0] {s} p mem)))
	// cond: i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (ROLWconst <v.Type> [8] (MOVWload [i0] {s} p mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x1 := v.Args[_i0]
			if x1.Op != OpAMD64MOVBload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLLconst || sh.AuxInt != 8 {
				continue
			}
			x0 := sh.Args[0]
			if x0.Op != OpAMD64MOVBload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpAMD64ROLWconst, v.Type)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = 8
			v1 := b.NewValue0(x0.Pos, OpAMD64MOVWload, typ.UInt16)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (ORL r1:(ROLWconst [8] x1:(MOVWload [i1] {s} p mem)) sh:(SHLLconst [16] r0:(ROLWconst [8] x0:(MOVWload [i0] {s} p mem))))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (BSWAPL <v.Type> (MOVLload [i0] {s} p mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			r1 := v.Args[_i0]
			if r1.Op != OpAMD64ROLWconst || r1.AuxInt != 8 {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpAMD64MOVWload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLLconst || sh.AuxInt != 16 {
				continue
			}
			r0 := sh.Args[0]
			if r0.Op != OpAMD64ROLWconst || r0.AuxInt != 8 {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpAMD64MOVWload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpAMD64BSWAPL, v.Type)
			v.reset(OpCopy)
			v.AddArg(v0)
			v1 := b.NewValue0(x0.Pos, OpAMD64MOVLload, typ.UInt32)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (ORL s0:(SHLLconst [j0] x0:(MOVBload [i0] {s} p mem)) or:(ORL s1:(SHLLconst [j1] x1:(MOVBload [i1] {s} p mem)) y))
	// cond: i1 == i0+1 && j1 == j0-8 && j1 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORL <v.Type> (SHLLconst <v.Type> [j1] (ROLWconst <typ.UInt16> [8] (MOVWload [i0] {s} p mem))) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s0 := v.Args[_i0]
			if s0.Op != OpAMD64SHLLconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpAMD64MOVBload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			or := v.Args[1^_i0]
			if or.Op != OpAMD64ORL {
				continue
			}
			_ = or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				s1 := or.Args[_i1]
				if s1.Op != OpAMD64SHLLconst {
					continue
				}
				j1 := s1.AuxInt
				x1 := s1.Args[0]
				if x1.Op != OpAMD64MOVBload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				y := or.Args[1^_i1]
				if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x1.Pos, OpAMD64ORL, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x1.Pos, OpAMD64SHLLconst, v.Type)
				v1.AuxInt = j1
				v2 := b.NewValue0(x1.Pos, OpAMD64ROLWconst, typ.UInt16)
				v2.AuxInt = 8
				v3 := b.NewValue0(x1.Pos, OpAMD64MOVWload, typ.UInt16)
				v3.AuxInt = i0
				v3.Aux = s
				v3.AddArg(p)
				v3.AddArg(mem)
				v2.AddArg(v3)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORL x1:(MOVBloadidx1 [i1] {s} p idx mem) sh:(SHLLconst [8] x0:(MOVBloadidx1 [i0] {s} p idx mem)))
	// cond: i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (ROLWconst <v.Type> [8] (MOVWloadidx1 [i0] {s} p idx mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x1 := v.Args[_i0]
			if x1.Op != OpAMD64MOVBloadidx1 {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x1.Args[_i1]
				idx := x1.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLLconst || sh.AuxInt != 8 {
					continue
				}
				x0 := sh.Args[0]
				if x0.Op != OpAMD64MOVBloadidx1 {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x0.Args[_i2] || idx != x0.Args[1^_i2] || mem != x0.Args[2] || !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64ROLWconst, v.Type)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = 8
					v1 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, typ.UInt16)
					v1.AuxInt = i0
					v1.Aux = s
					v1.AddArg(p)
					v1.AddArg(idx)
					v1.AddArg(mem)
					v0.AddArg(v1)
					return true
				}
			}
		}
		break
	}
	// match: (ORL r1:(ROLWconst [8] x1:(MOVWloadidx1 [i1] {s} p idx mem)) sh:(SHLLconst [16] r0:(ROLWconst [8] x0:(MOVWloadidx1 [i0] {s} p idx mem))))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (BSWAPL <v.Type> (MOVLloadidx1 [i0] {s} p idx mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			r1 := v.Args[_i0]
			if r1.Op != OpAMD64ROLWconst || r1.AuxInt != 8 {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpAMD64MOVWloadidx1 {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x1.Args[_i1]
				idx := x1.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLLconst || sh.AuxInt != 16 {
					continue
				}
				r0 := sh.Args[0]
				if r0.Op != OpAMD64ROLWconst || r0.AuxInt != 8 {
					continue
				}
				x0 := r0.Args[0]
				if x0.Op != OpAMD64MOVWloadidx1 {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x0.Args[_i2] || idx != x0.Args[1^_i2] || mem != x0.Args[2] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64BSWAPL, v.Type)
					v.reset(OpCopy)
					v.AddArg(v0)
					v1 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
					v1.AuxInt = i0
					v1.Aux = s
					v1.AddArg(p)
					v1.AddArg(idx)
					v1.AddArg(mem)
					v0.AddArg(v1)
					return true
				}
			}
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORL_30(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORL s0:(SHLLconst [j0] x0:(MOVBloadidx1 [i0] {s} p idx mem)) or:(ORL s1:(SHLLconst [j1] x1:(MOVBloadidx1 [i1] {s} p idx mem)) y))
	// cond: i1 == i0+1 && j1 == j0-8 && j1 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORL <v.Type> (SHLLconst <v.Type> [j1] (ROLWconst <typ.UInt16> [8] (MOVWloadidx1 [i0] {s} p idx mem))) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s0 := v.Args[_i0]
			if s0.Op != OpAMD64SHLLconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpAMD64MOVBloadidx1 {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x0.Args[_i1]
				idx := x0.Args[1^_i1]
				or := v.Args[1^_i0]
				if or.Op != OpAMD64ORL {
					continue
				}
				_ = or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2++ {
					s1 := or.Args[_i2]
					if s1.Op != OpAMD64SHLLconst {
						continue
					}
					j1 := s1.AuxInt
					x1 := s1.Args[0]
					if x1.Op != OpAMD64MOVBloadidx1 {
						continue
					}
					i1 := x1.AuxInt
					if x1.Aux != s {
						continue
					}
					_ = x1.Args[2]
					for _i3 := 0; _i3 <= 1; _i3++ {
						if p != x1.Args[_i3] || idx != x1.Args[1^_i3] || mem != x1.Args[2] {
							continue
						}
						y := or.Args[1^_i2]
						if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpAMD64ORL, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpAMD64SHLLconst, v.Type)
						v1.AuxInt = j1
						v2 := b.NewValue0(v.Pos, OpAMD64ROLWconst, typ.UInt16)
						v2.AuxInt = 8
						v3 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, typ.UInt16)
						v3.AuxInt = i0
						v3.Aux = s
						v3.AddArg(p)
						v3.AddArg(idx)
						v3.AddArg(mem)
						v2.AddArg(v3)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (ORL x l:(MOVLload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (ORLload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVLload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ORLload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORLconst_0(v *Value) bool {
	// match: (ORLconst [c] x)
	// cond: isUint32PowerOfTwo(c) && uint64(c) >= 128
	// result: (BTSLconst [log2uint32(c)] x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isUint32PowerOfTwo(c) && uint64(c) >= 128) {
			break
		}
		v.reset(OpAMD64BTSLconst)
		v.AuxInt = log2uint32(c)
		v.AddArg(x)
		return true
	}
	// match: (ORLconst [c] (ORLconst [d] x))
	// result: (ORLconst [c | d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ORLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ORLconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	// match: (ORLconst [c] (BTSLconst [d] x))
	// result: (ORLconst [c | 1<<uint32(d)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTSLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ORLconst)
		v.AuxInt = c | 1<<uint32(d)
		v.AddArg(x)
		return true
	}
	// match: (ORLconst [c] x)
	// cond: int32(c)==0
	// result: x
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORLconst [c] _)
	// cond: int32(c)==-1
	// result: (MOVLconst [-1])
	for {
		c := v.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORLconst [c] (MOVLconst [d]))
	// result: (MOVLconst [c|d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = c | d
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORLconstmodify_0(v *Value) bool {
	// match: (ORLconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (ORLconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64ORLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ORLconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (ORLconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ORLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORLload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORLload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (ORLload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ORLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ORLload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ORLload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ORLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: ( ORLload x [off] {sym} ptr (MOVSSstore [off] {sym} ptr y _))
	// result: ( ORL x (MOVLf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSSstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64ORL)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVLf2i, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORLmodify_0(v *Value) bool {
	// match: (ORLmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (ORLmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ORLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (ORLmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ORLmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ORLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORQ_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORQ (SHLQ (MOVQconst [1]) y) x)
	// result: (BTSQ x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLQ {
				continue
			}
			y := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64MOVQconst || v_0_0.AuxInt != 1 {
				continue
			}
			x := v.Args[1^_i0]
			v.reset(OpAMD64BTSQ)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ORQ (MOVQconst [c]) x)
	// cond: isUint64PowerOfTwo(c) && uint64(c) >= 128
	// result: (BTSQconst [log2(c)] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			if !(isUint64PowerOfTwo(c) && uint64(c) >= 128) {
				continue
			}
			v.reset(OpAMD64BTSQconst)
			v.AuxInt = log2(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORQ x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (ORQconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpAMD64ORQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORQ (SHLQconst x [c]) (SHRQconst x [d]))
	// cond: d==64-c
	// result: (ROLQconst x [c])
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLQconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRQconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpAMD64ROLQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ORQ (SHLQ x y) (ANDQ (SHRQ x (NEGQ y)) (SBBQcarrymask (CMPQconst (NEGQ (ADDQconst (ANDQconst y [63]) [-64])) [64]))))
	// result: (ROLQ x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLQ {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDQ {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHRQ {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGQ || y != v_1_0_1.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBQcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPQconst || v_1_1_0.AuxInt != 64 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGQ {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDQconst || v_1_1_0_0_0.AuxInt != -64 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDQconst || v_1_1_0_0_0_0.AuxInt != 63 || y != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v.reset(OpAMD64ROLQ)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORQ (SHLQ x y) (ANDQ (SHRQ x (NEGL y)) (SBBQcarrymask (CMPLconst (NEGL (ADDLconst (ANDLconst y [63]) [-64])) [64]))))
	// result: (ROLQ x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLQ {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDQ {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHRQ {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGL || y != v_1_0_1.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBQcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPLconst || v_1_1_0.AuxInt != 64 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGL {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDLconst || v_1_1_0_0_0.AuxInt != -64 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDLconst || v_1_1_0_0_0_0.AuxInt != 63 || y != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v.reset(OpAMD64ROLQ)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORQ (SHRQ x y) (ANDQ (SHLQ x (NEGQ y)) (SBBQcarrymask (CMPQconst (NEGQ (ADDQconst (ANDQconst y [63]) [-64])) [64]))))
	// result: (RORQ x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHRQ {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDQ {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHLQ {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGQ || y != v_1_0_1.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBQcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPQconst || v_1_1_0.AuxInt != 64 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGQ {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDQconst || v_1_1_0_0_0.AuxInt != -64 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDQconst || v_1_1_0_0_0_0.AuxInt != 63 || y != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v.reset(OpAMD64RORQ)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORQ (SHRQ x y) (ANDQ (SHLQ x (NEGL y)) (SBBQcarrymask (CMPLconst (NEGL (ADDLconst (ANDLconst y [63]) [-64])) [64]))))
	// result: (RORQ x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHRQ {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64ANDQ {
				continue
			}
			_ = v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				v_1_0 := v_1.Args[_i1]
				if v_1_0.Op != OpAMD64SHLQ {
					continue
				}
				_ = v_1_0.Args[1]
				if x != v_1_0.Args[0] {
					continue
				}
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpAMD64NEGL || y != v_1_0_1.Args[0] {
					continue
				}
				v_1_1 := v_1.Args[1^_i1]
				if v_1_1.Op != OpAMD64SBBQcarrymask {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAMD64CMPLconst || v_1_1_0.AuxInt != 64 {
					continue
				}
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAMD64NEGL {
					continue
				}
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpAMD64ADDLconst || v_1_1_0_0_0.AuxInt != -64 {
					continue
				}
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpAMD64ANDLconst || v_1_1_0_0_0_0.AuxInt != 63 || y != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v.reset(OpAMD64RORQ)
				v.AddArg(x)
				v.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORQ x x)
	// result: x
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORQ x0:(MOVBload [i0] {s} p mem) sh:(SHLQconst [8] x1:(MOVBload [i1] {s} p mem)))
	// cond: i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVWload [i0] {s} p mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVBload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 8 {
				continue
			}
			x1 := sh.Args[0]
			if x1.Op != OpAMD64MOVBload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpAMD64MOVWload, typ.UInt16)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORQ_10(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORQ x0:(MOVWload [i0] {s} p mem) sh:(SHLQconst [16] x1:(MOVWload [i1] {s} p mem)))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVLload [i0] {s} p mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVWload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 16 {
				continue
			}
			x1 := sh.Args[0]
			if x1.Op != OpAMD64MOVWload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpAMD64MOVLload, typ.UInt32)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (ORQ x0:(MOVLload [i0] {s} p mem) sh:(SHLQconst [32] x1:(MOVLload [i1] {s} p mem)))
	// cond: i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVQload [i0] {s} p mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVLload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 32 {
				continue
			}
			x1 := sh.Args[0]
			if x1.Op != OpAMD64MOVLload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpAMD64MOVQload, typ.UInt64)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (ORQ s1:(SHLQconst [j1] x1:(MOVBload [i1] {s} p mem)) or:(ORQ s0:(SHLQconst [j0] x0:(MOVBload [i0] {s} p mem)) y))
	// cond: i1 == i0+1 && j1 == j0+8 && j0 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORQ <v.Type> (SHLQconst <v.Type> [j0] (MOVWload [i0] {s} p mem)) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s1 := v.Args[_i0]
			if s1.Op != OpAMD64SHLQconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpAMD64MOVBload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			or := v.Args[1^_i0]
			if or.Op != OpAMD64ORQ {
				continue
			}
			_ = or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				s0 := or.Args[_i1]
				if s0.Op != OpAMD64SHLQconst {
					continue
				}
				j0 := s0.AuxInt
				x0 := s0.Args[0]
				if x0.Op != OpAMD64MOVBload {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[1]
				if p != x0.Args[0] || mem != x0.Args[1] {
					continue
				}
				y := or.Args[1^_i1]
				if !(i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x0.Pos, OpAMD64ORQ, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x0.Pos, OpAMD64SHLQconst, v.Type)
				v1.AuxInt = j0
				v2 := b.NewValue0(x0.Pos, OpAMD64MOVWload, typ.UInt16)
				v2.AuxInt = i0
				v2.Aux = s
				v2.AddArg(p)
				v2.AddArg(mem)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORQ s1:(SHLQconst [j1] x1:(MOVWload [i1] {s} p mem)) or:(ORQ s0:(SHLQconst [j0] x0:(MOVWload [i0] {s} p mem)) y))
	// cond: i1 == i0+2 && j1 == j0+16 && j0 % 32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORQ <v.Type> (SHLQconst <v.Type> [j0] (MOVLload [i0] {s} p mem)) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s1 := v.Args[_i0]
			if s1.Op != OpAMD64SHLQconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpAMD64MOVWload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			or := v.Args[1^_i0]
			if or.Op != OpAMD64ORQ {
				continue
			}
			_ = or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				s0 := or.Args[_i1]
				if s0.Op != OpAMD64SHLQconst {
					continue
				}
				j0 := s0.AuxInt
				x0 := s0.Args[0]
				if x0.Op != OpAMD64MOVWload {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[1]
				if p != x0.Args[0] || mem != x0.Args[1] {
					continue
				}
				y := or.Args[1^_i1]
				if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x0.Pos, OpAMD64ORQ, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x0.Pos, OpAMD64SHLQconst, v.Type)
				v1.AuxInt = j0
				v2 := b.NewValue0(x0.Pos, OpAMD64MOVLload, typ.UInt32)
				v2.AuxInt = i0
				v2.Aux = s
				v2.AddArg(p)
				v2.AddArg(mem)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORQ x0:(MOVBloadidx1 [i0] {s} p idx mem) sh:(SHLQconst [8] x1:(MOVBloadidx1 [i1] {s} p idx mem)))
	// cond: i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVWloadidx1 <v.Type> [i0] {s} p idx mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVBloadidx1 {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x0.Args[_i1]
				idx := x0.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 8 {
					continue
				}
				x1 := sh.Args[0]
				if x1.Op != OpAMD64MOVBloadidx1 {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x1.Args[_i2] || idx != x1.Args[1^_i2] || mem != x1.Args[2] || !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, v.Type)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (ORQ x0:(MOVWloadidx1 [i0] {s} p idx mem) sh:(SHLQconst [16] x1:(MOVWloadidx1 [i1] {s} p idx mem)))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVLloadidx1 [i0] {s} p idx mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVWloadidx1 {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x0.Args[_i1]
				idx := x0.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 16 {
					continue
				}
				x1 := sh.Args[0]
				if x1.Op != OpAMD64MOVWloadidx1 {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x1.Args[_i2] || idx != x1.Args[1^_i2] || mem != x1.Args[2] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (ORQ x0:(MOVLloadidx1 [i0] {s} p idx mem) sh:(SHLQconst [32] x1:(MOVLloadidx1 [i1] {s} p idx mem)))
	// cond: i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (MOVQloadidx1 [i0] {s} p idx mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x0 := v.Args[_i0]
			if x0.Op != OpAMD64MOVLloadidx1 {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x0.Args[_i1]
				idx := x0.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 32 {
					continue
				}
				x1 := sh.Args[0]
				if x1.Op != OpAMD64MOVLloadidx1 {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x1.Args[_i2] || idx != x1.Args[1^_i2] || mem != x1.Args[2] || !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64MOVQloadidx1, typ.UInt64)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = i0
					v0.Aux = s
					v0.AddArg(p)
					v0.AddArg(idx)
					v0.AddArg(mem)
					return true
				}
			}
		}
		break
	}
	// match: (ORQ s1:(SHLQconst [j1] x1:(MOVBloadidx1 [i1] {s} p idx mem)) or:(ORQ s0:(SHLQconst [j0] x0:(MOVBloadidx1 [i0] {s} p idx mem)) y))
	// cond: i1 == i0+1 && j1 == j0+8 && j0 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORQ <v.Type> (SHLQconst <v.Type> [j0] (MOVWloadidx1 [i0] {s} p idx mem)) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s1 := v.Args[_i0]
			if s1.Op != OpAMD64SHLQconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpAMD64MOVBloadidx1 {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x1.Args[_i1]
				idx := x1.Args[1^_i1]
				or := v.Args[1^_i0]
				if or.Op != OpAMD64ORQ {
					continue
				}
				_ = or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2++ {
					s0 := or.Args[_i2]
					if s0.Op != OpAMD64SHLQconst {
						continue
					}
					j0 := s0.AuxInt
					x0 := s0.Args[0]
					if x0.Op != OpAMD64MOVBloadidx1 {
						continue
					}
					i0 := x0.AuxInt
					if x0.Aux != s {
						continue
					}
					_ = x0.Args[2]
					for _i3 := 0; _i3 <= 1; _i3++ {
						if p != x0.Args[_i3] || idx != x0.Args[1^_i3] || mem != x0.Args[2] {
							continue
						}
						y := or.Args[1^_i2]
						if !(i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpAMD64ORQ, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
						v1.AuxInt = j0
						v2 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, typ.UInt16)
						v2.AuxInt = i0
						v2.Aux = s
						v2.AddArg(p)
						v2.AddArg(idx)
						v2.AddArg(mem)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (ORQ s1:(SHLQconst [j1] x1:(MOVWloadidx1 [i1] {s} p idx mem)) or:(ORQ s0:(SHLQconst [j0] x0:(MOVWloadidx1 [i0] {s} p idx mem)) y))
	// cond: i1 == i0+2 && j1 == j0+16 && j0 % 32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORQ <v.Type> (SHLQconst <v.Type> [j0] (MOVLloadidx1 [i0] {s} p idx mem)) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s1 := v.Args[_i0]
			if s1.Op != OpAMD64SHLQconst {
				continue
			}
			j1 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpAMD64MOVWloadidx1 {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x1.Args[_i1]
				idx := x1.Args[1^_i1]
				or := v.Args[1^_i0]
				if or.Op != OpAMD64ORQ {
					continue
				}
				_ = or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2++ {
					s0 := or.Args[_i2]
					if s0.Op != OpAMD64SHLQconst {
						continue
					}
					j0 := s0.AuxInt
					x0 := s0.Args[0]
					if x0.Op != OpAMD64MOVWloadidx1 {
						continue
					}
					i0 := x0.AuxInt
					if x0.Aux != s {
						continue
					}
					_ = x0.Args[2]
					for _i3 := 0; _i3 <= 1; _i3++ {
						if p != x0.Args[_i3] || idx != x0.Args[1^_i3] || mem != x0.Args[2] {
							continue
						}
						y := or.Args[1^_i2]
						if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpAMD64ORQ, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
						v1.AuxInt = j0
						v2 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
						v2.AuxInt = i0
						v2.Aux = s
						v2.AddArg(p)
						v2.AddArg(idx)
						v2.AddArg(mem)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (ORQ x1:(MOVBload [i1] {s} p mem) sh:(SHLQconst [8] x0:(MOVBload [i0] {s} p mem)))
	// cond: i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (ROLWconst <v.Type> [8] (MOVWload [i0] {s} p mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x1 := v.Args[_i0]
			if x1.Op != OpAMD64MOVBload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 8 {
				continue
			}
			x0 := sh.Args[0]
			if x0.Op != OpAMD64MOVBload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpAMD64ROLWconst, v.Type)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = 8
			v1 := b.NewValue0(x0.Pos, OpAMD64MOVWload, typ.UInt16)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORQ_20(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORQ r1:(ROLWconst [8] x1:(MOVWload [i1] {s} p mem)) sh:(SHLQconst [16] r0:(ROLWconst [8] x0:(MOVWload [i0] {s} p mem))))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (BSWAPL <v.Type> (MOVLload [i0] {s} p mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			r1 := v.Args[_i0]
			if r1.Op != OpAMD64ROLWconst || r1.AuxInt != 8 {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpAMD64MOVWload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 16 {
				continue
			}
			r0 := sh.Args[0]
			if r0.Op != OpAMD64ROLWconst || r0.AuxInt != 8 {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpAMD64MOVWload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpAMD64BSWAPL, v.Type)
			v.reset(OpCopy)
			v.AddArg(v0)
			v1 := b.NewValue0(x0.Pos, OpAMD64MOVLload, typ.UInt32)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (ORQ r1:(BSWAPL x1:(MOVLload [i1] {s} p mem)) sh:(SHLQconst [32] r0:(BSWAPL x0:(MOVLload [i0] {s} p mem))))
	// cond: i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (BSWAPQ <v.Type> (MOVQload [i0] {s} p mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			r1 := v.Args[_i0]
			if r1.Op != OpAMD64BSWAPL {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpAMD64MOVLload {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[1]
			p := x1.Args[0]
			sh := v.Args[1^_i0]
			if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 32 {
				continue
			}
			r0 := sh.Args[0]
			if r0.Op != OpAMD64BSWAPL {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpAMD64MOVLload {
				continue
			}
			i0 := x0.AuxInt
			if x0.Aux != s {
				continue
			}
			_ = x0.Args[1]
			if p != x0.Args[0] || mem != x0.Args[1] || !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x0.Pos, OpAMD64BSWAPQ, v.Type)
			v.reset(OpCopy)
			v.AddArg(v0)
			v1 := b.NewValue0(x0.Pos, OpAMD64MOVQload, typ.UInt64)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (ORQ s0:(SHLQconst [j0] x0:(MOVBload [i0] {s} p mem)) or:(ORQ s1:(SHLQconst [j1] x1:(MOVBload [i1] {s} p mem)) y))
	// cond: i1 == i0+1 && j1 == j0-8 && j1 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORQ <v.Type> (SHLQconst <v.Type> [j1] (ROLWconst <typ.UInt16> [8] (MOVWload [i0] {s} p mem))) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s0 := v.Args[_i0]
			if s0.Op != OpAMD64SHLQconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpAMD64MOVBload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			or := v.Args[1^_i0]
			if or.Op != OpAMD64ORQ {
				continue
			}
			_ = or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				s1 := or.Args[_i1]
				if s1.Op != OpAMD64SHLQconst {
					continue
				}
				j1 := s1.AuxInt
				x1 := s1.Args[0]
				if x1.Op != OpAMD64MOVBload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				y := or.Args[1^_i1]
				if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x1.Pos, OpAMD64ORQ, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x1.Pos, OpAMD64SHLQconst, v.Type)
				v1.AuxInt = j1
				v2 := b.NewValue0(x1.Pos, OpAMD64ROLWconst, typ.UInt16)
				v2.AuxInt = 8
				v3 := b.NewValue0(x1.Pos, OpAMD64MOVWload, typ.UInt16)
				v3.AuxInt = i0
				v3.Aux = s
				v3.AddArg(p)
				v3.AddArg(mem)
				v2.AddArg(v3)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORQ s0:(SHLQconst [j0] r0:(ROLWconst [8] x0:(MOVWload [i0] {s} p mem))) or:(ORQ s1:(SHLQconst [j1] r1:(ROLWconst [8] x1:(MOVWload [i1] {s} p mem))) y))
	// cond: i1 == i0+2 && j1 == j0-16 && j1 % 32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORQ <v.Type> (SHLQconst <v.Type> [j1] (BSWAPL <typ.UInt32> (MOVLload [i0] {s} p mem))) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s0 := v.Args[_i0]
			if s0.Op != OpAMD64SHLQconst {
				continue
			}
			j0 := s0.AuxInt
			r0 := s0.Args[0]
			if r0.Op != OpAMD64ROLWconst || r0.AuxInt != 8 {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpAMD64MOVWload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			or := v.Args[1^_i0]
			if or.Op != OpAMD64ORQ {
				continue
			}
			_ = or.Args[1]
			for _i1 := 0; _i1 <= 1; _i1++ {
				s1 := or.Args[_i1]
				if s1.Op != OpAMD64SHLQconst {
					continue
				}
				j1 := s1.AuxInt
				r1 := s1.Args[0]
				if r1.Op != OpAMD64ROLWconst || r1.AuxInt != 8 {
					continue
				}
				x1 := r1.Args[0]
				if x1.Op != OpAMD64MOVWload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				y := or.Args[1^_i1]
				if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
					continue
				}
				b = mergePoint(b, x0, x1, y)
				v0 := b.NewValue0(x1.Pos, OpAMD64ORQ, v.Type)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x1.Pos, OpAMD64SHLQconst, v.Type)
				v1.AuxInt = j1
				v2 := b.NewValue0(x1.Pos, OpAMD64BSWAPL, typ.UInt32)
				v3 := b.NewValue0(x1.Pos, OpAMD64MOVLload, typ.UInt32)
				v3.AuxInt = i0
				v3.Aux = s
				v3.AddArg(p)
				v3.AddArg(mem)
				v2.AddArg(v3)
				v1.AddArg(v2)
				v0.AddArg(v1)
				v0.AddArg(y)
				return true
			}
		}
		break
	}
	// match: (ORQ x1:(MOVBloadidx1 [i1] {s} p idx mem) sh:(SHLQconst [8] x0:(MOVBloadidx1 [i0] {s} p idx mem)))
	// cond: i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (ROLWconst <v.Type> [8] (MOVWloadidx1 [i0] {s} p idx mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x1 := v.Args[_i0]
			if x1.Op != OpAMD64MOVBloadidx1 {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x1.Args[_i1]
				idx := x1.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 8 {
					continue
				}
				x0 := sh.Args[0]
				if x0.Op != OpAMD64MOVBloadidx1 {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x0.Args[_i2] || idx != x0.Args[1^_i2] || mem != x0.Args[2] || !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64ROLWconst, v.Type)
					v.reset(OpCopy)
					v.AddArg(v0)
					v0.AuxInt = 8
					v1 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, typ.UInt16)
					v1.AuxInt = i0
					v1.Aux = s
					v1.AddArg(p)
					v1.AddArg(idx)
					v1.AddArg(mem)
					v0.AddArg(v1)
					return true
				}
			}
		}
		break
	}
	// match: (ORQ r1:(ROLWconst [8] x1:(MOVWloadidx1 [i1] {s} p idx mem)) sh:(SHLQconst [16] r0:(ROLWconst [8] x0:(MOVWloadidx1 [i0] {s} p idx mem))))
	// cond: i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (BSWAPL <v.Type> (MOVLloadidx1 [i0] {s} p idx mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			r1 := v.Args[_i0]
			if r1.Op != OpAMD64ROLWconst || r1.AuxInt != 8 {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpAMD64MOVWloadidx1 {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x1.Args[_i1]
				idx := x1.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 16 {
					continue
				}
				r0 := sh.Args[0]
				if r0.Op != OpAMD64ROLWconst || r0.AuxInt != 8 {
					continue
				}
				x0 := r0.Args[0]
				if x0.Op != OpAMD64MOVWloadidx1 {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x0.Args[_i2] || idx != x0.Args[1^_i2] || mem != x0.Args[2] || !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64BSWAPL, v.Type)
					v.reset(OpCopy)
					v.AddArg(v0)
					v1 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
					v1.AuxInt = i0
					v1.Aux = s
					v1.AddArg(p)
					v1.AddArg(idx)
					v1.AddArg(mem)
					v0.AddArg(v1)
					return true
				}
			}
		}
		break
	}
	// match: (ORQ r1:(BSWAPL x1:(MOVLloadidx1 [i1] {s} p idx mem)) sh:(SHLQconst [32] r0:(BSWAPL x0:(MOVLloadidx1 [i0] {s} p idx mem))))
	// cond: i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b,x0,x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)
	// result: @mergePoint(b,x0,x1) (BSWAPQ <v.Type> (MOVQloadidx1 [i0] {s} p idx mem))
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			r1 := v.Args[_i0]
			if r1.Op != OpAMD64BSWAPL {
				continue
			}
			x1 := r1.Args[0]
			if x1.Op != OpAMD64MOVLloadidx1 {
				continue
			}
			i1 := x1.AuxInt
			s := x1.Aux
			mem := x1.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x1.Args[_i1]
				idx := x1.Args[1^_i1]
				sh := v.Args[1^_i0]
				if sh.Op != OpAMD64SHLQconst || sh.AuxInt != 32 {
					continue
				}
				r0 := sh.Args[0]
				if r0.Op != OpAMD64BSWAPL {
					continue
				}
				x0 := r0.Args[0]
				if x0.Op != OpAMD64MOVLloadidx1 {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[2]
				for _i2 := 0; _i2 <= 1; _i2++ {
					if p != x0.Args[_i2] || idx != x0.Args[1^_i2] || mem != x0.Args[2] || !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
						continue
					}
					b = mergePoint(b, x0, x1)
					v0 := b.NewValue0(v.Pos, OpAMD64BSWAPQ, v.Type)
					v.reset(OpCopy)
					v.AddArg(v0)
					v1 := b.NewValue0(v.Pos, OpAMD64MOVQloadidx1, typ.UInt64)
					v1.AuxInt = i0
					v1.Aux = s
					v1.AddArg(p)
					v1.AddArg(idx)
					v1.AddArg(mem)
					v0.AddArg(v1)
					return true
				}
			}
		}
		break
	}
	// match: (ORQ s0:(SHLQconst [j0] x0:(MOVBloadidx1 [i0] {s} p idx mem)) or:(ORQ s1:(SHLQconst [j1] x1:(MOVBloadidx1 [i1] {s} p idx mem)) y))
	// cond: i1 == i0+1 && j1 == j0-8 && j1 % 16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORQ <v.Type> (SHLQconst <v.Type> [j1] (ROLWconst <typ.UInt16> [8] (MOVWloadidx1 [i0] {s} p idx mem))) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s0 := v.Args[_i0]
			if s0.Op != OpAMD64SHLQconst {
				continue
			}
			j0 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpAMD64MOVBloadidx1 {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x0.Args[_i1]
				idx := x0.Args[1^_i1]
				or := v.Args[1^_i0]
				if or.Op != OpAMD64ORQ {
					continue
				}
				_ = or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2++ {
					s1 := or.Args[_i2]
					if s1.Op != OpAMD64SHLQconst {
						continue
					}
					j1 := s1.AuxInt
					x1 := s1.Args[0]
					if x1.Op != OpAMD64MOVBloadidx1 {
						continue
					}
					i1 := x1.AuxInt
					if x1.Aux != s {
						continue
					}
					_ = x1.Args[2]
					for _i3 := 0; _i3 <= 1; _i3++ {
						if p != x1.Args[_i3] || idx != x1.Args[1^_i3] || mem != x1.Args[2] {
							continue
						}
						y := or.Args[1^_i2]
						if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpAMD64ORQ, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
						v1.AuxInt = j1
						v2 := b.NewValue0(v.Pos, OpAMD64ROLWconst, typ.UInt16)
						v2.AuxInt = 8
						v3 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, typ.UInt16)
						v3.AuxInt = i0
						v3.Aux = s
						v3.AddArg(p)
						v3.AddArg(idx)
						v3.AddArg(mem)
						v2.AddArg(v3)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (ORQ s0:(SHLQconst [j0] r0:(ROLWconst [8] x0:(MOVWloadidx1 [i0] {s} p idx mem))) or:(ORQ s1:(SHLQconst [j1] r1:(ROLWconst [8] x1:(MOVWloadidx1 [i1] {s} p idx mem))) y))
	// cond: i1 == i0+2 && j1 == j0-16 && j1 % 32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b,x0,x1,y) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)
	// result: @mergePoint(b,x0,x1,y) (ORQ <v.Type> (SHLQconst <v.Type> [j1] (BSWAPL <typ.UInt32> (MOVLloadidx1 [i0] {s} p idx mem))) y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			s0 := v.Args[_i0]
			if s0.Op != OpAMD64SHLQconst {
				continue
			}
			j0 := s0.AuxInt
			r0 := s0.Args[0]
			if r0.Op != OpAMD64ROLWconst || r0.AuxInt != 8 {
				continue
			}
			x0 := r0.Args[0]
			if x0.Op != OpAMD64MOVWloadidx1 {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[2]
			for _i1 := 0; _i1 <= 1; _i1++ {
				p := x0.Args[_i1]
				idx := x0.Args[1^_i1]
				or := v.Args[1^_i0]
				if or.Op != OpAMD64ORQ {
					continue
				}
				_ = or.Args[1]
				for _i2 := 0; _i2 <= 1; _i2++ {
					s1 := or.Args[_i2]
					if s1.Op != OpAMD64SHLQconst {
						continue
					}
					j1 := s1.AuxInt
					r1 := s1.Args[0]
					if r1.Op != OpAMD64ROLWconst || r1.AuxInt != 8 {
						continue
					}
					x1 := r1.Args[0]
					if x1.Op != OpAMD64MOVWloadidx1 {
						continue
					}
					i1 := x1.AuxInt
					if x1.Aux != s {
						continue
					}
					_ = x1.Args[2]
					for _i3 := 0; _i3 <= 1; _i3++ {
						if p != x1.Args[_i3] || idx != x1.Args[1^_i3] || mem != x1.Args[2] {
							continue
						}
						y := or.Args[1^_i2]
						if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1, y) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
							continue
						}
						b = mergePoint(b, x0, x1, y)
						v0 := b.NewValue0(v.Pos, OpAMD64ORQ, v.Type)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(v.Pos, OpAMD64SHLQconst, v.Type)
						v1.AuxInt = j1
						v2 := b.NewValue0(v.Pos, OpAMD64BSWAPL, typ.UInt32)
						v3 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
						v3.AuxInt = i0
						v3.Aux = s
						v3.AddArg(p)
						v3.AddArg(idx)
						v3.AddArg(mem)
						v2.AddArg(v3)
						v1.AddArg(v2)
						v0.AddArg(v1)
						v0.AddArg(y)
						return true
					}
				}
			}
		}
		break
	}
	// match: (ORQ x l:(MOVQload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (ORQload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVQload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64ORQload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORQconst_0(v *Value) bool {
	// match: (ORQconst [c] x)
	// cond: isUint64PowerOfTwo(c) && uint64(c) >= 128
	// result: (BTSQconst [log2(c)] x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isUint64PowerOfTwo(c) && uint64(c) >= 128) {
			break
		}
		v.reset(OpAMD64BTSQconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (ORQconst [c] (ORQconst [d] x))
	// result: (ORQconst [c | d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ORQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ORQconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	// match: (ORQconst [c] (BTSQconst [d] x))
	// result: (ORQconst [c | 1<<uint32(d)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTSQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ORQconst)
		v.AuxInt = c | 1<<uint32(d)
		v.AddArg(x)
		return true
	}
	// match: (ORQconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORQconst [-1] _)
	// result: (MOVQconst [-1])
	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [c|d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = c | d
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORQconstmodify_0(v *Value) bool {
	// match: (ORQconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (ORQconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64ORQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ORQconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (ORQconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ORQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORQload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORQload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (ORQload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ORQload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (ORQload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ORQload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ORQload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: ( ORQload x [off] {sym} ptr (MOVSDstore [off] {sym} ptr y _))
	// result: ( ORQ x (MOVQf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64ORQ)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVQf2i, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ORQmodify_0(v *Value) bool {
	// match: (ORQmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (ORQmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64ORQmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (ORQmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (ORQmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64ORQmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ROLB_0(v *Value) bool {
	// match: (ROLB x (NEGQ y))
	// result: (RORB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64RORB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ROLB x (NEGL y))
	// result: (RORB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64RORB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ROLB x (MOVQconst [c]))
	// result: (ROLBconst [c&7 ] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLBconst)
		v.AuxInt = c & 7
		v.AddArg(x)
		return true
	}
	// match: (ROLB x (MOVLconst [c]))
	// result: (ROLBconst [c&7 ] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLBconst)
		v.AuxInt = c & 7
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ROLBconst_0(v *Value) bool {
	// match: (ROLBconst [c] (ROLBconst [d] x))
	// result: (ROLBconst [(c+d)& 7] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ROLBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ROLBconst)
		v.AuxInt = (c + d) & 7
		v.AddArg(x)
		return true
	}
	// match: (ROLBconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ROLL_0(v *Value) bool {
	// match: (ROLL x (NEGQ y))
	// result: (RORL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64RORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ROLL x (NEGL y))
	// result: (RORL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64RORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ROLL x (MOVQconst [c]))
	// result: (ROLLconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (ROLL x (MOVLconst [c]))
	// result: (ROLLconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ROLLconst_0(v *Value) bool {
	// match: (ROLLconst [c] (ROLLconst [d] x))
	// result: (ROLLconst [(c+d)&31] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ROLLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ROLLconst)
		v.AuxInt = (c + d) & 31
		v.AddArg(x)
		return true
	}
	// match: (ROLLconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ROLQ_0(v *Value) bool {
	// match: (ROLQ x (NEGQ y))
	// result: (RORQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64RORQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ROLQ x (NEGL y))
	// result: (RORQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64RORQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ROLQ x (MOVQconst [c]))
	// result: (ROLQconst [c&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLQconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (ROLQ x (MOVLconst [c]))
	// result: (ROLQconst [c&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLQconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ROLQconst_0(v *Value) bool {
	// match: (ROLQconst [c] (ROLQconst [d] x))
	// result: (ROLQconst [(c+d)&63] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ROLQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ROLQconst)
		v.AuxInt = (c + d) & 63
		v.AddArg(x)
		return true
	}
	// match: (ROLQconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ROLW_0(v *Value) bool {
	// match: (ROLW x (NEGQ y))
	// result: (RORW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64RORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ROLW x (NEGL y))
	// result: (RORW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64RORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ROLW x (MOVQconst [c]))
	// result: (ROLWconst [c&15] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLWconst)
		v.AuxInt = c & 15
		v.AddArg(x)
		return true
	}
	// match: (ROLW x (MOVLconst [c]))
	// result: (ROLWconst [c&15] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLWconst)
		v.AuxInt = c & 15
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64ROLWconst_0(v *Value) bool {
	// match: (ROLWconst [c] (ROLWconst [d] x))
	// result: (ROLWconst [(c+d)&15] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ROLWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64ROLWconst)
		v.AuxInt = (c + d) & 15
		v.AddArg(x)
		return true
	}
	// match: (ROLWconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64RORB_0(v *Value) bool {
	// match: (RORB x (NEGQ y))
	// result: (ROLB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64ROLB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RORB x (NEGL y))
	// result: (ROLB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64ROLB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RORB x (MOVQconst [c]))
	// result: (ROLBconst [(-c)&7 ] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLBconst)
		v.AuxInt = (-c) & 7
		v.AddArg(x)
		return true
	}
	// match: (RORB x (MOVLconst [c]))
	// result: (ROLBconst [(-c)&7 ] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLBconst)
		v.AuxInt = (-c) & 7
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64RORL_0(v *Value) bool {
	// match: (RORL x (NEGQ y))
	// result: (ROLL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64ROLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RORL x (NEGL y))
	// result: (ROLL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64ROLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RORL x (MOVQconst [c]))
	// result: (ROLLconst [(-c)&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLLconst)
		v.AuxInt = (-c) & 31
		v.AddArg(x)
		return true
	}
	// match: (RORL x (MOVLconst [c]))
	// result: (ROLLconst [(-c)&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLLconst)
		v.AuxInt = (-c) & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64RORQ_0(v *Value) bool {
	// match: (RORQ x (NEGQ y))
	// result: (ROLQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64ROLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RORQ x (NEGL y))
	// result: (ROLQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64ROLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RORQ x (MOVQconst [c]))
	// result: (ROLQconst [(-c)&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLQconst)
		v.AuxInt = (-c) & 63
		v.AddArg(x)
		return true
	}
	// match: (RORQ x (MOVLconst [c]))
	// result: (ROLQconst [(-c)&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLQconst)
		v.AuxInt = (-c) & 63
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64RORW_0(v *Value) bool {
	// match: (RORW x (NEGQ y))
	// result: (ROLW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64ROLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RORW x (NEGL y))
	// result: (ROLW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(OpAMD64ROLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RORW x (MOVQconst [c]))
	// result: (ROLWconst [(-c)&15] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLWconst)
		v.AuxInt = (-c) & 15
		v.AddArg(x)
		return true
	}
	// match: (RORW x (MOVLconst [c]))
	// result: (ROLWconst [(-c)&15] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64ROLWconst)
		v.AuxInt = (-c) & 15
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SARB_0(v *Value) bool {
	// match: (SARB x (MOVQconst [c]))
	// result: (SARBconst [min(c&31,7)] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SARBconst)
		v.AuxInt = min(c&31, 7)
		v.AddArg(x)
		return true
	}
	// match: (SARB x (MOVLconst [c]))
	// result: (SARBconst [min(c&31,7)] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SARBconst)
		v.AuxInt = min(c&31, 7)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SARBconst_0(v *Value) bool {
	// match: (SARBconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SARBconst [c] (MOVQconst [d]))
	// result: (MOVQconst [int64(int8(d))>>uint64(c)])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = int64(int8(d)) >> uint64(c)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SARL_0(v *Value) bool {
	b := v.Block
	// match: (SARL x (MOVQconst [c]))
	// result: (SARLconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SARLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SARL x (MOVLconst [c]))
	// result: (SARLconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SARLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SARL x (ADDQconst [c] y))
	// cond: c & 31 == 0
	// result: (SARL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SARL x (NEGQ <t> (ADDQconst [c] y)))
	// cond: c & 31 == 0
	// result: (SARL x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SARL x (ANDQconst [c] y))
	// cond: c & 31 == 31
	// result: (SARL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SARL x (NEGQ <t> (ANDQconst [c] y)))
	// cond: c & 31 == 31
	// result: (SARL x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SARL x (ADDLconst [c] y))
	// cond: c & 31 == 0
	// result: (SARL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SARL x (NEGL <t> (ADDLconst [c] y)))
	// cond: c & 31 == 0
	// result: (SARL x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SARL x (ANDLconst [c] y))
	// cond: c & 31 == 31
	// result: (SARL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SARL x (NEGL <t> (ANDLconst [c] y)))
	// cond: c & 31 == 31
	// result: (SARL x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SARLconst_0(v *Value) bool {
	// match: (SARLconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SARLconst [c] (MOVQconst [d]))
	// result: (MOVQconst [int64(int32(d))>>uint64(c)])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = int64(int32(d)) >> uint64(c)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SARQ_0(v *Value) bool {
	b := v.Block
	// match: (SARQ x (MOVQconst [c]))
	// result: (SARQconst [c&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SARQconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SARQ x (MOVLconst [c]))
	// result: (SARQconst [c&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SARQconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SARQ x (ADDQconst [c] y))
	// cond: c & 63 == 0
	// result: (SARQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SARQ x (NEGQ <t> (ADDQconst [c] y)))
	// cond: c & 63 == 0
	// result: (SARQ x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SARQ x (ANDQconst [c] y))
	// cond: c & 63 == 63
	// result: (SARQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SARQ x (NEGQ <t> (ANDQconst [c] y)))
	// cond: c & 63 == 63
	// result: (SARQ x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SARQ x (ADDLconst [c] y))
	// cond: c & 63 == 0
	// result: (SARQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SARQ x (NEGL <t> (ADDLconst [c] y)))
	// cond: c & 63 == 0
	// result: (SARQ x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SARQ x (ANDLconst [c] y))
	// cond: c & 63 == 63
	// result: (SARQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SARQ x (NEGL <t> (ANDLconst [c] y)))
	// cond: c & 63 == 63
	// result: (SARQ x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SARQconst_0(v *Value) bool {
	// match: (SARQconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SARQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [d>>uint64(c)])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = d >> uint64(c)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SARW_0(v *Value) bool {
	// match: (SARW x (MOVQconst [c]))
	// result: (SARWconst [min(c&31,15)] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SARWconst)
		v.AuxInt = min(c&31, 15)
		v.AddArg(x)
		return true
	}
	// match: (SARW x (MOVLconst [c]))
	// result: (SARWconst [min(c&31,15)] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SARWconst)
		v.AuxInt = min(c&31, 15)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SARWconst_0(v *Value) bool {
	// match: (SARWconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SARWconst [c] (MOVQconst [d]))
	// result: (MOVQconst [int64(int16(d))>>uint64(c)])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = int64(int16(d)) >> uint64(c)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SBBLcarrymask_0(v *Value) bool {
	// match: (SBBLcarrymask (FlagEQ))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SBBLcarrymask (FlagLT_ULT))
	// result: (MOVLconst [-1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = -1
		return true
	}
	// match: (SBBLcarrymask (FlagLT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SBBLcarrymask (FlagGT_ULT))
	// result: (MOVLconst [-1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = -1
		return true
	}
	// match: (SBBLcarrymask (FlagGT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SBBQ_0(v *Value) bool {
	// match: (SBBQ x (MOVQconst [c]) borrow)
	// cond: is32Bit(c)
	// result: (SBBQconst x [c] borrow)
	for {
		borrow := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpAMD64SBBQconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(borrow)
		return true
	}
	// match: (SBBQ x y (FlagEQ))
	// result: (SUBQborrow x y)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64SUBQborrow)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SBBQcarrymask_0(v *Value) bool {
	// match: (SBBQcarrymask (FlagEQ))
	// result: (MOVQconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = 0
		return true
	}
	// match: (SBBQcarrymask (FlagLT_ULT))
	// result: (MOVQconst [-1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = -1
		return true
	}
	// match: (SBBQcarrymask (FlagLT_UGT))
	// result: (MOVQconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = 0
		return true
	}
	// match: (SBBQcarrymask (FlagGT_ULT))
	// result: (MOVQconst [-1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = -1
		return true
	}
	// match: (SBBQcarrymask (FlagGT_UGT))
	// result: (MOVQconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SBBQconst_0(v *Value) bool {
	// match: (SBBQconst x [c] (FlagEQ))
	// result: (SUBQconstborrow x [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64SUBQconstborrow)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETA_0(v *Value) bool {
	// match: (SETA (InvertFlags x))
	// result: (SETB x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETB)
		v.AddArg(x)
		return true
	}
	// match: (SETA (FlagEQ))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETA (FlagLT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETA (FlagLT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETA (FlagGT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETA (FlagGT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETAE_0(v *Value) bool {
	// match: (SETAE (InvertFlags x))
	// result: (SETBE x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETBE)
		v.AddArg(x)
		return true
	}
	// match: (SETAE (FlagEQ))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETAE (FlagLT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETAE (FlagLT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETAE (FlagGT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETAE (FlagGT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETAEstore_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETAEstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETBEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETBEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETAEstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETAEstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETAEstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETAEstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETAEstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETAEstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETAEstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETAEstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETAEstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETAEstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETAEstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETAstore_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETAstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETAstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETAstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETAstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETAstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETAstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETAstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETAstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETAstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETAstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETAstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETAstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETB_0(v *Value) bool {
	// match: (SETB (InvertFlags x))
	// result: (SETA x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETA)
		v.AddArg(x)
		return true
	}
	// match: (SETB (FlagEQ))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETB (FlagLT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETB (FlagLT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETB (FlagGT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETB (FlagGT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETBE_0(v *Value) bool {
	// match: (SETBE (InvertFlags x))
	// result: (SETAE x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETAE)
		v.AddArg(x)
		return true
	}
	// match: (SETBE (FlagEQ))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETBE (FlagLT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETBE (FlagLT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETBE (FlagGT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETBE (FlagGT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETBEstore_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETBEstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETAEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETAEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETBEstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETBEstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETBEstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETBEstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETBEstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETBEstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETBEstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETBEstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETBEstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETBEstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETBEstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETBstore_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETBstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETAstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETAstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETBstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETBstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETBstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETBstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETBstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETBstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETBstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETBstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETBstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETEQ_0(v *Value) bool {
	b := v.Block
	// match: (SETEQ (TESTL (SHLL (MOVLconst [1]) x) y))
	// result: (SETAE (BTL x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0_0 := v_0.Args[_i0]
			if v_0_0.Op != OpAMD64SHLL {
				continue
			}
			x := v_0_0.Args[1]
			v_0_0_0 := v_0_0.Args[0]
			if v_0_0_0.Op != OpAMD64MOVLconst || v_0_0_0.AuxInt != 1 {
				continue
			}
			y := v_0.Args[1^_i0]
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTL, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETEQ (TESTQ (SHLQ (MOVQconst [1]) x) y))
	// result: (SETAE (BTQ x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0_0 := v_0.Args[_i0]
			if v_0_0.Op != OpAMD64SHLQ {
				continue
			}
			x := v_0_0.Args[1]
			v_0_0_0 := v_0_0.Args[0]
			if v_0_0_0.Op != OpAMD64MOVQconst || v_0_0_0.AuxInt != 1 {
				continue
			}
			y := v_0.Args[1^_i0]
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETEQ (TESTLconst [c] x))
	// cond: isUint32PowerOfTwo(c)
	// result: (SETAE (BTLconst [log2uint32(c)] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(isUint32PowerOfTwo(c)) {
			break
		}
		v.reset(OpAMD64SETAE)
		v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
		v0.AuxInt = log2uint32(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SETEQ (TESTQconst [c] x))
	// cond: isUint64PowerOfTwo(c)
	// result: (SETAE (BTQconst [log2(c)] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(isUint64PowerOfTwo(c)) {
			break
		}
		v.reset(OpAMD64SETAE)
		v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SETEQ (TESTQ (MOVQconst [c]) x))
	// cond: isUint64PowerOfTwo(c)
	// result: (SETAE (BTQconst [log2(c)] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0_0 := v_0.Args[_i0]
			if v_0_0.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_0_0.AuxInt
			x := v_0.Args[1^_i0]
			if !(isUint64PowerOfTwo(c)) {
				continue
			}
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = log2(c)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETEQ (CMPLconst [1] s:(ANDLconst [1] _)))
	// result: (SETNE (CMPLconst [0] s))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64CMPLconst || v_0.AuxInt != 1 {
			break
		}
		s := v_0.Args[0]
		if s.Op != OpAMD64ANDLconst || s.AuxInt != 1 {
			break
		}
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(s)
		v.AddArg(v0)
		return true
	}
	// match: (SETEQ (CMPQconst [1] s:(ANDQconst [1] _)))
	// result: (SETNE (CMPQconst [0] s))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64CMPQconst || v_0.AuxInt != 1 {
			break
		}
		s := v_0.Args[0]
		if s.Op != OpAMD64ANDQconst || s.AuxInt != 1 {
			break
		}
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(s)
		v.AddArg(v0)
		return true
	}
	// match: (SETEQ (TESTQ z1:(SHLQconst [63] (SHRQconst [63] x)) z2))
	// cond: z1==z2
	// result: (SETAE (BTQconst [63] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHLQconst || z1.AuxInt != 63 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 63 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 63
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETEQ (TESTL z1:(SHLLconst [31] (SHRQconst [31] x)) z2))
	// cond: z1==z2
	// result: (SETAE (BTQconst [31] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHLLconst || z1.AuxInt != 31 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 31 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 31
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETEQ (TESTQ z1:(SHRQconst [63] (SHLQconst [63] x)) z2))
	// cond: z1==z2
	// result: (SETAE (BTQconst [0] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHLQconst || z1_0.AuxInt != 63 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETEQ_10(v *Value) bool {
	b := v.Block
	// match: (SETEQ (TESTL z1:(SHRLconst [31] (SHLLconst [31] x)) z2))
	// cond: z1==z2
	// result: (SETAE (BTLconst [0] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHLLconst || z1_0.AuxInt != 31 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETEQ (TESTQ z1:(SHRQconst [63] x) z2))
	// cond: z1==z2
	// result: (SETAE (BTQconst [63] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
				continue
			}
			x := z1.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 63
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETEQ (TESTL z1:(SHRLconst [31] x) z2))
	// cond: z1==z2
	// result: (SETAE (BTLconst [31] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
				continue
			}
			x := z1.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAE)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 31
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETEQ (InvertFlags x))
	// result: (SETEQ x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETEQ)
		v.AddArg(x)
		return true
	}
	// match: (SETEQ (FlagEQ))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETEQ (FlagLT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETEQ (FlagLT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETEQ (FlagGT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETEQ (FlagGT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETEQstore_0(v *Value) bool {
	b := v.Block
	// match: (SETEQstore [off] {sym} ptr (TESTL (SHLL (MOVLconst [1]) x) y) mem)
	// result: (SETAEstore [off] {sym} ptr (BTL x y) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTL {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_1_0 := v_1.Args[_i0]
			if v_1_0.Op != OpAMD64SHLL {
				continue
			}
			x := v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAMD64MOVLconst || v_1_0_0.AuxInt != 1 {
				continue
			}
			y := v_1.Args[1^_i0]
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTL, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETEQstore [off] {sym} ptr (TESTQ (SHLQ (MOVQconst [1]) x) y) mem)
	// result: (SETAEstore [off] {sym} ptr (BTQ x y) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_1_0 := v_1.Args[_i0]
			if v_1_0.Op != OpAMD64SHLQ {
				continue
			}
			x := v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAMD64MOVQconst || v_1_0_0.AuxInt != 1 {
				continue
			}
			y := v_1.Args[1^_i0]
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETEQstore [off] {sym} ptr (TESTLconst [c] x) mem)
	// cond: isUint32PowerOfTwo(c)
	// result: (SETAEstore [off] {sym} ptr (BTLconst [log2uint32(c)] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTLconst {
			break
		}
		c := v_1.AuxInt
		x := v_1.Args[0]
		if !(isUint32PowerOfTwo(c)) {
			break
		}
		v.reset(OpAMD64SETAEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
		v0.AuxInt = log2uint32(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off] {sym} ptr (TESTQconst [c] x) mem)
	// cond: isUint64PowerOfTwo(c)
	// result: (SETAEstore [off] {sym} ptr (BTQconst [log2(c)] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQconst {
			break
		}
		c := v_1.AuxInt
		x := v_1.Args[0]
		if !(isUint64PowerOfTwo(c)) {
			break
		}
		v.reset(OpAMD64SETAEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off] {sym} ptr (TESTQ (MOVQconst [c]) x) mem)
	// cond: isUint64PowerOfTwo(c)
	// result: (SETAEstore [off] {sym} ptr (BTQconst [log2(c)] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_1_0 := v_1.Args[_i0]
			if v_1_0.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1_0.AuxInt
			x := v_1.Args[1^_i0]
			if !(isUint64PowerOfTwo(c)) {
				continue
			}
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = log2(c)
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETEQstore [off] {sym} ptr (CMPLconst [1] s:(ANDLconst [1] _)) mem)
	// result: (SETNEstore [off] {sym} ptr (CMPLconst [0] s) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64CMPLconst || v_1.AuxInt != 1 {
			break
		}
		s := v_1.Args[0]
		if s.Op != OpAMD64ANDLconst || s.AuxInt != 1 {
			break
		}
		v.reset(OpAMD64SETNEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(s)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off] {sym} ptr (CMPQconst [1] s:(ANDQconst [1] _)) mem)
	// result: (SETNEstore [off] {sym} ptr (CMPQconst [0] s) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64CMPQconst || v_1.AuxInt != 1 {
			break
		}
		s := v_1.Args[0]
		if s.Op != OpAMD64ANDQconst || s.AuxInt != 1 {
			break
		}
		v.reset(OpAMD64SETNEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(s)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off] {sym} ptr (TESTQ z1:(SHLQconst [63] (SHRQconst [63] x)) z2) mem)
	// cond: z1==z2
	// result: (SETAEstore [off] {sym} ptr (BTQconst [63] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHLQconst || z1.AuxInt != 63 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 63 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 63
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETEQstore [off] {sym} ptr (TESTL z1:(SHLLconst [31] (SHRLconst [31] x)) z2) mem)
	// cond: z1==z2
	// result: (SETAEstore [off] {sym} ptr (BTLconst [31] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTL {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHLLconst || z1.AuxInt != 31 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHRLconst || z1_0.AuxInt != 31 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 31
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETEQstore [off] {sym} ptr (TESTQ z1:(SHRQconst [63] (SHLQconst [63] x)) z2) mem)
	// cond: z1==z2
	// result: (SETAEstore [off] {sym} ptr (BTQconst [0] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHLQconst || z1_0.AuxInt != 63 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETEQstore_10(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETEQstore [off] {sym} ptr (TESTL z1:(SHRLconst [31] (SHLLconst [31] x)) z2) mem)
	// cond: z1==z2
	// result: (SETAEstore [off] {sym} ptr (BTLconst [0] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTL {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHLLconst || z1_0.AuxInt != 31 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETEQstore [off] {sym} ptr (TESTQ z1:(SHRQconst [63] x) z2) mem)
	// cond: z1==z2
	// result: (SETAEstore [off] {sym} ptr (BTQconst [63] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
				continue
			}
			x := z1.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 63
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETEQstore [off] {sym} ptr (TESTL z1:(SHRLconst [31] x) z2) mem)
	// cond: z1==z2
	// result: (SETAEstore [off] {sym} ptr (BTLconst [31] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTL {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
				continue
			}
			x := z1.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETAEstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 31
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETEQstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETEQstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETEQstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETEQstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETEQstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETEQstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETEQstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETEQstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETEQstore_20(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETEQstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETG_0(v *Value) bool {
	// match: (SETG (InvertFlags x))
	// result: (SETL x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETL)
		v.AddArg(x)
		return true
	}
	// match: (SETG (FlagEQ))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETG (FlagLT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETG (FlagLT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETG (FlagGT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETG (FlagGT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETGE_0(v *Value) bool {
	// match: (SETGE (InvertFlags x))
	// result: (SETLE x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETLE)
		v.AddArg(x)
		return true
	}
	// match: (SETGE (FlagEQ))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETGE (FlagLT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETGE (FlagLT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETGE (FlagGT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETGE (FlagGT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETGEstore_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETGEstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETLEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETLEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETGEstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETGEstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETGEstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETGEstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETGEstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETGEstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETGEstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETGEstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETGEstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETGEstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETGEstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETGstore_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETGstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETLstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETLstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETGstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETGstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETGstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETGstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETGstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETGstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETGstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETGstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETGstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETGstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETGstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETL_0(v *Value) bool {
	// match: (SETL (InvertFlags x))
	// result: (SETG x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETG)
		v.AddArg(x)
		return true
	}
	// match: (SETL (FlagEQ))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETL (FlagLT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETL (FlagLT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETL (FlagGT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETL (FlagGT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETLE_0(v *Value) bool {
	// match: (SETLE (InvertFlags x))
	// result: (SETGE x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETGE)
		v.AddArg(x)
		return true
	}
	// match: (SETLE (FlagEQ))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETLE (FlagLT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETLE (FlagLT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETLE (FlagGT_ULT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETLE (FlagGT_UGT))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETLEstore_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETLEstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETGEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETGEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETLEstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETLEstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETLEstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETLEstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETLEstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETLEstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETLEstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETLEstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETLEstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETLEstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETLEstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETLstore_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETLstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETGstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETGstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETLstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETLstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETLstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETLstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETLstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETLstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETLstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETLstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETLstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETLstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETLstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETNE_0(v *Value) bool {
	b := v.Block
	// match: (SETNE (TESTL (SHLL (MOVLconst [1]) x) y))
	// result: (SETB (BTL x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0_0 := v_0.Args[_i0]
			if v_0_0.Op != OpAMD64SHLL {
				continue
			}
			x := v_0_0.Args[1]
			v_0_0_0 := v_0_0.Args[0]
			if v_0_0_0.Op != OpAMD64MOVLconst || v_0_0_0.AuxInt != 1 {
				continue
			}
			y := v_0.Args[1^_i0]
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTL, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETNE (TESTQ (SHLQ (MOVQconst [1]) x) y))
	// result: (SETB (BTQ x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0_0 := v_0.Args[_i0]
			if v_0_0.Op != OpAMD64SHLQ {
				continue
			}
			x := v_0_0.Args[1]
			v_0_0_0 := v_0_0.Args[0]
			if v_0_0_0.Op != OpAMD64MOVQconst || v_0_0_0.AuxInt != 1 {
				continue
			}
			y := v_0.Args[1^_i0]
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETNE (TESTLconst [c] x))
	// cond: isUint32PowerOfTwo(c)
	// result: (SETB (BTLconst [log2uint32(c)] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(isUint32PowerOfTwo(c)) {
			break
		}
		v.reset(OpAMD64SETB)
		v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
		v0.AuxInt = log2uint32(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SETNE (TESTQconst [c] x))
	// cond: isUint64PowerOfTwo(c)
	// result: (SETB (BTQconst [log2(c)] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(isUint64PowerOfTwo(c)) {
			break
		}
		v.reset(OpAMD64SETB)
		v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SETNE (TESTQ (MOVQconst [c]) x))
	// cond: isUint64PowerOfTwo(c)
	// result: (SETB (BTQconst [log2(c)] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0_0 := v_0.Args[_i0]
			if v_0_0.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_0_0.AuxInt
			x := v_0.Args[1^_i0]
			if !(isUint64PowerOfTwo(c)) {
				continue
			}
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = log2(c)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETNE (CMPLconst [1] s:(ANDLconst [1] _)))
	// result: (SETEQ (CMPLconst [0] s))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64CMPLconst || v_0.AuxInt != 1 {
			break
		}
		s := v_0.Args[0]
		if s.Op != OpAMD64ANDLconst || s.AuxInt != 1 {
			break
		}
		v.reset(OpAMD64SETEQ)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(s)
		v.AddArg(v0)
		return true
	}
	// match: (SETNE (CMPQconst [1] s:(ANDQconst [1] _)))
	// result: (SETEQ (CMPQconst [0] s))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64CMPQconst || v_0.AuxInt != 1 {
			break
		}
		s := v_0.Args[0]
		if s.Op != OpAMD64ANDQconst || s.AuxInt != 1 {
			break
		}
		v.reset(OpAMD64SETEQ)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(s)
		v.AddArg(v0)
		return true
	}
	// match: (SETNE (TESTQ z1:(SHLQconst [63] (SHRQconst [63] x)) z2))
	// cond: z1==z2
	// result: (SETB (BTQconst [63] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHLQconst || z1.AuxInt != 63 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 63 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 63
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETNE (TESTL z1:(SHLLconst [31] (SHRQconst [31] x)) z2))
	// cond: z1==z2
	// result: (SETB (BTQconst [31] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHLLconst || z1.AuxInt != 31 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 31 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 31
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETNE (TESTQ z1:(SHRQconst [63] (SHLQconst [63] x)) z2))
	// cond: z1==z2
	// result: (SETB (BTQconst [0] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHLQconst || z1_0.AuxInt != 63 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETNE_10(v *Value) bool {
	b := v.Block
	// match: (SETNE (TESTL z1:(SHRLconst [31] (SHLLconst [31] x)) z2))
	// cond: z1==z2
	// result: (SETB (BTLconst [0] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHLLconst || z1_0.AuxInt != 31 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETNE (TESTQ z1:(SHRQconst [63] x) z2))
	// cond: z1==z2
	// result: (SETB (BTQconst [63] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTQ {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
				continue
			}
			x := z1.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 63
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETNE (TESTL z1:(SHRLconst [31] x) z2))
	// cond: z1==z2
	// result: (SETB (BTLconst [31] x))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64TESTL {
			break
		}
		_ = v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_0.Args[_i0]
			if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
				continue
			}
			x := z1.Args[0]
			z2 := v_0.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETB)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 31
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (SETNE (InvertFlags x))
	// result: (SETNE x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETNE)
		v.AddArg(x)
		return true
	}
	// match: (SETNE (FlagEQ))
	// result: (MOVLconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SETNE (FlagLT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETNE (FlagLT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETNE (FlagGT_ULT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	// match: (SETNE (FlagGT_UGT))
	// result: (MOVLconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETNEstore_0(v *Value) bool {
	b := v.Block
	// match: (SETNEstore [off] {sym} ptr (TESTL (SHLL (MOVLconst [1]) x) y) mem)
	// result: (SETBstore [off] {sym} ptr (BTL x y) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTL {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_1_0 := v_1.Args[_i0]
			if v_1_0.Op != OpAMD64SHLL {
				continue
			}
			x := v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAMD64MOVLconst || v_1_0_0.AuxInt != 1 {
				continue
			}
			y := v_1.Args[1^_i0]
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTL, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETNEstore [off] {sym} ptr (TESTQ (SHLQ (MOVQconst [1]) x) y) mem)
	// result: (SETBstore [off] {sym} ptr (BTQ x y) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_1_0 := v_1.Args[_i0]
			if v_1_0.Op != OpAMD64SHLQ {
				continue
			}
			x := v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAMD64MOVQconst || v_1_0_0.AuxInt != 1 {
				continue
			}
			y := v_1.Args[1^_i0]
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETNEstore [off] {sym} ptr (TESTLconst [c] x) mem)
	// cond: isUint32PowerOfTwo(c)
	// result: (SETBstore [off] {sym} ptr (BTLconst [log2uint32(c)] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTLconst {
			break
		}
		c := v_1.AuxInt
		x := v_1.Args[0]
		if !(isUint32PowerOfTwo(c)) {
			break
		}
		v.reset(OpAMD64SETBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
		v0.AuxInt = log2uint32(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off] {sym} ptr (TESTQconst [c] x) mem)
	// cond: isUint64PowerOfTwo(c)
	// result: (SETBstore [off] {sym} ptr (BTQconst [log2(c)] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQconst {
			break
		}
		c := v_1.AuxInt
		x := v_1.Args[0]
		if !(isUint64PowerOfTwo(c)) {
			break
		}
		v.reset(OpAMD64SETBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off] {sym} ptr (TESTQ (MOVQconst [c]) x) mem)
	// cond: isUint64PowerOfTwo(c)
	// result: (SETBstore [off] {sym} ptr (BTQconst [log2(c)] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_1_0 := v_1.Args[_i0]
			if v_1_0.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1_0.AuxInt
			x := v_1.Args[1^_i0]
			if !(isUint64PowerOfTwo(c)) {
				continue
			}
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = log2(c)
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETNEstore [off] {sym} ptr (CMPLconst [1] s:(ANDLconst [1] _)) mem)
	// result: (SETEQstore [off] {sym} ptr (CMPLconst [0] s) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64CMPLconst || v_1.AuxInt != 1 {
			break
		}
		s := v_1.Args[0]
		if s.Op != OpAMD64ANDLconst || s.AuxInt != 1 {
			break
		}
		v.reset(OpAMD64SETEQstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(s)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off] {sym} ptr (CMPQconst [1] s:(ANDQconst [1] _)) mem)
	// result: (SETEQstore [off] {sym} ptr (CMPQconst [0] s) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64CMPQconst || v_1.AuxInt != 1 {
			break
		}
		s := v_1.Args[0]
		if s.Op != OpAMD64ANDQconst || s.AuxInt != 1 {
			break
		}
		v.reset(OpAMD64SETEQstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(s)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off] {sym} ptr (TESTQ z1:(SHLQconst [63] (SHRQconst [63] x)) z2) mem)
	// cond: z1==z2
	// result: (SETBstore [off] {sym} ptr (BTQconst [63] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHLQconst || z1.AuxInt != 63 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 63 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 63
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETNEstore [off] {sym} ptr (TESTL z1:(SHLLconst [31] (SHRLconst [31] x)) z2) mem)
	// cond: z1==z2
	// result: (SETBstore [off] {sym} ptr (BTLconst [31] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTL {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHLLconst || z1.AuxInt != 31 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHRLconst || z1_0.AuxInt != 31 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 31
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETNEstore [off] {sym} ptr (TESTQ z1:(SHRQconst [63] (SHLQconst [63] x)) z2) mem)
	// cond: z1==z2
	// result: (SETBstore [off] {sym} ptr (BTQconst [0] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHLQconst || z1_0.AuxInt != 63 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETNEstore_10(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETNEstore [off] {sym} ptr (TESTL z1:(SHRLconst [31] (SHLLconst [31] x)) z2) mem)
	// cond: z1==z2
	// result: (SETBstore [off] {sym} ptr (BTLconst [0] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTL {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
				continue
			}
			z1_0 := z1.Args[0]
			if z1_0.Op != OpAMD64SHLLconst || z1_0.AuxInt != 31 {
				continue
			}
			x := z1_0.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETNEstore [off] {sym} ptr (TESTQ z1:(SHRQconst [63] x) z2) mem)
	// cond: z1==z2
	// result: (SETBstore [off] {sym} ptr (BTQconst [63] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTQ {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
				continue
			}
			x := z1.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = 63
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETNEstore [off] {sym} ptr (TESTL z1:(SHRLconst [31] x) z2) mem)
	// cond: z1==z2
	// result: (SETBstore [off] {sym} ptr (BTLconst [31] x) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64TESTL {
			break
		}
		_ = v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			z1 := v_1.Args[_i0]
			if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
				continue
			}
			x := z1.Args[0]
			z2 := v_1.Args[1^_i0]
			if !(z1 == z2) {
				continue
			}
			v.reset(OpAMD64SETBstore)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = 31
			v0.AddArg(x)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		break
	}
	// match: (SETNEstore [off] {sym} ptr (InvertFlags x) mem)
	// result: (SETNEstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64InvertFlags {
			break
		}
		x := v_1.Args[0]
		v.reset(OpAMD64SETNEstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SETNEstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SETNEstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SETNEstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SETNEstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off] {sym} ptr (FlagEQ) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [0]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagEQ {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off] {sym} ptr (FlagLT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off] {sym} ptr (FlagLT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagLT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (SETNEstore [off] {sym} ptr (FlagGT_ULT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_ULT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SETNEstore_20(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETNEstore [off] {sym} ptr (FlagGT_UGT) mem)
	// result: (MOVBstore [off] {sym} ptr (MOVLconst <typ.UInt8> [1]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64FlagGT_UGT {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLconst, typ.UInt8)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHLL_0(v *Value) bool {
	b := v.Block
	// match: (SHLL x (MOVQconst [c]))
	// result: (SHLLconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SHLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SHLL x (MOVLconst [c]))
	// result: (SHLLconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SHLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SHLL x (ADDQconst [c] y))
	// cond: c & 31 == 0
	// result: (SHLL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHLL x (NEGQ <t> (ADDQconst [c] y)))
	// cond: c & 31 == 0
	// result: (SHLL x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHLL x (ANDQconst [c] y))
	// cond: c & 31 == 31
	// result: (SHLL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHLL x (NEGQ <t> (ANDQconst [c] y)))
	// cond: c & 31 == 31
	// result: (SHLL x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHLL x (ADDLconst [c] y))
	// cond: c & 31 == 0
	// result: (SHLL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHLL x (NEGL <t> (ADDLconst [c] y)))
	// cond: c & 31 == 0
	// result: (SHLL x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHLL x (ANDLconst [c] y))
	// cond: c & 31 == 31
	// result: (SHLL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHLL x (NEGL <t> (ANDLconst [c] y)))
	// cond: c & 31 == 31
	// result: (SHLL x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHLLconst_0(v *Value) bool {
	// match: (SHLLconst [1] (SHRLconst [1] x))
	// result: (BTRLconst [0] x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHRLconst || v_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTRLconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (SHLLconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHLQ_0(v *Value) bool {
	b := v.Block
	// match: (SHLQ x (MOVQconst [c]))
	// result: (SHLQconst [c&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SHLQconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SHLQ x (MOVLconst [c]))
	// result: (SHLQconst [c&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SHLQconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SHLQ x (ADDQconst [c] y))
	// cond: c & 63 == 0
	// result: (SHLQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHLQ x (NEGQ <t> (ADDQconst [c] y)))
	// cond: c & 63 == 0
	// result: (SHLQ x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHLQ x (ANDQconst [c] y))
	// cond: c & 63 == 63
	// result: (SHLQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHLQ x (NEGQ <t> (ANDQconst [c] y)))
	// cond: c & 63 == 63
	// result: (SHLQ x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHLQ x (ADDLconst [c] y))
	// cond: c & 63 == 0
	// result: (SHLQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHLQ x (NEGL <t> (ADDLconst [c] y)))
	// cond: c & 63 == 0
	// result: (SHLQ x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHLQ x (ANDLconst [c] y))
	// cond: c & 63 == 63
	// result: (SHLQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHLQ x (NEGL <t> (ANDLconst [c] y)))
	// cond: c & 63 == 63
	// result: (SHLQ x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHLQconst_0(v *Value) bool {
	// match: (SHLQconst [1] (SHRQconst [1] x))
	// result: (BTRQconst [0] x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHRQconst || v_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTRQconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (SHLQconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHRB_0(v *Value) bool {
	// match: (SHRB x (MOVQconst [c]))
	// cond: c&31 < 8
	// result: (SHRBconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 < 8) {
			break
		}
		v.reset(OpAMD64SHRBconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SHRB x (MOVLconst [c]))
	// cond: c&31 < 8
	// result: (SHRBconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 < 8) {
			break
		}
		v.reset(OpAMD64SHRBconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SHRB _ (MOVQconst [c]))
	// cond: c&31 >= 8
	// result: (MOVLconst [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 >= 8) {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SHRB _ (MOVLconst [c]))
	// cond: c&31 >= 8
	// result: (MOVLconst [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 >= 8) {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHRBconst_0(v *Value) bool {
	// match: (SHRBconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHRL_0(v *Value) bool {
	b := v.Block
	// match: (SHRL x (MOVQconst [c]))
	// result: (SHRLconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SHRLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SHRL x (MOVLconst [c]))
	// result: (SHRLconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SHRLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SHRL x (ADDQconst [c] y))
	// cond: c & 31 == 0
	// result: (SHRL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHRL x (NEGQ <t> (ADDQconst [c] y)))
	// cond: c & 31 == 0
	// result: (SHRL x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHRL x (ANDQconst [c] y))
	// cond: c & 31 == 31
	// result: (SHRL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHRL x (NEGQ <t> (ANDQconst [c] y)))
	// cond: c & 31 == 31
	// result: (SHRL x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHRL x (ADDLconst [c] y))
	// cond: c & 31 == 0
	// result: (SHRL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHRL x (NEGL <t> (ADDLconst [c] y)))
	// cond: c & 31 == 0
	// result: (SHRL x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 0) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHRL x (ANDLconst [c] y))
	// cond: c & 31 == 31
	// result: (SHRL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHRL x (NEGL <t> (ANDLconst [c] y)))
	// cond: c & 31 == 31
	// result: (SHRL x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&31 == 31) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHRLconst_0(v *Value) bool {
	// match: (SHRLconst [1] (SHLLconst [1] x))
	// result: (BTRLconst [31] x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHLLconst || v_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTRLconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	// match: (SHRLconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHRQ_0(v *Value) bool {
	b := v.Block
	// match: (SHRQ x (MOVQconst [c]))
	// result: (SHRQconst [c&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SHRQconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SHRQ x (MOVLconst [c]))
	// result: (SHRQconst [c&63] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SHRQconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (SHRQ x (ADDQconst [c] y))
	// cond: c & 63 == 0
	// result: (SHRQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHRQ x (NEGQ <t> (ADDQconst [c] y)))
	// cond: c & 63 == 0
	// result: (SHRQ x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHRQ x (ANDQconst [c] y))
	// cond: c & 63 == 63
	// result: (SHRQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHRQ x (NEGQ <t> (ANDQconst [c] y)))
	// cond: c & 63 == 63
	// result: (SHRQ x (NEGQ <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGQ {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDQconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHRQ x (ADDLconst [c] y))
	// cond: c & 63 == 0
	// result: (SHRQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHRQ x (NEGL <t> (ADDLconst [c] y)))
	// cond: c & 63 == 0
	// result: (SHRQ x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ADDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 0) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SHRQ x (ANDLconst [c] y))
	// cond: c & 63 == 63
	// result: (SHRQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SHRQ x (NEGL <t> (ANDLconst [c] y)))
	// cond: c & 63 == 63
	// result: (SHRQ x (NEGL <t> y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64NEGL {
			break
		}
		t := v_1.Type
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpAMD64ANDLconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1_0.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64NEGL, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHRQconst_0(v *Value) bool {
	// match: (SHRQconst [1] (SHLQconst [1] x))
	// result: (BTRQconst [63] x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SHLQconst || v_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64BTRQconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}
	// match: (SHRQconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHRW_0(v *Value) bool {
	// match: (SHRW x (MOVQconst [c]))
	// cond: c&31 < 16
	// result: (SHRWconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 < 16) {
			break
		}
		v.reset(OpAMD64SHRWconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SHRW x (MOVLconst [c]))
	// cond: c&31 < 16
	// result: (SHRWconst [c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 < 16) {
			break
		}
		v.reset(OpAMD64SHRWconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (SHRW _ (MOVQconst [c]))
	// cond: c&31 >= 16
	// result: (MOVLconst [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 >= 16) {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SHRW _ (MOVLconst [c]))
	// cond: c&31 >= 16
	// result: (MOVLconst [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 >= 16) {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SHRWconst_0(v *Value) bool {
	// match: (SHRWconst x [0])
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBL_0(v *Value) bool {
	b := v.Block
	// match: (SUBL x (MOVLconst [c]))
	// result: (SUBLconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAMD64SUBLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUBL (MOVLconst [c]) x)
	// result: (NEGL (SUBLconst <v.Type> x [c]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpAMD64NEGL)
		v0 := b.NewValue0(v.Pos, OpAMD64SUBLconst, v.Type)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBL x x)
	// result: (MOVLconst [0])
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUBL x l:(MOVLload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (SUBLload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
			break
		}
		v.reset(OpAMD64SUBLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBLconst_0(v *Value) bool {
	// match: (SUBLconst [c] x)
	// cond: int32(c) == 0
	// result: x
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SUBLconst [c] x)
	// result: (ADDLconst [int64(int32(-c))] x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		v.reset(OpAMD64ADDLconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpAMD64SUBLload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SUBLload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (SUBLload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SUBLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (SUBLload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SUBLload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SUBLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (SUBLload x [off] {sym} ptr (MOVSSstore [off] {sym} ptr y _))
	// result: (SUBL x (MOVLf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSSstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64SUBL)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVLf2i, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBLmodify_0(v *Value) bool {
	// match: (SUBLmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SUBLmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SUBLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SUBLmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SUBLmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SUBLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBQ_0(v *Value) bool {
	b := v.Block
	// match: (SUBQ x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (SUBQconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpAMD64SUBQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUBQ (MOVQconst [c]) x)
	// cond: is32Bit(c)
	// result: (NEGQ (SUBQconst <v.Type> x [c]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		c := v_0.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpAMD64NEGQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SUBQconst, v.Type)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBQ x x)
	// result: (MOVQconst [0])
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUBQ x l:(MOVQload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (SUBQload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVQload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
			break
		}
		v.reset(OpAMD64SUBQload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBQborrow_0(v *Value) bool {
	// match: (SUBQborrow x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (SUBQconstborrow x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64MOVQconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpAMD64SUBQconstborrow)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBQconst_0(v *Value) bool {
	// match: (SUBQconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SUBQconst [c] x)
	// cond: c != -(1<<31)
	// result: (ADDQconst [-c] x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c != -(1 << 31)) {
			break
		}
		v.reset(OpAMD64ADDQconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	// match: (SUBQconst (MOVQconst [d]) [c])
	// result: (MOVQconst [d-c])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = d - c
		return true
	}
	// match: (SUBQconst (SUBQconst x [d]) [c])
	// cond: is32Bit(-c-d)
	// result: (ADDQconst [-c-d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SUBQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c - d)) {
			break
		}
		v.reset(OpAMD64ADDQconst)
		v.AuxInt = -c - d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBQload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SUBQload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (SUBQload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SUBQload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (SUBQload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SUBQload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SUBQload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (SUBQload x [off] {sym} ptr (MOVSDstore [off] {sym} ptr y _))
	// result: (SUBQ x (MOVQf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64SUBQ)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVQf2i, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBQmodify_0(v *Value) bool {
	// match: (SUBQmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (SUBQmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SUBQmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (SUBQmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SUBQmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SUBQmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBSD_0(v *Value) bool {
	// match: (SUBSD x l:(MOVSDload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (SUBSDload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVSDload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
			break
		}
		v.reset(OpAMD64SUBSDload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBSDload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SUBSDload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (SUBSDload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SUBSDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (SUBSDload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SUBSDload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SUBSDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (SUBSDload x [off] {sym} ptr (MOVQstore [off] {sym} ptr y _))
	// result: (SUBSD x (MOVQi2f y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVQstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64SUBSD)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVQi2f, typ.Float64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBSS_0(v *Value) bool {
	// match: (SUBSS x l:(MOVSSload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (SUBSSload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != OpAMD64MOVSSload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		mem := l.Args[1]
		ptr := l.Args[0]
		if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
			break
		}
		v.reset(OpAMD64SUBSSload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64SUBSSload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SUBSSload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (SUBSSload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64SUBSSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (SUBSSload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (SUBSSload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64SUBSSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (SUBSSload x [off] {sym} ptr (MOVLstore [off] {sym} ptr y _))
	// result: (SUBSS x (MOVLi2f y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVLstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64SUBSS)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVLi2f, typ.Float32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64TESTB_0(v *Value) bool {
	b := v.Block
	// match: (TESTB (MOVLconst [c]) x)
	// result: (TESTBconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			v.reset(OpAMD64TESTBconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (TESTB l:(MOVBload {sym} [off] ptr mem) l2)
	// cond: l == l2 && l.Uses == 2 && validValAndOff(0,off) && clobber(l)
	// result: @l.Block (CMPBconstload {sym} [makeValAndOff(0,off)] ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := v.Args[_i0]
			if l.Op != OpAMD64MOVBload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			l2 := v.Args[1^_i0]
			if !(l == l2 && l.Uses == 2 && validValAndOff(0, off) && clobber(l)) {
				continue
			}
			b = l.Block
			v0 := b.NewValue0(l.Pos, OpAMD64CMPBconstload, types.TypeFlags)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = makeValAndOff(0, off)
			v0.Aux = sym
			v0.AddArg(ptr)
			v0.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64TESTBconst_0(v *Value) bool {
	// match: (TESTBconst [-1] x)
	// cond: x.Op != OpAMD64MOVLconst
	// result: (TESTB x x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		if !(x.Op != OpAMD64MOVLconst) {
			break
		}
		v.reset(OpAMD64TESTB)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64TESTL_0(v *Value) bool {
	b := v.Block
	// match: (TESTL (MOVLconst [c]) x)
	// result: (TESTLconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			v.reset(OpAMD64TESTLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (TESTL l:(MOVLload {sym} [off] ptr mem) l2)
	// cond: l == l2 && l.Uses == 2 && validValAndOff(0,off) && clobber(l)
	// result: @l.Block (CMPLconstload {sym} [makeValAndOff(0,off)] ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := v.Args[_i0]
			if l.Op != OpAMD64MOVLload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			l2 := v.Args[1^_i0]
			if !(l == l2 && l.Uses == 2 && validValAndOff(0, off) && clobber(l)) {
				continue
			}
			b = l.Block
			v0 := b.NewValue0(l.Pos, OpAMD64CMPLconstload, types.TypeFlags)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = makeValAndOff(0, off)
			v0.Aux = sym
			v0.AddArg(ptr)
			v0.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64TESTLconst_0(v *Value) bool {
	// match: (TESTLconst [-1] x)
	// cond: x.Op != OpAMD64MOVLconst
	// result: (TESTL x x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		if !(x.Op != OpAMD64MOVLconst) {
			break
		}
		v.reset(OpAMD64TESTL)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64TESTQ_0(v *Value) bool {
	b := v.Block
	// match: (TESTQ (MOVQconst [c]) x)
	// cond: is32Bit(c)
	// result: (TESTQconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpAMD64TESTQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (TESTQ l:(MOVQload {sym} [off] ptr mem) l2)
	// cond: l == l2 && l.Uses == 2 && validValAndOff(0,off) && clobber(l)
	// result: @l.Block (CMPQconstload {sym} [makeValAndOff(0,off)] ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := v.Args[_i0]
			if l.Op != OpAMD64MOVQload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			l2 := v.Args[1^_i0]
			if !(l == l2 && l.Uses == 2 && validValAndOff(0, off) && clobber(l)) {
				continue
			}
			b = l.Block
			v0 := b.NewValue0(l.Pos, OpAMD64CMPQconstload, types.TypeFlags)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = makeValAndOff(0, off)
			v0.Aux = sym
			v0.AddArg(ptr)
			v0.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64TESTQconst_0(v *Value) bool {
	// match: (TESTQconst [-1] x)
	// cond: x.Op != OpAMD64MOVQconst
	// result: (TESTQ x x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		if !(x.Op != OpAMD64MOVQconst) {
			break
		}
		v.reset(OpAMD64TESTQ)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64TESTW_0(v *Value) bool {
	b := v.Block
	// match: (TESTW (MOVLconst [c]) x)
	// result: (TESTWconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			v.reset(OpAMD64TESTWconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (TESTW l:(MOVWload {sym} [off] ptr mem) l2)
	// cond: l == l2 && l.Uses == 2 && validValAndOff(0,off) && clobber(l)
	// result: @l.Block (CMPWconstload {sym} [makeValAndOff(0,off)] ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			l := v.Args[_i0]
			if l.Op != OpAMD64MOVWload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			l2 := v.Args[1^_i0]
			if !(l == l2 && l.Uses == 2 && validValAndOff(0, off) && clobber(l)) {
				continue
			}
			b = l.Block
			v0 := b.NewValue0(l.Pos, OpAMD64CMPWconstload, types.TypeFlags)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = makeValAndOff(0, off)
			v0.Aux = sym
			v0.AddArg(ptr)
			v0.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64TESTWconst_0(v *Value) bool {
	// match: (TESTWconst [-1] x)
	// cond: x.Op != OpAMD64MOVLconst
	// result: (TESTW x x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		if !(x.Op != OpAMD64MOVLconst) {
			break
		}
		v.reset(OpAMD64TESTW)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XADDLlock_0(v *Value) bool {
	// match: (XADDLlock [off1] {sym} val (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (XADDLlock [off1+off2] {sym} val ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64XADDLlock)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XADDQlock_0(v *Value) bool {
	// match: (XADDQlock [off1] {sym} val (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (XADDQlock [off1+off2] {sym} val ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64XADDQlock)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XCHGL_0(v *Value) bool {
	// match: (XCHGL [off1] {sym} val (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (XCHGL [off1+off2] {sym} val ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64XCHGL)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (XCHGL [off1] {sym1} val (LEAQ [off2] {sym2} ptr) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && ptr.Op != OpSB
	// result: (XCHGL [off1+off2] {mergeSym(sym1,sym2)} val ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpAMD64XCHGL)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XCHGQ_0(v *Value) bool {
	// match: (XCHGQ [off1] {sym} val (ADDQconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (XCHGQ [off1+off2] {sym} val ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64XCHGQ)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (XCHGQ [off1] {sym1} val (LEAQ [off2] {sym2} ptr) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2) && ptr.Op != OpSB
	// result: (XCHGQ [off1+off2] {mergeSym(sym1,sym2)} val ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpAMD64XCHGQ)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORL_0(v *Value) bool {
	// match: (XORL (SHLL (MOVLconst [1]) y) x)
	// result: (BTCL x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLL {
				continue
			}
			y := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64MOVLconst || v_0_0.AuxInt != 1 {
				continue
			}
			x := v.Args[1^_i0]
			v.reset(OpAMD64BTCL)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (XORL (MOVLconst [c]) x)
	// cond: isUint32PowerOfTwo(c) && uint64(c) >= 128
	// result: (BTCLconst [log2uint32(c)] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			if !(isUint32PowerOfTwo(c) && uint64(c) >= 128) {
				continue
			}
			v.reset(OpAMD64BTCLconst)
			v.AuxInt = log2uint32(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORL x (MOVLconst [c]))
	// result: (XORLconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVLconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpAMD64XORLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORL (SHLLconst x [c]) (SHRLconst x [d]))
	// cond: d==32-c
	// result: (ROLLconst x [c])
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRLconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpAMD64ROLLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORL <t> (SHLLconst x [c]) (SHRWconst x [d]))
	// cond: d==16-c && c < 16 && t.Size() == 2
	// result: (ROLWconst x [c])
	for {
		t := v.Type
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 16-c && c < 16 && t.Size() == 2) {
				continue
			}
			v.reset(OpAMD64ROLWconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORL <t> (SHLLconst x [c]) (SHRBconst x [d]))
	// cond: d==8-c && c < 8 && t.Size() == 1
	// result: (ROLBconst x [c])
	for {
		t := v.Type
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLLconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRBconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 8-c && c < 8 && t.Size() == 1) {
				continue
			}
			v.reset(OpAMD64ROLBconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORL x x)
	// result: (MOVLconst [0])
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = 0
		return true
	}
	// match: (XORL x l:(MOVLload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (XORLload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVLload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64XORLload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORLconst_0(v *Value) bool {
	// match: (XORLconst [c] x)
	// cond: isUint32PowerOfTwo(c) && uint64(c) >= 128
	// result: (BTCLconst [log2uint32(c)] x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isUint32PowerOfTwo(c) && uint64(c) >= 128) {
			break
		}
		v.reset(OpAMD64BTCLconst)
		v.AuxInt = log2uint32(c)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETNE x))
	// result: (SETEQ x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETNE {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETEQ)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETEQ x))
	// result: (SETNE x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETEQ {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETNE)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETL x))
	// result: (SETGE x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETL {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETGE)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETGE x))
	// result: (SETL x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETGE {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETL)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETLE x))
	// result: (SETG x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETLE {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETG)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETG x))
	// result: (SETLE x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETG {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETLE)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETB x))
	// result: (SETAE x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETB {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETAE)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETAE x))
	// result: (SETB x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETAE {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETB)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [1] (SETBE x))
	// result: (SETA x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETBE {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETA)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORLconst_10(v *Value) bool {
	// match: (XORLconst [1] (SETA x))
	// result: (SETBE x)
	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64SETA {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAMD64SETBE)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [c] (XORLconst [d] x))
	// result: (XORLconst [c ^ d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64XORLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64XORLconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [c] (BTCLconst [d] x))
	// result: (XORLconst [c ^ 1<<uint32(d)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTCLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64XORLconst)
		v.AuxInt = c ^ 1<<uint32(d)
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [c] x)
	// cond: int32(c)==0
	// result: x
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (XORLconst [c] (MOVLconst [d]))
	// result: (MOVLconst [c^d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = c ^ d
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORLconstmodify_0(v *Value) bool {
	// match: (XORLconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (XORLconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64XORLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (XORLconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (XORLconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64XORLconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORLload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (XORLload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (XORLload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64XORLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (XORLload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (XORLload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64XORLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (XORLload x [off] {sym} ptr (MOVSSstore [off] {sym} ptr y _))
	// result: (XORL x (MOVLf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSSstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64XORL)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVLf2i, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORLmodify_0(v *Value) bool {
	// match: (XORLmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (XORLmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64XORLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (XORLmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (XORLmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64XORLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORQ_0(v *Value) bool {
	// match: (XORQ (SHLQ (MOVQconst [1]) y) x)
	// result: (BTCQ x y)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLQ {
				continue
			}
			y := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64MOVQconst || v_0_0.AuxInt != 1 {
				continue
			}
			x := v.Args[1^_i0]
			v.reset(OpAMD64BTCQ)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (XORQ (MOVQconst [c]) x)
	// cond: isUint64PowerOfTwo(c) && uint64(c) >= 128
	// result: (BTCQconst [log2(c)] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_0.AuxInt
			x := v.Args[1^_i0]
			if !(isUint64PowerOfTwo(c) && uint64(c) >= 128) {
				continue
			}
			v.reset(OpAMD64BTCQconst)
			v.AuxInt = log2(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORQ x (MOVQconst [c]))
	// cond: is32Bit(c)
	// result: (XORQconst [c] x)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64MOVQconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpAMD64XORQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORQ (SHLQconst x [c]) (SHRQconst x [d]))
	// cond: d==64-c
	// result: (ROLQconst x [c])
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			v_0 := v.Args[_i0]
			if v_0.Op != OpAMD64SHLQconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			v_1 := v.Args[1^_i0]
			if v_1.Op != OpAMD64SHRQconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpAMD64ROLQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XORQ x x)
	// result: (MOVQconst [0])
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = 0
		return true
	}
	// match: (XORQ x l:(MOVQload [off] {sym} ptr mem))
	// cond: canMergeLoadClobber(v, l, x) && clobber(l)
	// result: (XORQload x [off] {sym} ptr mem)
	for {
		_ = v.Args[1]
		for _i0 := 0; _i0 <= 1; _i0++ {
			x := v.Args[_i0]
			l := v.Args[1^_i0]
			if l.Op != OpAMD64MOVQload {
				continue
			}
			off := l.AuxInt
			sym := l.Aux
			mem := l.Args[1]
			ptr := l.Args[0]
			if !(canMergeLoadClobber(v, l, x) && clobber(l)) {
				continue
			}
			v.reset(OpAMD64XORQload)
			v.AuxInt = off
			v.Aux = sym
			v.AddArg(x)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		break
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORQconst_0(v *Value) bool {
	// match: (XORQconst [c] x)
	// cond: isUint64PowerOfTwo(c) && uint64(c) >= 128
	// result: (BTCQconst [log2(c)] x)
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isUint64PowerOfTwo(c) && uint64(c) >= 128) {
			break
		}
		v.reset(OpAMD64BTCQconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (XORQconst [c] (XORQconst [d] x))
	// result: (XORQconst [c ^ d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64XORQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64XORQconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	// match: (XORQconst [c] (BTCQconst [d] x))
	// result: (XORQconst [c ^ 1<<uint32(d)] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64BTCQconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpAMD64XORQconst)
		v.AuxInt = c ^ 1<<uint32(d)
		v.AddArg(x)
		return true
	}
	// match: (XORQconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (XORQconst [c] (MOVQconst [d]))
	// result: (MOVQconst [c^d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64MOVQconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = c ^ d
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORQconstmodify_0(v *Value) bool {
	// match: (XORQconstmodify [valoff1] {sym} (ADDQconst [off2] base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2)
	// result: (XORQconstmodify [ValAndOff(valoff1).add(off2)] {sym} base mem)
	for {
		valoff1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2)) {
			break
		}
		v.reset(OpAMD64XORQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (XORQconstmodify [valoff1] {sym1} (LEAQ [off2] {sym2} base) mem)
	// cond: ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)
	// result: (XORQconstmodify [ValAndOff(valoff1).add(off2)] {mergeSym(sym1,sym2)} base mem)
	for {
		valoff1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		if !(ValAndOff(valoff1).canAdd(off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64XORQconstmodify)
		v.AuxInt = ValAndOff(valoff1).add(off2)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORQload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (XORQload [off1] {sym} val (ADDQconst [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (XORQload [off1+off2] {sym} val base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64XORQload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (XORQload [off1] {sym1} val (LEAQ [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (XORQload [off1+off2] {mergeSym(sym1,sym2)} val base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64XORQload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (XORQload x [off] {sym} ptr (MOVSDstore [off] {sym} ptr y _))
	// result: (XORQ x (MOVQf2i y))
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64MOVSDstore || v_2.AuxInt != off || v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		if ptr != v_2.Args[0] {
			break
		}
		y := v_2.Args[1]
		v.reset(OpAMD64XORQ)
		v.AddArg(x)
		v0 := b.NewValue0(v_2.Pos, OpAMD64MOVQf2i, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAMD64XORQmodify_0(v *Value) bool {
	// match: (XORQmodify [off1] {sym} (ADDQconst [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (XORQmodify [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64ADDQconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpAMD64XORQmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (XORQmodify [off1] {sym1} (LEAQ [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (XORQmodify [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64LEAQ {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpAMD64XORQmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpAdd16_0(v *Value) bool {
	// match: (Add16 x y)
	// result: (ADDL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ADDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAdd32_0(v *Value) bool {
	// match: (Add32 x y)
	// result: (ADDL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ADDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAdd32F_0(v *Value) bool {
	// match: (Add32F x y)
	// result: (ADDSS x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ADDSS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAdd64_0(v *Value) bool {
	// match: (Add64 x y)
	// result: (ADDQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ADDQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAdd64F_0(v *Value) bool {
	// match: (Add64F x y)
	// result: (ADDSD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ADDSD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAdd8_0(v *Value) bool {
	// match: (Add8 x y)
	// result: (ADDL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ADDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAddPtr_0(v *Value) bool {
	// match: (AddPtr x y)
	// result: (ADDQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ADDQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAddr_0(v *Value) bool {
	// match: (Addr {sym} base)
	// result: (LEAQ {sym} base)
	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpAMD64LEAQ)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueAMD64_OpAnd16_0(v *Value) bool {
	// match: (And16 x y)
	// result: (ANDL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ANDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAnd32_0(v *Value) bool {
	// match: (And32 x y)
	// result: (ANDL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ANDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAnd64_0(v *Value) bool {
	// match: (And64 x y)
	// result: (ANDQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ANDQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAnd8_0(v *Value) bool {
	// match: (And8 x y)
	// result: (ANDL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ANDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAndB_0(v *Value) bool {
	// match: (AndB x y)
	// result: (ANDL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ANDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpAtomicAdd32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicAdd32 ptr val mem)
	// result: (AddTupleFirst32 val (XADDLlock val ptr mem))
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpAMD64AddTupleFirst32)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpAMD64XADDLlock, types.NewTuple(typ.UInt32, types.TypeMem))
		v0.AddArg(val)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpAtomicAdd64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicAdd64 ptr val mem)
	// result: (AddTupleFirst64 val (XADDQlock val ptr mem))
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpAMD64AddTupleFirst64)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpAMD64XADDQlock, types.NewTuple(typ.UInt64, types.TypeMem))
		v0.AddArg(val)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpAtomicAnd8_0(v *Value) bool {
	// match: (AtomicAnd8 ptr val mem)
	// result: (ANDBlock ptr val mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpAMD64ANDBlock)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicCompareAndSwap32_0(v *Value) bool {
	// match: (AtomicCompareAndSwap32 ptr old new_ mem)
	// result: (CMPXCHGLlock ptr old new_ mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		v.reset(OpAMD64CMPXCHGLlock)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicCompareAndSwap64_0(v *Value) bool {
	// match: (AtomicCompareAndSwap64 ptr old new_ mem)
	// result: (CMPXCHGQlock ptr old new_ mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		v.reset(OpAMD64CMPXCHGQlock)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicExchange32_0(v *Value) bool {
	// match: (AtomicExchange32 ptr val mem)
	// result: (XCHGL val ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpAMD64XCHGL)
		v.AddArg(val)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicExchange64_0(v *Value) bool {
	// match: (AtomicExchange64 ptr val mem)
	// result: (XCHGQ val ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpAMD64XCHGQ)
		v.AddArg(val)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicLoad32_0(v *Value) bool {
	// match: (AtomicLoad32 ptr mem)
	// result: (MOVLatomicload ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64MOVLatomicload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicLoad64_0(v *Value) bool {
	// match: (AtomicLoad64 ptr mem)
	// result: (MOVQatomicload ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64MOVQatomicload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicLoad8_0(v *Value) bool {
	// match: (AtomicLoad8 ptr mem)
	// result: (MOVBatomicload ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64MOVBatomicload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicLoadPtr_0(v *Value) bool {
	// match: (AtomicLoadPtr ptr mem)
	// result: (MOVQatomicload ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64MOVQatomicload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicOr8_0(v *Value) bool {
	// match: (AtomicOr8 ptr val mem)
	// result: (ORBlock ptr val mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpAMD64ORBlock)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpAtomicStore32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicStore32 ptr val mem)
	// result: (Select1 (XCHGL <types.NewTuple(typ.UInt32,types.TypeMem)> val ptr mem))
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64XCHGL, types.NewTuple(typ.UInt32, types.TypeMem))
		v0.AddArg(val)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpAtomicStore64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicStore64 ptr val mem)
	// result: (Select1 (XCHGQ <types.NewTuple(typ.UInt64,types.TypeMem)> val ptr mem))
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64XCHGQ, types.NewTuple(typ.UInt64, types.TypeMem))
		v0.AddArg(val)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpAtomicStore8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicStore8 ptr val mem)
	// result: (Select1 (XCHGB <types.NewTuple(typ.UInt8,types.TypeMem)> val ptr mem))
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64XCHGB, types.NewTuple(typ.UInt8, types.TypeMem))
		v0.AddArg(val)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpAtomicStorePtrNoWB_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AtomicStorePtrNoWB ptr val mem)
	// result: (Select1 (XCHGQ <types.NewTuple(typ.BytePtr,types.TypeMem)> val ptr mem))
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64XCHGQ, types.NewTuple(typ.BytePtr, types.TypeMem))
		v0.AddArg(val)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpAvg64u_0(v *Value) bool {
	// match: (Avg64u x y)
	// result: (AVGQU x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64AVGQU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpBitLen16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen16 x)
	// result: (BSRL (LEAL1 <typ.UInt32> [1] (MOVWQZX <typ.UInt32> x) (MOVWQZX <typ.UInt32> x)))
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSRL)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL1, typ.UInt32)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpAMD64MOVWQZX, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVWQZX, typ.UInt32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpBitLen32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen32 x)
	// result: (Select0 (BSRQ (LEAQ1 <typ.UInt64> [1] (MOVLQZX <typ.UInt64> x) (MOVLQZX <typ.UInt64> x))))
	for {
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64BSRQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v1 := b.NewValue0(v.Pos, OpAMD64LEAQ1, typ.UInt64)
		v1.AuxInt = 1
		v2 := b.NewValue0(v.Pos, OpAMD64MOVLQZX, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpAMD64MOVLQZX, typ.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpBitLen64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen64 <t> x)
	// result: (ADDQconst [1] (CMOVQEQ <t> (Select0 <t> (BSRQ x)) (MOVQconst <t> [-1]) (Select1 <types.TypeFlags> (BSRQ x))))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpAMD64ADDQconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpAMD64CMOVQEQ, t)
		v1 := b.NewValue0(v.Pos, OpSelect0, t)
		v2 := b.NewValue0(v.Pos, OpAMD64BSRQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v2.AddArg(x)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpAMD64MOVQconst, t)
		v3.AuxInt = -1
		v0.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v5 := b.NewValue0(v.Pos, OpAMD64BSRQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v5.AddArg(x)
		v4.AddArg(v5)
		v0.AddArg(v4)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpBitLen8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen8 x)
	// result: (BSRL (LEAL1 <typ.UInt32> [1] (MOVBQZX <typ.UInt32> x) (MOVBQZX <typ.UInt32> x)))
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSRL)
		v0 := b.NewValue0(v.Pos, OpAMD64LEAL1, typ.UInt32)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpAMD64MOVBQZX, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVBQZX, typ.UInt32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpBswap32_0(v *Value) bool {
	// match: (Bswap32 x)
	// result: (BSWAPL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSWAPL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpBswap64_0(v *Value) bool {
	// match: (Bswap64 x)
	// result: (BSWAPQ x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSWAPQ)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCeil_0(v *Value) bool {
	// match: (Ceil x)
	// result: (ROUNDSD [2] x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64ROUNDSD)
		v.AuxInt = 2
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpClosureCall_0(v *Value) bool {
	// match: (ClosureCall [argwid] entry closure mem)
	// result: (CALLclosure [argwid] entry closure mem)
	for {
		argwid := v.AuxInt
		mem := v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		v.reset(OpAMD64CALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpCom16_0(v *Value) bool {
	// match: (Com16 x)
	// result: (NOTL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64NOTL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCom32_0(v *Value) bool {
	// match: (Com32 x)
	// result: (NOTL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64NOTL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCom64_0(v *Value) bool {
	// match: (Com64 x)
	// result: (NOTQ x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64NOTQ)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCom8_0(v *Value) bool {
	// match: (Com8 x)
	// result: (NOTL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64NOTL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCondSelect_0(v *Value) bool {
	// match: (CondSelect <t> x y (SETEQ cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQEQ y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETEQ {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQEQ)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETNE cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQNE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETNE {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQNE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETL cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQLT y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETL {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQLT)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETG cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQGT y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETG {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQGT)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETLE cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQLE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETLE {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQLE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETGE cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQGE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGE {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQGE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETA cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQHI y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETA {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQHI)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETB cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQCS y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETB {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQCS)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETAE cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQCC y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETAE {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQCC)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETBE cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQLS y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETBE {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQLS)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	return false
}
func rewriteValueAMD64_OpCondSelect_10(v *Value) bool {
	// match: (CondSelect <t> x y (SETEQF cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQEQF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETEQF {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQEQF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETNEF cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQNEF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETNEF {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQNEF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETGF cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQGTF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGF {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQGTF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETGEF cond))
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (CMOVQGEF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGEF {
			break
		}
		cond := v_2.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64CMOVQGEF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETEQ cond))
	// cond: is32BitInt(t)
	// result: (CMOVLEQ y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETEQ {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLEQ)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETNE cond))
	// cond: is32BitInt(t)
	// result: (CMOVLNE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETNE {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLNE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETL cond))
	// cond: is32BitInt(t)
	// result: (CMOVLLT y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETL {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLLT)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETG cond))
	// cond: is32BitInt(t)
	// result: (CMOVLGT y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETG {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLGT)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETLE cond))
	// cond: is32BitInt(t)
	// result: (CMOVLLE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETLE {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLLE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETGE cond))
	// cond: is32BitInt(t)
	// result: (CMOVLGE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGE {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLGE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	return false
}
func rewriteValueAMD64_OpCondSelect_20(v *Value) bool {
	// match: (CondSelect <t> x y (SETA cond))
	// cond: is32BitInt(t)
	// result: (CMOVLHI y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETA {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLHI)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETB cond))
	// cond: is32BitInt(t)
	// result: (CMOVLCS y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETB {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLCS)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETAE cond))
	// cond: is32BitInt(t)
	// result: (CMOVLCC y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETAE {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLCC)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETBE cond))
	// cond: is32BitInt(t)
	// result: (CMOVLLS y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETBE {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLLS)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETEQF cond))
	// cond: is32BitInt(t)
	// result: (CMOVLEQF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETEQF {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLEQF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETNEF cond))
	// cond: is32BitInt(t)
	// result: (CMOVLNEF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETNEF {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLNEF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETGF cond))
	// cond: is32BitInt(t)
	// result: (CMOVLGTF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGF {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLGTF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETGEF cond))
	// cond: is32BitInt(t)
	// result: (CMOVLGEF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGEF {
			break
		}
		cond := v_2.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLGEF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETEQ cond))
	// cond: is16BitInt(t)
	// result: (CMOVWEQ y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETEQ {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWEQ)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETNE cond))
	// cond: is16BitInt(t)
	// result: (CMOVWNE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETNE {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWNE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	return false
}
func rewriteValueAMD64_OpCondSelect_30(v *Value) bool {
	// match: (CondSelect <t> x y (SETL cond))
	// cond: is16BitInt(t)
	// result: (CMOVWLT y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETL {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWLT)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETG cond))
	// cond: is16BitInt(t)
	// result: (CMOVWGT y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETG {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWGT)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETLE cond))
	// cond: is16BitInt(t)
	// result: (CMOVWLE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETLE {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWLE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETGE cond))
	// cond: is16BitInt(t)
	// result: (CMOVWGE y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGE {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWGE)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETA cond))
	// cond: is16BitInt(t)
	// result: (CMOVWHI y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETA {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWHI)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETB cond))
	// cond: is16BitInt(t)
	// result: (CMOVWCS y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETB {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWCS)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETAE cond))
	// cond: is16BitInt(t)
	// result: (CMOVWCC y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETAE {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWCC)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETBE cond))
	// cond: is16BitInt(t)
	// result: (CMOVWLS y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETBE {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWLS)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETEQF cond))
	// cond: is16BitInt(t)
	// result: (CMOVWEQF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETEQF {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWEQF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETNEF cond))
	// cond: is16BitInt(t)
	// result: (CMOVWNEF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETNEF {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWNEF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	return false
}
func rewriteValueAMD64_OpCondSelect_40(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CondSelect <t> x y (SETGF cond))
	// cond: is16BitInt(t)
	// result: (CMOVWGTF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGF {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWGTF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y (SETGEF cond))
	// cond: is16BitInt(t)
	// result: (CMOVWGEF y x cond)
	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpAMD64SETGEF {
			break
		}
		cond := v_2.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWGEF)
		v.AddArg(y)
		v.AddArg(x)
		v.AddArg(cond)
		return true
	}
	// match: (CondSelect <t> x y check)
	// cond: !check.Type.IsFlags() && check.Type.Size() == 1
	// result: (CondSelect <t> x y (MOVBQZX <typ.UInt64> check))
	for {
		t := v.Type
		check := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(!check.Type.IsFlags() && check.Type.Size() == 1) {
			break
		}
		v.reset(OpCondSelect)
		v.Type = t
		v.AddArg(x)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBQZX, typ.UInt64)
		v0.AddArg(check)
		v.AddArg(v0)
		return true
	}
	// match: (CondSelect <t> x y check)
	// cond: !check.Type.IsFlags() && check.Type.Size() == 2
	// result: (CondSelect <t> x y (MOVWQZX <typ.UInt64> check))
	for {
		t := v.Type
		check := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(!check.Type.IsFlags() && check.Type.Size() == 2) {
			break
		}
		v.reset(OpCondSelect)
		v.Type = t
		v.AddArg(x)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWQZX, typ.UInt64)
		v0.AddArg(check)
		v.AddArg(v0)
		return true
	}
	// match: (CondSelect <t> x y check)
	// cond: !check.Type.IsFlags() && check.Type.Size() == 4
	// result: (CondSelect <t> x y (MOVLQZX <typ.UInt64> check))
	for {
		t := v.Type
		check := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(!check.Type.IsFlags() && check.Type.Size() == 4) {
			break
		}
		v.reset(OpCondSelect)
		v.Type = t
		v.AddArg(x)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLQZX, typ.UInt64)
		v0.AddArg(check)
		v.AddArg(v0)
		return true
	}
	// match: (CondSelect <t> x y check)
	// cond: !check.Type.IsFlags() && check.Type.Size() == 8 && (is64BitInt(t) || isPtr(t))
	// result: (CMOVQNE y x (CMPQconst [0] check))
	for {
		t := v.Type
		check := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(!check.Type.IsFlags() && check.Type.Size() == 8 && (is64BitInt(t) || isPtr(t))) {
			break
		}
		v.reset(OpAMD64CMOVQNE)
		v.AddArg(y)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(check)
		v.AddArg(v0)
		return true
	}
	// match: (CondSelect <t> x y check)
	// cond: !check.Type.IsFlags() && check.Type.Size() == 8 && is32BitInt(t)
	// result: (CMOVLNE y x (CMPQconst [0] check))
	for {
		t := v.Type
		check := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(!check.Type.IsFlags() && check.Type.Size() == 8 && is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVLNE)
		v.AddArg(y)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(check)
		v.AddArg(v0)
		return true
	}
	// match: (CondSelect <t> x y check)
	// cond: !check.Type.IsFlags() && check.Type.Size() == 8 && is16BitInt(t)
	// result: (CMOVWNE y x (CMPQconst [0] check))
	for {
		t := v.Type
		check := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(!check.Type.IsFlags() && check.Type.Size() == 8 && is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64CMOVWNE)
		v.AddArg(y)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(check)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpConst16_0(v *Value) bool {
	// match: (Const16 [val])
	// result: (MOVLconst [val])
	for {
		val := v.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueAMD64_OpConst32_0(v *Value) bool {
	// match: (Const32 [val])
	// result: (MOVLconst [val])
	for {
		val := v.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueAMD64_OpConst32F_0(v *Value) bool {
	// match: (Const32F [val])
	// result: (MOVSSconst [val])
	for {
		val := v.AuxInt
		v.reset(OpAMD64MOVSSconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueAMD64_OpConst64_0(v *Value) bool {
	// match: (Const64 [val])
	// result: (MOVQconst [val])
	for {
		val := v.AuxInt
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueAMD64_OpConst64F_0(v *Value) bool {
	// match: (Const64F [val])
	// result: (MOVSDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpAMD64MOVSDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueAMD64_OpConst8_0(v *Value) bool {
	// match: (Const8 [val])
	// result: (MOVLconst [val])
	for {
		val := v.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueAMD64_OpConstBool_0(v *Value) bool {
	// match: (ConstBool [b])
	// result: (MOVLconst [b])
	for {
		b := v.AuxInt
		v.reset(OpAMD64MOVLconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueAMD64_OpConstNil_0(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVQconst [0])
	for {
		v.reset(OpAMD64MOVQconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueAMD64_OpCtz16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz16 x)
	// result: (BSFL (BTSLconst <typ.UInt32> [16] x))
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSFL)
		v0 := b.NewValue0(v.Pos, OpAMD64BTSLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpCtz16NonZero_0(v *Value) bool {
	// match: (Ctz16NonZero x)
	// result: (BSFL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSFL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCtz32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz32 x)
	// result: (Select0 (BSFQ (BTSQconst <typ.UInt64> [32] x)))
	for {
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64BSFQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v1 := b.NewValue0(v.Pos, OpAMD64BTSQconst, typ.UInt64)
		v1.AuxInt = 32
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpCtz32NonZero_0(v *Value) bool {
	// match: (Ctz32NonZero x)
	// result: (BSFL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSFL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCtz64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz64 <t> x)
	// result: (CMOVQEQ (Select0 <t> (BSFQ x)) (MOVQconst <t> [64]) (Select1 <types.TypeFlags> (BSFQ x)))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpAMD64CMOVQEQ)
		v0 := b.NewValue0(v.Pos, OpSelect0, t)
		v1 := b.NewValue0(v.Pos, OpAMD64BSFQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQconst, t)
		v2.AuxInt = 64
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v4 := b.NewValue0(v.Pos, OpAMD64BSFQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v4.AddArg(x)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueAMD64_OpCtz64NonZero_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz64NonZero x)
	// result: (Select0 (BSFQ x))
	for {
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64BSFQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpCtz8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz8 x)
	// result: (BSFL (BTSLconst <typ.UInt32> [ 8] x))
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSFL)
		v0 := b.NewValue0(v.Pos, OpAMD64BTSLconst, typ.UInt32)
		v0.AuxInt = 8
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpCtz8NonZero_0(v *Value) bool {
	// match: (Ctz8NonZero x)
	// result: (BSFL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64BSFL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt32Fto32_0(v *Value) bool {
	// match: (Cvt32Fto32 x)
	// result: (CVTTSS2SL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTTSS2SL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt32Fto64_0(v *Value) bool {
	// match: (Cvt32Fto64 x)
	// result: (CVTTSS2SQ x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTTSS2SQ)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt32Fto64F_0(v *Value) bool {
	// match: (Cvt32Fto64F x)
	// result: (CVTSS2SD x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTSS2SD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt32to32F_0(v *Value) bool {
	// match: (Cvt32to32F x)
	// result: (CVTSL2SS x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTSL2SS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt32to64F_0(v *Value) bool {
	// match: (Cvt32to64F x)
	// result: (CVTSL2SD x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTSL2SD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt64Fto32_0(v *Value) bool {
	// match: (Cvt64Fto32 x)
	// result: (CVTTSD2SL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTTSD2SL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt64Fto32F_0(v *Value) bool {
	// match: (Cvt64Fto32F x)
	// result: (CVTSD2SS x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTSD2SS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt64Fto64_0(v *Value) bool {
	// match: (Cvt64Fto64 x)
	// result: (CVTTSD2SQ x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTTSD2SQ)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt64to32F_0(v *Value) bool {
	// match: (Cvt64to32F x)
	// result: (CVTSQ2SS x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTSQ2SS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpCvt64to64F_0(v *Value) bool {
	// match: (Cvt64to64F x)
	// result: (CVTSQ2SD x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64CVTSQ2SD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpDiv128u_0(v *Value) bool {
	// match: (Div128u xhi xlo y)
	// result: (DIVQU2 xhi xlo y)
	for {
		y := v.Args[2]
		xhi := v.Args[0]
		xlo := v.Args[1]
		v.reset(OpAMD64DIVQU2)
		v.AddArg(xhi)
		v.AddArg(xlo)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpDiv16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 [a] x y)
	// result: (Select0 (DIVW [a] x y))
	for {
		a := v.AuxInt
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVW, types.NewTuple(typ.Int16, typ.Int16))
		v0.AuxInt = a
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpDiv16u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (Select0 (DIVWU x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVWU, types.NewTuple(typ.UInt16, typ.UInt16))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpDiv32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 [a] x y)
	// result: (Select0 (DIVL [a] x y))
	for {
		a := v.AuxInt
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVL, types.NewTuple(typ.Int32, typ.Int32))
		v0.AuxInt = a
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpDiv32F_0(v *Value) bool {
	// match: (Div32F x y)
	// result: (DIVSS x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64DIVSS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpDiv32u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (Select0 (DIVLU x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVLU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpDiv64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div64 [a] x y)
	// result: (Select0 (DIVQ [a] x y))
	for {
		a := v.AuxInt
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVQ, types.NewTuple(typ.Int64, typ.Int64))
		v0.AuxInt = a
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpDiv64F_0(v *Value) bool {
	// match: (Div64F x y)
	// result: (DIVSD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64DIVSD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpDiv64u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div64u x y)
	// result: (Select0 (DIVQU x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVQU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpDiv8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (Select0 (DIVW (SignExt8to16 x) (SignExt8to16 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVW, types.NewTuple(typ.Int16, typ.Int16))
		v1 := b.NewValue0(v.Pos, OpSignExt8to16, typ.Int16)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to16, typ.Int16)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpDiv8u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (Select0 (DIVWU (ZeroExt8to16 x) (ZeroExt8to16 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVWU, types.NewTuple(typ.UInt16, typ.UInt16))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to16, typ.UInt16)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to16, typ.UInt16)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpEq16_0(v *Value) bool {
	b := v.Block
	// match: (Eq16 x y)
	// result: (SETEQ (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETEQ)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpEq32_0(v *Value) bool {
	b := v.Block
	// match: (Eq32 x y)
	// result: (SETEQ (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETEQ)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpEq32F_0(v *Value) bool {
	b := v.Block
	// match: (Eq32F x y)
	// result: (SETEQF (UCOMISS x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETEQF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISS, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpEq64_0(v *Value) bool {
	b := v.Block
	// match: (Eq64 x y)
	// result: (SETEQ (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETEQ)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpEq64F_0(v *Value) bool {
	b := v.Block
	// match: (Eq64F x y)
	// result: (SETEQF (UCOMISD x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETEQF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpEq8_0(v *Value) bool {
	b := v.Block
	// match: (Eq8 x y)
	// result: (SETEQ (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETEQ)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpEqB_0(v *Value) bool {
	b := v.Block
	// match: (EqB x y)
	// result: (SETEQ (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETEQ)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpEqPtr_0(v *Value) bool {
	b := v.Block
	// match: (EqPtr x y)
	// result: (SETEQ (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETEQ)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpFMA_0(v *Value) bool {
	// match: (FMA x y z)
	// result: (VFMADD231SD z x y)
	for {
		z := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpAMD64VFMADD231SD)
		v.AddArg(z)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpFloor_0(v *Value) bool {
	// match: (Floor x)
	// result: (ROUNDSD [1] x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64ROUNDSD)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpGeq16_0(v *Value) bool {
	b := v.Block
	// match: (Geq16 x y)
	// result: (SETGE (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq16U_0(v *Value) bool {
	b := v.Block
	// match: (Geq16U x y)
	// result: (SETAE (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETAE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq32_0(v *Value) bool {
	b := v.Block
	// match: (Geq32 x y)
	// result: (SETGE (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq32F_0(v *Value) bool {
	b := v.Block
	// match: (Geq32F x y)
	// result: (SETGEF (UCOMISS x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGEF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISS, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq32U_0(v *Value) bool {
	b := v.Block
	// match: (Geq32U x y)
	// result: (SETAE (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETAE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq64_0(v *Value) bool {
	b := v.Block
	// match: (Geq64 x y)
	// result: (SETGE (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq64F_0(v *Value) bool {
	b := v.Block
	// match: (Geq64F x y)
	// result: (SETGEF (UCOMISD x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGEF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq64U_0(v *Value) bool {
	b := v.Block
	// match: (Geq64U x y)
	// result: (SETAE (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETAE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq8_0(v *Value) bool {
	b := v.Block
	// match: (Geq8 x y)
	// result: (SETGE (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGeq8U_0(v *Value) bool {
	b := v.Block
	// match: (Geq8U x y)
	// result: (SETAE (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETAE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGetCallerPC_0(v *Value) bool {
	// match: (GetCallerPC)
	// result: (LoweredGetCallerPC)
	for {
		v.reset(OpAMD64LoweredGetCallerPC)
		return true
	}
}
func rewriteValueAMD64_OpGetCallerSP_0(v *Value) bool {
	// match: (GetCallerSP)
	// result: (LoweredGetCallerSP)
	for {
		v.reset(OpAMD64LoweredGetCallerSP)
		return true
	}
}
func rewriteValueAMD64_OpGetClosurePtr_0(v *Value) bool {
	// match: (GetClosurePtr)
	// result: (LoweredGetClosurePtr)
	for {
		v.reset(OpAMD64LoweredGetClosurePtr)
		return true
	}
}
func rewriteValueAMD64_OpGetG_0(v *Value) bool {
	// match: (GetG mem)
	// result: (LoweredGetG mem)
	for {
		mem := v.Args[0]
		v.reset(OpAMD64LoweredGetG)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpGreater16_0(v *Value) bool {
	b := v.Block
	// match: (Greater16 x y)
	// result: (SETG (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETG)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater16U_0(v *Value) bool {
	b := v.Block
	// match: (Greater16U x y)
	// result: (SETA (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETA)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater32_0(v *Value) bool {
	b := v.Block
	// match: (Greater32 x y)
	// result: (SETG (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETG)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater32F_0(v *Value) bool {
	b := v.Block
	// match: (Greater32F x y)
	// result: (SETGF (UCOMISS x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISS, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater32U_0(v *Value) bool {
	b := v.Block
	// match: (Greater32U x y)
	// result: (SETA (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETA)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater64_0(v *Value) bool {
	b := v.Block
	// match: (Greater64 x y)
	// result: (SETG (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETG)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater64F_0(v *Value) bool {
	b := v.Block
	// match: (Greater64F x y)
	// result: (SETGF (UCOMISD x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater64U_0(v *Value) bool {
	b := v.Block
	// match: (Greater64U x y)
	// result: (SETA (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETA)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater8_0(v *Value) bool {
	b := v.Block
	// match: (Greater8 x y)
	// result: (SETG (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETG)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpGreater8U_0(v *Value) bool {
	b := v.Block
	// match: (Greater8U x y)
	// result: (SETA (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETA)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpHmul32_0(v *Value) bool {
	// match: (Hmul32 x y)
	// result: (HMULL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64HMULL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpHmul32u_0(v *Value) bool {
	// match: (Hmul32u x y)
	// result: (HMULLU x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64HMULLU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpHmul64_0(v *Value) bool {
	// match: (Hmul64 x y)
	// result: (HMULQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64HMULQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpHmul64u_0(v *Value) bool {
	// match: (Hmul64u x y)
	// result: (HMULQU x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64HMULQU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpInterCall_0(v *Value) bool {
	// match: (InterCall [argwid] entry mem)
	// result: (CALLinter [argwid] entry mem)
	for {
		argwid := v.AuxInt
		mem := v.Args[1]
		entry := v.Args[0]
		v.reset(OpAMD64CALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpIsInBounds_0(v *Value) bool {
	b := v.Block
	// match: (IsInBounds idx len)
	// result: (SETB (CMPQ idx len))
	for {
		len := v.Args[1]
		idx := v.Args[0]
		v.reset(OpAMD64SETB)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	// match: (IsNonNil p)
	// result: (SETNE (TESTQ p p))
	for {
		p := v.Args[0]
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64TESTQ, types.TypeFlags)
		v0.AddArg(p)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	// match: (IsSliceInBounds idx len)
	// result: (SETBE (CMPQ idx len))
	for {
		len := v.Args[1]
		idx := v.Args[0]
		v.reset(OpAMD64SETBE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq16_0(v *Value) bool {
	b := v.Block
	// match: (Leq16 x y)
	// result: (SETLE (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETLE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq16U_0(v *Value) bool {
	b := v.Block
	// match: (Leq16U x y)
	// result: (SETBE (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETBE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq32_0(v *Value) bool {
	b := v.Block
	// match: (Leq32 x y)
	// result: (SETLE (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETLE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq32F_0(v *Value) bool {
	b := v.Block
	// match: (Leq32F x y)
	// result: (SETGEF (UCOMISS y x))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGEF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISS, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq32U_0(v *Value) bool {
	b := v.Block
	// match: (Leq32U x y)
	// result: (SETBE (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETBE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq64_0(v *Value) bool {
	b := v.Block
	// match: (Leq64 x y)
	// result: (SETLE (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETLE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq64F_0(v *Value) bool {
	b := v.Block
	// match: (Leq64F x y)
	// result: (SETGEF (UCOMISD y x))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGEF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISD, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq64U_0(v *Value) bool {
	b := v.Block
	// match: (Leq64U x y)
	// result: (SETBE (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETBE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq8_0(v *Value) bool {
	b := v.Block
	// match: (Leq8 x y)
	// result: (SETLE (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETLE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLeq8U_0(v *Value) bool {
	b := v.Block
	// match: (Leq8U x y)
	// result: (SETBE (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETBE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess16_0(v *Value) bool {
	b := v.Block
	// match: (Less16 x y)
	// result: (SETL (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETL)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess16U_0(v *Value) bool {
	b := v.Block
	// match: (Less16U x y)
	// result: (SETB (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETB)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess32_0(v *Value) bool {
	b := v.Block
	// match: (Less32 x y)
	// result: (SETL (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETL)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess32F_0(v *Value) bool {
	b := v.Block
	// match: (Less32F x y)
	// result: (SETGF (UCOMISS y x))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISS, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess32U_0(v *Value) bool {
	b := v.Block
	// match: (Less32U x y)
	// result: (SETB (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETB)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess64_0(v *Value) bool {
	b := v.Block
	// match: (Less64 x y)
	// result: (SETL (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETL)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess64F_0(v *Value) bool {
	b := v.Block
	// match: (Less64F x y)
	// result: (SETGF (UCOMISD y x))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETGF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISD, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess64U_0(v *Value) bool {
	b := v.Block
	// match: (Less64U x y)
	// result: (SETB (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETB)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess8_0(v *Value) bool {
	b := v.Block
	// match: (Less8 x y)
	// result: (SETL (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETL)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLess8U_0(v *Value) bool {
	b := v.Block
	// match: (Less8U x y)
	// result: (SETB (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETB)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpLoad_0(v *Value) bool {
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVQload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpAMD64MOVQload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t)
	// result: (MOVLload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is32BitInt(t)) {
			break
		}
		v.reset(OpAMD64MOVLload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t)
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is16BitInt(t)) {
			break
		}
		v.reset(OpAMD64MOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (t.IsBoolean() || is8BitInt(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(t.IsBoolean() || is8BitInt(t)) {
			break
		}
		v.reset(OpAMD64MOVBload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (MOVSSload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpAMD64MOVSSload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (MOVSDload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpAMD64MOVSDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLocalAddr_0(v *Value) bool {
	// match: (LocalAddr {sym} base _)
	// result: (LEAQ {sym} base)
	for {
		sym := v.Aux
		_ = v.Args[1]
		base := v.Args[0]
		v.reset(OpAMD64LEAQ)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueAMD64_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	// match: (Lsh16x16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPWconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh16x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	// match: (Lsh16x32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPLconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh16x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	// match: (Lsh16x64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPQconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh16x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	// match: (Lsh16x8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPBconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh16x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	// match: (Lsh32x16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPWconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh32x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	// match: (Lsh32x32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPLconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh32x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	// match: (Lsh32x64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPQconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh32x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	// match: (Lsh32x8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPBconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh32x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	// match: (Lsh64x16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDQ (SHLQ <t> x y) (SBBQcarrymask <t> (CMPWconst y [64])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQ, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh64x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	// match: (Lsh64x32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDQ (SHLQ <t> x y) (SBBQcarrymask <t> (CMPLconst y [64])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQ, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh64x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	// match: (Lsh64x64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDQ (SHLQ <t> x y) (SBBQcarrymask <t> (CMPQconst y [64])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQ, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	// match: (Lsh64x8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDQ (SHLQ <t> x y) (SBBQcarrymask <t> (CMPBconst y [64])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLQ, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh64x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	// match: (Lsh8x16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPWconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh8x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	// match: (Lsh8x32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPLconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh8x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	// match: (Lsh8x64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPQconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh8x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	// match: (Lsh8x8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHLL <t> x y) (SBBLcarrymask <t> (CMPBconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Lsh8x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SHLL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpMod16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 [a] x y)
	// result: (Select1 (DIVW [a] x y))
	for {
		a := v.AuxInt
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVW, types.NewTuple(typ.Int16, typ.Int16))
		v0.AuxInt = a
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpMod16u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Select1 (DIVWU x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVWU, types.NewTuple(typ.UInt16, typ.UInt16))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpMod32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 [a] x y)
	// result: (Select1 (DIVL [a] x y))
	for {
		a := v.AuxInt
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVL, types.NewTuple(typ.Int32, typ.Int32))
		v0.AuxInt = a
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpMod32u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (Select1 (DIVLU x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVLU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpMod64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64 [a] x y)
	// result: (Select1 (DIVQ [a] x y))
	for {
		a := v.AuxInt
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVQ, types.NewTuple(typ.Int64, typ.Int64))
		v0.AuxInt = a
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpMod64u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64u x y)
	// result: (Select1 (DIVQU x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVQU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpMod8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Select1 (DIVW (SignExt8to16 x) (SignExt8to16 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVW, types.NewTuple(typ.Int16, typ.Int16))
		v1 := b.NewValue0(v.Pos, OpSignExt8to16, typ.Int16)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to16, typ.Int16)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpMod8u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Select1 (DIVWU (ZeroExt8to16 x) (ZeroExt8to16 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpAMD64DIVWU, types.NewTuple(typ.UInt16, typ.UInt16))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to16, typ.UInt16)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to16, typ.UInt16)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpMove_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBload, typ.UInt8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWload, typ.UInt16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVLstore dst (MOVLload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVLstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [8] dst src mem)
	// result: (MOVQstore dst (MOVQload src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVQstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [16] dst src mem)
	// cond: config.useSSE
	// result: (MOVOstore dst (MOVOload src mem) mem)
	for {
		if v.AuxInt != 16 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVOstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVOload, types.TypeInt128)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [16] dst src mem)
	// cond: !config.useSSE
	// result: (MOVQstore [8] dst (MOVQload [8] src mem) (MOVQstore dst (MOVQload src mem) mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(!config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [32] dst src mem)
	// result: (Move [16] (OffPtr <dst.Type> dst [16]) (OffPtr <src.Type> src [16]) (Move [16] dst src mem))
	for {
		if v.AuxInt != 32 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpMove)
		v.AuxInt = 16
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = 16
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = 16
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMove, types.TypeMem)
		v2.AuxInt = 16
		v2.AddArg(dst)
		v2.AddArg(src)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}
	// match: (Move [48] dst src mem)
	// cond: config.useSSE
	// result: (Move [32] (OffPtr <dst.Type> dst [16]) (OffPtr <src.Type> src [16]) (Move [16] dst src mem))
	for {
		if v.AuxInt != 48 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(config.useSSE) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = 16
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = 16
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMove, types.TypeMem)
		v2.AuxInt = 16
		v2.AddArg(dst)
		v2.AddArg(src)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}
	// match: (Move [64] dst src mem)
	// cond: config.useSSE
	// result: (Move [32] (OffPtr <dst.Type> dst [32]) (OffPtr <src.Type> src [32]) (Move [32] dst src mem))
	for {
		if v.AuxInt != 64 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(config.useSSE) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = 32
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = 32
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMove, types.TypeMem)
		v2.AuxInt = 32
		v2.AddArg(dst)
		v2.AddArg(src)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueAMD64_OpMove_10(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [3] dst src mem)
	// result: (MOVBstore [2] dst (MOVBload [2] src mem) (MOVWstore dst (MOVWload src mem) mem))
	for {
		if v.AuxInt != 3 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVWstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVWload, typ.UInt16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [5] dst src mem)
	// result: (MOVBstore [4] dst (MOVBload [4] src mem) (MOVLstore dst (MOVLload src mem) mem))
	for {
		if v.AuxInt != 5 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBload, typ.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVLstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [6] dst src mem)
	// result: (MOVWstore [4] dst (MOVWload [4] src mem) (MOVLstore dst (MOVLload src mem) mem))
	for {
		if v.AuxInt != 6 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWload, typ.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVLstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [7] dst src mem)
	// result: (MOVLstore [3] dst (MOVLload [3] src mem) (MOVLstore dst (MOVLload src mem) mem))
	for {
		if v.AuxInt != 7 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVLstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [9] dst src mem)
	// result: (MOVBstore [8] dst (MOVBload [8] src mem) (MOVQstore dst (MOVQload src mem) mem))
	for {
		if v.AuxInt != 9 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVBstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBload, typ.UInt8)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [10] dst src mem)
	// result: (MOVWstore [8] dst (MOVWload [8] src mem) (MOVQstore dst (MOVQload src mem) mem))
	for {
		if v.AuxInt != 10 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVWstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWload, typ.UInt16)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [12] dst src mem)
	// result: (MOVLstore [8] dst (MOVLload [8] src mem) (MOVQstore dst (MOVQload src mem) mem))
	for {
		if v.AuxInt != 12 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpAMD64MOVLstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s == 11 || s >= 13 && s <= 15
	// result: (MOVQstore [s-8] dst (MOVQload [s-8] src mem) (MOVQstore dst (MOVQload src mem) mem))
	for {
		s := v.AuxInt
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(s == 11 || s >= 13 && s <= 15) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AuxInt = s - 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v0.AuxInt = s - 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 16 && s%16 != 0 && s%16 <= 8
	// result: (Move [s-s%16] (OffPtr <dst.Type> dst [s%16]) (OffPtr <src.Type> src [s%16]) (MOVQstore dst (MOVQload src mem) mem))
	for {
		s := v.AuxInt
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(s > 16 && s%16 != 0 && s%16 <= 8) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s - s%16
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = s % 16
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = s % 16
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQstore, types.TypeMem)
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v3.AddArg(src)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 16 && s%16 != 0 && s%16 > 8 && config.useSSE
	// result: (Move [s-s%16] (OffPtr <dst.Type> dst [s%16]) (OffPtr <src.Type> src [s%16]) (MOVOstore dst (MOVOload src mem) mem))
	for {
		s := v.AuxInt
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(s > 16 && s%16 != 0 && s%16 > 8 && config.useSSE) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s - s%16
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = s % 16
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = s % 16
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVOstore, types.TypeMem)
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpAMD64MOVOload, types.TypeInt128)
		v3.AddArg(src)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueAMD64_OpMove_20(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [s] dst src mem)
	// cond: s > 16 && s%16 != 0 && s%16 > 8 && !config.useSSE
	// result: (Move [s-s%16] (OffPtr <dst.Type> dst [s%16]) (OffPtr <src.Type> src [s%16]) (MOVQstore [8] dst (MOVQload [8] src mem) (MOVQstore dst (MOVQload src mem) mem)))
	for {
		s := v.AuxInt
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(s > 16 && s%16 != 0 && s%16 > 8 && !config.useSSE) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s - s%16
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = s % 16
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = s % 16
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQstore, types.TypeMem)
		v2.AuxInt = 8
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v3.AuxInt = 8
		v3.AddArg(src)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpAMD64MOVQstore, types.TypeMem)
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v5.AddArg(src)
		v5.AddArg(mem)
		v4.AddArg(v5)
		v4.AddArg(mem)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 64 && s <= 16*64 && s%16 == 0 && !config.noDuffDevice
	// result: (DUFFCOPY [14*(64-s/16)] dst src mem)
	for {
		s := v.AuxInt
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(s > 64 && s <= 16*64 && s%16 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpAMD64DUFFCOPY)
		v.AuxInt = 14 * (64 - s/16)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: (s > 16*64 || config.noDuffDevice) && s%8 == 0
	// result: (REPMOVSQ dst src (MOVQconst [s/8]) mem)
	for {
		s := v.AuxInt
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !((s > 16*64 || config.noDuffDevice) && s%8 == 0) {
			break
		}
		v.reset(OpAMD64REPMOVSQ)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQconst, typ.UInt64)
		v0.AuxInt = s / 8
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpMul16_0(v *Value) bool {
	// match: (Mul16 x y)
	// result: (MULL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64MULL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpMul32_0(v *Value) bool {
	// match: (Mul32 x y)
	// result: (MULL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64MULL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpMul32F_0(v *Value) bool {
	// match: (Mul32F x y)
	// result: (MULSS x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64MULSS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpMul64_0(v *Value) bool {
	// match: (Mul64 x y)
	// result: (MULQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64MULQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpMul64F_0(v *Value) bool {
	// match: (Mul64F x y)
	// result: (MULSD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64MULSD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpMul64uhilo_0(v *Value) bool {
	// match: (Mul64uhilo x y)
	// result: (MULQU2 x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64MULQU2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpMul8_0(v *Value) bool {
	// match: (Mul8 x y)
	// result: (MULL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64MULL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpNeg16_0(v *Value) bool {
	// match: (Neg16 x)
	// result: (NEGL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64NEGL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpNeg32_0(v *Value) bool {
	// match: (Neg32 x)
	// result: (NEGL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64NEGL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpNeg32F_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg32F x)
	// result: (PXOR x (MOVSSconst <typ.Float32> [auxFrom32F(float32(math.Copysign(0, -1)))]))
	for {
		x := v.Args[0]
		v.reset(OpAMD64PXOR)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVSSconst, typ.Float32)
		v0.AuxInt = auxFrom32F(float32(math.Copysign(0, -1)))
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeg64_0(v *Value) bool {
	// match: (Neg64 x)
	// result: (NEGQ x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64NEGQ)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpNeg64F_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg64F x)
	// result: (PXOR x (MOVSDconst <typ.Float64> [auxFrom64F(math.Copysign(0, -1))]))
	for {
		x := v.Args[0]
		v.reset(OpAMD64PXOR)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVSDconst, typ.Float64)
		v0.AuxInt = auxFrom64F(math.Copysign(0, -1))
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeg8_0(v *Value) bool {
	// match: (Neg8 x)
	// result: (NEGL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64NEGL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpNeq16_0(v *Value) bool {
	b := v.Block
	// match: (Neq16 x y)
	// result: (SETNE (CMPW x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeq32_0(v *Value) bool {
	b := v.Block
	// match: (Neq32 x y)
	// result: (SETNE (CMPL x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPL, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeq32F_0(v *Value) bool {
	b := v.Block
	// match: (Neq32F x y)
	// result: (SETNEF (UCOMISS x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETNEF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISS, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeq64_0(v *Value) bool {
	b := v.Block
	// match: (Neq64 x y)
	// result: (SETNE (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeq64F_0(v *Value) bool {
	b := v.Block
	// match: (Neq64F x y)
	// result: (SETNEF (UCOMISD x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETNEF)
		v0 := b.NewValue0(v.Pos, OpAMD64UCOMISD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeq8_0(v *Value) bool {
	b := v.Block
	// match: (Neq8 x y)
	// result: (SETNE (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeqB_0(v *Value) bool {
	b := v.Block
	// match: (NeqB x y)
	// result: (SETNE (CMPB x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPB, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	// match: (NeqPtr x y)
	// result: (SETNE (CMPQ x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SETNE)
		v0 := b.NewValue0(v.Pos, OpAMD64CMPQ, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpNilCheck_0(v *Value) bool {
	// match: (NilCheck ptr mem)
	// result: (LoweredNilCheck ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64LoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpNot_0(v *Value) bool {
	// match: (Not x)
	// result: (XORLconst [1] x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64XORLconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpOffPtr_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (OffPtr [off] ptr)
	// cond: is32Bit(off)
	// result: (ADDQconst [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if !(is32Bit(off)) {
			break
		}
		v.reset(OpAMD64ADDQconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADDQ (MOVQconst [off]) ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpAMD64ADDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQconst, typ.UInt64)
		v0.AuxInt = off
		v.AddArg(v0)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueAMD64_OpOr16_0(v *Value) bool {
	// match: (Or16 x y)
	// result: (ORL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpOr32_0(v *Value) bool {
	// match: (Or32 x y)
	// result: (ORL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpOr64_0(v *Value) bool {
	// match: (Or64 x y)
	// result: (ORQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ORQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpOr8_0(v *Value) bool {
	// match: (Or8 x y)
	// result: (ORL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpOrB_0(v *Value) bool {
	// match: (OrB x y)
	// result: (ORL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64ORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpPanicBounds_0(v *Value) bool {
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpAMD64LoweredPanicBoundsA)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpAMD64LoweredPanicBoundsB)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpAMD64LoweredPanicBoundsC)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpPopCount16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount16 x)
	// result: (POPCNTL (MOVWQZX <typ.UInt32> x))
	for {
		x := v.Args[0]
		v.reset(OpAMD64POPCNTL)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWQZX, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpPopCount32_0(v *Value) bool {
	// match: (PopCount32 x)
	// result: (POPCNTL x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64POPCNTL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpPopCount64_0(v *Value) bool {
	// match: (PopCount64 x)
	// result: (POPCNTQ x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64POPCNTQ)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpPopCount8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount8 x)
	// result: (POPCNTL (MOVBQZX <typ.UInt32> x))
	for {
		x := v.Args[0]
		v.reset(OpAMD64POPCNTL)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBQZX, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpRotateLeft16_0(v *Value) bool {
	// match: (RotateLeft16 a b)
	// result: (ROLW a b)
	for {
		b := v.Args[1]
		a := v.Args[0]
		v.reset(OpAMD64ROLW)
		v.AddArg(a)
		v.AddArg(b)
		return true
	}
}
func rewriteValueAMD64_OpRotateLeft32_0(v *Value) bool {
	// match: (RotateLeft32 a b)
	// result: (ROLL a b)
	for {
		b := v.Args[1]
		a := v.Args[0]
		v.reset(OpAMD64ROLL)
		v.AddArg(a)
		v.AddArg(b)
		return true
	}
}
func rewriteValueAMD64_OpRotateLeft64_0(v *Value) bool {
	// match: (RotateLeft64 a b)
	// result: (ROLQ a b)
	for {
		b := v.Args[1]
		a := v.Args[0]
		v.reset(OpAMD64ROLQ)
		v.AddArg(a)
		v.AddArg(b)
		return true
	}
}
func rewriteValueAMD64_OpRotateLeft8_0(v *Value) bool {
	// match: (RotateLeft8 a b)
	// result: (ROLB a b)
	for {
		b := v.Args[1]
		a := v.Args[0]
		v.reset(OpAMD64ROLB)
		v.AddArg(a)
		v.AddArg(b)
		return true
	}
}
func rewriteValueAMD64_OpRound32F_0(v *Value) bool {
	// match: (Round32F x)
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpRound64F_0(v *Value) bool {
	// match: (Round64F x)
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpRoundToEven_0(v *Value) bool {
	// match: (RoundToEven x)
	// result: (ROUNDSD [0] x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64ROUNDSD)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	// match: (Rsh16Ux16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRW <t> x y) (SBBLcarrymask <t> (CMPWconst y [16])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v2.AuxInt = 16
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh16Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRW x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh16Ux32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRW <t> x y) (SBBLcarrymask <t> (CMPLconst y [16])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v2.AuxInt = 16
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh16Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRW x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	// match: (Rsh16Ux64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRW <t> x y) (SBBLcarrymask <t> (CMPQconst y [16])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v2.AuxInt = 16
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh16Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRW x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	// match: (Rsh16Ux8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRW <t> x y) (SBBLcarrymask <t> (CMPBconst y [16])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v2.AuxInt = 16
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh16Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRW x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	// match: (Rsh16x16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARW <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPWconst y [16])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARW)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v3.AuxInt = 16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SARW x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh16x32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARW <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPLconst y [16])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARW)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v3.AuxInt = 16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SARW x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	// match: (Rsh16x64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARW <t> x (ORQ <y.Type> y (NOTQ <y.Type> (SBBQcarrymask <y.Type> (CMPQconst y [16])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARW)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORQ, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTQ, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v3.AuxInt = 16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SARW x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	// match: (Rsh16x8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARW <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPBconst y [16])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARW)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v3.AuxInt = 16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SARW x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32Ux16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRL <t> x y) (SBBLcarrymask <t> (CMPWconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh32Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32Ux32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRL <t> x y) (SBBLcarrymask <t> (CMPLconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh32Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32Ux64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRL <t> x y) (SBBLcarrymask <t> (CMPQconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh32Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32Ux8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRL <t> x y) (SBBLcarrymask <t> (CMPBconst y [32])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh32Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32x16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARL <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPWconst y [32])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARL)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v3.AuxInt = 32
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SARL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32x32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARL <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPLconst y [32])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARL)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v3.AuxInt = 32
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SARL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32x64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARL <t> x (ORQ <y.Type> y (NOTQ <y.Type> (SBBQcarrymask <y.Type> (CMPQconst y [32])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARL)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORQ, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTQ, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v3.AuxInt = 32
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SARL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32x8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARL <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPBconst y [32])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARL)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v3.AuxInt = 32
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SARL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	// match: (Rsh64Ux16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDQ (SHRQ <t> x y) (SBBQcarrymask <t> (CMPWconst y [64])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRQ, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh64Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh64Ux32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDQ (SHRQ <t> x y) (SBBQcarrymask <t> (CMPLconst y [64])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRQ, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh64Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	// match: (Rsh64Ux64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDQ (SHRQ <t> x y) (SBBQcarrymask <t> (CMPQconst y [64])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRQ, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh64Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	// match: (Rsh64Ux8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDQ (SHRQ <t> x y) (SBBQcarrymask <t> (CMPBconst y [64])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDQ)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRQ, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh64Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	// match: (Rsh64x16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARQ <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPWconst y [64])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SARQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh64x32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARQ <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPLconst y [64])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SARQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	// match: (Rsh64x64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARQ <t> x (ORQ <y.Type> y (NOTQ <y.Type> (SBBQcarrymask <y.Type> (CMPQconst y [64])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORQ, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTQ, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SARQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	// match: (Rsh64x8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARQ <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPBconst y [64])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SARQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	// match: (Rsh8Ux16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRB <t> x y) (SBBLcarrymask <t> (CMPWconst y [8])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v2.AuxInt = 8
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh8Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh8Ux32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRB <t> x y) (SBBLcarrymask <t> (CMPLconst y [8])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v2.AuxInt = 8
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh8Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	// match: (Rsh8Ux64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRB <t> x y) (SBBLcarrymask <t> (CMPQconst y [8])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v2.AuxInt = 8
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh8Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	// match: (Rsh8Ux8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (ANDL (SHRB <t> x y) (SBBLcarrymask <t> (CMPBconst y [8])))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64ANDL)
		v0 := b.NewValue0(v.Pos, OpAMD64SHRB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v2.AuxInt = 8
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
	// match: (Rsh8Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SHRB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SHRB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	// match: (Rsh8x16 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARB <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPWconst y [8])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARB)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPWconst, types.TypeFlags)
		v3.AuxInt = 8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SARB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh8x32 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARB <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPLconst y [8])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARB)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPLconst, types.TypeFlags)
		v3.AuxInt = 8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SARB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	// match: (Rsh8x64 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARB <t> x (ORQ <y.Type> y (NOTQ <y.Type> (SBBQcarrymask <y.Type> (CMPQconst y [8])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARB)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORQ, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTQ, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPQconst, types.TypeFlags)
		v3.AuxInt = 8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SARB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	// match: (Rsh8x8 <t> x y)
	// cond: !shiftIsBounded(v)
	// result: (SARB <t> x (ORL <y.Type> y (NOTL <y.Type> (SBBLcarrymask <y.Type> (CMPBconst y [8])))))
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		if !(!shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARB)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAMD64ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpAMD64NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, OpAMD64CMPBconst, types.TypeFlags)
		v3.AuxInt = 8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SARB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpAMD64SARB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueAMD64_OpSelect0_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Select0 (Mul64uover x y))
	// result: (Select0 <typ.UInt64> (MULQU x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMul64uover {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpSelect0)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpAMD64MULQU, types.NewTuple(typ.UInt64, types.TypeFlags))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 (Mul32uover x y))
	// result: (Select0 <typ.UInt32> (MULLU x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMul32uover {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpSelect0)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpAMD64MULLU, types.NewTuple(typ.UInt32, types.TypeFlags))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 (Add64carry x y c))
	// result: (Select0 <typ.UInt64> (ADCQ x y (Select1 <types.TypeFlags> (NEGLflags c))))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64carry {
			break
		}
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSelect0)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpAMD64ADCQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v0.AddArg(x)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v2 := b.NewValue0(v.Pos, OpAMD64NEGLflags, types.NewTuple(typ.UInt32, types.TypeFlags))
		v2.AddArg(c)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 (Sub64borrow x y c))
	// result: (Select0 <typ.UInt64> (SBBQ x y (Select1 <types.TypeFlags> (NEGLflags c))))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub64borrow {
			break
		}
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSelect0)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpAMD64SBBQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v0.AddArg(x)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v2 := b.NewValue0(v.Pos, OpAMD64NEGLflags, types.NewTuple(typ.UInt32, types.TypeFlags))
		v2.AddArg(c)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 <t> (AddTupleFirst32 val tuple))
	// result: (ADDL val (Select0 <t> tuple))
	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64AddTupleFirst32 {
			break
		}
		tuple := v_0.Args[1]
		val := v_0.Args[0]
		v.reset(OpAMD64ADDL)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpSelect0, t)
		v0.AddArg(tuple)
		v.AddArg(v0)
		return true
	}
	// match: (Select0 <t> (AddTupleFirst64 val tuple))
	// result: (ADDQ val (Select0 <t> tuple))
	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64AddTupleFirst64 {
			break
		}
		tuple := v_0.Args[1]
		val := v_0.Args[0]
		v.reset(OpAMD64ADDQ)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpSelect0, t)
		v0.AddArg(tuple)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueAMD64_OpSelect1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Select1 (Mul64uover x y))
	// result: (SETO (Select1 <types.TypeFlags> (MULQU x y)))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMul64uover {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpAMD64SETO)
		v0 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpAMD64MULQU, types.NewTuple(typ.UInt64, types.TypeFlags))
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select1 (Mul32uover x y))
	// result: (SETO (Select1 <types.TypeFlags> (MULLU x y)))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMul32uover {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpAMD64SETO)
		v0 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpAMD64MULLU, types.NewTuple(typ.UInt32, types.TypeFlags))
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select1 (Add64carry x y c))
	// result: (NEGQ <typ.UInt64> (SBBQcarrymask <typ.UInt64> (Select1 <types.TypeFlags> (ADCQ x y (Select1 <types.TypeFlags> (NEGLflags c))))))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64carry {
			break
		}
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpAMD64NEGQ)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v2 := b.NewValue0(v.Pos, OpAMD64ADCQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v2.AddArg(x)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v4 := b.NewValue0(v.Pos, OpAMD64NEGLflags, types.NewTuple(typ.UInt32, types.TypeFlags))
		v4.AddArg(c)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select1 (Sub64borrow x y c))
	// result: (NEGQ <typ.UInt64> (SBBQcarrymask <typ.UInt64> (Select1 <types.TypeFlags> (SBBQ x y (Select1 <types.TypeFlags> (NEGLflags c))))))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub64borrow {
			break
		}
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpAMD64NEGQ)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpAMD64SBBQcarrymask, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v2 := b.NewValue0(v.Pos, OpAMD64SBBQ, types.NewTuple(typ.UInt64, types.TypeFlags))
		v2.AddArg(x)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSelect1, types.TypeFlags)
		v4 := b.NewValue0(v.Pos, OpAMD64NEGLflags, types.NewTuple(typ.UInt32, types.TypeFlags))
		v4.AddArg(c)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Select1 (NEGLflags (MOVQconst [0])))
	// result: (FlagEQ)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64NEGLflags {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAMD64MOVQconst || v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpAMD64FlagEQ)
		return true
	}
	// match: (Select1 (NEGLflags (NEGQ (SBBQcarrymask x))))
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64NEGLflags {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAMD64NEGQ {
			break
		}
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpAMD64SBBQcarrymask {
			break
		}
		x := v_0_0_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (Select1 (AddTupleFirst32 _ tuple))
	// result: (Select1 tuple)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64AddTupleFirst32 {
			break
		}
		tuple := v_0.Args[1]
		v.reset(OpSelect1)
		v.AddArg(tuple)
		return true
	}
	// match: (Select1 (AddTupleFirst64 _ tuple))
	// result: (Select1 tuple)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAMD64AddTupleFirst64 {
			break
		}
		tuple := v_0.Args[1]
		v.reset(OpSelect1)
		v.AddArg(tuple)
		return true
	}
	return false
}
func rewriteValueAMD64_OpSignExt16to32_0(v *Value) bool {
	// match: (SignExt16to32 x)
	// result: (MOVWQSX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVWQSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpSignExt16to64_0(v *Value) bool {
	// match: (SignExt16to64 x)
	// result: (MOVWQSX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVWQSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpSignExt32to64_0(v *Value) bool {
	// match: (SignExt32to64 x)
	// result: (MOVLQSX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVLQSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpSignExt8to16_0(v *Value) bool {
	// match: (SignExt8to16 x)
	// result: (MOVBQSX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVBQSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpSignExt8to32_0(v *Value) bool {
	// match: (SignExt8to32 x)
	// result: (MOVBQSX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVBQSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpSignExt8to64_0(v *Value) bool {
	// match: (SignExt8to64 x)
	// result: (MOVBQSX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVBQSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpSlicemask_0(v *Value) bool {
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SARQconst (NEGQ <t> x) [63])
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpAMD64SARQconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpAMD64NEGQ, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64_OpSqrt_0(v *Value) bool {
	// match: (Sqrt x)
	// result: (SQRTSD x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64SQRTSD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpStaticCall_0(v *Value) bool {
	// match: (StaticCall [argwid] {target} mem)
	// result: (CALLstatic [argwid] {target} mem)
	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpAMD64CALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpStore_0(v *Value) bool {
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (MOVSDstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpAMD64MOVSDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (MOVSSstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpAMD64MOVSSstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8
	// result: (MOVQstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 8) {
			break
		}
		v.reset(OpAMD64MOVQstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4
	// result: (MOVLstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 4) {
			break
		}
		v.reset(OpAMD64MOVLstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpAMD64MOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpAMD64MOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpSub16_0(v *Value) bool {
	// match: (Sub16 x y)
	// result: (SUBL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpSub32_0(v *Value) bool {
	// match: (Sub32 x y)
	// result: (SUBL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpSub32F_0(v *Value) bool {
	// match: (Sub32F x y)
	// result: (SUBSS x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SUBSS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpSub64_0(v *Value) bool {
	// match: (Sub64 x y)
	// result: (SUBQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SUBQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpSub64F_0(v *Value) bool {
	// match: (Sub64F x y)
	// result: (SUBSD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SUBSD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpSub8_0(v *Value) bool {
	// match: (Sub8 x y)
	// result: (SUBL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpSubPtr_0(v *Value) bool {
	// match: (SubPtr x y)
	// result: (SUBQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64SUBQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpTrunc_0(v *Value) bool {
	// match: (Trunc x)
	// result: (ROUNDSD [3] x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64ROUNDSD)
		v.AuxInt = 3
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpTrunc16to8_0(v *Value) bool {
	// match: (Trunc16to8 x)
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpTrunc32to16_0(v *Value) bool {
	// match: (Trunc32to16 x)
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpTrunc32to8_0(v *Value) bool {
	// match: (Trunc32to8 x)
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpTrunc64to16_0(v *Value) bool {
	// match: (Trunc64to16 x)
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpTrunc64to32_0(v *Value) bool {
	// match: (Trunc64to32 x)
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpTrunc64to8_0(v *Value) bool {
	// match: (Trunc64to8 x)
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpWB_0(v *Value) bool {
	// match: (WB {fn} destptr srcptr mem)
	// result: (LoweredWB {fn} destptr srcptr mem)
	for {
		fn := v.Aux
		mem := v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		v.reset(OpAMD64LoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueAMD64_OpXor16_0(v *Value) bool {
	// match: (Xor16 x y)
	// result: (XORL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64XORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpXor32_0(v *Value) bool {
	// match: (Xor32 x y)
	// result: (XORL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64XORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpXor64_0(v *Value) bool {
	// match: (Xor64 x y)
	// result: (XORQ x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64XORQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpXor8_0(v *Value) bool {
	// match: (Xor8 x y)
	// result: (XORL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpAMD64XORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueAMD64_OpZero_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v.Args[1]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Zero [1] destptr mem)
	// result: (MOVBstoreconst [0] destptr mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] destptr mem)
	// result: (MOVWstoreconst [0] destptr mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [4] destptr mem)
	// result: (MOVLstoreconst [0] destptr mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [8] destptr mem)
	// result: (MOVQstoreconst [0] destptr mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [3] destptr mem)
	// result: (MOVBstoreconst [makeValAndOff(0,2)] destptr (MOVWstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 3 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = makeValAndOff(0, 2)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [5] destptr mem)
	// result: (MOVBstoreconst [makeValAndOff(0,4)] destptr (MOVLstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 5 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		v.reset(OpAMD64MOVBstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [6] destptr mem)
	// result: (MOVWstoreconst [makeValAndOff(0,4)] destptr (MOVLstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 6 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		v.reset(OpAMD64MOVWstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [7] destptr mem)
	// result: (MOVLstoreconst [makeValAndOff(0,3)] destptr (MOVLstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 7 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		v.reset(OpAMD64MOVLstoreconst)
		v.AuxInt = makeValAndOff(0, 3)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: s%8 != 0 && s > 8 && !config.useSSE
	// result: (Zero [s-s%8] (OffPtr <destptr.Type> destptr [s%8]) (MOVQstoreconst [0] destptr mem))
	for {
		s := v.AuxInt
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(s%8 != 0 && s > 8 && !config.useSSE) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = s - s%8
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = s % 8
		v0.AddArg(destptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueAMD64_OpZero_10(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	// match: (Zero [16] destptr mem)
	// cond: !config.useSSE
	// result: (MOVQstoreconst [makeValAndOff(0,8)] destptr (MOVQstoreconst [0] destptr mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(!config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = makeValAndOff(0, 8)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [24] destptr mem)
	// cond: !config.useSSE
	// result: (MOVQstoreconst [makeValAndOff(0,16)] destptr (MOVQstoreconst [makeValAndOff(0,8)] destptr (MOVQstoreconst [0] destptr mem)))
	for {
		if v.AuxInt != 24 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(!config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = makeValAndOff(0, 16)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v0.AuxInt = makeValAndOff(0, 8)
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [32] destptr mem)
	// cond: !config.useSSE
	// result: (MOVQstoreconst [makeValAndOff(0,24)] destptr (MOVQstoreconst [makeValAndOff(0,16)] destptr (MOVQstoreconst [makeValAndOff(0,8)] destptr (MOVQstoreconst [0] destptr mem))))
	for {
		if v.AuxInt != 32 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(!config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = makeValAndOff(0, 24)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v0.AuxInt = makeValAndOff(0, 16)
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v1.AuxInt = makeValAndOff(0, 8)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v2.AuxInt = 0
		v2.AddArg(destptr)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: s > 8 && s < 16 && config.useSSE
	// result: (MOVQstoreconst [makeValAndOff(0,s-8)] destptr (MOVQstoreconst [0] destptr mem))
	for {
		s := v.AuxInt
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(s > 8 && s < 16 && config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVQstoreconst)
		v.AuxInt = makeValAndOff(0, s-8)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: s%16 != 0 && s > 16 && s%16 > 8 && config.useSSE
	// result: (Zero [s-s%16] (OffPtr <destptr.Type> destptr [s%16]) (MOVOstore destptr (MOVOconst [0]) mem))
	for {
		s := v.AuxInt
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(s%16 != 0 && s > 16 && s%16 > 8 && config.useSSE) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = s - s%16
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = s % 16
		v0.AddArg(destptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVOstore, types.TypeMem)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: s%16 != 0 && s > 16 && s%16 <= 8 && config.useSSE
	// result: (Zero [s-s%16] (OffPtr <destptr.Type> destptr [s%16]) (MOVQstoreconst [0] destptr mem))
	for {
		s := v.AuxInt
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(s%16 != 0 && s > 16 && s%16 <= 8 && config.useSSE) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = s - s%16
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = s % 16
		v0.AddArg(destptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQstoreconst, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [16] destptr mem)
	// cond: config.useSSE
	// result: (MOVOstore destptr (MOVOconst [0]) mem)
	for {
		if v.AuxInt != 16 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVOstore)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [32] destptr mem)
	// cond: config.useSSE
	// result: (MOVOstore (OffPtr <destptr.Type> destptr [16]) (MOVOconst [0]) (MOVOstore destptr (MOVOconst [0]) mem))
	for {
		if v.AuxInt != 32 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVOstore)
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = 16
		v0.AddArg(destptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVOstore, types.TypeMem)
		v2.AddArg(destptr)
		v3 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v3.AuxInt = 0
		v2.AddArg(v3)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}
	// match: (Zero [48] destptr mem)
	// cond: config.useSSE
	// result: (MOVOstore (OffPtr <destptr.Type> destptr [32]) (MOVOconst [0]) (MOVOstore (OffPtr <destptr.Type> destptr [16]) (MOVOconst [0]) (MOVOstore destptr (MOVOconst [0]) mem)))
	for {
		if v.AuxInt != 48 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVOstore)
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = 32
		v0.AddArg(destptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVOstore, types.TypeMem)
		v3 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v3.AuxInt = 16
		v3.AddArg(destptr)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v4.AuxInt = 0
		v2.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpAMD64MOVOstore, types.TypeMem)
		v5.AddArg(destptr)
		v6 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v2.AddArg(v5)
		v.AddArg(v2)
		return true
	}
	// match: (Zero [64] destptr mem)
	// cond: config.useSSE
	// result: (MOVOstore (OffPtr <destptr.Type> destptr [48]) (MOVOconst [0]) (MOVOstore (OffPtr <destptr.Type> destptr [32]) (MOVOconst [0]) (MOVOstore (OffPtr <destptr.Type> destptr [16]) (MOVOconst [0]) (MOVOstore destptr (MOVOconst [0]) mem))))
	for {
		if v.AuxInt != 64 {
			break
		}
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(config.useSSE) {
			break
		}
		v.reset(OpAMD64MOVOstore)
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = 48
		v0.AddArg(destptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpAMD64MOVOstore, types.TypeMem)
		v3 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v3.AuxInt = 32
		v3.AddArg(destptr)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v4.AuxInt = 0
		v2.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpAMD64MOVOstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v6.AuxInt = 16
		v6.AddArg(destptr)
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v7.AuxInt = 0
		v5.AddArg(v7)
		v8 := b.NewValue0(v.Pos, OpAMD64MOVOstore, types.TypeMem)
		v8.AddArg(destptr)
		v9 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v9.AuxInt = 0
		v8.AddArg(v9)
		v8.AddArg(mem)
		v5.AddArg(v8)
		v2.AddArg(v5)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueAMD64_OpZero_20(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [s] destptr mem)
	// cond: s > 64 && s <= 1024 && s%16 == 0 && !config.noDuffDevice
	// result: (DUFFZERO [s] destptr (MOVOconst [0]) mem)
	for {
		s := v.AuxInt
		mem := v.Args[1]
		destptr := v.Args[0]
		if !(s > 64 && s <= 1024 && s%16 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpAMD64DUFFZERO)
		v.AuxInt = s
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVOconst, types.TypeInt128)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: (s > 1024 || (config.noDuffDevice && s > 64 || !config.useSSE && s > 32)) && s%8 == 0
	// result: (REPSTOSQ destptr (MOVQconst [s/8]) (MOVQconst [0]) mem)
	for {
		s := v.AuxInt
		mem := v.Args[1]
		destptr := v.Args[0]
		if !((s > 1024 || (config.noDuffDevice && s > 64 || !config.useSSE && s > 32)) && s%8 == 0) {
			break
		}
		v.reset(OpAMD64REPSTOSQ)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQconst, typ.UInt64)
		v0.AuxInt = s / 8
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpAMD64MOVQconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64_OpZeroExt16to32_0(v *Value) bool {
	// match: (ZeroExt16to32 x)
	// result: (MOVWQZX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVWQZX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpZeroExt16to64_0(v *Value) bool {
	// match: (ZeroExt16to64 x)
	// result: (MOVWQZX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVWQZX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpZeroExt32to64_0(v *Value) bool {
	// match: (ZeroExt32to64 x)
	// result: (MOVLQZX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVLQZX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpZeroExt8to16_0(v *Value) bool {
	// match: (ZeroExt8to16 x)
	// result: (MOVBQZX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpZeroExt8to32_0(v *Value) bool {
	// match: (ZeroExt8to32 x)
	// result: (MOVBQZX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64_OpZeroExt8to64_0(v *Value) bool {
	// match: (ZeroExt8to64 x)
	// result: (MOVBQZX x)
	for {
		x := v.Args[0]
		v.reset(OpAMD64MOVBQZX)
		v.AddArg(x)
		return true
	}
}
func rewriteBlockAMD64(b *Block) bool {
	switch b.Kind {
	case BlockAMD64EQ:
		// match: (EQ (TESTL (SHLL (MOVLconst [1]) x) y))
		// result: (UGE (BTL x y))
		for b.Controls[0].Op == OpAMD64TESTL {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				v_0_0 := v_0.Args[_i0]
				if v_0_0.Op != OpAMD64SHLL {
					continue
				}
				x := v_0_0.Args[1]
				v_0_0_0 := v_0_0.Args[0]
				if v_0_0_0.Op != OpAMD64MOVLconst || v_0_0_0.AuxInt != 1 {
					continue
				}
				y := v_0.Args[1^_i0]
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTL, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (TESTQ (SHLQ (MOVQconst [1]) x) y))
		// result: (UGE (BTQ x y))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				v_0_0 := v_0.Args[_i0]
				if v_0_0.Op != OpAMD64SHLQ {
					continue
				}
				x := v_0_0.Args[1]
				v_0_0_0 := v_0_0.Args[0]
				if v_0_0_0.Op != OpAMD64MOVQconst || v_0_0_0.AuxInt != 1 {
					continue
				}
				y := v_0.Args[1^_i0]
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQ, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (TESTLconst [c] x))
		// cond: isUint32PowerOfTwo(c)
		// result: (UGE (BTLconst [log2uint32(c)] x))
		for b.Controls[0].Op == OpAMD64TESTLconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			x := v_0.Args[0]
			if !(isUint32PowerOfTwo(c)) {
				break
			}
			b.Reset(BlockAMD64UGE)
			v0 := b.NewValue0(v_0.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = log2uint32(c)
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (TESTQconst [c] x))
		// cond: isUint64PowerOfTwo(c)
		// result: (UGE (BTQconst [log2(c)] x))
		for b.Controls[0].Op == OpAMD64TESTQconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			x := v_0.Args[0]
			if !(isUint64PowerOfTwo(c)) {
				break
			}
			b.Reset(BlockAMD64UGE)
			v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = log2(c)
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (TESTQ (MOVQconst [c]) x))
		// cond: isUint64PowerOfTwo(c)
		// result: (UGE (BTQconst [log2(c)] x))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				v_0_0 := v_0.Args[_i0]
				if v_0_0.Op != OpAMD64MOVQconst {
					continue
				}
				c := v_0_0.AuxInt
				x := v_0.Args[1^_i0]
				if !(isUint64PowerOfTwo(c)) {
					continue
				}
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = log2(c)
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (TESTQ z1:(SHLQconst [63] (SHRQconst [63] x)) z2))
		// cond: z1==z2
		// result: (UGE (BTQconst [63] x))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHLQconst || z1.AuxInt != 63 {
					continue
				}
				z1_0 := z1.Args[0]
				if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 63 {
					continue
				}
				x := z1_0.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = 63
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (TESTL z1:(SHLLconst [31] (SHRQconst [31] x)) z2))
		// cond: z1==z2
		// result: (UGE (BTQconst [31] x))
		for b.Controls[0].Op == OpAMD64TESTL {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHLLconst || z1.AuxInt != 31 {
					continue
				}
				z1_0 := z1.Args[0]
				if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 31 {
					continue
				}
				x := z1_0.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = 31
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (TESTQ z1:(SHRQconst [63] (SHLQconst [63] x)) z2))
		// cond: z1==z2
		// result: (UGE (BTQconst [0] x))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
					continue
				}
				z1_0 := z1.Args[0]
				if z1_0.Op != OpAMD64SHLQconst || z1_0.AuxInt != 63 {
					continue
				}
				x := z1_0.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = 0
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (TESTL z1:(SHRLconst [31] (SHLLconst [31] x)) z2))
		// cond: z1==z2
		// result: (UGE (BTLconst [0] x))
		for b.Controls[0].Op == OpAMD64TESTL {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
					continue
				}
				z1_0 := z1.Args[0]
				if z1_0.Op != OpAMD64SHLLconst || z1_0.AuxInt != 31 {
					continue
				}
				x := z1_0.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTLconst, types.TypeFlags)
				v0.AuxInt = 0
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (TESTQ z1:(SHRQconst [63] x) z2))
		// cond: z1==z2
		// result: (UGE (BTQconst [63] x))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
					continue
				}
				x := z1.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = 63
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (TESTL z1:(SHRLconst [31] x) z2))
		// cond: z1==z2
		// result: (UGE (BTLconst [31] x))
		for b.Controls[0].Op == OpAMD64TESTL {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
					continue
				}
				x := z1.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64UGE)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTLconst, types.TypeFlags)
				v0.AuxInt = 31
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (InvertFlags cmp) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64EQ)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (EQ (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockAMD64GE:
		// match: (GE (InvertFlags cmp) yes no)
		// result: (LE cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64LE)
			b.AddControl(cmp)
			return true
		}
		// match: (GE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GE (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GE (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GE (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GE (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
	case BlockAMD64GT:
		// match: (GT (InvertFlags cmp) yes no)
		// result: (LT cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64LT)
			b.AddControl(cmp)
			return true
		}
		// match: (GT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GT (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
	case BlockIf:
		// match: (If (SETL cmp) yes no)
		// result: (LT cmp yes no)
		for b.Controls[0].Op == OpAMD64SETL {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64LT)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETLE cmp) yes no)
		// result: (LE cmp yes no)
		for b.Controls[0].Op == OpAMD64SETLE {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64LE)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETG cmp) yes no)
		// result: (GT cmp yes no)
		for b.Controls[0].Op == OpAMD64SETG {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64GT)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETGE cmp) yes no)
		// result: (GE cmp yes no)
		for b.Controls[0].Op == OpAMD64SETGE {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64GE)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETEQ cmp) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpAMD64SETEQ {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64EQ)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETNE cmp) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpAMD64SETNE {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64NE)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETB cmp) yes no)
		// result: (ULT cmp yes no)
		for b.Controls[0].Op == OpAMD64SETB {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64ULT)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETBE cmp) yes no)
		// result: (ULE cmp yes no)
		for b.Controls[0].Op == OpAMD64SETBE {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64ULE)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETA cmp) yes no)
		// result: (UGT cmp yes no)
		for b.Controls[0].Op == OpAMD64SETA {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64UGT)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETAE cmp) yes no)
		// result: (UGE cmp yes no)
		for b.Controls[0].Op == OpAMD64SETAE {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64UGE)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETO cmp) yes no)
		// result: (OS cmp yes no)
		for b.Controls[0].Op == OpAMD64SETO {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64OS)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETGF cmp) yes no)
		// result: (UGT cmp yes no)
		for b.Controls[0].Op == OpAMD64SETGF {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64UGT)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETGEF cmp) yes no)
		// result: (UGE cmp yes no)
		for b.Controls[0].Op == OpAMD64SETGEF {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64UGE)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETEQF cmp) yes no)
		// result: (EQF cmp yes no)
		for b.Controls[0].Op == OpAMD64SETEQF {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64EQF)
			b.AddControl(cmp)
			return true
		}
		// match: (If (SETNEF cmp) yes no)
		// result: (NEF cmp yes no)
		for b.Controls[0].Op == OpAMD64SETNEF {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64NEF)
			b.AddControl(cmp)
			return true
		}
		// match: (If cond yes no)
		// result: (NE (TESTB cond cond) yes no)
		for {
			cond := b.Controls[0]
			b.Reset(BlockAMD64NE)
			v0 := b.NewValue0(cond.Pos, OpAMD64TESTB, types.TypeFlags)
			v0.AddArg(cond)
			v0.AddArg(cond)
			b.AddControl(v0)
			return true
		}
	case BlockAMD64LE:
		// match: (LE (InvertFlags cmp) yes no)
		// result: (GE cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64GE)
			b.AddControl(cmp)
			return true
		}
		// match: (LE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LE (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockAMD64LT:
		// match: (LT (InvertFlags cmp) yes no)
		// result: (GT cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64GT)
			b.AddControl(cmp)
			return true
		}
		// match: (LT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LT (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LT (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LT (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LT (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockAMD64NE:
		// match: (NE (TESTB (SETL cmp) (SETL cmp)) yes no)
		// result: (LT cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETL {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETL || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64LT)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETLE cmp) (SETLE cmp)) yes no)
		// result: (LE cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETLE {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETLE || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64LE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETG cmp) (SETG cmp)) yes no)
		// result: (GT cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETG {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETG || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64GT)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETGE cmp) (SETGE cmp)) yes no)
		// result: (GE cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETGE {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETGE || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64GE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETEQ cmp) (SETEQ cmp)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETEQ {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETEQ || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64EQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETNE cmp) (SETNE cmp)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETNE {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETNE || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64NE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETB cmp) (SETB cmp)) yes no)
		// result: (ULT cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETB {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETB || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64ULT)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETBE cmp) (SETBE cmp)) yes no)
		// result: (ULE cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETBE {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETBE || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64ULE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETA cmp) (SETA cmp)) yes no)
		// result: (UGT cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETA {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETA || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64UGT)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETAE cmp) (SETAE cmp)) yes no)
		// result: (UGE cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETAE {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETAE || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64UGE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETO cmp) (SETO cmp)) yes no)
		// result: (OS cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETO {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETO || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64OS)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTL (SHLL (MOVLconst [1]) x) y))
		// result: (ULT (BTL x y))
		for b.Controls[0].Op == OpAMD64TESTL {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				v_0_0 := v_0.Args[_i0]
				if v_0_0.Op != OpAMD64SHLL {
					continue
				}
				x := v_0_0.Args[1]
				v_0_0_0 := v_0_0.Args[0]
				if v_0_0_0.Op != OpAMD64MOVLconst || v_0_0_0.AuxInt != 1 {
					continue
				}
				y := v_0.Args[1^_i0]
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTL, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTQ (SHLQ (MOVQconst [1]) x) y))
		// result: (ULT (BTQ x y))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				v_0_0 := v_0.Args[_i0]
				if v_0_0.Op != OpAMD64SHLQ {
					continue
				}
				x := v_0_0.Args[1]
				v_0_0_0 := v_0_0.Args[0]
				if v_0_0_0.Op != OpAMD64MOVQconst || v_0_0_0.AuxInt != 1 {
					continue
				}
				y := v_0.Args[1^_i0]
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQ, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTLconst [c] x))
		// cond: isUint32PowerOfTwo(c)
		// result: (ULT (BTLconst [log2uint32(c)] x))
		for b.Controls[0].Op == OpAMD64TESTLconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			x := v_0.Args[0]
			if !(isUint32PowerOfTwo(c)) {
				break
			}
			b.Reset(BlockAMD64ULT)
			v0 := b.NewValue0(v_0.Pos, OpAMD64BTLconst, types.TypeFlags)
			v0.AuxInt = log2uint32(c)
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (TESTQconst [c] x))
		// cond: isUint64PowerOfTwo(c)
		// result: (ULT (BTQconst [log2(c)] x))
		for b.Controls[0].Op == OpAMD64TESTQconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			x := v_0.Args[0]
			if !(isUint64PowerOfTwo(c)) {
				break
			}
			b.Reset(BlockAMD64ULT)
			v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
			v0.AuxInt = log2(c)
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (TESTQ (MOVQconst [c]) x))
		// cond: isUint64PowerOfTwo(c)
		// result: (ULT (BTQconst [log2(c)] x))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				v_0_0 := v_0.Args[_i0]
				if v_0_0.Op != OpAMD64MOVQconst {
					continue
				}
				c := v_0_0.AuxInt
				x := v_0.Args[1^_i0]
				if !(isUint64PowerOfTwo(c)) {
					continue
				}
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = log2(c)
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTQ z1:(SHLQconst [63] (SHRQconst [63] x)) z2))
		// cond: z1==z2
		// result: (ULT (BTQconst [63] x))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHLQconst || z1.AuxInt != 63 {
					continue
				}
				z1_0 := z1.Args[0]
				if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 63 {
					continue
				}
				x := z1_0.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = 63
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTL z1:(SHLLconst [31] (SHRQconst [31] x)) z2))
		// cond: z1==z2
		// result: (ULT (BTQconst [31] x))
		for b.Controls[0].Op == OpAMD64TESTL {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHLLconst || z1.AuxInt != 31 {
					continue
				}
				z1_0 := z1.Args[0]
				if z1_0.Op != OpAMD64SHRQconst || z1_0.AuxInt != 31 {
					continue
				}
				x := z1_0.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = 31
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTQ z1:(SHRQconst [63] (SHLQconst [63] x)) z2))
		// cond: z1==z2
		// result: (ULT (BTQconst [0] x))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
					continue
				}
				z1_0 := z1.Args[0]
				if z1_0.Op != OpAMD64SHLQconst || z1_0.AuxInt != 63 {
					continue
				}
				x := z1_0.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = 0
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTL z1:(SHRLconst [31] (SHLLconst [31] x)) z2))
		// cond: z1==z2
		// result: (ULT (BTLconst [0] x))
		for b.Controls[0].Op == OpAMD64TESTL {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
					continue
				}
				z1_0 := z1.Args[0]
				if z1_0.Op != OpAMD64SHLLconst || z1_0.AuxInt != 31 {
					continue
				}
				x := z1_0.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTLconst, types.TypeFlags)
				v0.AuxInt = 0
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTQ z1:(SHRQconst [63] x) z2))
		// cond: z1==z2
		// result: (ULT (BTQconst [63] x))
		for b.Controls[0].Op == OpAMD64TESTQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHRQconst || z1.AuxInt != 63 {
					continue
				}
				x := z1.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTQconst, types.TypeFlags)
				v0.AuxInt = 63
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTL z1:(SHRLconst [31] x) z2))
		// cond: z1==z2
		// result: (ULT (BTLconst [31] x))
		for b.Controls[0].Op == OpAMD64TESTL {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0++ {
				z1 := v_0.Args[_i0]
				if z1.Op != OpAMD64SHRLconst || z1.AuxInt != 31 {
					continue
				}
				x := z1.Args[0]
				z2 := v_0.Args[1^_i0]
				if !(z1 == z2) {
					continue
				}
				b.Reset(BlockAMD64ULT)
				v0 := b.NewValue0(v_0.Pos, OpAMD64BTLconst, types.TypeFlags)
				v0.AuxInt = 31
				v0.AddArg(x)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (TESTB (SETGF cmp) (SETGF cmp)) yes no)
		// result: (UGT cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETGF {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETGF || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64UGT)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETGEF cmp) (SETGEF cmp)) yes no)
		// result: (UGE cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETGEF {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETGEF || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64UGE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETEQF cmp) (SETEQF cmp)) yes no)
		// result: (EQF cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETEQF {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETEQF || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64EQF)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (TESTB (SETNEF cmp) (SETNEF cmp)) yes no)
		// result: (NEF cmp yes no)
		for b.Controls[0].Op == OpAMD64TESTB {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAMD64SETNEF {
				break
			}
			cmp := v_0_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpAMD64SETNEF || cmp != v_0_1.Args[0] {
				break
			}
			b.Reset(BlockAMD64NEF)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (InvertFlags cmp) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64NE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (NE (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
	case BlockAMD64UGE:
		// match: (UGE (InvertFlags cmp) yes no)
		// result: (ULE cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64ULE)
			b.AddControl(cmp)
			return true
		}
		// match: (UGE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGE (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGE (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGE (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGE (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
	case BlockAMD64UGT:
		// match: (UGT (InvertFlags cmp) yes no)
		// result: (ULT cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64ULT)
			b.AddControl(cmp)
			return true
		}
		// match: (UGT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGT (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
	case BlockAMD64ULE:
		// match: (ULE (InvertFlags cmp) yes no)
		// result: (UGE cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64UGE)
			b.AddControl(cmp)
			return true
		}
		// match: (ULE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULE (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULE (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULE (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULE (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockAMD64ULT:
		// match: (ULT (InvertFlags cmp) yes no)
		// result: (UGT cmp yes no)
		for b.Controls[0].Op == OpAMD64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockAMD64UGT)
			b.AddControl(cmp)
			return true
		}
		// match: (ULT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULT (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULT (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULT (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpAMD64FlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULT (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpAMD64FlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	}
	return false
}
