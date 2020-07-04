package evaluate

import "github.com/p2pquake/userquake-aggregator/pkg/epsp"

type Confidence float64

type Result struct {
	StartedAt      epsp.EPSPTime
	Confidence     Confidence
	AreaConfidence map[epsp.AreaCode]AreaResult
}

type AreaResult struct {
	Confidence Confidence
	Count      int
}

func (ar AreaResult) Display() string {
	if ar.Confidence < 0 {
		return "F"
	}

	index := int(ar.Confidence * 5)
	return []string{"E", "D", "C", "B", "A", "A"}[index]
}
