package evaluate

import (
	"log"

	"github.com/p2pquake/userquake-aggregator/pkg/aggregate"
)

type CompatibleEvaluator struct {
}

func (c CompatibleEvaluator) Evaluate(records []aggregate.Record) Result {
	log.Fatalln("Not implemented")
	return Result{}
}
