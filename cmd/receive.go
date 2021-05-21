package cmd

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var rcvCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive a msg from another device",
	Long: `
	This command allows you to receive a message from another device.
	To run: 
		gomsg receive [hostname]
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 || len(args) < 1 {
			fmt.Println("Error: 1 arg should be provided, the hostname")
			return
		}
		var hostname = args[0]
		resp, err := http.Get(hostname)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		buf := new(strings.Builder)
		io.Copy(buf, resp.Body)
		// check errors
		fmt.Println(buf.String())
	},
}

func init() {
	rootCmd.AddCommand(rcvCmd)
}
