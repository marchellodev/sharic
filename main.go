package main

import (
	"github.com/marchello/sharik-cli/cmd"
	"net"
)

func main() {
	cmd.Execute()
	//
	//log.Println("started")
	//log.Println(lib.GetLocalIp())
	//lib.RunDiscoveryDaemon(1*time.Second, func(peers []lib.Peer) {
	//	for _, p := range peers {
	//		// todo dont repeat
	//		fmt.Println("Discovered sharik: http://" + p.String())
	//	}
	//})
}

func consists(ips []net.IP, ip net.IP) bool {
	for _, el := range ips {
		if el.Equal(ip) {
			return true
		}
	}
	return false
}
