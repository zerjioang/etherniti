	.file	"strings.c"
	.intel_syntax noprefix
	.text
	.p2align 4,,15
	.globl	isDigit
	.type	isDigit, @function
isDigit:
	sub	edi, 48
	cmp	dil, 9
	setbe	al
	ret
	.size	isDigit, .-isDigit
	.p2align 4,,15
	.globl	isNumericArray
	.type	isNumericArray, @function
isNumericArray:
	mov	BYTE PTR [rdx], 1
	test	esi, esi
	jle	.L10
	movzx	eax, BYTE PTR [rdi]
	sub	eax, 48
	cmp	al, 9
	ja	.L6
	add	rdi, 1
	lea	ecx, -1[rsi]
	add	rcx, rdi
	jmp	.L8
	.p2align 4,,10
	.p2align 3
.L7:
	movzx	eax, BYTE PTR [rdi]
	add	rdi, 1
	sub	eax, 48
	cmp	al, 9
	ja	.L6
.L8:
	cmp	rdi, rcx
	jne	.L7
	ret
	.p2align 4,,10
	.p2align 3
.L6:
	mov	BYTE PTR [rdx], 0
	ret
	.p2align 4,,10
	.p2align 3
.L10:
	ret
	.size	isNumericArray, .-isNumericArray
	.p2align 4,,15
	.globl	lowerCase
	.type	lowerCase, @function
lowerCase:
	test	esi, esi
	jle	.L16
	lea	r8d, -1[rsi]
	xor	eax, eax
	add	r8, 1
	.p2align 4,,10
	.p2align 3
.L14:
	movzx	ecx, BYTE PTR [rdi+rax]
	lea	esi, -65[rcx]
	cmp	sil, 25
	ja	.L13
	or	ecx, 32
	mov	BYTE PTR [rdx+rax], cl
.L13:
	add	rax, 1
	cmp	r8, rax
	jne	.L14
.L16:
	ret
	.size	lowerCase, .-lowerCase
	.ident	"GCC: (Ubuntu 7.4.0-1ubuntu1~18.04) 7.4.0"
	.section	.note.GNU-stack,"",@progbits
