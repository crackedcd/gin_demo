package LoggerUtils

import (
	"../DatetimeUtils"
	"fmt"
	"runtime"
	"runtime/debug"
)

func Info(info string) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(DatetimeUtils.GetTimeStrDuration("0s"), " [INFO] ", file, line, info)
}

func Error(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(DatetimeUtils.GetTimeStrDuration("0s"), " [ERROR] ", file, line, err)
	}
}

func DebugError(err error) {
	Error(err)
	debug.PrintStack()
}

func FatalError(err error) {
	Error(err)
	runtime.Goexit()
}
