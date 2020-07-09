package cmd

import (
	"github.com/p2pquake/userquake-aggregator/pkg/supplier"
	"github.com/spf13/cobra"
)

func init() {
	serverCmd.Flags().IntVarP(&Port, "port", "p", 8080, "listen port")
	rootCmd.AddCommand(serverCmd)
}

var Port int
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start as server",
	Run: func(cmd *cobra.Command, args []string) {
		supplier.Server{}.Start(Port)
	},
}
