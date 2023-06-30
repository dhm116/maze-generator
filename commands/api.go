package commands

import (
	"net"

	"github.com/spf13/cobra"

	"github.com/dhm116/maze-generator/api"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the mazebot API",
	Run:   runAPI,
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().Int("port", 8080, "Changes the server port")
	apiCmd.Flags().IP("host", net.ParseIP("127.0.0.1"), "Specifies the host addresses to respond to")
}

func runAPI(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetIP("host")
	port, _ := cmd.Flags().GetInt("port")

	api.RunServer(host.String(), port)
}
