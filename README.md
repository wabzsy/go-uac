用来练手的项目,练习Golang和OLE/COM交互.

利用的是UACME的41和74,这两个在实战用的比较多,不依赖dll劫持对go比较友好

目前配合donut/gonut可过360,火绒,Microsoft Defender

参考:

- https://github.com/go-ole/go-ole
- https://github.com/hfiref0x/UACME
- https://learn.microsoft.com/zh-cn/windows/win32/learnwin32/creating-an-object-in-com
- https://learn.microsoft.com/en-us/windows/win32/api/objbase/nf-objbase-coinitialize
- https://github.com/zcgonvh/TaskSchedulerMisc
- https://github.com/0xlane/BypassUAC

```
41:
    Author: Oddvar Moe
    Type: Elevated COM interface
    Method: ICMLuaUtil
    Target(s): Attacker defined
    Component(s): Attacker defined
    Implementation: ucmCMLuaUtilShellExecMethod
    Works from: Windows 7 (7600)
    Fixed in: unfixed 🙈
    How: -
    Code status: added in v2.7.9
74:
    Author: zcgonvh
    Type: Elevated COM interface
    Method: IElevatedFactoryServer
    Target(s): Attacker defined
    Component(s): Attacker defined
    Implementation: ucmVFServerTaskSchedMethod
    Works from: Windows 8.1 (9600)
    Fixed in: unfixed 🙈
    How: -
    Code status: added in v3.6.1
```

TODO: 现在都是白加黑或者注入所以暂时没写进程伪装的部分,有空把进程伪装也写了