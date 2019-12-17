package types

import (
	"fmt"
	"net"
	"testing"
)

func TestModule(t *testing.T) {
	var module Module
	testIP := "0.0.0.0"

	module.IPAddr = net.ParseIP(testIP)

	if module.IPAddr.String() == "0.0.0.0" {
		t.Log("0.0.0.0 ==", module.IPAddr.String())
	} else {
		t.Fatal("Wrong IP address", testIP)
	}

}

func TestGetMyIP(t *testing.T) {
	var module Module

	userLocalIP := module.MyIP()

	fmt.Println(userLocalIP)

}
