package windows

import (
	"unsafe"

	"github.com/mackerelio/mackerel-agent/util/windows"
	"github.com/yusufpapurcu/wmi"
)

func GetProcessorArchitectureFromApi() uint16 {
	var systemInfo windows.SYSTEM_INFO
	windows.GetSystemInfo.Call(uintptr(unsafe.Pointer(&systemInfo)))
	return systemInfo.ProcessorArchitecture
}

type Win32_Processor struct {
	Architecture uint16
}

func GetProcessorArchitectureFromWmi() uint16 {
	var dst []Win32_Processor
	// q := wmi.CreateQuery(&dst, "")
	q := `SELECT Architecture FROM Win32_Processor`
	wmi.Query(q, &dst)
	return dst[0].Architecture
}

func GetProcessorArchitectureFromReg() string {
	arch, _, _ := windows.RegGetString(
		windows.HKEY_LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, `PROCESSOR_ARCHITECTURE`)
	return arch
}
