package aggregate

import (
	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
)

type Results map[epsp.EPSPTime]Result

type Result struct {
	StartedAt  epsp.EPSPTime
	Areapeers  epsp.Areapeers
	Userquakes []epsp.Userquake
}
