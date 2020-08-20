package wrapper

import (
	"github.com/marchellodev/sharic/lib"
)

func GetIp() string {
	ip, _ := lib.GetLocalIp()
	return ip.String()
}
