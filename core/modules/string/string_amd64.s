// +build !noasm !appengine
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

TEXT ·_isDigit(SB), $0-16

	MOVQ b+0(FP), DI

	QUAD $0x0aff8040d0c78040
	WORD $0x920f; BYTE $0xd0
	MOVQ AX, result+8(FP)
	RET

TEXT ·_isNumericArray(SB), $0-24

	MOVQ buf+0(FP), DI
	MOVQ len+8(FP), SI
	MOVQ res+16(FP), DX

	LONG $0x850102c6; BYTE $0xf6
	JLE  LBB1_5
	LONG $0x31f66348; BYTE $0xc9

LBB1_3:
	QUAD $0x0a3cd0040f04b60f
	JAE  LBB1_4
	LONG $0x01c18348; WORD $0x3948; BYTE $0xf1
	JL   LBB1_3

LBB1_5:
	RET

LBB1_4:
	QUAD $0xc35dec89480002c6

lowerCase:
	QUAD $0xf8e48348e5894855
	WORD $0xf685
	JLE  LBB2_6
	QUAD $0x8341c88945f18941
	LONG $0xfe8301e0; BYTE $0x01
	JNE  LBB2_7
	LONG $0x854dc031; BYTE $0xc0
	JNE  LBB2_4
	JMP  LBB2_6

LBB2_7:
	LONG $0x31c1294d; BYTE $0xc0

LBB2_8:
	QUAD $0xc180f1890734b60f
	LONG $0x19f980bf
	JA   LBB2_10
	QUAD $0x0234884020ce8040

LBB2_10:
	QUAD $0x80f189010774b60f
	LONG $0xf980bfc1; BYTE $0x1a
	JAE  LBB2_12
	QUAD $0x0274884020ce8040
	BYTE $0x01

LBB2_12:
	LONG $0x02c08348; WORD $0x3949; BYTE $0xc1
	JNE  LBB2_8
	WORD $0x854d; BYTE $0xc0
	JE   LBB2_6

LBB2_4:
	QUAD $0xc180f18907348a40
	LONG $0x19f980bf
	JA   LBB2_6
	QUAD $0x0234884020ce8040

LBB2_6:
	LONG $0x5dec8948; BYTE $0xc3

TEXT ·_lowerCase(SB), $0-24

	MOVQ buf+0(FP), DI
	MOVQ len+8(FP), SI
	MOVQ res+16(FP), DX

	WORD $0xf685
	JLE  LBB2_6
	QUAD $0x8341c88945f18941
	LONG $0xfe8301e0; BYTE $0x01
	JNE  LBB2_7
	LONG $0x854dc031; BYTE $0xc0
	JNE  LBB2_4
	JMP  LBB2_6

LBB2_7:
	LONG $0x31c1294d; BYTE $0xc0

LBB2_8:
	QUAD $0xc180f1890734b60f
	LONG $0x19f980bf
	JA   LBB2_10
	QUAD $0x0234884020ce8040

LBB2_10:
	QUAD $0x80f189010774b60f
	LONG $0xf980bfc1; BYTE $0x1a
	JAE  LBB2_12
	QUAD $0x0274884020ce8040
	BYTE $0x01

LBB2_12:
	LONG $0x02c08348; WORD $0x3949; BYTE $0xc1
	JNE  LBB2_8
	WORD $0x854d; BYTE $0xc0
	JE   LBB2_6

LBB2_4:
	QUAD $0xc180f18907348a40
	LONG $0x19f980bf
	JA   LBB2_6
	QUAD $0x0234884020ce8040

LBB2_6:
	RET
