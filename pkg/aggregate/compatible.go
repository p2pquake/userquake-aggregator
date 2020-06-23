package aggregate

import (
	"time"

	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
)

type CompatibleAggregator struct {
}

func (c CompatibleAggregator) Aggregate(records []epsp.Record) Results {
	results := map[epsp.EPSPTime]Result{}

	var areapeers *epsp.Areapeers
	for _, record := range records {
		if record.Areapeers != nil {
			areapeers = record.Areapeers
			break
		}
	}

	result := Result{Userquakes: []epsp.Userquake{}}
	for _, record := range records {
		if record.Userquake != nil {
			if len(result.Userquakes) > 0 && record.Userquake.Time.Time.Sub(*result.Userquakes[len(result.Userquakes)-1].Time.Time) > 40*time.Second {
				results[result.Userquakes[0].Time] = result
				result = Result{Userquakes: []epsp.Userquake{}}
			}
			if len(result.Userquakes) == 0 {
				result.StartedAt = record.Time
				result.Areapeers = *areapeers
			}
			result.Userquakes = append(result.Userquakes, *record.Userquake)
		}
		if record.Areapeers != nil {
			areapeers = record.Areapeers
		}
	}

	if len(result.Userquakes) > 0 {
		results[result.Userquakes[0].Time] = result
	}

	return results
}
