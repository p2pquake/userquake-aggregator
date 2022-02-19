package supplier

import (
	"testing"

	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
)

func TestAggregateAndEvaluate(t *testing.T) {
	r := aggregateAndEvaluate([]epsp.Record{})

	if len(r) != 0 {
		t.Errorf("Length got %v; want %v", len(r), 0)
	}
}
