package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	startUpType := os.Args[1]
	switch startUpType {
	case "web":
		port, error := strconv.Atoi(os.Args[2])
		if error != nil {
			fmt.Println("Error while reading port variable, application is exiting")
			break
		}
		if port == 0 {
			fmt.Println("No port specified. Falling to default 8080")
			port = 8080
		}
		startWebServer(port)
	case "console":
		startApplication()
	}
}

func startWebServer(port int) {

}

func startApplication() {

}
