package main

import (
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Peer struct {
	ip   net.IP
	port int
}

var (
	// discovered peers
	peers []Peer
)

func discoveryDaemon() {
	go func() {
		for {
			// should be j == 256, but that doesn't work for some reason

			ips := strings.Split(getIp().String(), ".")
			ip := ips[0] + "." + ips[1] + "." + ips[2] + "."

			log.Println("discovered ", len(peers), " peers")

			peers = []Peer{}
			run(ip, 50500)

			time.Sleep(5 * time.Second)
		}

	}()
}

func getIp() net.IP {
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

func run(ip string, port int) {
	if !strings.HasSuffix(ip, ".") {
		return
	}
	log.Println("starting...")

	wg := sync.WaitGroup{}
	wg.Add(256)

	for i := 0; i <= 255; i++ {
		_ip := ip + strconv.Itoa(i)
		go func() {

			if portExists(_ip, port) {
				peers = append(peers, Peer{
					ip:   net.ParseIP(_ip),
					port: port,
				})
			}
			time.Sleep(1 * time.Second)

			wg.Done()

		}()

	}
	wg.Wait()

}

func portExists(ip string, p int) bool {
	timeout := time.Second
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort(ip, strconv.Itoa(p)), timeout)

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}
