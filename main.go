package main

import (
	"./router"
	"fmt"
)

func main() {
	var err error

	/** 1.init router*/
	server := router.InitRouter()

	/** 2.server run*/
	err = server.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println("the error reason is", err)
		e := recover()
		fmt.Println("the server is run err", e)
	}
}
