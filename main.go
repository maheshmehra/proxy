package main

import (
	"fmt"
	"github.com/maheshmehra/proxy/connectivity"
	"github.com/maheshmehra/proxy/lib"
	"github.com/maheshmehra/proxy/utils"
	"time"
)

func main() {
	// handling exceptions
	defer lib.HandlePanic()

	// loading config data
	status := utils.ConfigSetter("config.json")

	// checking status
	if !status {
		fmt.Println("Unable to read from config...")
	}

	// checking if port is opened
	if !connectivity.HostHttp() {
		return
	}

	for {
		time.Sleep(24 * time.Hour)
	}
}
