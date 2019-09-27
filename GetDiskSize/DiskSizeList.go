package GetDiskSize

import (
	"syscall"
	"GetMountDisk"
)
//定义常量
const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type DiskStatus struct {
	TotalSize uint64
	UsedSize uint64
	FreeSize uint64
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.TotalSize = fs.Blocks * uint64(fs.Bsize)
	disk.FreeSize = fs.Bfree * uint64(fs.Bsize)
	disk.UsedSize = disk.TotalSize - disk.FreeSize
	return
}

//获取磁盘使用百分比
func DetachDisk() map[string]float64 {
	DiskPercent := make(map[string]float64)

	for _,value := range GetMountDisk.AccessDiskList() {
		percent := float64(DiskUsage(value).UsedSize)/float64(DiskUsage(value).TotalSize)
		if percent > float64(0.07) {
			DiskPercent[value] = percent
		}
	}
	return DiskPercent
}