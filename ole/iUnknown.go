package ole

import (
	"syscall"
	"unsafe"
)

type UnknownLike interface {
	QueryInterface(iid *GUID) (disp *IDispatch, hr HResult)
	AddRef() HResult
	Release() HResult
}

type IUnknown struct {
	RawVTable *interface{}
}

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

func (v *IUnknown) VTable() *IUnknownVtbl {
	return (*IUnknownVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUnknown) QueryInterface(iid *GUID) (disp *IDispatch, hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().QueryInterface,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&disp)),
	)

	hr = HResult(r1)
	return

}

func (v *IUnknown) AddRef() (hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().AddRef,
		uintptr(unsafe.Pointer(v)),
		0,
		0,
	)

	hr = HResult(r1)
	return
}

func (v *IUnknown) Release() (hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().Release,
		uintptr(unsafe.Pointer(v)),
		0,
		0,
	)

	hr = HResult(r1)
	return
}
