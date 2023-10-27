package common

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetAppPath 获取运行目录的绝对路径
func GetAppPath() string {
	if path, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		return path
	}
	return os.Args[0]
}

// IsWindows 判断当前系统是否为Windows系统？
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// PathIsExis 判断目录或文件是否存在
func PathIsExis(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
