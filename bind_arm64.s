#include "textflag.h"

TEXT ·sqlite3_open(SB),NOSPLIT,$0
	MOVD name+0(FP), R0
	MOVD handle+8(FP), R1

	CALL sqlite3_open(SB)

	MOVD R0, retval+16(FP)

	RET

TEXT ·sqlite3_close(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0

	CALL sqlite3_close(SB)

	RET

TEXT ·sqlite3_prepare_v2(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0
	MOVD sql+8(FP), R1
	MOVD bytes+16(FP), R2
	MOVD statement+24(FP), R3
	MOVD tail+32(FP), R4

	CALL sqlite3_prepare_v2(SB)

	MOVD R0, retval+40(FP)

	RET

TEXT ·sqlite3_step(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0

	CALL sqlite3_step(SB)

	MOVD R0, retval+8(FP)

	RET


TEXT ·sqlite3_errmsg(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0

	CALL sqlite3_errmsg(SB)

	MOVD R0, retval+8(FP)

	RET

TEXT ·sqlite3_finalize(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0

	CALL sqlite3_finalize(SB)

	MOVD R0, handle+8(FP)

	RET

TEXT ·sqlite3_bind_int(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0
	MOVD column+8(FP), R1
	MOVD value+16(FP), R2

	CALL sqlite3_bind_int(SB)

	MOVD R0, retval+24(FP)

	RET

TEXT ·sqlite3_bind_text(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0
	MOVD column+8(FP), R1
	MOVD value+16(FP), R2
	MOVD length+24(FP), R3
	MOVD length+32(FP), R4

	CALL sqlite3_bind_text(SB)

	MOVD R0, retval+40(FP)

	RET

TEXT ·sqlite3_bind_blob(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0
	MOVD column+8(FP), R1
	MOVD value+16(FP), R2
	MOVD length+24(FP), R3
	MOVD length+32(FP), R4

	CALL sqlite3_bind_blob(SB)

	MOVD R0, retval+40(FP)

	RET

TEXT ·sqlite3_reset(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0

	CALL sqlite3_reset(SB)

	MOVD R0, retval+8(FP)

	RET

TEXT ·sqlite3_column_int(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0
	MOVD column+8(FP), R1

	CALL sqlite3_column_int(SB)

	MOVD R0, retval+16(FP)

	RET

TEXT ·sqlite3_column_text(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0
	MOVD column+8(FP), R1

	CALL sqlite3_column_text(SB)

	MOVD R0, retval+16(FP)

	RET

TEXT ·sqlite3_column_blob(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0
	MOVD column+8(FP), R1

	CALL sqlite3_column_blob(SB)

	MOVD R0, retval+16(FP)

	RET

TEXT ·sqlite3_column_bytes(SB),NOSPLIT,$0
	MOVD handle+0(FP), R0
	MOVD column+8(FP), R1

	CALL sqlite3_column_bytes(SB)

	MOVD R0, retval+16(FP)

	RET
