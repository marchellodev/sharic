package wrapper

import (
	"github.com/marchello/sharik-cli/lib"
)

func GetIp() string {
	ip, ping := lib.GetLocalIp()
	if ping == false {
		//fmt.Println("Fetched the ugly way")
	}
	//fmt.Println("Fetched the good way")
	return ip.String()
}
