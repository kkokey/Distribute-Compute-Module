package types

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

type DistributeModule struct {
	IPAddr net.IP
}

var instance *DistributeModule
var once sync.Once

func GetModule() *DistributeModule {
	return getRunningModuleInstance()
}

func getRunningModuleInstance() *DistributeModule {
	ip, err := findLocalIP()
	if err != nil {
		ip = nil
	}
	setLocalIP(ip)

	return instance
}

func setLocalIP(ip net.IP) {
	if instance == nil {
		once.Do(func() {
			instance = &DistributeModule{IPAddr: ip}
			fmt.Print("once.Do\n")
		})
	}
	if instance.IPAddr.String() == ip.String() {
		return
	}
	instance.IPAddr = ip
}

func findLocalIP() (net.IP, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, ifVal := range interfaces {
		addrs, err := ifVal.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			inet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			v4 := inet.IP.To4()
			if v4 == nil || v4[0] == 127 { // loopback address
				continue
			}
			return v4, nil
		}
	}
	return nil, errors.New("cannot find local IP address")
}
