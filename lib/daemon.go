package lib

import (
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Peer struct {
	Ip   net.IP
	Port int
}

func (p Peer) String() string {
	return p.Ip.String() + ":" + strconv.Itoa(p.Port)
}

// todo check sharik.json
// todo discover few ports at a time
// todo use Ping to form a list of devices in the network
// todo reuse local ip, dont just fetch it every time needed
func RunDiscoveryDaemon(sleep time.Duration, feedback func([]Peer)) {
	for {
		local, _ := GetLocalIp()
		ips := strings.Split(local.String(), ".")
		ip := ips[0] + "." + ips[1] + "." + ips[2] + "."

		feedback(Run(ip, 50500, local.String()))

		time.Sleep(sleep)
	}

}

func Run(ip string, port int, exclude string) []Peer {
	if !strings.HasSuffix(ip, ".") {
		return []Peer{}
	}

	wg := sync.WaitGroup{}
	wg.Add(253)

	var peers []Peer

	for i := 1; i < 255; i++ {
		_ip := ip + strconv.Itoa(i)

		if _ip == exclude {
			continue
		}

		go func() {

			if DoesPortExist(_ip, port) {
				peers = append(peers, Peer{
					Ip:   net.ParseIP(_ip),
					Port: port,
				})
			}

			wg.Done()

		}()

	}
	wg.Wait()
	return peers
}
