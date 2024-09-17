#include "textflag.h"

TEXT ·sqlite3_open(SB),NOSPLIT,$0
	MOVQ name+0(FP), DI
	MOVQ handle+8(FP), SI
	
	CALL sqlite3_open(SB)
	
	MOVQ AX, retval+16(FP)

	RET

TEXT ·sqlite3_close(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	
	CALL sqlite3_close(SB)

	RET

TEXT ·sqlite3_prepare_v2(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	MOVQ sql+8(FP), SI
	MOVQ bytes+16(FP), DX
	MOVQ statement+24(FP), CX
	MOVQ tail+32(FP), R8
	
	CALL sqlite3_prepare_v2(SB)
	
	MOVQ AX, retval+40(FP)

	RET

TEXT ·sqlite3_step(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	
	CALL sqlite3_step(SB)
	
	MOVQ AX, retval+8(FP)

	RET

TEXT ·sqlite3_errmsg(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	
	CALL sqlite3_errmsg(SB)
	
	MOVQ AX, retval+8(FP)

	RET

TEXT ·sqlite3_finalize(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	
	CALL sqlite3_finalize(SB)
	
	MOVQ AX, handle+8(FP)

	RET

TEXT ·sqlite3_bind_int(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	MOVQ column+8(FP), SI
	MOVQ value+16(FP), DX

	CALL sqlite3_bind_int(SB)

	MOVQ AX, retval+24(FP)

	RET

TEXT ·sqlite3_bind_text(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	MOVQ column+8(FP), SI
	MOVQ value+16(FP), DX
	MOVQ length+24(FP), CX
	MOVQ length+32(FP), R8

	CALL sqlite3_bind_text(SB)

	MOVQ AX, retval+40(FP)

	RET

TEXT ·sqlite3_bind_blob(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	MOVQ column+8(FP), SI
	MOVQ value+16(FP), DX
	MOVQ length+24(FP), CX
	MOVQ length+32(FP), R8

	CALL sqlite3_bind_blob(SB)

	MOVQ AX, retval+40(FP)

	RET

TEXT ·sqlite3_reset(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI

	CALL sqlite3_reset(SB)

	MOVQ AX, retval+8(FP)

	RET

TEXT ·sqlite3_column_int(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	MOVQ column+8(FP), SI

	CALL sqlite3_column_int(SB)

	MOVQ AX, retval+16(FP)

	RET

TEXT ·sqlite3_column_text(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	MOVQ column+8(FP), SI

	CALL sqlite3_column_text(SB)

	MOVQ AX, retval+16(FP)

	RET

TEXT ·sqlite3_column_blob(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	MOVQ column+8(FP), SI

	CALL sqlite3_column_blob(SB)

	MOVQ AX, retval+16(FP)

	RET

TEXT ·sqlite3_column_bytes(SB),NOSPLIT,$0
	MOVQ handle+0(FP), DI
	MOVQ column+8(FP), SI

	CALL sqlite3_column_bytes(SB)

	MOVQ AX, retval+16(FP)

	RET
