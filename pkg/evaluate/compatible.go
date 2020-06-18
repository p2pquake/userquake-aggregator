package evaluate

import (
	"log"

	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
)

type CompatibleEvaluator struct {
}

func (c CompatibleEvaluator) Evaluate(records []epsp.Record) Result {
	log.Fatalln("Not implemented")
	return Result{}
}
