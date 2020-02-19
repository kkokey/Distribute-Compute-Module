package types

import (
	"errors"
	"net"
)

type DistributeModule struct {
	IPAddr net.IP
}

func NewModule() *DistributeModule {
	ip, err := DistributeModule.findLocalIP(DistributeModule{IPAddr: nil})
	if err != nil {
		return &DistributeModule{IPAddr: nil}
	}
	return &DistributeModule{ip}
}

func (m DistributeModule) findLocalIP() (net.IP, error) {
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
