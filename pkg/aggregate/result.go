package aggregate

import "github.com/p2pquake/userquake-aggregator/pkg/epsp"

type StartedAt string
type Results map[StartedAt]Result

type Result struct {
	StartedAt  StartedAt
	Areapeers  epsp.Areapeers
	Userquakes []epsp.Userquake
}
