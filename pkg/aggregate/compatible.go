package aggregate

import "log"

type CompatibleAggregator struct {
}

func (c CompatibleAggregator) Aggregate(records []Record) Result {
	log.Fatalln("Not implemented")
	return Result{}
}
