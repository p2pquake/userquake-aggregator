package aggregate

import (
	"log"

	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
)

type CompatibleAggregator struct {
}

func (c CompatibleAggregator) Aggregate(records []epsp.Record) Result {
	log.Fatalln("Not implemented")
	return Result{}
}
