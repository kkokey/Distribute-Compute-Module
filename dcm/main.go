package main

import (
	"kkokey/distribute-compute-module/dcm/core/daemon"

	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	for {
		go daemon.Runner(msg)
		time.Sleep(time.Second * 2)

		daemon.PrintMemUsage()

		msgReceiver := <-msg
		fmt.Println(msgReceiver)
	}
}
