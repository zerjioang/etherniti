	.file	"hex2.c"
	.intel_syntax noprefix
	.text
	.p2align 4,,15
	.globl	_Z7bin2hexPhi
	.type	_Z7bin2hexPhi, @function
_Z7bin2hexPhi:
	push	rbp
	push	rbx
	mov	rbx, rdi
	lea	edi, 1[rsi+rsi]
	mov	ebp, esi
	sub	rsp, 8
	movsx	rdi, edi
	call	malloc@PLT
	test	rbx, rbx
	mov	rcx, rax
	je	.L2
.L11:
	test	ebp, ebp
	je	.L2
.L3:
	movzx	edx, BYTE PTR [rbx]
	mov	r8d, edx
	shr	r8b, 4
	lea	esi, 48[r8]
	lea	edi, 55[r8]
	cmp	r8b, 9
	cmovbe	edi, esi
	and	edx, 15
	cmp	dl, 9
	mov	BYTE PTR [rcx], dil
	jg	.L6
	add	edx, 48
	add	rcx, 2
	add	rbx, 1
	mov	BYTE PTR -1[rcx], dl
	sub	ebp, 1
	je	.L2
	test	rbx, rbx
	jne	.L3
.L2:
	mov	BYTE PTR [rcx], 0
	add	rsp, 8
	pop	rbx
	pop	rbp
	ret
	.p2align 4,,10
	.p2align 3
.L6:
	add	edx, 55
	add	rcx, 2
	sub	ebp, 1
	mov	BYTE PTR -1[rcx], dl
	add	rbx, 1
	jne	.L11
	mov	BYTE PTR [rcx], 0
	add	rsp, 8
	pop	rbx
	pop	rbp
	ret
	.size	_Z7bin2hexPhi, .-_Z7bin2hexPhi
	.p2align 4,,15
	.globl	_Z7hex2binPKc
	.type	_Z7hex2binPKc, @function
_Z7hex2binPKc:
	push	r13
	push	r12
	push	rbp
	push	rbx
	mov	rbp, rdi
	mov	edi, 1
	sub	rsp, 8
	call	malloc@PLT
	test	rbp, rbp
	mov	r13, rax
	mov	BYTE PTR [rax], 0
	je	.L33
	movzx	ebx, BYTE PTR 0[rbp]
	test	bl, bl
	je	.L33
	lea	rax, 1[rbp]
	xor	r12d, r12d
	.p2align 4,,10
	.p2align 3
.L16:
	add	rax, 1
	add	r12d, 1
	cmp	BYTE PTR -1[rax], 0
	jne	.L16
	mov	edi, r12d
	sar	edi
	add	edi, 1
	movsx	rdi, edi
	call	malloc@PLT
	xor	ecx, ecx
	and	r12d, 1
	mov	BYTE PTR [rax], 0
	sete	cl
	mov	rsi, rax
	sal	ecx, 2
	jmp	.L18
	.p2align 4,,10
	.p2align 3
.L35:
	sub	ebx, 48
	sal	ebx, cl
	add	BYTE PTR [rsi], bl
.L20:
	add	rbp, 1
	test	ecx, ecx
	movzx	ebx, BYTE PTR 0[rbp]
	jne	.L25
.L36:
	mov	BYTE PTR 1[rsi], 0
	add	rsi, 1
	test	bl, bl
	mov	ecx, 4
	je	.L24
.L18:
	lea	edx, -48[rbx]
	cmp	dl, 9
	jbe	.L35
	lea	edx, -65[rbx]
	cmp	dl, 5
	ja	.L21
	sub	ebx, 55
	add	rbp, 1
	sal	ebx, cl
	add	BYTE PTR [rsi], bl
	test	ecx, ecx
	movzx	ebx, BYTE PTR 0[rbp]
	je	.L36
.L25:
	xor	ecx, ecx
	test	bl, bl
	jne	.L18
.L24:
	mov	r13, rax
.L33:
	add	rsp, 8
	mov	rax, r13
	pop	rbx
	pop	rbp
	pop	r12
	pop	r13
	ret
	.p2align 4,,10
	.p2align 3
.L21:
	lea	edx, -97[rbx]
	cmp	dl, 5
	ja	.L33
	sub	ebx, 87
	sal	ebx, cl
	add	BYTE PTR [rsi], bl
	jmp	.L20
	.size	_Z7hex2binPKc, .-_Z7hex2binPKc
	.ident	"GCC: (Ubuntu 7.4.0-1ubuntu1~18.04) 7.4.0"
	.section	.note.GNU-stack,"",@progbits
