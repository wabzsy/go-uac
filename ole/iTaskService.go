package ole

import (
	"syscall"
	"unsafe"
)

var (
	CLSID_TaskScheduler = NewGUID("{0F87369F-A4E5-4CFC-BD3E-73E6154572DD}")
	IID_ITaskService    = NewGUID("{2FABA4C7-4DA9-4013-9697-20CC3FD40F85}")
)

/*
ITaskService
[Guid("2faba4c7-4da9-4013-9697-20cc3fd40f85")]

	interface ITaskService : IDispatch {
		HRESULT Proc7( [In] BSTR p0,  [Out] ITaskFolder** p1);
		HRESULT Proc8( [In] int p0,  [Out] IRunningTaskCollection** p1);
		HRESULT Proc9( [In] int p0,  [Out]  IUnknown** p1);
		HRESULT Proc10( [In] VARIANT* p0,  [In] VARIANT* p1,  [In] VARIANT* p2,  [In] VARIANT* p3);
		HRESULT Proc11( [Out] short* p0);
		HRESULT Proc12( [Out] BSTR* p0);
		HRESULT Proc13( [Out] BSTR* p0);
		HRESULT Proc14( [Out] BSTR* p0);
		HRESULT Proc15( [Out] int* p0);
	}
*/
type ITaskService struct {
	IDispatch
}

type ITaskServiceVtbl struct {
	IDispatchVtbl
	GetFolder           uintptr
	GetRunningTasks     uintptr
	NewTask             uintptr
	Connect             uintptr
	get_Connected       uintptr
	get_TargetServer    uintptr
	get_ConnectedUser   uintptr
	get_ConnectedDomain uintptr
	get_HighestVersion  uintptr
}

func (v *ITaskService) Connect(serverName, user, domain, password VARIANT) (hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().Connect,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&serverName)),
		uintptr(unsafe.Pointer(&user)),
		uintptr(unsafe.Pointer(&domain)),
		uintptr(unsafe.Pointer(&password)),
	)

	hr = HResult(r1)
	return
}

func (v *ITaskService) VTable() *ITaskServiceVtbl {
	return (*ITaskServiceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ITaskService) GetFolder(path *uint16) (taskFolder *ITaskFolder, hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().GetFolder,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&taskFolder)),
	)

	hr = HResult(r1)
	return
}

/*
ITaskFolder
[Guid("8cfac062-a080-4c15-9a88-aa7c2af80dfc")]

	interface ITaskFolder : IDispatch {
		HRESULT Proc7( [Out] BSTR* p0); get_Name
		HRESULT Proc8( [Out] BSTR* p0); get_Path
		HRESULT Proc9( [In] BSTR p0,  [Out] ITaskFolder** p1); GetFolder
		HRESULT Proc10( [In] int p0,  [Out] ITaskFolderCollection** p1); GetFolders
		HRESULT Proc11( [In] BSTR p0,  [In] VARIANT* p1,  [Out] ITaskFolder** p2); CreateFolder
		HRESULT Proc12( [In] BSTR p0,  [In] int p1); DeleteFolder
		HRESULT Proc13( [In] BSTR p0,  [Out] IRegisteredTask** p1); GetTask
		HRESULT Proc14( [In] int p0,  [Out] IRegisteredTaskCollection** p1); GetTasks
		HRESULT Proc15( [In] BSTR p0,  [In] int p1); DeleteTask
		HRESULT Proc16( [In] BSTR p0,  [In] BSTR p1,  [In] int p2,  [In] VARIANT* p3,  [In] VARIANT* p4,  [In]  int p5,  [In] VARIANT* p6,  [Out] IRegisteredTask** p7); RegisterTask
		HRESULT Proc17( [In] BSTR p0,  [In]  IUnknown* p1,  [In] int p2,  [In] VARIANT* p3,  [In] VARIANT* p4,  [In]  int p5,  [In] VARIANT* p6,  [Out] IRegisteredTask** p7); RegisterTaskDefinition
		HRESULT Proc18( [In] int p0,  [Out] BSTR* p1); GetSecurityDescriptor
		HRESULT Proc19( [In] BSTR p0,  [In] int p1); SetSecurityDescriptor
	}
*/
type ITaskFolder struct {
	IDispatch
}

type ITaskFolderVtbl struct {
	IDispatchVtbl
	get_Name               uintptr
	get_Path               uintptr
	GetFolder              uintptr
	GetFolders             uintptr
	CreateFolder           uintptr
	DeleteFolder           uintptr
	GetTask                uintptr
	GetTasks               uintptr
	DeleteTask             uintptr
	RegisterTask           uintptr
	RegisterTaskDefinition uintptr
	GetSecurityDescriptor  uintptr
	SetSecurityDescriptor  uintptr
}

func (v *ITaskFolder) VTable() *ITaskFolderVtbl {
	return (*ITaskFolderVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ITaskFolder) GetTask(bstrPath *uint16) (registeredTask *IRegisteredTask, hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().GetTask,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(bstrPath)),
		uintptr(unsafe.Pointer(&registeredTask)),
	)

	hr = HResult(r1)
	return
}

func (v *ITaskFolder) DeleteTask(bstrName *uint16, flags int32) (hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().DeleteTask,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(bstrName)),
		uintptr(flags),
	)

	hr = HResult(r1)
	return
}

func (v *ITaskFolder) RegisterTask(
	bstrPath, bstrXmlText *uint16, flags int, userId, password VARIANT, logonType int, sddl VARIANT,
) (
	registeredTask *IRegisteredTask, hr HResult,
) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().RegisterTask,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(bstrPath)),
		uintptr(unsafe.Pointer(bstrXmlText)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&userId)),
		uintptr(unsafe.Pointer(&password)),
		uintptr(logonType),
		uintptr(unsafe.Pointer(&sddl)),
		uintptr(unsafe.Pointer(&registeredTask)),
	)

	hr = HResult(r1)
	return
}

/*
IRegisteredTask
[Guid("9c86f320-dee3-4dd1-b972-a303f26b061e")]

	interface IRegisteredTask : IDispatch {
	    HRESULT Proc7( [Out] BSTR* p0); get_Name
	    HRESULT Proc8( [Out] BSTR* p0); get_Path
	    HRESULT Proc9( [Out]  int* p0); get_State
		HRESULT Proc10( [Out] short* p0); get_Enabled
		HRESULT Proc11( [In] short p0); set_Enabled
		HRESULT Proc12( [In] VARIANT* p0,  [Out] IRunningTask** p1); Run
		HRESULT Proc13( [In] VARIANT* p0,  [In] int p1,  [In] int p2,  [In] BSTR p3,  [Out] IRunningTask** p4); RunEx
		HRESULT Proc14( [In] int p0,  [Out] IRunningTaskCollection** p1); GetInstances
		HRESULT Proc15( [Out] double* p0); get_LastRunTime
		HRESULT Proc16( [Out] int* p0); get_LastTaskResult
		HRESULT Proc17( [Out] int* p0); get_NumberOfMissedRuns
		HRESULT Proc18( [Out] double* p0); get_NextRunTime
		HRESULT Proc19( [Out]  IUnknown** p0); get_Definition
		HRESULT Proc20( [Out] BSTR* p0); get_Xml
		HRESULT Proc21( [In] int p0,  [Out] BSTR* p1); GetSecurityDescriptor
		HRESULT Proc22( [In] BSTR p0,  [In] int p1); SetSecurityDescriptor
		HRESULT Proc23( [In] int p0); Stop
		HRESULT Proc24( [In] struct Struct_20* p0,  [In] struct Struct_20* p1,  [In, Out] int* p2,  [Out] struct Struct_20** p3); GetRunTimes
	}
*/
type IRegisteredTask struct {
	IDispatch
}

type IRegisteredTaskVtbl struct {
	IDispatchVtbl
	get_Name               uintptr
	get_Path               uintptr
	get_State              uintptr
	get_Enabled            uintptr
	set_Enabled            uintptr
	Run                    uintptr
	RunEx                  uintptr
	GetInstances           uintptr
	get_LastRunTime        uintptr
	get_LastTaskResult     uintptr
	get_NumberOfMissedRuns uintptr
	get_NextRunTime        uintptr
	get_Definition         uintptr
	get_Xml                uintptr
	GetSecurityDescriptor  uintptr
	SetSecurityDescriptor  uintptr
	Stop                   uintptr
	GetRunTimes            uintptr
}

func (v *IRegisteredTask) VTable() *IRegisteredTaskVtbl {
	return (*IRegisteredTaskVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IRegisteredTask) Run(params VARIANT) (runningTask *IRunningTask, hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().Run,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&params)),
		uintptr(unsafe.Pointer(&runningTask)),
	)
	hr = HResult(r1)
	return
}

func (v *IRegisteredTask) Stop(flags int32) (hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().Stop,
		uintptr(unsafe.Pointer(v)),
		uintptr(flags),
	)
	hr = HResult(r1)
	return
}

/*
IRunningTask
[Guid("653758fb-7b9a-4f1e-a471-beeb8e9b834e")]

	interface IRunningTask : IDispatch {
		HRESULT Proc7( [Out] BSTR* p0);
		HRESULT Proc8( [Out] BSTR* p0);
		HRESULT Proc9( [Out] BSTR* p0);
		HRESULT Proc10( [Out] int* p0);
		HRESULT Proc11( [Out] BSTR* p0);
		HRESULT Proc12();
		HRESULT Proc13();
		HRESULT Proc14( [Out] int* p0);
	}
*/
type IRunningTask struct {
	IDispatch
}

type IRunningTaskVtbl struct {
	IDispatchVtbl
	get_Name          uintptr
	get_InstanceGuid  uintptr
	get_Path          uintptr
	get_State         uintptr
	get_CurrentAction uintptr
	Stop              uintptr
	Refresh           uintptr
	get_EnginePID     uintptr
}

func (v *IRunningTask) VTable() *IRunningTaskVtbl {
	return (*IRunningTaskVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IRunningTask) GetState() (state TaskState, hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().get_State,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&state)),
	)
	hr = HResult(r1)
	return
}

func (v *IRunningTask) Stop() (hr HResult) {
	r1, _, _ := syscall.SyscallN(
		v.VTable().Stop,
		uintptr(unsafe.Pointer(v)),
	)
	hr = HResult(r1)
	return
}

type TaskState int

const (
	TASK_STATE_UNKNOWN  TaskState = 0
	TASK_STATE_DISABLED TaskState = 1
	TASK_STATE_QUEUED   TaskState = 2
	TASK_STATE_READY    TaskState = 3
	TASK_STATE_RUNNING  TaskState = 4
)

func (x TaskState) String() string {
	switch x {
	case TASK_STATE_DISABLED:
		return "Disabled"
	case TASK_STATE_QUEUED:
		return "Queued"
	case TASK_STATE_READY:
		return "Ready"
	case TASK_STATE_RUNNING:
		return "Running"
	default:
		return "Unknown"
	}
}
