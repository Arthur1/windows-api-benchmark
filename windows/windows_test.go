package windows

import (
	"testing"
)

func BenchmarkGetProcessorArchitectureFromApi(b *testing.B) {
	arch := GetProcessorArchitectureFromApi()
	b.Logf("%d", arch)
}

func BenchmarkGetProcessorArchitectureFromWmi(b *testing.B) {
	arch := GetProcessorArchitectureFromWmi()
	b.Logf("%d", arch)
}

func BenchmarkGetProcessorArchitectureFromReg(b *testing.B) {
	arch := GetProcessorArchitectureFromReg()
	b.Logf("%s", arch)
}
