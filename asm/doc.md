# CPU Flags

## Example Flags

```bash
fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault epb invpcid_single pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid rdseed adx smap intel_pt xsaveopt dtherm ida arat pln pts md_clear flush_l1d
```

## Flags description

* `fpu`: Floating point unit
* `lm`: 64-bit (x86_64/AMD64/Intel64)
* `vme`: a flag in a modern x86 CPU indicating support of Virtual 8086 mode.
* `de`: no description.
* `pse`: no description.
* `tsc`: no description.
* `msr`: no description.
* `pae`: no description.
* `mce`: no description.
* `cx8`: no description.
* `apic`: no description.
* `sep`: no description.
* `mtrr`: no description.
* `pge`: no description.
* `mca`: no description.
* `cmov`: no description.
* `pat`: no description.
* `pse36`: no description.
* `clflush`: no description.
* `dts`: no description.
* `acpi`: no description.
* `mmx`: no description.
* `fxsr`: no description.
* `sse`: no description.
* `sse2`: no description.
* `ss`: no description.
* `ht`: no description.
* `tm`: no description.
* `pbe`: no description.
* `syscall`: no description.
* `nx`: no description.
* `pdpe1gb`: no description.
* `rdtscp`: no description.
* `lm`: no description.
* `constant_tsc`: no description.
* `arch_perfmon`: no description.
* `pebs`: no description.
* `bts`: no description.
* `rep_good`: no description.
* `nopl`: no description.
* `xtopology`: no description.
* `nonstop_tsc`: no description.
* `cpuid`: no description.
* `aperfmperf`: no description.
* `pni`: no description.
* `pclmulqdq`: no description.
* `dtes64`: no description.
* `monitor`: no description.
* `ds_cpl`: no description.
* `vmx`: no description.
* `est`: no description.
* `tm2`: no description.
* `ssse3`: no description.
* `sdbg`: no description.
* `fma`: no description.
* `cx16`: no description.
* `xtpr`: no description.
* `pdcm`: no description.
* `pcid`: no description.
* `sse4_1`: no description.
* `sse4_2`: no description.
* `x2apic`: no description.
* `movbe`: no description.
* `popcnt`: no description.
* `tsc_deadline_timer`: no description.
* `aes`: no description.
* `xsave`: no description.
* `avx`: no description.
* `f16c`: no description.
* `rdrand`: no description.
* `lahf_lm`: no description.
* `abm`: no description.
* `3dnowprefetch`: no description.
* `cpuid_fault`: no description.
* `epb`: no description.
* `invpcid_single`: no description.
* `pti`: no description.
* `ssbd`: no description.
* `ibrs`: no description.
* `ibpb`: no description.
* `stibp`: no description.
* `tpr_shadow`: no description.
* `vnmi`: no description.
* `flexpriority`: no description.
* `ept`: no description.
* `vpid`: no description.
* `fsgsbase`: no description.
* `tsc_adjust`: no description.
* `bmi1`: no description.
* `avx2`: no description.
* `smep`: no description.
* `bmi2`: no description.
* `erms`: no description.
* `invpcid`: no description.
* `rdseed`: no description.
* `adx`: no description.
* `smap`: no description.
* `intel_pt`: no description.
* `xsaveopt`: no description.
* `dtherm`: no description.
* `ida`: no description.
* `arat`: no description.
* `pln`: no description.
* `pts`: no description.
* `md_clear`: no description.
* `flush_l1d`: no description.

## SIMD

### Introduction
SIMD (Single Instruction, Multiple Data). SIMD describes any extension to microprocessors that allow it to operate on data in parallel. Some common SIMD extensions are MMX, 3DNow!, SSE, and AltiVec (related to VMX). There are many others, but these are the most common ones found in ordinary PCs.

Most SIMD instruction sets have gone through a few revisions since their initial implementation. This gives us extended sets of each variety, including MMX, extended MMX, 3DNow!, 3DNow!2 (sometimes called 3DNow! Professional or 3DNow!+), SSE (also known as Katmai New Instructions or simply KNI), SSE2 (also known as Willamette New Instructions or simply WNI), SSE3 (also known as Prescott New Instructions or simply PNI), and SSE4 (also known as Tejas New Instructions or simply TNI). Later AVX grew as another extension for more parallelism.

In 2008, Intel added AES NI, CLMUL (which is a subset of AES NI), and AVX, a radical departure from SSE.

In 1997, Cyrix extended the MMX Instruction set and called it EMMX.

With their XScale line of mobile processors, Intel developed a variation of MMX Called WMMX, and later WMMX2.

ARM added NEON as their processors became more media-centric with consumers.

## SSE
### An Overview
SSE is a newer SIMD extension to the Intel Pentium III and AMD AthlonXP microprocessors. Unlike MMX and 3DNow! extensions, which occupy the same register space as the normal FPU registers, SSE adds a separate register space to the microprocessor. Because of this, SSE can only be used on operating systems that support it. Fortunately, most recent operating systems have support built in. All versions of Windows since Windows98 support SSE, as do Linux kernels since 2.2.

SSE was introduced in 1999, and was also known as "Katmai New Instructions" (or KNI) after the Pentium III's core codename.

SSE adds 8 new 128-bit registers, divided into 4 32-bit (single precision) floating point values. These registers are called XMM0 - XMM7. An additional control register, MXCSR, is also available to control and check the status of SSE instructions.

SSE gives us access to 70 new instructions that operate on these 128bit registers, MMX registers, and sometimes even regular 32bit registers.

## CPUID

`cpuid` is an instruction added to Intel Pentiums (and some later 80486's) that enables programmers to determine what kind of features the current CPU supports, who made it, various extensions and abilities, and cache information.

This article will show you how to get information using `cpuid`, and how to interpret that information to detect support for MMX and its extensions, 3DNow! and its extensions, SSE and its extensions, and some other useful features.

### Using CPUID
There are many ways to use cpuid, depending on where you are working. There is inline assembly for GCC and MSVC (both different), as well as using it in plain assembly files (for NASM et.al.)

cpuid uses the value in eax, and returns data into eax, ebx, ecx, and edx. The eax input is known as the "Function" input. It can have values from 0x00000000 to 0x00000001, and 0x80000000 to 0x80000008.

To use cpuid in GCC, it is normally easiest to just define a macro, like this:

```c
#define cpuid(func,ax,bx,cx,dx)\
	__asm__ __volatile__ ("cpuid":\
	"=a" (ax), "=b" (bx), "=c" (cx), "=d" (dx) : "a" (func));
```


In the above, you simply put whatever function number you want in for func, and put in 4 variables that will get the output values of eax, ebx, ecx, and edx, respectively.

```c
int a,b,c,d;
...
cpuid(0,a,b,c,d);
...
```

In NASM, you simply use the instruction cpuid, and handle the outputs as desired. In MSVC, it's a bit longer, but still not too crazy. (Please keep in mind that I don't use MSVC much at all anymore, so this isn't tested).

```c
#define cpuid(func,a,b,c,d)\
	asm {\
	mov	eax, func\
	cpuid\
	mov	a, eax\
	mov	b, ebx\
	mov	c, ecx\
	mov	d, edx\
	}
```

And then you can call it with the same above snippet.

More information at: http://softpixel.com/~cwright/programming/simd/cpuid.php
https://unix.stackexchange.com/questions/43539/what-do-the-flags-in-proc-cpuinfo-mean
https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux.git/tree/arch/x86/include/asm/cpufeatures.h

## ARM

http://www.keil.com/support/man/docs/armclang_ref/armclang_ref_chr1392632801932.htm

http://www.novalis.org/documents/mmx.html#note1
https://software.intel.com/sites/landingpage/IntrinsicsGuide/#expand=91,555,124,127,826,2812,2813&cats=OS-Targeted

## Internals

The basic process is to (in the prologue) setup the stack and registers as how the C code expects this to be the case, and upon exiting the subroutine (in the epilogue) to revert back to the golang world and pass a return value back if required. In more details:

* Define assembly subroutine with proper golang decoration in terms of needed stack space and overall size of arguments plus return value.
* Function arguments are loaded from the golang stack into registers and prior to starting the C code any arguments beyond 6 are stored in C stack space.
* Stack space is reserved and setup for the C code. Depending on the C code, the stack pointer maybe aligned on a certain boundary (especially needed for code that takes advantages of SIMD instructions such as AVX etc.).
* A constants table is generated (if needed) and anyrip-based references are replaced with proper offsets to where Go will put the table.

## Limitations
* Arguments need (for now) to be 64-bit size, meaning either a value or a pointer (this requirement will be lifted)
* Maximum number of 14 arguments (hard limit — if you hit this maybe you should rethink your api anyway…)
* Generally nocallstatements (thus inline your C code) with a couple of exceptions for functions such asmemsetandmemcpy(seeclib_amd64.s)
