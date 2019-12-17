package types

import (
	"fmt"
	"net"
)

type Module struct {
	IPAddr net.IP
}

func (m Module) MyIP() net.IP {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Something error...")
	}

	var ip net.IP
	for _, i := range interfaces {
		addresses, err := i.Addrs()
		if err != nil {
			fmt.Println("Something error...")
		}

		for _, addr := range addresses {

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			fmt.Println("This is your ip", ip)
			// process IP address
		}
	}
	return ip
}
