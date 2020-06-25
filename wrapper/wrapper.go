package wrapper

import (
	"fmt"
	"github.com/marchello/sharic/lib"
	"net/http"
	"runtime"
)

func GetIp() string {
	ip, ping := lib.GetLocalIp()
	if ping == false {
		_, _ = http.Get("https://marchello.cf/shas/debug?app=sharic&version=0.1&platform=" + runtime.GOOS)

		fmt.Println("IP was fetched the ugly way")
	}
	return ip.String()
}
