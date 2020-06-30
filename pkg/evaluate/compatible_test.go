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

var confidences []Confidence = []Confidence{0.97015, 0.96774, 0.97024, 0.98052}

func TestNotTruly(t *testing.T) {
	confidence(
		[]epsp.Userquake{
			{Area: 101, Time: genTime("2020/01/05 18:00:00.050")},
			{Area: 101, Time: genTime("2020/01/05 18:00:00.100")},
		},
		0,
		t,
	)
}

func TestType1(t *testing.T) {
	confidence(
		[]epsp.Userquake{
			{Area: 101, Time: genTime("2020/01/05 18:00:00.000")},
			{Area: 201, Time: genTime("2020/01/05 18:00:20.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:24.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:24.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:24.000")},
			{Area: 101, Time: genTime("2020/01/05 18:00:24.000")},
		},
		2,
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
		0,
		t,
	)
}

func TestType2(t *testing.T) {
	uqs := []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 28; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:03:19.000")})
	}
	uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/05 18:03:20.000")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:03:20.000")})
	confidence(uqs, 2, t)

	uqs = []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 27; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:03:19.000")})
	}
	uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/05 18:03:20.000")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:03:20.000")})
	confidence(uqs, 1, t)
}

func TestType3(t *testing.T) {
	uqs := []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 94; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 202, Time: genTime("2020/01/06 18:00:00.000")})
	}
	for i := 0; i < 11; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/06 18:00:00.000")})
	}
	confidence(uqs, 2, t)

	uqs = []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 94; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 202, Time: genTime("2020/01/06 18:00:00.000")})
	}
	for i := 0; i < 10; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/06 18:00:00.000")})
	}
	confidence(uqs, 1, t)
}

func TestType4(t *testing.T) {
	uqs := []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 11; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/06 18:00:00.000")})
	}
	for i := 0; i < 52; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 202, Time: genTime("2020/01/06 18:00:00.000")})
	}
	confidence(uqs, 2, t)

	uqs = []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 10; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/06 18:00:00.000")})
	}
	for i := 0; i < 53; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 202, Time: genTime("2020/01/06 18:00:00.000")})
	}
	confidence(uqs, 1, t)
}

func TestType5(t *testing.T) {
	uqs := []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	confidence(uqs, 2, t)

	uqs = []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	uqs = append(uqs, epsp.Userquake{Area: 102, Time: genTime("2020/01/05 18:00:22.200")})
	confidence(uqs, 0, t)
}

func genTime(t string) epsp.EPSPTime {
	e := epsp.EPSPTime{}
	e.UnmarshalJSON([]byte("\"" + t + "\""))
	return e
}

func confidence(uqs []epsp.Userquake, level int, t *testing.T) {
	r := aggregate.Result{
		StartedAt:  uqs[0].Time,
		Areapeers:  aps,
		Userquakes: uqs,
	}

	result := CompatibleEvaluator{}.Evaluate(r)

	if level <= 0 && result.Confidence > 0 {
		t.Errorf("Confidence got %v; want %v", result.Confidence, 0)
	}

	if level > 0 && result.Confidence < confidences[level-1] {
		t.Errorf("Confidence got %v; want >= %v", result.Confidence, confidences[level-1])
	}
}
