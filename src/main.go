package main

import (
	"distribute/core/types"
	"fmt"
	"strings"
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
	strBuilder := strings.Builder{}
	_, err := strBuilder.WriteString("daemon running... local ip addr: ")
	if err != nil {
		fmt.Print(err)
	}

	distributeModule := types.NewModule()
	_, err = strBuilder.WriteString(distributeModule.IPAddr.String());
	if err != nil {
		fmt.Print(err)
	}

	msg <- strBuilder.String()
}
