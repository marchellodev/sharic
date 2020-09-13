package lib

//
//import (
//	"fmt"
//	"github.com/fatedier/frp/client"
//	"github.com/fatedier/frp/models/config"
//	"github.com/fatedier/frp/utils/log"
//	"github.com/fatedier/golib/crypto"
//	"math/rand"
//	"strconv"
//	"time"
//)
//
//// todo other domains
//func RunFrp(port int) string {
//	crypto.DefaultSalt = "frp"
//	rand.Seed(time.Now().UnixNano())
//
//	code := strconv.Itoa(rand.Intn(999999)) + "_" + strconv.Itoa(rand.Intn(999999))
//	fmt.Println(code)
//
//	go func() {
//		err := runClient(port, code)
//
//		// todo properly handle errors
//		if err != nil {
//			fmt.Println(err)
//			//os.Exit(1)
//		}
//	}()
//
//	return "https://marchello.cf/sharik/" + code
//}
//func runClient(port int, domain string) (err error) {
//	var content string
//	content = "[common]\nserver_addr = 35.246.234.109\nserver_port = 7000\n\n[" + domain + "]\ntype = http\nlocal_port = " + strconv.Itoa(port) + "\ncustom_domains = " + domain + "\n"
//
//	cfg, err := config.UnmarshalClientConfFromIni(content)
//	if err != nil {
//		return err
//	}
//
//	pxyConf, visitorConf, err := config.LoadAllConfFromIni(cfg.User, content, cfg.Start)
//	if err != nil {
//		return err
//	}
//
//	// todo handle logs
//	log.InitLog("", "", "",
//		0, false)
//
//	svr, errRet := client.NewService(cfg, pxyConf, visitorConf, "")
//	if errRet != nil {
//		err = errRet
//		return
//	}
//
//	err = svr.Run()
//
//	return
//}
