package main

import (
	"./common/utils/LoggerUtils"
	"./controller"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	router := controller.GetRouter()
	err := router.Run(":54103")
	LoggerUtils.Error(err)
}
