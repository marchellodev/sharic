package lib

import (
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

func (p Peer) String() string {
	return p.ip.String() + ":" + strconv.Itoa(p.port)
}

// todo check sharik.json
// todo discover other ports
func RunDiscoveryDaemon(sleep time.Duration, feedback func([]Peer)) {
	for {
		ips := strings.Split(GetLocalIp().String(), ".")
		ip := ips[0] + "." + ips[1] + "." + ips[2] + "."

		feedback(run(ip, 50500))

		time.Sleep(sleep)
	}

}

// help needed, there is probably a better approach
func run(ip string, port int) []Peer {
	if !strings.HasSuffix(ip, ".") {
		return []Peer{}
	}

	wg := sync.WaitGroup{}
	wg.Add(256)

	var peers []Peer

	for i := 0; i <= 255; i++ {
		_ip := ip + strconv.Itoa(i)
		go func() {

			if DoesPortExist(_ip, port) {
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
	return peers
}
