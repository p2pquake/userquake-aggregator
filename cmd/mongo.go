package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"github.com/p2pquake/userquake-aggregator/pkg/supplier"
	"github.com/spf13/cobra"
)

type Config struct {
	MongoURI        string `envconfig:"mongo_uri" required:"true"`
	MongoDatabase   string `envconfig:"mongo_database" required:"true"`
	MongoCollection string `envconfig:"mongo_collection" required:"true"`
}

func init() {
	rootCmd.AddCommand(mongoCmd)
}

var mongoCmd = &cobra.Command{
	Use:   "mongo",
	Short: "Follow MongoDB and insert evaluation results",
	Run:   mongoRun,
}

func mongoRun(cmd *cobra.Command, args []string) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("env: %v", config)

	log.Println("Starting...")
	ctx, cancel := context.WithCancel(context.Background())

	m := supplier.Mongo{}
	m.Start(ctx, config.MongoURI, config.MongoDatabase, config.MongoCollection)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

L:
	for {
		select {
		case <-quit:
			break L
		}
	}

	log.Println("Exiting...")
	cancel()
	<-m.Done
	log.Println("Bye!")
}
