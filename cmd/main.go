package main

import (
	"fmt"
	"github.com/wangsin/paul/router"
)

func main() {
	fmt.Println("Hello, Paul!")
	err := router.InitRouter().Run(":8998")
	if err != nil {
		return
	}
}
