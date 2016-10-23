package main

import (
	"fmt"
	"net"
)

// Returns IPv4 of the host for the eth0 interface
func GetIP() net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, i := range ifaces {
		if i.Name == "eth0" {
			addrs, err := i.Addrs()
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP
					}
				}
			}
		}
	}
	return nil
}

/*
func main() {
	ip := GetIP()
	if ip != nil {
		fmt.Println(ip)
	}
}
*/
