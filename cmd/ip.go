package cmd

import (
	"fmt"
	"github.com/marchellodev/sharic/lib"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ipCmd)
}

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Prints local ip",
	Long:  `Prints your local ip. You can be accessed by it withing your local network`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getting your ip...")
		ip, _ := lib.GetLocalIp()

		//if ping {
		//	fmt.Println("ip was fetched the ping way")
		//} else {
		//	fmt.Println("ip was fetched the dumb, probably not-working way, please file an issue")
		//}

		fmt.Println("your ip is:", ip)

	},
}
