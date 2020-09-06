# remote-OHM

Fetch data from the remote web server of [OpenHardwareMonitor](https://openhardwaremonitor.org/).

Currently only support two HDDs and one NVIDIA GPU.

> Reason why not supporting more GPU, HDD or AMD GPU is that I don't have them!
> But it's pretty easy to support them if you have, so take it easy to PR!

I hope this will help those who want to make a raspberrypi into a hardware monitor (just as what I am going to do)

## How to use

Edit params in `conf.yml`:
```yml
ip: 192.168.1.100
port: "8085"
```

Use this command to start:
```PowerShell
go run ./remote-OHM.go
```

Then you will get a `data.yml` like this: [data.yml](data.yml)

If you run into error like this:
```
Error while fetching data :
        Get "http://192.168.1.100:8085/data.json": dial tcp 192.168.1.100:8085: i/o timeout
```
You should check your firewall.

## How to build

Building for linux_arm on Windows_amd64:
```PowerShell
$env:CGO_ENABLED = 0
$env:GOOS = "linux"
$env:GOARCH = "arm"
$env:GOARM = 6
go build -o remote-OHM_0.1_linux_armv6 .\remote-OHM.go
```

Building for linux_arm on Windows_amd64:
```PowerShell
$env:CGO_ENABLED = 0
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o .\remote-OHM_0.1_windows_amd64.exe .\remote-OHM.go
```
