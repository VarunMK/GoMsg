package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send a message to another device",
		Long: `
		This command allows you to send a message to another device.
		To run: 
			gomsg send [hostname] [msg]
		`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Enter [user ip] and [msg]")
				return
			}
			if len(args) > 2 {
				fmt.Println("More arguments than expected")
				return
			}
			userip := args[0]
			msg := args[1]
			sendmsg(userip, msg)
			// var flag bool
			// flag = true
			// for flag {
			// 	fmt.Scanln(msg)
			// 	if msg == "E" {
			// 		flag = false
			// 	}
			// 	sendmsg(userip, msg)
			// }
		},
	}
)

func sendmsg(userip string, msg string) {
	// postBody, _ := json.Marshal(map[string]string{
	// 	"msg": msg,
	// })
	// responseBody := bytes.NewBuffer(postBody)
	h := http.NewServeMux()
	s := http.Server{Addr: ":3000", Handler: h}
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, msg)
	}
	h.HandleFunc("/", h1)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
