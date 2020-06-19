package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	ip := getIp().String()

	fmt.Println(ip)

	discoveryDaemon()

	//fmt.Println(portExists("127.0.0.1", 50500))
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
