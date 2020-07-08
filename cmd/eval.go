package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/p2pquake/userquake-aggregator/pkg/aggregate"
	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
	"github.com/p2pquake/userquake-aggregator/pkg/evaluate"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(evalCmd)
}

var evalCmd = &cobra.Command{
	Use:   "eval",
	Short: "Parse data from standard input and output the evaluation result",
	Run:   execute,
}

func execute(cmd *cobra.Command, args []string) {
	// read
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Read error: %v", err)
	}

	// parse
	epspRecords := []epsp.Record{}
	err = json.Unmarshal(body, &epspRecords)
	if err != nil {
		log.Fatalf("JSON unmarshal error: %v from %v", err, string(body))
	}

	sort.Slice(epspRecords, func(i, j int) bool { return epspRecords[i].Time.Time.Before(*epspRecords[j].Time.Time) })

	// aggregate & evaluate
	aggregationResults := aggregate.CompatibleAggregator{}.Aggregate(epspRecords)

	evaluationResults := []evaluate.Result{}

	for _, r := range aggregationResults {
		result := evaluate.CompatibleEvaluator{}.Evaluate(r)
		evaluationResults = append(evaluationResults, result)
	}

	json, err := json.Marshal(evaluationResults)
	if err != nil {
		log.Fatalf("JSON marshal error: %v", err)
	}

	fmt.Println(string(json))
}
