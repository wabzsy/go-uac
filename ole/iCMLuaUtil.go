package ole

import (
	"syscall"
	"unsafe"
)

var (
	CLSID_CMSTPLUA = NewGUID("{3E5FC7F9-9A51-4367-9063-A120244FBEC7}")
	IID_ICMLuaUtil = NewGUID("{6EDD6D74-C007-4E75-B76A-E5740995E24C}")
)

/*
[Guid("6edd6d74-c007-4e75-b76a-e5740995e24c")]

	interface ICMLuaUtil : IUnknown {
	    HRESULT Proc3( [In] wchar_t* p0,  [In] wchar_t* p1,  [In] wchar_t* p2,  [In] int p3);
	    HRESULT Proc4( [In] wchar_t* p0,  [In] wchar_t* p1,  [In, Out] wchar_t** p2,  [In] int p3);
	    HRESULT Proc5( [In] wchar_t* p0,  [In] wchar_t* p1);
	    HRESULT Proc6( [In] wchar_t* p0,  [In] wchar_t* p1,  [In] wchar_t* p2,  [In] int p3);
	    HRESULT Proc7( [In] wchar_t* p0,  [In] wchar_t* p1,  [In] int p2);
	    HRESULT Proc8( [In] wchar_t* p0);
	    HRESULT Proc9( [In] wchar_t* p0,  [In] wchar_t* p1,  [In] wchar_t* p2,  [In] int p3,  [In] int p4);
	    HRESULT Proc10( [In] int p0,  [In] wchar_t* p1,  [In] wchar_t* p2,  [In] wchar_t* p3);
	    HRESULT Proc11( [In] int p0,  [In] wchar_t* p1,  [In] wchar_t* p2);
	    HRESULT Proc12( [In] int p0,  [In] wchar_t* p1,  [In] int p2);
	    HRESULT Proc13( [In] int p0,  [In] wchar_t* p1);
	    HRESULT Proc14();
	    HRESULT Proc15( [In] wchar_t* p0);
	    HRESULT Proc16( [In] wchar_t* p0,  [In] int p1,  [In] int p2,  [In] int p3,  [In] int p4);
	    HRESULT Proc17( [In] wchar_t* p0);
	    HRESULT Proc18( [In] wchar_t* p0,  [In] wchar_t* p1,  [In] wchar_t* p2,  [In] wchar_t* p3,  [In, Out] int* p4);
	    HRESULT Proc19( [In] wchar_t* p0,  [In] wchar_t* p1,  [Out] wchar_t** p2);
	}
*/
type ICMLuaUtil struct {
	IUnknown
}

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

func (x *ICMLuaUtil) ShellExec(lpFile, lpParameters, lpDirectory *uint16, fMask, nShow uint32) (hr error) {
	r1, _, _ := syscall.SyscallN(
		x.VTable().ShellExec,
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(lpFile)),
		uintptr(unsafe.Pointer(lpParameters)),
		uintptr(unsafe.Pointer(lpDirectory)),
		uintptr(fMask),
		uintptr(nShow),
	)

	hr = HResult(r1)
	return
}

func (v *ICMLuaUtil) VTable() *ICMLuaUtilVtbl {
	return (*ICMLuaUtilVtbl)(unsafe.Pointer(v.IUnknown.RawVTable))
}
