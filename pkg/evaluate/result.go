package evaluate

import "github.com/p2pquake/userquake-aggregator/pkg/epsp"

type Confidence float64

type Result struct {
	StartedAt      epsp.EPSPTime
	Confidence     Confidence
	AreaConfidence map[epsp.AreaCode]Confidence
}

type AreaResult struct {
	Confidence Confidence
	Count      int
	Display    string
}
