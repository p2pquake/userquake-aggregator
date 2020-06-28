package evaluate

import (
	"testing"

	"github.com/p2pquake/userquake-aggregator/pkg/aggregate"
	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
)

var aps epsp.Areapeers = epsp.Areapeers{
	Areas: []epsp.Areapeer{
		{Id: 101, Peer: 100},
		{Id: 102, Peer: 200},
		{Id: 201, Peer: 300},
		{Id: 202, Peer: 10000},
	},
}

func TestNotTruly(t *testing.T) {
	confidence(
		[]epsp.Userquake{
			{Area: 101, Time: genTime("2020/01/05 18:00:00.050")},
			{Area: 101, Time: genTime("2020/01/05 18:00:00.100")},
		},
		false,
		t,
	)
}

func TestTruly(t *testing.T) {
	confidence(
		[]epsp.Userquake{
			{Area: 101, Time: genTime("2020/01/05 18:00:00.000")},
			{Area: 201, Time: genTime("2020/01/05 18:00:20.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:24.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:24.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:24.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:24.000")},
		},
		true,
		t,
	)

	confidence(
		[]epsp.Userquake{
			{Area: 101, Time: genTime("2020/01/05 18:00:00.000")},
			{Area: 201, Time: genTime("2020/01/05 18:00:18.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:20.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:20.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:20.000")},
		},
		false,
		t,
	)
}

func genTime(t string) epsp.EPSPTime {
	e := epsp.EPSPTime{}
	e.UnmarshalJSON([]byte("\"" + t + "\""))
	return e
}

// FIXME: レベルごとに判定できるようになっていない.
func confidence(uqs []epsp.Userquake, expect bool, t *testing.T) {
	r := aggregate.Result{
		StartedAt:  uqs[0].Time,
		Areapeers:  aps,
		Userquakes: uqs,
	}

	result := CompatibleEvaluator{}.Evaluate(r)

	if (result.Confidence > 0 && !expect) ||
		(result.Confidence <= 0 && expect) {
		t.Errorf("Confidence got %v; want %v", result.Confidence, expect)
	}
}
