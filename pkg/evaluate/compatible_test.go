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
		aps,
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
		aps,
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
		aps,
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
	confidence(aps, uqs, 2, t)

	uqs = []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 27; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:03:19.000")})
	}
	uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/05 18:03:20.000")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:03:20.000")})
	confidence(aps, uqs, 1, t)
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
	confidence(aps, uqs, 2, t)

	uqs = []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 94; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 202, Time: genTime("2020/01/06 18:00:00.000")})
	}
	for i := 0; i < 10; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/06 18:00:00.000")})
	}
	confidence(aps, uqs, 1, t)
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
	confidence(aps, uqs, 2, t)

	uqs = []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/05 18:00:00.000")})
	for i := 0; i < 10; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 201, Time: genTime("2020/01/06 18:00:00.000")})
	}
	for i := 0; i < 53; i++ {
		uqs = append(uqs, epsp.Userquake{Area: 202, Time: genTime("2020/01/06 18:00:00.000")})
	}
	confidence(aps, uqs, 1, t)
}

func TestType5(t *testing.T) {
	uqs := []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	confidence(aps, uqs, 2, t)

	uqs = []epsp.Userquake{}
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:00.000")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	uqs = append(uqs, epsp.Userquake{Area: 101, Time: genTime("2020/01/05 18:00:22.200")})
	uqs = append(uqs, epsp.Userquake{Area: 102, Time: genTime("2020/01/05 18:00:22.200")})
	confidence(aps, uqs, 0, t)
}

func TestAreaConfidence(t *testing.T) {
	peers := epsp.Areapeers{
		Areas: []epsp.Areapeer{
			{Id: 10, Peer: 63},
			{Id: 15, Peer: 6},
			{Id: 25, Peer: 2},
			{Id: 30, Peer: 5},
			{Id: 35, Peer: 3},
			{Id: 40, Peer: 1},
			{Id: 45, Peer: 1},
			{Id: 50, Peer: 3},
			{Id: 55, Peer: 6},
			{Id: 65, Peer: 10},
			{Id: 70, Peer: 5},
			{Id: 75, Peer: 1},
			{Id: 100, Peer: 20},
			{Id: 105, Peer: 15},
			{Id: 106, Peer: 2},
			{Id: 110, Peer: 5},
			{Id: 111, Peer: 3},
			{Id: 115, Peer: 19},
			{Id: 120, Peer: 21},
			{Id: 125, Peer: 36},
			{Id: 130, Peer: 21},
			{Id: 135, Peer: 10},
			{Id: 140, Peer: 5},
			{Id: 141, Peer: 1},
			{Id: 142, Peer: 10},
			{Id: 143, Peer: 5},
			{Id: 150, Peer: 24},
			{Id: 151, Peer: 7},
			{Id: 152, Peer: 6},
			{Id: 200, Peer: 28},
			{Id: 205, Peer: 30},
			{Id: 210, Peer: 7},
			{Id: 215, Peer: 32},
			{Id: 220, Peer: 5},
			{Id: 225, Peer: 38},
			{Id: 230, Peer: 28},
			{Id: 231, Peer: 149},
			{Id: 232, Peer: 3},
			{Id: 240, Peer: 16},
			{Id: 241, Peer: 108},
			{Id: 242, Peer: 9},
			{Id: 250, Peer: 465},
			{Id: 270, Peer: 222},
			{Id: 275, Peer: 50},
			{Id: 300, Peer: 3},
			{Id: 301, Peer: 10},
			{Id: 302, Peer: 18},
			{Id: 305, Peer: 1},
			{Id: 310, Peer: 7},
			{Id: 315, Peer: 4},
			{Id: 320, Peer: 2},
			{Id: 325, Peer: 10},
			{Id: 330, Peer: 4},
			{Id: 340, Peer: 6},
			{Id: 345, Peer: 7},
			{Id: 350, Peer: 10},
			{Id: 351, Peer: 16},
			{Id: 355, Peer: 6},
			{Id: 400, Peer: 1},
			{Id: 405, Peer: 22},
			{Id: 410, Peer: 2},
			{Id: 411, Peer: 31},
			{Id: 415, Peer: 26},
			{Id: 416, Peer: 23},
			{Id: 420, Peer: 16},
			{Id: 425, Peer: 77},
			{Id: 430, Peer: 18},
			{Id: 435, Peer: 3},
			{Id: 440, Peer: 3},
			{Id: 445, Peer: 3},
			{Id: 455, Peer: 21},
			{Id: 460, Peer: 66},
			{Id: 465, Peer: 37},
			{Id: 470, Peer: 2},
			{Id: 475, Peer: 49},
			{Id: 480, Peer: 15},
			{Id: 490, Peer: 9},
			{Id: 495, Peer: 5},
			{Id: 500, Peer: 2},
			{Id: 505, Peer: 2},
			{Id: 510, Peer: 7},
			{Id: 520, Peer: 2},
			{Id: 525, Peer: 12},
			{Id: 530, Peer: 2},
			{Id: 535, Peer: 19},
			{Id: 541, Peer: 8},
			{Id: 545, Peer: 4},
			{Id: 550, Peer: 6},
			{Id: 555, Peer: 2},
			{Id: 560, Peer: 23},
			{Id: 570, Peer: 8},
			{Id: 575, Peer: 5},
			{Id: 576, Peer: 3},
			{Id: 581, Peer: 7},
			{Id: 600, Peer: 17},
			{Id: 601, Peer: 9},
			{Id: 602, Peer: 2},
			{Id: 605, Peer: 6},
			{Id: 610, Peer: 1},
			{Id: 615, Peer: 1},
			{Id: 620, Peer: 1},
			{Id: 625, Peer: 2},
			{Id: 641, Peer: 9},
			{Id: 646, Peer: 1},
			{Id: 650, Peer: 2},
			{Id: 651, Peer: 7},
			{Id: 655, Peer: 1},
			{Id: 656, Peer: 1},
			{Id: 660, Peer: 3},
			{Id: 665, Peer: 7},
			{Id: 670, Peer: 7},
			{Id: 675, Peer: 4},
			{Id: 701, Peer: 7},
			{Id: 710, Peer: 1},
			{Id: 900, Peer: 77},
			{Id: 901, Peer: 10},
			{Id: 905, Peer: 2},
		}}

	result := confidence(
		peers,
		[]epsp.Userquake{
			{Area: 205, Time: genTime("2020/07/02 23:29:31.900")},
			{Area: 215, Time: genTime("2020/07/02 23:29:35.961")},
			{Area: 200, Time: genTime("2020/07/02 23:29:36.378")},
			{Area: 231, Time: genTime("2020/07/02 23:29:37.989")},
			{Area: 215, Time: genTime("2020/07/02 23:29:39.182")},
			{Area: 241, Time: genTime("2020/07/02 23:29:42.570")},
			{Area: 205, Time: genTime("2020/07/02 23:29:42.900")},
			{Area: 200, Time: genTime("2020/07/02 23:29:43.512")},
			{Area: 215, Time: genTime("2020/07/02 23:29:43.667")},
			{Area: 240, Time: genTime("2020/07/02 23:29:43.719")},
			{Area: 241, Time: genTime("2020/07/02 23:29:44.004")},
			{Area: 231, Time: genTime("2020/07/02 23:29:44.466")},
			{Area: 205, Time: genTime("2020/07/02 23:29:44.563")},
			{Area: 241, Time: genTime("2020/07/02 23:29:45.352")},
			{Area: 205, Time: genTime("2020/07/02 23:29:46.006")},
			{Area: 241, Time: genTime("2020/07/02 23:29:46.706")},
			{Area: 250, Time: genTime("2020/07/02 23:29:47.442")},
			{Area: 250, Time: genTime("2020/07/02 23:29:47.522")},
			{Area: 230, Time: genTime("2020/07/02 23:29:47.537")},
			{Area: 250, Time: genTime("2020/07/02 23:29:48.233")},
			{Area: 205, Time: genTime("2020/07/02 23:29:48.338")},
			{Area: 231, Time: genTime("2020/07/02 23:29:48.465")},
			{Area: 215, Time: genTime("2020/07/02 23:29:48.744")},
			{Area: 900, Time: genTime("2020/07/02 23:29:49.021")},
			{Area: 241, Time: genTime("2020/07/02 23:29:49.103")},
			{Area: 230, Time: genTime("2020/07/02 23:29:49.153")},
			{Area: 270, Time: genTime("2020/07/02 23:29:49.955")},
			{Area: 151, Time: genTime("2020/07/02 23:29:51.255")},
			{Area: 270, Time: genTime("2020/07/02 23:29:51.871")},
			{Area: 270, Time: genTime("2020/07/02 23:29:52.070")},
			{Area: 231, Time: genTime("2020/07/02 23:29:52.165")},
			{Area: 241, Time: genTime("2020/07/02 23:29:53.130")},
			{Area: 231, Time: genTime("2020/07/02 23:29:53.880")},
			{Area: 250, Time: genTime("2020/07/02 23:29:56.442")},
			{Area: 215, Time: genTime("2020/07/02 23:29:56.672")},
			{Area: 240, Time: genTime("2020/07/02 23:29:57.253")},
			{Area: 241, Time: genTime("2020/07/02 23:29:57.857")},
			{Area: 250, Time: genTime("2020/07/02 23:29:58.707")},
			{Area: 250, Time: genTime("2020/07/02 23:29:59.075")},
			{Area: 241, Time: genTime("2020/07/02 23:29:59.181")},
			{Area: 225, Time: genTime("2020/07/02 23:29:59.678")},
			{Area: 250, Time: genTime("2020/07/02 23:30:04.106")},
			{Area: 275, Time: genTime("2020/07/02 23:30:05.225")},
			{Area: 250, Time: genTime("2020/07/02 23:30:07.287")},
			{Area: 220, Time: genTime("2020/07/02 23:30:09.626")},
			{Area: 220, Time: genTime("2020/07/02 23:30:11.352")},
			{Area: 275, Time: genTime("2020/07/02 23:30:12.950")},
			{Area: 250, Time: genTime("2020/07/02 23:30:13.663")},
			{Area: 225, Time: genTime("2020/07/02 23:30:15.389")},
			{Area: 150, Time: genTime("2020/07/02 23:30:21.496")},
			{Area: 231, Time: genTime("2020/07/02 23:30:23.117")},
			{Area: 410, Time: genTime("2020/07/02 23:30:29.460")},
			{Area: 241, Time: genTime("2020/07/02 23:30:30.746")},
			{Area: 231, Time: genTime("2020/07/02 23:30:34.520")},
			{Area: 232, Time: genTime("2020/07/02 23:30:36.176")},
			{Area: 270, Time: genTime("2020/07/02 23:30:39.519")},
			{Area: 125, Time: genTime("2020/07/02 23:30:46.665")},
			{Area: 241, Time: genTime("2020/07/02 23:30:51.433")},
			{Area: 250, Time: genTime("2020/07/02 23:30:51.540")},
			{Area: 241, Time: genTime("2020/07/02 23:30:55.627")},
			{Area: 241, Time: genTime("2020/07/02 23:31:01.814")},
		},
		2,
		t,
	)

	confidenceAreas := []epsp.AreaCode{220, 225, 230, 231, 232, 240, 241, 250, 150, 151, 200, 205, 215, 125, 270, 275, 410}

	if len(confidenceAreas) != len(result.AreaConfidence) {
		t.Errorf("result.AreaConfidence length got %v; want %v", len(result.AreaConfidence), len(confidenceAreas))
	}

	// TODO: 信頼度の検証をする
	for _, area := range confidenceAreas {
		if _, ok := result.AreaConfidence[area]; !ok {
			t.Errorf("result(%v) not exist; want exist", area)
		}
	}
}

func genTime(t string) epsp.EPSPTime {
	e := epsp.EPSPTime{}
	e.UnmarshalJSON([]byte("\"" + t + "\""))
	return e
}

func confidence(aps epsp.Areapeers, uqs []epsp.Userquake, level int, t *testing.T) Result {
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

	return result
}
