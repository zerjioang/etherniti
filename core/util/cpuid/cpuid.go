package cpuid

// assembler functions

// cpuid executes the CPUID instruction to obtain processor identification and
// feature information.
// cpuid executes the CPUID instruction with the given EAX, ECX inputs.
func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)

// xgetbv executes the XGETBV instruction.
func xgetbv() (eax, edx uint32)

type CpuFeatures struct {
	HashBMI2      bool
	HasOSXSAVE    bool
	HasAES        bool
	HasVAES       bool
	HasAVX        bool
	HasAVX512F    bool
	HasAVX512VL   bool
	HasAVX512DQ   bool
	EnabledAVX    bool
	EnabledAVX512 bool
}

var (
	// useAsm flag determines whether the assembly version of EncodeInt will be
	// used. By Default we fall back to encodeInt.
	useAsm bool
	// cpu contains feature flags relevant to selecting a Meow implementation.
	cpuFeatures CpuFeatures
)

// depetermine cpu features
// init determines whether to use assembly version by performing CPU feature
// check.
func init() {
	determineCPUFeatures()

	switch {
	case cpuFeatures.HasVAES && cpuFeatures.HasAVX512F && cpuFeatures.EnabledAVX512:
		break
	case cpuFeatures.HasAES && cpuFeatures.HasAVX && cpuFeatures.EnabledAVX:
		// AVX required for VEX-encoded AES instruction, which allows non-aligned memory addresses.
		break
	}
	cpuFeatures.HashBMI2 = hasBMI2()
}

func GetCpuFeatures() CpuFeatures {
	return cpuFeatures
}

// hasBMI2 returns whether the CPU supports Bit Manipulation Instruction Set 2.
func hasBMI2() bool {
	_, ebx, _, _ := cpuid(7, 0)
	s := ebx&(1<<8) != 0
	return s
}

// determineCPUFeatures populates flags in global cpu variable by querying CPUID.
func determineCPUFeatures() {
	maxID, _, _, _ := cpuid(0, 0)
	if maxID < 1 {
		return
	}

	_, _, ecx1, _ := cpuid(1, 0)
	cpuFeatures.HasOSXSAVE = isSet(ecx1, 27)
	cpuFeatures.HasAES = isSet(ecx1, 25)
	cpuFeatures.HasAVX = isSet(ecx1, 28)

	if cpuFeatures.HasOSXSAVE {
		eax, _ := xgetbv()
		cpuFeatures.EnabledAVX = (eax & 0x6) == 0x6
		cpuFeatures.EnabledAVX512 = (eax & 0xe0) == 0xe0
	}

	if maxID < 7 {
		return
	}
	_, ebx7, ecx7, _ := cpuid(7, 0)
	cpuFeatures.HasVAES = isSet(ecx7, 9)
	cpuFeatures.HasAVX512F = isSet(ebx7, 16)
	cpuFeatures.HasAVX512VL = isSet(ebx7, 31)
	cpuFeatures.HasAVX512DQ = isSet(ebx7, 17)
}

// isSet determines if bit i of x is set.
func isSet(x uint32, i uint) bool {
	return (x>>i)&1 == 1
}
