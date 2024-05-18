#include "textflag.h"

TEXT ·wasmFloor(SB),NOSPLIT,$0
	Get SP
	F32Load x+0(FP)
	F32Floor
	F32Store ret+8(FP)
	RET

TEXT ·wasmCeil(SB),NOSPLIT,$0
	Get SP
	F32Load x+0(FP)
	F32Ceil
	F32Store ret+8(FP)
	RET

TEXT ·wasmTrunc(SB),NOSPLIT,$0
	Get SP
	F32Load x+0(FP)
	F32Trunc
	F32Store ret+8(FP)
	RET
