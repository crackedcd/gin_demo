package main

import (
	"./controller"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	router := controller.GetRouter()
	err := router.Run(":54103"); if err != nil {
		fmt.Println("listen failed.")
	}
}
