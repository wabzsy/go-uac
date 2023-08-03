package ole

import (
	"syscall"
	"unsafe"
)

type IElevatedFactoryServerVtbl struct {
	IUnknownVtbl
	ServerCreateElevatedObject uintptr
}

type IElevatedFactoryServer[T UnknownLike] struct {
	IUnknown
}

func (x *IElevatedFactoryServer[T]) ServerCreateElevatedObject(clsid, iid *GUID) (unk T, hr HResult) {
	if iid == nil {
		iid = IID_IUnknown
	}

	r1, _, _ := syscall.SyscallN(
		x.VTable().ServerCreateElevatedObject,
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(clsid)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&unk)),
	)

	hr = HResult(r1)
	return
}

func (v *IElevatedFactoryServer[T]) VTable() *IElevatedFactoryServerVtbl {
	return (*IElevatedFactoryServerVtbl)(unsafe.Pointer(v.RawVTable))
}
