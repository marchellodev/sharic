package lib

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func RunServer(fileName string, port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {

		file, err := os.Open(fileName)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		stats, _ := file.Stat()

		w.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
		w.Header().Set("Content-Transfer-Encoding", "Binary")
		w.Header().Set("Content-disposition", "attachment; filename=f.txt")
		w.Header().Set("Content-length", strconv.FormatInt(stats.Size(), 10))

		if _, err = io.Copy(w, file); err != nil {
			return
		}
	})
	_ = http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
