package cmd

import (
	"errors"
	"fmt"
	"github.com/marchellodev/sharic/lib"
	"github.com/spf13/cobra"
	"os"
	"sync"
	"time"
)

// todo complete cmd
var rootCmd = &cobra.Command{
	Use:   "sharic [file] [-tunnel]",
	Short: "Sharic is a cli version of Sharik",
	Long:  "Sharic shares files in your local network",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a file argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := lib.GetLocalIp()
		port := lib.GetLocalPort()
		go func() {
			lib.RunServer(args[0], port)
		}()
		fmt.Printf("Serving at http://%s:%d\n", ip, port)

		if len(args) == 2 {
			url := lib.RunFrp(port)
			fmt.Println("Tunnel is active: " + url)
		}

		go func() {
			lib.RunDiscoveryDaemon(2*time.Second, func(peer lib.Peer, status int) {
				// todo interactive list
				if status == lib.PeerAdd {
					fmt.Println("Discovered sharik: http://" + peer.String())
				}
				if status == lib.PeerRemove {
					fmt.Print("\033[H\033[2J")
					fmt.Println("Sharik was closed: http://" + peer.String())
				}
			})
		}()

		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
