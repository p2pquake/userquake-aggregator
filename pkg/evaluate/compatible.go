package evaluate

import (
	"github.com/p2pquake/userquake-aggregator/pkg/aggregate"
	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
)

type position struct {
	x int
	y int
}

const allowXRange = 35
const allowYRange = 45

var areaPositions map[epsp.AreaCode]position = map[epsp.AreaCode]position{
	10:  {x: 459, y: 88},
	15:  {x: 426, y: 120},
	20:  {x: 398, y: 118},
	25:  {x: 439, y: 88},
	30:  {x: 480, y: 79},
	35:  {x: 504, y: 74},
	40:  {x: 476, y: 55},
	45:  {x: 480, y: 21},
	50:  {x: 530, y: 68},
	55:  {x: 444, y: 101},
	60:  {x: 503, y: 110},
	65:  {x: 524, y: 93},
	70:  {x: 557, y: 88},
	75:  {x: 576, y: 80},
	100: {x: 421, y: 163},
	105: {x: 444, y: 172},
	106: {x: 444, y: 140},
	110: {x: 458, y: 197},
	111: {x: 456, y: 216},
	115: {x: 438, y: 202},
	120: {x: 430, y: 229},
	125: {x: 420, y: 245},
	130: {x: 409, y: 193},
	135: {x: 423, y: 182},
	140: {x: 396, y: 223},
	141: {x: 411, y: 223},
	142: {x: 407, y: 236},
	143: {x: 396, y: 246},
	150: {x: 411, y: 266},
	151: {x: 427, y: 276},
	152: {x: 389, y: 267},
	200: {x: 412, y: 296},
	205: {x: 399, y: 308},
	210: {x: 388, y: 283},
	215: {x: 391, y: 297},
	220: {x: 357, y: 285},
	225: {x: 354, y: 298},
	230: {x: 375, y: 305},
	231: {x: 384, y: 312},
	232: {x: 357, y: 308},
	240: {x: 409, y: 323},
	241: {x: 396, y: 320},
	242: {x: 391, y: 336},
	250: {x: 373, y: 318},
	255: {x: 374, y: 359},
	260: {x: 387, y: 394},
	265: {x: 427, y: 435},
	270: {x: 378, y: 329},
	275: {x: 364, y: 325},
	300: {x: 332, y: 270},
	301: {x: 357, y: 268},
	302: {x: 378, y: 250},
	305: {x: 332, y: 227},
	310: {x: 315, y: 281},
	315: {x: 299, y: 283},
	320: {x: 297, y: 256},
	325: {x: 283, y: 284},
	330: {x: 274, y: 300},
	335: {x: 258, y: 311},
	340: {x: 351, y: 319},
	345: {x: 339, y: 312},
	350: {x: 332, y: 284},
	351: {x: 333, y: 298},
	355: {x: 321, y: 312},
	400: {x: 303, y: 297},
	405: {x: 287, y: 313},
	410: {x: 355, y: 344},
	411: {x: 349, y: 330},
	415: {x: 332, y: 332},
	416: {x: 320, y: 342},
	420: {x: 303, y: 344},
	425: {x: 294, y: 328},
	430: {x: 270, y: 341},
	435: {x: 272, y: 356},
	440: {x: 268, y: 320},
	445: {x: 264, y: 331},
	450: {x: 242, y: 312},
	455: {x: 249, y: 327},
	460: {x: 246, y: 337},
	465: {x: 244, y: 346},
	470: {x: 225, y: 311},
	475: {x: 227, y: 328},
	480: {x: 253, y: 356},
	490: {x: 235, y: 359},
	495: {x: 247, y: 375},
	500: {x: 209, y: 308},
	505: {x: 189, y: 307},
	510: {x: 165, y: 311},
	514: {x: 163, y: 276},
	515: {x: 140, y: 320},
	520: {x: 193, y: 319},
	525: {x: 190, y: 333},
	530: {x: 165, y: 324},
	535: {x: 161, y: 337},
	540: {x: 116, y: 337},
	541: {x: 130, y: 344},
	545: {x: 103, y: 346},
	550: {x: 200, y: 357},
	555: {x: 210, y: 365},
	560: {x: 196, y: 347},
	570: {x: 173, y: 355},
	575: {x: 154, y: 359},
	576: {x: 145, y: 372},
	580: {x: 196, y: 375},
	581: {x: 177, y: 367},
	582: {x: 158, y: 384},
	600: {x: 77, y: 354},
	601: {x: 97, y: 355},
	602: {x: 90, y: 364},
	605: {x: 85, y: 374},
	610: {x: 65, y: 363},
	615: {x: 72, y: 374},
	620: {x: 52, y: 370},
	625: {x: 63, y: 385},
	630: {x: 59, y: 336},
	635: {x: 21, y: 369},
	640: {x: 99, y: 388},
	641: {x: 87, y: 389},
	645: {x: 91, y: 405},
	646: {x: 74, y: 399},
	650: {x: 114, y: 363},
	651: {x: 117, y: 375},
	655: {x: 100, y: 376},
	656: {x: 123, y: 385},
	660: {x: 111, y: 408},
	661: {x: 105, y: 398},
	665: {x: 106, y: 423},
	666: {x: 97, y: 414},
	670: {x: 75, y: 416},
	675: {x: 89, y: 434},
	680: {x: 26, y: 427},
	685: {x: 155, y: 421},
	700: {x: 313, y: 385},
	701: {x: 310, y: 411},
	702: {x: 276, y: 404},
	705: {x: 237, y: 423},
	706: {x: 272, y: 437},
	710: {x: 324, y: 426},
}

type CompatibleEvaluator struct {
}

func (c CompatibleEvaluator) Evaluate(r aggregate.Result) Result {
	result := Result{Confidence: 0, AreaConfidence: map[epsp.AreaCode]Confidence{}}
	result.StartedAt = r.StartedAt

	for i := 3; i <= len(r.Userquakes); i++ {
		u := r.Userquakes[0:i]
		if c := calcConfidence(r.Areapeers, u); c > result.Confidence {
			result.Confidence = c
		}
	}

	result.AreaConfidence = calcAreaConfidence(r.Areapeers, r.Userquakes)

	return result
}

func calcAreaConfidence(p epsp.Areapeers, us []epsp.Userquake) (result map[epsp.AreaCode]Confidence) {
	result = map[epsp.AreaCode]Confidence{}

	// 先頭 2 件はかならず表示対象とする
	for _, u := range us[0:2] {
		result[u.Area] = 0
	}

	// 表示判定
	for i := 3; i <= len(us); i++ {
		u := us[0:i]

		// 座標による
		xs := []int{}
		ys := []int{}
		for k := range result {
			xs = append(xs, areaPositions[k].x)
			ys = append(ys, areaPositions[k].y)
		}

		left := min(xs)
		right := max(xs)
		top := min(ys)
		bottom := max(ys)

		if pos, ok := areaPositions[u[len(u)-1].Area]; ok {
			if pos.x >= left-allowXRange && pos.x <= right+allowXRange &&
				pos.y >= top-allowYRange && pos.y <= bottom+allowYRange {
				result[u[len(u)-1].Area] = 0
			}
		}
		// 発信数による
		areas, uqs := toMap(p, u, 1)
		for k, count := range uqs {
			peers, ok := areas[k]
			if !ok {
				continue
			}

			if _, ok := result[k]; !ok {
				continue
			}

			if (count >= 3 && float64(count)/float64(peers) >= 0.5) ||
				(count >= 5 && float64(count)/float64(peers) >= 0.1) {
				result[k] = 0
			}
		}
	}

	for i := 3; i <= len(us); i++ {
		pArea, uArea := toMap(p, us[0:i], 1)
		pPref, uPref := toMap(p, us[0:i], 10)
		pRegion, uRegion := toMap(p, us[0:i], 100)

		for area, count := range uArea {
			if _, ok := pArea[area]; !ok {
				continue
			}

			if _, ok := result[area]; !ok {
				continue
			}

			// 信頼度
			pc := float64(count) / float64(pArea[area]) * 100
			if float64(count)/float64(sum(p)) < 0.01 {
				pc *= float64(count) / float64(sum(p)) * 100
			} else {
				pc *= 1.2
			}

			pc *= float64(uPref[area/10*10])/float64(pPref[area/10*10])*5 + 1
			pc *= float64(uRegion[area/100*100])/float64(pRegion[area/100*100])*5 + 1
			pc = minF(pc, 100)
			if pc < 0 {
				pc = 0
			}

			result[area] = Confidence(pc / 100)
		}
	}

	return
}

func calcConfidence(p epsp.Areapeers, u []epsp.Userquake) Confidence {
	speed := float64(len(u)) / (u[len(u)-1].Time.Time.Sub(*u[0].Time.Time)).Seconds()
	rate := float64(len(u)) / float64(sum(p))
	areaRate := calcMaxAreaRate(p, u)
	regionRate := calcMaxRegionRate(p, u)

	confidences := []Confidence{0.97015, 0.96774, 0.97024, 0.98052}
	factors := []float64{0.875, 1.0, 1.2, 1.4}

	for i := 3; i >= 0; i-- {
		factor := factors[i]
		confidence := confidences[i]

		if speed >= 0.25*factor && areaRate >= 0.05*factor {
			return confidence
		}

		if speed >= 0.15*factor && areaRate >= 0.3*factor {
			return confidence
		}

		if rate >= 0.01*factor && areaRate >= 0.035*factor {
			return confidence
		}

		if rate >= 0.006*factor && areaRate >= 0.04*factor && regionRate >= minF(1*factor, 1.0) {
			return confidence
		}

		if speed >= 0.18*factor && areaRate >= 0.04*factor && regionRate >= minF(1*factor, 1.0) {
			return confidence
		}
	}

	return 0
}

func max(values []int) int {
	v := values[0]
	for _, e := range values {
		if e > v {
			v = e
		}
	}
	return v
}

func min(values []int) int {
	v := values[0]
	for _, e := range values {
		if v > e {
			v = e
		}
	}
	return v
}

func minF(values ...float64) float64 {
	min := values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return min
}

func sum(p epsp.Areapeers) int {
	peers := 0
	for _, a := range p.Areas {
		peers += a.Peer
	}
	return peers
}

func calcMaxAreaRate(p epsp.Areapeers, us []epsp.Userquake) (rate float64) {
	rate = 0

	areas, uqs := toMap(p, us, 1)
	for area, count := range uqs {
		if area/100 == 9 {
			continue
		}

		if peers, ok := areas[area]; ok {
			if r := float64(count) / float64(peers); r > rate {
				rate = r
			}
		}
	}

	return
}

func calcMaxRegionRate(p epsp.Areapeers, us []epsp.Userquake) (rate float64) {
	rate = 0

	_, uqs := toMap(p, us, 100)
	for area, count := range uqs {
		if area/100 == 9 {
			continue
		}

		if r := float64(count) / float64(len(us)); r > rate {
			rate = r
		}
	}

	return
}

func toMap(p epsp.Areapeers, us []epsp.Userquake, mp int) (areas map[epsp.AreaCode]int, uqs map[epsp.AreaCode]int) {
	areas = map[epsp.AreaCode]int{}
	for _, a := range p.Areas {
		id := a.Id / epsp.AreaCode(mp) * epsp.AreaCode(mp)
		if _, ok := areas[id]; !ok {
			areas[id] = 0
		}
		areas[id] += a.Peer
	}

	uqs = map[epsp.AreaCode]int{}
	for _, u := range us {
		id := u.Area / epsp.AreaCode(mp) * epsp.AreaCode(mp)
		if _, ok := uqs[id]; !ok {
			uqs[id] = 0
		}
		uqs[id]++
	}

	return
}
