package main

import (
	"fmt"
	"time"
)

func main() {

	msg := make(chan string)
	for {
		go daemonRunner(msg)
		time.Sleep(time.Second * 2)

		msgReceiver := <-msg
		fmt.Println(msgReceiver)
	}
}

func daemonRunner(msg chan string) {

	msg <- "daemon running..."
}
