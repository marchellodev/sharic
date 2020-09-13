package lib

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// todo .gitignore

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// todo improve performance (goroutine management)
// todo cache network name
/// if bool is true, ip was fetched the ping way
func GetLocalIp() (net.IP, bool) {
	arr := getLocalIps()

	fmt.Println(arr)

	var result net.IP

	for _, el := range arr {
		if result != nil {
			break
		}

		ips := strings.Split(el.String(), ".")
		if ips[3] == "1" || ips[0] == "10" {
			continue
		}

		ip := ips[0] + "." + ips[1] + "." + ips[2] + "."

		fmt.Println("Starting: " + ip)

		wg := sync.WaitGroup{}
		wg.Add(254)

		for j := 1; j < 255; j++ {

			if result != nil {
				break
			}

			_ip := ip + strconv.Itoa(j)

			if _ip == el.String() {
				wg.Done()
				continue
			}

			go func() {
				if Ping(_ip) {
					result = el
				}

				wg.Done()
			}()

		}

		wg.Wait()

		//for i := 1; i < 255; i++ {
		//
		//}
		//
		//wg.Wait()
	}

	if result != nil {
		return result, true
	}

	return GetOutboundIP(), true
	//for _, el := range arr {
	//	if strings.HasPrefix(el.String(), "192.168") {
	//		return el, false
	//	}
	//}
	//
	//for _, el := range arr {
	//	if strings.HasPrefix(el.String(), "172") {
	//		return el, false
	//	}
	//}
	//
	//for _, el := range arr {
	//	if strings.HasPrefix(el.String(), "10.") {
	//		return el, false
	//	}
	//}
	//
	//return nil, false
}

func getLocalIps() []net.IP {
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

	return arr
}

func GetLocalPort() int {
	result, _ := GetLocalIp()
	ip := result.String()

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

//
//func ValidateLocalIp(input net.IP) bool {
//
//	ips := strings.Split(input.String(), ".")
//	ip := ips[0] + "." + ips[1] + "." + ips[2] + "."
//
//	wg := sync.WaitGroup{}
//	wg.Add(253)
//
//	found := false
//	for i := 1; i < 255; i++ {
//		_ip := ip + strconv.Itoa(i)
//
//		if _ip == el.String() {
//			continue
//		}
//
//		go func() {
//			if Ping(_ip) {
//				found = true
//			}
//
//			wg.Done()
//
//		}()
//
//	}
//	wg.Wait()
//	return found
//}

/*


func GetLocalIp() (net.IP, bool) {
	arr := getLocalIps()

	fmt.Println(arr)

	var result net.IP

	for _, el := range arr {
		if result != nil {
			break
		}

		ips := strings.Split(el.String(), ".")
		if ips[3] == "1" {
			continue
		}

		ip := ips[0] + "." + ips[1] + "." + ips[2] + "."

		fmt.Println("Starting: " + ip)
		for j := 0; j < 51; j++ {

			if result != nil {
				break
			}

			wg := sync.WaitGroup{}
			wg.Add(5)

			for i := 1; i < 6; i++ {
				if j*5+i == 255 {
					wg.Done()
					continue
				}

				fmt.Println(j*5 + i)

				_ip := ip + strconv.Itoa(j*5+i)
				fmt.Println(_ip)

				if _ip == el.String() {
					wg.Done()
					continue
				}

				go func() {
					if Ping(_ip) {
						result = el
					}

					wg.Done()
				}()

			}

			wg.Wait()
		}

		//for i := 1; i < 255; i++ {
		//
		//}
		//
		//wg.Wait()
	}

	if result != nil {
		return result, true
	}

	return GetOutboundIP(), true
	//for _, el := range arr {
	//	if strings.HasPrefix(el.String(), "192.168") {
	//		return el, false
	//	}
	//}
	//
	//for _, el := range arr {
	//	if strings.HasPrefix(el.String(), "172") {
	//		return el, false
	//	}
	//}
	//
	//for _, el := range arr {
	//	if strings.HasPrefix(el.String(), "10.") {
	//		return el, false
	//	}
	//}
	//
	//return nil, false
}



*/
