	.text
	.intel_syntax noprefix
	.file	"ip.c"
	.globl	ip_to_int               # -- Begin function ip_to_int
	.p2align	4, 0x90
	.type	ip_to_int,@function
ip_to_int:                              # @ip_to_int
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	cl, byte ptr [rdi]
	add	rdi, 1
	mov	edx, ecx
	add	dl, -48
	xor	eax, eax
	xor	esi, esi
	cmp	dl, 9
	ja	.LBB0_2
	.p2align	4, 0x90
.LBB0_1:                                # =>This Inner Loop Header: Depth=1
	movsx	edx, cl
	lea	ecx, [rsi + 4*rsi]
	lea	esi, [rdx + 2*rcx]
	add	esi, -48
	movzx	ecx, byte ptr [rdi]
	add	rdi, 1
	mov	edx, ecx
	add	dl, -48
	cmp	dl, 10
	jb	.LBB0_1
.LBB0_2:
	cmp	esi, 255
	jg	.LBB0_17
# %bb.3:
	cmp	cl, 46
	jne	.LBB0_17
# %bb.4:
	mov	cl, byte ptr [rdi]
	add	rdi, 1
	mov	edx, ecx
	add	dl, -48
	xor	eax, eax
	xor	r10d, r10d
	cmp	dl, 9
	ja	.LBB0_6
	.p2align	4, 0x90
.LBB0_5:                                # =>This Inner Loop Header: Depth=1
	movsx	edx, cl
	lea	ecx, [r10 + 4*r10]
	lea	r10d, [rdx + 2*rcx]
	add	r10d, -48
	movzx	ecx, byte ptr [rdi]
	add	rdi, 1
	mov	edx, ecx
	add	dl, -48
	cmp	dl, 10
	jb	.LBB0_5
.LBB0_6:
	cmp	r10d, 255
	jg	.LBB0_17
# %bb.7:
	cmp	cl, 46
	jne	.LBB0_17
# %bb.8:
	mov	dl, byte ptr [rdi]
	add	rdi, 1
	mov	ecx, edx
	add	cl, -48
	xor	eax, eax
	xor	r9d, r9d
	cmp	cl, 9
	ja	.LBB0_10
	.p2align	4, 0x90
.LBB0_9:                                # =>This Inner Loop Header: Depth=1
	movsx	ecx, dl
	lea	edx, [r9 + 4*r9]
	lea	r9d, [rcx + 2*rdx]
	add	r9d, -48
	movzx	edx, byte ptr [rdi]
	add	rdi, 1
	mov	ecx, edx
	add	cl, -48
	cmp	cl, 10
	jb	.LBB0_9
.LBB0_10:
	cmp	r9d, 255
	jg	.LBB0_17
# %bb.11:
	cmp	dl, 46
	jne	.LBB0_17
# %bb.12:
	mov	al, byte ptr [rdi]
	mov	ecx, eax
	add	cl, -48
	xor	r8d, r8d
	cmp	cl, 9
	ja	.LBB0_16
# %bb.13:
	add	rdi, 1
	xor	r8d, r8d
	.p2align	4, 0x90
.LBB0_14:                               # =>This Inner Loop Header: Depth=1
	movsx	eax, al
	lea	ecx, [r8 + 4*r8]
	lea	r8d, [rax + 2*rcx]
	add	r8d, -48
	movzx	eax, byte ptr [rdi]
	mov	ecx, eax
	add	cl, -48
	add	rdi, 1
	cmp	cl, 10
	jb	.LBB0_14
# %bb.15:
	xor	eax, eax
	cmp	r8d, 255
	jg	.LBB0_17
.LBB0_16:
	shl	esi, 8
	add	r10d, esi
	shl	r10d, 8
	add	r9d, r10d
	shl	r9d, 8
	add	r9d, r8d
	mov	eax, r9d
.LBB0_17:
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end0:
	.size	ip_to_int, .Lfunc_end0-ip_to_int
                                        # -- End function

	.ident	"clang version 7.0.0-3~ubuntu0.18.04.1 (tags/RELEASE_700/final)"
	.section	".note.GNU-stack","",@progbits
	.addrsig
