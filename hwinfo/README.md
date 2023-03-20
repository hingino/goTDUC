# hwinfo

A package that returns hardware information.


### GetChassisInfo() bool

Returns true if device is laptop (untested on laptop)


### GetGpuData() []GpuInfo

Returns an array of gpus with the following information:
* vendor name
* product name
* local driver version


### GetHwInfo() DeviceInfo

Returns a struct containing all device information 
quaryable by this package.