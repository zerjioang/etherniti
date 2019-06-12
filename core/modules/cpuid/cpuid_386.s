// Copyright (c) 2015 Klaus Post, released under MIT License. See LICENSE file.

// +build 386,!gccgo

// https://github.com/klauspost/cpuid/blob/master/cpuid_amd64.s
// func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)
TEXT ·cpuid(SB), 7, $0
	MOVL op+0(FP), AX
	MOVL op2+4(FP), CX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

// func xgetbv() (eax, edx uint32)
TEXT ·xgetbv(SB), 7, $0
	MOVL index+0(FP), CX
	BYTE $0x0f; BYTE $0x01; BYTE $0xd0 // XGETBV
	MOVL AX, eax+4(FP)
	MOVL DX, edx+8(FP)
	RET

// func asmCpuid(op uint32) (eax, ebx, ecx, edx uint32)
TEXT ·asmCpuid(SB), 7, $0
	XORL CX, CX
	MOVL op+0(FP), AX
	CPUID
	MOVL AX, eax+4(FP)
	MOVL BX, ebx+8(FP)
	MOVL CX, ecx+12(FP)
	MOVL DX, edx+16(FP)
	RET

// func asmRdtscpAsm() (eax, ebx, ecx, edx uint32)
TEXT ·asmRdtscpAsm(SB), 7, $0
	BYTE $0x0F; BYTE $0x01; BYTE $0xF9 // RDTSCP
	MOVL AX, eax+0(FP)
	MOVL BX, ebx+4(FP)
	MOVL CX, ecx+8(FP)
	MOVL DX, edx+12(FP)
	RET
