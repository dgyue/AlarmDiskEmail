package GetMountDisk

import (
	gofstab "github.com/deniswernert/go-fstab"
)

func AccessDiskList() []string {
	diskList := make([]string,0)

	mounts,_ := gofstab.ParseSystem()

	for _,value := range mounts{
		diskList = append(diskList,value.File)
	}
	return diskList
}