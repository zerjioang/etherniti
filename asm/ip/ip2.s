	.text
	.intel_syntax noprefix
	.file	"ip2.c"
	.globl	ip_to_int2              # -- Begin function ip_to_int2
	.p2align	4, 0x90
	.type	ip_to_int2,@function
ip_to_int2:                             # @ip_to_int2
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -16
	sub	rsp, 16
	test	esi, esi
	jle	.LBB0_1
# %bb.2:
	mov	r8d, esi
	xor	esi, esi
	xor	edx, edx
	xor	r9d, r9d
	.p2align	4, 0x90
.LBB0_3:                                # =>This Inner Loop Header: Depth=1
	movzx	ecx, byte ptr [rdi + rsi]
	cmp	cl, 46
	jne	.LBB0_5
# %bb.4:                                #   in Loop: Header=BB0_3 Depth=1
	movzx	ecx, r9b
	mov	byte ptr [rsp + 4*rcx + 3], dl
	add	r9b, 1
	xor	edx, edx
	add	rsi, 1
	cmp	r8, rsi
	jne	.LBB0_3
	jmp	.LBB0_7
	.p2align	4, 0x90
.LBB0_5:                                #   in Loop: Header=BB0_3 Depth=1
	add	cl, -48
	movzx	eax, r9b
	movzx	edx, dl
	lea	rax, [rsp + 4*rax]
	mov	byte ptr [rdx + rax], cl
	add	dl, 1
	add	rsi, 1
	cmp	r8, rsi
	jne	.LBB0_3
	jmp	.LBB0_7
.LBB0_1:
	xor	r9d, r9d
	xor	edx, edx
.LBB0_7:
	movzx	eax, r9b
	mov	byte ptr [rsp + 4*rax + 3], dl
	mov	al, byte ptr [rsp + 3]
	cmp	al, 3
	je	.LBB0_13
# %bb.8:
	cmp	al, 2
	je	.LBB0_12
# %bb.9:
	cmp	al, 1
	jne	.LBB0_10
# %bb.11:
	mov	cl, byte ptr [rsp]
	mov	al, byte ptr [rsp + 7]
	cmp	al, 1
	jne	.LBB0_15
	jmp	.LBB0_20
.LBB0_13:
	mov	al, byte ptr [rsp]
	mov	cl, byte ptr [rsp + 1]
	mov	dl, 100
	mul	dl
	mov	edx, eax
	mov	sil, 10
	mov	eax, ecx
	mul	sil
	mov	ecx, eax
	add	cl, dl
	add	cl, byte ptr [rsp + 2]
	mov	al, byte ptr [rsp + 7]
	cmp	al, 1
	je	.LBB0_20
.LBB0_15:
	cmp	al, 2
	je	.LBB0_19
# %bb.16:
	cmp	al, 3
	jne	.LBB0_17
# %bb.18:
	mov	al, byte ptr [rsp + 4]
	mov	dl, byte ptr [rsp + 5]
	mov	sil, 100
	mul	sil
	mov	esi, eax
	mov	dil, 10
	mov	eax, edx
	mul	dil
	mov	edx, eax
	add	dl, sil
	add	dl, byte ptr [rsp + 6]
	mov	al, byte ptr [rsp + 11]
	cmp	al, 1
	jne	.LBB0_22
	jmp	.LBB0_27
.LBB0_12:
	mov	al, byte ptr [rsp]
	mov	cl, 10
	mul	cl
	mov	ecx, eax
	add	cl, byte ptr [rsp + 1]
	mov	al, byte ptr [rsp + 7]
	cmp	al, 1
	jne	.LBB0_15
.LBB0_20:
	mov	dl, byte ptr [rsp + 4]
	mov	al, byte ptr [rsp + 11]
	cmp	al, 1
	je	.LBB0_27
.LBB0_22:
	cmp	al, 2
	je	.LBB0_26
# %bb.23:
	cmp	al, 3
	jne	.LBB0_24
# %bb.25:
	mov	al, byte ptr [rsp + 8]
	mov	sil, byte ptr [rsp + 9]
	mov	dil, 100
	mul	dil
	mov	edi, eax
	mov	r8b, 10
	mov	eax, esi
	mul	r8b
	mov	esi, eax
	add	sil, dil
	add	sil, byte ptr [rsp + 10]
	mov	al, byte ptr [rsp + 15]
	cmp	al, 1
	jne	.LBB0_29
	jmp	.LBB0_34
.LBB0_10:
	xor	ecx, ecx
	mov	al, byte ptr [rsp + 7]
	cmp	al, 1
	jne	.LBB0_15
	jmp	.LBB0_20
.LBB0_19:
	mov	al, byte ptr [rsp + 4]
	mov	dl, 10
	mul	dl
	mov	edx, eax
	add	dl, byte ptr [rsp + 5]
	mov	al, byte ptr [rsp + 11]
	cmp	al, 1
	jne	.LBB0_22
.LBB0_27:
	mov	sil, byte ptr [rsp + 8]
	mov	al, byte ptr [rsp + 15]
	cmp	al, 1
	je	.LBB0_34
.LBB0_29:
	cmp	al, 2
	je	.LBB0_33
# %bb.30:
	cmp	al, 3
	jne	.LBB0_31
# %bb.32:
	mov	al, byte ptr [rsp + 12]
	mov	dil, byte ptr [rsp + 13]
	mov	r8b, 100
	mul	r8b
	mov	r8d, eax
	mov	r9b, 10
	mov	eax, edi
	mul	r9b
                                        # kill: def $al killed $al def $eax
	add	al, r8b
	add	al, byte ptr [rsp + 14]
	jmp	.LBB0_35
.LBB0_17:
	xor	edx, edx
	mov	al, byte ptr [rsp + 11]
	cmp	al, 1
	jne	.LBB0_22
	jmp	.LBB0_27
.LBB0_26:
	mov	al, byte ptr [rsp + 8]
	mov	sil, 10
	mul	sil
	mov	esi, eax
	add	sil, byte ptr [rsp + 9]
	mov	al, byte ptr [rsp + 15]
	cmp	al, 1
	jne	.LBB0_29
.LBB0_34:
	mov	al, byte ptr [rsp + 12]
	jmp	.LBB0_35
.LBB0_24:
	xor	esi, esi
	mov	al, byte ptr [rsp + 15]
	cmp	al, 1
	jne	.LBB0_29
	jmp	.LBB0_34
.LBB0_33:
	mov	al, byte ptr [rsp + 12]
	mov	dil, 10
	mul	dil
                                        # kill: def $al killed $al def $eax
	add	al, byte ptr [rsp + 13]
	jmp	.LBB0_35
.LBB0_31:
	xor	eax, eax
.LBB0_35:
	movzx	eax, al
	movzx	esi, sil
	shl	esi, 8
	or	esi, eax
	movzx	edx, dl
	shl	edx, 16
	or	edx, esi
	movzx	eax, cl
	shl	eax, 24
	or	eax, edx
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end0:
	.size	ip_to_int2, .Lfunc_end0-ip_to_int2
                                        # -- End function

	.ident	"clang version 7.0.0-3~ubuntu0.18.04.1 (tags/RELEASE_700/final)"
	.section	".note.GNU-stack","",@progbits
	.addrsig
