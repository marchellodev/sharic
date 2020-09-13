package lib

import (
	"encoding/json"
	"net/http"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-disposition", "attachment; filename="+fileName)
		http.ServeFile(w, r, fileName)
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
