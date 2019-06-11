	.text
	.intel_syntax noprefix
	.file	"upper.c"
	.globl	toUpper                 # -- Begin function toUpper
	.p2align	4, 0x90
	.type	toUpper,@function
toUpper:                                # @toUpper
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	al, byte ptr [rdi]
	test	al, al
	je	.LBB0_5
# %bb.1:
	lea	rcx, [rdi + 1]
	.p2align	4, 0x90
.LBB0_2:                                # =>This Inner Loop Header: Depth=1
	mov	edx, eax
	add	dl, -97
	cmp	dl, 25
	ja	.LBB0_4
# %bb.3:                                #   in Loop: Header=BB0_2 Depth=1
	add	al, -32
	mov	byte ptr [rcx - 1], al
.LBB0_4:                                #   in Loop: Header=BB0_2 Depth=1
	movzx	eax, byte ptr [rcx]
	add	rcx, 1
	test	al, al
	jne	.LBB0_2
.LBB0_5:
	mov	rax, rdi
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end0:
	.size	toUpper, .Lfunc_end0-toUpper
                                        # -- End function
	.globl	toLower                 # -- Begin function toLower
	.p2align	4, 0x90
	.type	toLower,@function
toLower:                                # @toLower
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	al, byte ptr [rdi]
	test	al, al
	je	.LBB1_5
# %bb.1:
	lea	rcx, [rdi + 1]
	.p2align	4, 0x90
.LBB1_2:                                # =>This Inner Loop Header: Depth=1
	mov	edx, eax
	add	dl, -65
	cmp	dl, 25
	ja	.LBB1_4
# %bb.3:                                #   in Loop: Header=BB1_2 Depth=1
	add	al, 32
	mov	byte ptr [rcx - 1], al
.LBB1_4:                                #   in Loop: Header=BB1_2 Depth=1
	movzx	eax, byte ptr [rcx]
	add	rcx, 1
	test	al, al
	jne	.LBB1_2
.LBB1_5:
	mov	rax, rdi
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end1:
	.size	toLower, .Lfunc_end1-toLower
                                        # -- End function
	.globl	add                     # -- Begin function add
	.p2align	4, 0x90
	.type	add,@function
add:                                    # @add
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
                                        # kill: def %esi killed %esi def %rsi
                                        # kill: def %edi killed %edi def %rdi
	lea	eax, [rdi + rsi]
	mov	rsp, rbp
	pop	rbp
	ret
.Lfunc_end2:
	.size	add, .Lfunc_end2-add
                                        # -- End function

	.ident	"clang version 6.0.0-1ubuntu2 (tags/RELEASE_600/final)"
	.section	".note.GNU-stack","",@progbits
