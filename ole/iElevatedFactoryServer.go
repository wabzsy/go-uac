package ole

import (
	"syscall"
	"unsafe"
)

var (
	CLSID_VFServer            = NewGUID("{A6BFEA43-501F-456F-A845-983D3AD7B8F0}")
	IID_ElevatedFactoryServer = NewGUID("{804BD226-AF47-4D71-B492-443A57610B08}")
)

/*
[Guid("804bd226-af47-4d71-b492-443a57610b08")]

	interface IElevatedFactoryServer : IUnknown {
	    HRESULT Proc3( [In] GUID* p0,  [In] GUID* p1,  [Out]  IUnknown** p2);
	}
*/
type IElevatedFactoryServer[T UnknownLike] struct {
	IUnknown
}

type IElevatedFactoryServerVtbl struct {
	IUnknownVtbl
	ServerCreateElevatedObject uintptr
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
