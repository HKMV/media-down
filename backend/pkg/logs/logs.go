package logs

import (
	"context"
	"fmt"
	"io"
	"log"
	"media-down/backend/pkg/common"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func init() {
	//appPath := common.GetAppPath()
	//appPath=appPath+`log.log`
	// appPath := "logs.log"
	logFile, err := os.OpenFile(LogName(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)
}

func LogName() string {
	path := "logs"
	if !common.PathIsExis(path) {
		os.MkdirAll(path, os.ModePerm)
	}
	return fmt.Sprintf("%s/%s.log", path, time.Now().Format(time.DateOnly))
}

var ctx context.Context

func SetContext(c context.Context) {
	ctx = c
}

func Fatal(format string, v ...interface{}) {
	runtime.LogFatalf(ctx, format, v...)
}

// Error Log error level message.
func Error(format string, v ...interface{}) {
	runtime.LogErrorf(ctx, format, v...)
}

// Debug Log debug level message.
func Debug(format string, v ...interface{}) {
	runtime.LogDebugf(ctx, format, v...)
}

// Warn Log warn level message.
// compatibility alias for Warning()
func Warn(format string, v ...interface{}) {
	runtime.LogWarningf(ctx, format, v...)
}

// Info Log info level message.
// compatibility alias for Informational()
func Info(format string, v ...interface{}) {
	runtime.LogInfof(ctx, format, v...)
}

// Trace Log trace level message.
// compatibility alias for Debug()
func Trace(format string, v ...interface{}) {
	runtime.LogTracef(ctx, format, v...)
}
