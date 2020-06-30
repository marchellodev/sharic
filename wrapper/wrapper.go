package wrapper

import (
	"github.com/marchellodev/sharic/lib"
)

func GetIp() string {
	ip, _ := lib.GetLocalIp()
	//if ping == false {
	//	_, _ = http.Get("https://marchello.cf/shas/debug?app=sharic&version=0.1&platform=" + runtime.GOOS)
	//
	//	fmt.Println("IP was fetched the ugly way")
	//}
	return ip.String()
}
