package common

import (
	"fmt"
	"os"
	"path/filepath"
)

type AllEnv struct {
}

func (f *AllEnv) JudgeFile(dir string, file string) string {
	filePath := filepath.Join(dir, file)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println(err)
	}
	return filePath
}

func (f *AllEnv) CreateFile(dir string, file string) string {
	filePath := filepath.Join(dir, file)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create(filePath)
	} else {
		fmt.Println("文件存在")
	}
	return filePath
}
