package cmd

import (
	"errors"
	"fmt"
	"github.com/marchello/sharic/lib"
	"github.com/spf13/cobra"
	"os"
	"sync"
	"time"
)

// todo complete cmd
var rootCmd = &cobra.Command{
	Use:   "sharic [file]",
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

		go func() {
			lib.RunDiscoveryDaemon(10*time.Second, func(peers []lib.Peer) {
				for _, p := range peers {
					// todo dont repeat
					fmt.Println("Discovered sharik: http://" + p.String())
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
