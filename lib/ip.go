package lib

import (
	"net"
	"strconv"
	"strings"
	"time"
)

// todo works awfully (172.20.10.4)
func GetLocalIp() net.IP {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	var arr []net.IP

	for _, a := range addresses {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
			arr = append(arr, ip.IP)
		}
	}
	//fmt.Println(arr)

	for _, el := range arr {
		if strings.HasPrefix(el.String(), "192.168") {
			return el
		}
	}

	for _, el := range arr {
		if strings.HasPrefix(el.String(), "172") {
			return el
		}
	}

	for _, el := range arr {
		if strings.HasPrefix(el.String(), "10.") {
			return el
		}
	}

	return nil
}

func GetLocalPort() int {
	ip := GetLocalIp().String()

	if port := 50500; !DoesPortExist(ip, port) {
		return port
	}
	if port := 50050; !DoesPortExist(ip, port) {
		return port
	}
	if port := 56788; !DoesPortExist(ip, port) {
		return port
	}
	if port := 56788; !DoesPortExist(ip, port) {
		return port
	}

	listener, _ := net.Listen("tcp", ":0")

	port := listener.Addr().(*net.TCPAddr).Port
	_ = listener.Close()

	return port
}

func DoesPortExist(ip string, p int) bool {
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort(ip, strconv.Itoa(p)), time.Second)

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}
