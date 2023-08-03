package ole

import (
	"syscall"
	"unsafe"
)

type ICMLuaUtilVtbl struct {
	IUnknownVtbl
	Proc3     uintptr
	Proc4     uintptr
	Proc5     uintptr
	Proc6     uintptr
	Proc7     uintptr
	Proc8     uintptr
	ShellExec uintptr
}

type ICMLuaUtil struct {
	IUnknown
}

func (x *ICMLuaUtil) ShellExec(file, parameters, directory []byte, fMask, nShow uint32) (err error) {

	var lpFile, lpParameters, lpDirectory *uint16

	lpFile, err = syscall.UTF16PtrFromString(string(file))
	if err != nil {
		return err
	}

	if parameters != nil {
		lpParameters, err = syscall.UTF16PtrFromString(string(parameters))
		if err != nil {
			return err
		}
	}

	if directory != nil {
		lpDirectory, err = syscall.UTF16PtrFromString(string(directory))
		if err != nil {
			return err
		}
	}

	r1, _, _ := syscall.SyscallN(
		x.VTable().ShellExec,
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(lpFile)),
		uintptr(unsafe.Pointer(lpParameters)),
		uintptr(unsafe.Pointer(lpDirectory)),
		uintptr(fMask),
		uintptr(nShow),
	)

	if r1 != 0 {
		return HResult(r1)
	}

	return nil
}

func (v *ICMLuaUtil) VTable() *ICMLuaUtilVtbl {
	return (*ICMLuaUtilVtbl)(unsafe.Pointer(v.IUnknown.RawVTable))
}
