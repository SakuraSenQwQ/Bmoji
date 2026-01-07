package main

import (
	"fmt"
	"main/config"
	"main/method"
	"main/router"
	"os"
)

func main() {
	fmt.Println("Hello")
	config.GetConfig()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "meta":
			method.DownloadEmoteMeta()
		case "list":
			method.DownloadEmoteList()
		}
		return
	}
	router.StartUp()

}
