package logs

import (
	"io"
	"log"
	"media-down/backend/pkg/common"
	"os"
	appdebug "runtime/debug"
)

const (
	fatal = "[F] "
	error = "[E] "
	debug = "[D] "
	warn  = "[W] "
	info  = "[I] "
	trace = "[T] "
)

var myLog = log.New(os.Stdout, "", log.LstdFlags)

func init() {
	// 获取日志文件句柄
	// 以 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	appPath := common.GetAppPath()
	logFile, err := os.OpenFile(appPath+`/log.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	/*defer func() {
		err := logFile.Close()
		if err != nil {
			return
		}
	}()
	*/
	// 组合write，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	// 设置输出
	myLog.SetOutput(multiWriter)
}

func Fatal(format string, v ...interface{}) {
	f := fatal + format
	//f += string(appdebug.Stack())
	if v != nil {
		myLog.Fatalf(f, v...)
	} else {
		myLog.Fatal(f)
	}
}

// Error Log error level message.
func Error(format string, v ...interface{}) {
	f := error + format
	if v != nil {
		myLog.Printf(f, v...)
	} else {
		myLog.Print(f)
	}
}

// Debug Log debug level message.
func Debug(format string, v ...interface{}) {
	f := debug + format
	if v != nil {
		myLog.Printf(f, v...)
	} else {
		myLog.Print(f)
	}
}

// Warn Log warn level message.
// compatibility alias for Warning()
func Warn(format string, v ...interface{}) {
	f := warn + format
	if v != nil {
		myLog.Printf(f, v...)
	} else {
		myLog.Print(f)
	}
}

// Info Log info level message.
// compatibility alias for Informational()
func Info(format string, v ...interface{}) {
	f := info + format
	if v != nil {
		myLog.Printf(f, v...)
	} else {
		myLog.Print(f)
	}
}

// Trace Log trace level message.
// compatibility alias for Debug()
func Trace(format string, v ...interface{}) {
	f := trace + format
	f += string(appdebug.Stack())
	if v != nil {
		myLog.Printf(f, v...)
	} else {
		myLog.Print(f)
	}
}
