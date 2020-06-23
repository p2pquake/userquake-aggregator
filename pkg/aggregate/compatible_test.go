package aggregate

import (
	"testing"
	"time"

	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
)

func TestMain(t *testing.T) {
	timeFirst := epspTime(1000)
	timeSecond := epspTime(1096)
	timeThird := epspTime(1150)

	records := []epsp.Record{
		{Time: timeFirst, Userquake: &epsp.Userquake{Time: timeFirst, Area: 100}},
		{Time: timeFirst, Userquake: &epsp.Userquake{Time: timeFirst, Area: 100}},
		{Time: epspTime(1010), Areapeers: &epsp.Areapeers{Time: epspTime(1010), Areas: []epsp.Areapeer{{Id: 100, Peer: 5}, {Id: 101, Peer: 10}}}},
		{Time: epspTime(1015), Userquake: &epsp.Userquake{Time: epspTime(1015), Area: 101}},
		{Time: epspTime(1055), Userquake: &epsp.Userquake{Time: epspTime(1055), Area: 100}},
		{Time: timeSecond, Userquake: &epsp.Userquake{Time: timeSecond, Area: 100}},
		{Time: epspTime(1097), Userquake: &epsp.Userquake{Time: epspTime(1097), Area: 101}},
		{Time: epspTime(1100), Areapeers: &epsp.Areapeers{Time: epspTime(1100), Areas: []epsp.Areapeer{{Id: 110, Peer: 1}, {Id: 111, Peer: 2}}}},
		{Time: timeThird, Userquake: &epsp.Userquake{Time: timeThird, Area: 110}},
	}

	results := CompatibleAggregator{}.Aggregate(records)
	if len(results) != 3 {
		t.Errorf("len(results) want = %v, got = %v", 3, len(results))
	}

	r, ok := results[timeFirst]
	if !ok {
		t.Errorf("want = exist result[%v], got = not exist result[%v]", timeFirst, timeFirst)
	}
	if !r.Areapeers.Time.Time.Equal(*epspTime(1010).Time) {
		t.Errorf("Areapeers.Time want = %v, got = %v", epspTime(1010).Time, r.Areapeers.Time.Time)
	}
	if len(r.Userquakes) != 4 {
		t.Errorf("len(Userquakes) want = %v, got = %v", 4, len(r.Userquakes))
	}

	r, ok = results[timeSecond]
	if !ok {
		t.Errorf("want = exist result[%v], got = not exist result[%v]", timeSecond, timeSecond)
	}
	if !r.Areapeers.Time.Time.Equal(*epspTime(1010).Time) {
		t.Errorf("Areapeers.Time want = %v, got = %v", epspTime(1010).Time, r.Areapeers.Time.Time)
	}
	if len(r.Userquakes) != 2 {
		t.Errorf("len(Userquakes) want = %v, got = %v", 2, len(r.Userquakes))
	}

	r, ok = results[timeThird]
	if !ok {
		t.Errorf("want = exist result[%v], got = not exist result[%v]", timeThird, timeThird)
	}
	if !r.Areapeers.Time.Time.Equal(*epspTime(1100).Time) {
		t.Errorf("Areapeers.Time want = %v, got = %v", epspTime(1100).Time, r.Areapeers.Time.Time)
	}
	if len(r.Userquakes) != 1 {
		t.Errorf("len(Userquakes) want = %v, got = %v", 1, len(r.Userquakes))
	}

}

func epspTime(offset int64) epsp.EPSPTime {
	t := time.Date(2000, time.April, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(offset) * time.Second)
	return epsp.EPSPTime{Time: &t}
}
