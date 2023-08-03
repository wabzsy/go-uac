package ole

import (
	"syscall"
	"unsafe"
)

var (
	modole32    = syscall.NewLazyDLL("ole32.dll")
	modoleaut32 = syscall.NewLazyDLL("oleaut32.dll")

	procCoInitialize     = modole32.NewProc("CoInitialize")
	procCoInitializeEx   = modole32.NewProc("CoInitializeEx")
	procCoUninitialize   = modole32.NewProc("CoUninitialize")
	procCoCreateInstance = modole32.NewProc("CoCreateInstance")
	procCoGetObject      = modole32.NewProc("CoGetObject")
	procVariantInit      = modoleaut32.NewProc("VariantInit")
	procVariantClear     = modoleaut32.NewProc("VariantClear")
	procSysAllocString   = modoleaut32.NewProc("SysAllocString")
	procSysFreeString    = modoleaut32.NewProc("SysFreeString")
)

// CoGetObject retrieves pointer to active object.
func CoGetObject[T UnknownLike](programID string, bindOpts *BIND_OPTS3, iid *GUID) (unk T, hr HResult) {
	if bindOpts != nil {
		bindOpts.CbStruct = uint32(unsafe.Sizeof(BIND_OPTS3{}))
	}

	if iid == nil {
		iid = IID_IUnknown
	}

	r1, _, _ := procCoGetObject.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(programID))),
		uintptr(unsafe.Pointer(bindOpts)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&unk)),
	)

	hr = HResult(r1)
	return
}

// VariantInit initializes variant.
func VariantInit(v *VARIANT) (hr HResult) {
	r1, _, _ := procVariantInit.Call(uintptr(unsafe.Pointer(v)))
	hr = HResult(r1)
	return
}

// VariantClear clears value in Variant settings to VT_EMPTY.
func VariantClear(v *VARIANT) (hr HResult) {
	r1, _, _ := procVariantClear.Call(uintptr(unsafe.Pointer(v)))
	hr = HResult(r1)
	return
}

// SysAllocString allocates memory for string and copies string into memory.
func SysAllocString(v string) (ss *uint16) {
	pss, _, _ := procSysAllocString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v))))
	ss = (*uint16)(unsafe.Pointer(pss))
	return
}

// SysFreeString frees string system memory. This must be called with SysAllocString.
func SysFreeString(v *uint16) (hr HResult) {
	r1, _, _ := procSysFreeString.Call(uintptr(unsafe.Pointer(v)))
	hr = HResult(r1)
	return
}

// coInitialize initializes COM library on current thread.
//
// MSDN documentation suggests that this function should not be called. Call
// CoInitializeEx() instead. The reason has to do with threading and this
// function is only for single-threaded apartments.
//
// That said, most users of the library have gotten away with just this
// function. If you are experiencing threading issues, then use
// CoInitializeEx().
func coInitialize() (hr HResult) {
	// http://msdn.microsoft.com/en-us/library/windows/desktop/ms678543(v=vs.85).aspx
	// Suggests that no value should be passed to CoInitialized.
	// Could just be Call() since the parameter is optional. <-- Needs testing to be sure.
	r1, _, _ := procCoInitialize.Call(uintptr(0))

	hr = HResult(r1)
	return
}

// coInitializeEx initializes COM library with concurrency model.
func coInitializeEx(coinit uint32) (hr HResult) {
	// http://msdn.microsoft.com/en-us/library/windows/desktop/ms695279(v=vs.85).aspx
	// Suggests that the first parameter is not only optional but should always be NULL.
	r1, _, _ := procCoInitializeEx.Call(uintptr(0), uintptr(coinit))
	hr = HResult(r1)
	return
}

// CoInitialize initializes COM library on current thread.
//
// MSDN documentation suggests that this function should not be called. Call
// CoInitializeEx() instead. The reason has to do with threading and this
// function is only for single-threaded apartments.
//
// That said, most users of the library have gotten away with just this
// function. If you are experiencing threading issues, then use
// CoInitializeEx().
func CoInitialize(p uintptr) (hr HResult) {
	// p is ignored and won't be used.
	// Avoid any variable not used errors.
	p = uintptr(0)
	return coInitialize()
}

// CoInitializeEx initializes COM library with concurrency model.
func CoInitializeEx(p uintptr, coinit uint32) (hr HResult) {
	// Avoid any variable not used errors.
	p = uintptr(0)
	return coInitializeEx(coinit)
}

// CoUninitialize uninitializes COM Library.
func CoUninitialize() (hr HResult) {
	r1, _, _ := procCoUninitialize.Call()
	return HResult(r1)
}
