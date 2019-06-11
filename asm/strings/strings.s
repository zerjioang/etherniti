	.text
	.intel_syntax noprefix
	.file	"strings.c"
	.globl	isDigit                 # -- Begin function isDigit
	.p2align	4, 0x90
	.type	isDigit,@function
isDigit:                                # @isDigit
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	add	dil, -48
	cmp	dil, 10
	setb	al
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end0:
	.size	isDigit, .Lfunc_end0-isDigit
                                        # -- End function
	.globl	isNumericArray          # -- Begin function isNumericArray
	.p2align	4, 0x90
	.type	isNumericArray,@function
isNumericArray:                         # @isNumericArray
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	byte ptr [rdx], 1
	test	esi, esi
	jle	.LBB1_5
# %bb.1:
	movsxd	rsi, esi
	xor	ecx, ecx
	.p2align	4, 0x90
.LBB1_3:                                # =>This Inner Loop Header: Depth=1
	movzx	eax, byte ptr [rdi + rcx]
	add	al, -48
	cmp	al, 10
	jae	.LBB1_4
# %bb.2:                                #   in Loop: Header=BB1_3 Depth=1
	add	rcx, 1
	cmp	rcx, rsi
	jl	.LBB1_3
.LBB1_5:
	mov	rsp, rbp
	pop	rbp
	ret
.LBB1_4:
	mov	byte ptr [rdx], 0
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end1:
	.size	isNumericArray, .Lfunc_end1-isNumericArray
                                        # -- End function
	.globl	lowerCase               # -- Begin function lowerCase
	.p2align	4, 0x90
	.type	lowerCase,@function
lowerCase:                              # @lowerCase
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	test	esi, esi
	jle	.LBB2_6
# %bb.1:
	mov	r9d, esi
	mov	r8d, r9d
	and	r8d, 1
	cmp	esi, 1
	jne	.LBB2_7
# %bb.2:
	xor	eax, eax
	test	r8, r8
	jne	.LBB2_4
	jmp	.LBB2_6
.LBB2_7:
	sub	r9, r8
	xor	eax, eax
	.p2align	4, 0x90
.LBB2_8:                                # =>This Inner Loop Header: Depth=1
	movzx	esi, byte ptr [rdi + rax]
	mov	ecx, esi
	add	cl, -65
	cmp	cl, 25
	ja	.LBB2_10
# %bb.9:                                #   in Loop: Header=BB2_8 Depth=1
	or	sil, 32
	mov	byte ptr [rdx + rax], sil
.LBB2_10:                               #   in Loop: Header=BB2_8 Depth=1
	movzx	esi, byte ptr [rdi + rax + 1]
	mov	ecx, esi
	add	cl, -65
	cmp	cl, 26
	jae	.LBB2_12
# %bb.11:                               #   in Loop: Header=BB2_8 Depth=1
	or	sil, 32
	mov	byte ptr [rdx + rax + 1], sil
.LBB2_12:                               #   in Loop: Header=BB2_8 Depth=1
	add	rax, 2
	cmp	r9, rax
	jne	.LBB2_8
# %bb.3:
	test	r8, r8
	je	.LBB2_6
.LBB2_4:
	mov	sil, byte ptr [rdi + rax]
	mov	ecx, esi
	add	cl, -65
	cmp	cl, 25
	ja	.LBB2_6
# %bb.5:
	or	sil, 32
	mov	byte ptr [rdx + rax], sil
.LBB2_6:
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end2:
	.size	lowerCase, .Lfunc_end2-lowerCase
                                        # -- End function

	.ident	"clang version 6.0.0-1ubuntu2 (tags/RELEASE_600/final)"
	.section	".note.GNU-stack","",@progbits
