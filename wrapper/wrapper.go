package wrapper

import (
	"fmt"
	"github.com/marchello/sharik-cli/lib"
)

func GetIp() string {
	ip, ping := lib.GetLocalIp()
	if ping == false {
		fmt.Println("IP was fetched the ugly way")
	}
	return ip.String()
}
