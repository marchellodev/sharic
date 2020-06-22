package lib

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

type SharikJson struct {
	Version string `json:"sharik"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Os      string `json:"os"`
}

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
		w.Header().Set("Content-disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-length", strconv.FormatInt(stats.Size(), 10))

		if _, err = io.Copy(w, file); err != nil {
			return
		}
	})
	http.HandleFunc("/sharik.json", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		sharik := SharikJson{
			Version: "0.1c",
			Type:    "file",
			Name:    fileName,
			Os:      runtime.GOOS,
		}
		bytes, _ := json.Marshal(sharik)
		_, _ = w.Write(bytes)

	})
	_ = http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
