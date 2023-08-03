package ole

import "syscall"

type HResult int32

func (e HResult) IsSuccesed() bool {
	return e >= 0
}

func (e HResult) IsFailed() bool {
	return e < 0
}

func (e HResult) Error() string {
	return e.String()
}

func (e HResult) String() string {
	return syscall.Errno(uintptr(e)).Error()
}

func (e HResult) IsErr(errno syscall.Errno) bool {
	return syscall.Errno(e&0xFFFF) == errno
}

func HResultFrom(x uintptr) HResult {
	return HResult(x)
}
