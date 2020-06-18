package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getIp() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, a := range addresses {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() {

			if ip.IP.To4() != nil {
				// todo as switch
				if strings.HasPrefix(ip.IP.String(), "192.168") {
					return ip.IP.String()
				}
				if strings.HasPrefix(ip.IP.String(), "172") {
					return ip.IP.String()
				}
				if strings.HasPrefix(ip.IP.String(), "10.") {
					return ip.IP.String()
				}
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(getIp() + ":8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {

	FileName := "f.txt"
	file, err := os.Open(FileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	stats, _ := file.Stat()

	w.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	w.Header().Set("Content-Transfer-Encoding", "Binary")
	w.Header().Set("Content-disposition", "attachment; filename=f.txt")
	w.Header().Set("Content-length", strconv.FormatInt(stats.Size(), 10))

	fmt.Println(strconv.FormatInt(stats.Size(), 10))

	if _, err = io.Copy(w, file); err != nil {
		return
	}

}
