package main

import (
	"AlarmDisk"
	"GetDiskSize"
)

func main() {
	AlarmDisk.AlarmDiskSizeEmail(GetDiskSize.DetachDisk())
}