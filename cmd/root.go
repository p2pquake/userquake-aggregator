package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "userquake-aggregator",
	Short: "Aggregate userquake and add supplements",
	Long: "Aggregate userquake (earthquake detection information) and add supplements such as reliability\n" +
		"https://github.com/p2pquake/userquake-aggregator",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
