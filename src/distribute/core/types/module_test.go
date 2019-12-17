package types

import (
	"net"
	"testing"
)

func TestCreatesModule(t *testing.T) {
	var module DistributeModule
	testIP := "0.0.0.0"

	module.IPAddr = net.ParseIP(testIP)

	if module.IPAddr.String() != testIP {
		t.Fatal("Wrong IP address", testIP)
	} else {
		t.Logf("\"0.0.0.0\" == %q", module.IPAddr.String())
	}
}

func TestGetsMyIPSuccess(t *testing.T) {
	var module DistributeModule

	userLocalIP, err := module.findLocalIP()
	if err != nil {
		t.Errorf("Can not find IP Error Message: [%q]", err.Error())
	} else {
		t.Logf("Get Local IP: [%q]", userLocalIP)
	}
}
