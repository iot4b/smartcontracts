package utils

import (
	"net"
	"os"
)

// MyIp - возвращаем текущий ip ноды
func MyIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

// Hostname - hostname ноды
func Hostname() string {
	h, err := os.Hostname()
	if err != nil {
		panic("Can't get hostname, err: " + err.Error())
	}
	return h
}
