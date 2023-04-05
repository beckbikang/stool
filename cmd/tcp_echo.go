package cmd

import (
	"log"

	"stool/component"

	"github.com/spf13/cobra"
)

var ip string
var port int
var echo_str string

var tcpEchoClientCmd = &cobra.Command{
	Use:   "tec",
	Short: "tcp echo client",
	Long:  "tcp echo client",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("send data to an echo server")
		log.Printf("%s:%d str=%s", ip, port, echo_str)
		component.TcpClientDail(ip, port, echo_str)
	},
}

func init() {
	tcpEchoClientCmd.Flags().StringVarP(&ip, "ip", "i", "", `server addr`)
	tcpEchoClientCmd.Flags().IntVarP(&port, "port", "p", 8181, "server port")
	tcpEchoClientCmd.Flags().StringVarP(&echo_str, "str", "s", "", `echo str`)
}
