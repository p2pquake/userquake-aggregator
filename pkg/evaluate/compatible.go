package evaluate

import (
	"time"

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

// peerCtx は 1 クラスタの areapeers 由来の集計を保持する。
// クラスタ開始時点のピア状況はクラスタの生存期間中ずっと不変なので、一度だけ作って共有する。
type peerCtx struct {
	total  int
	area   map[epsp.AreaCode]int // mp=1
	pref   map[epsp.AreaCode]int // mp=10
	region map[epsp.AreaCode]int // mp=100
}

func newPeerCtx(p epsp.Areapeers) peerCtx {
	ctx := peerCtx{
		area:   map[epsp.AreaCode]int{},
		pref:   map[epsp.AreaCode]int{},
		region: map[epsp.AreaCode]int{},
	}
	for _, a := range p.Areas {
		ctx.total += a.Peer
		ctx.area[a.Id] += a.Peer
		ctx.pref[a.Id/10*10] += a.Peer
		ctx.region[a.Id/100*100] += a.Peer
	}
	return ctx
}

func (c CompatibleEvaluator) Evaluate(r aggregate.Result) Result {
	ie := NewIncrementalEvaluator(r.StartedAt, r.Areapeers)
	ie.Advance(r.Userquakes)
	return ie.Result()
}

// IncrementalEvaluator は 1 クラスタの評価を増分的に行う。
//
// クラスタは時刻昇順に末尾追記（append-only）で伸びる前提で、Userquake が増えた分だけを処理する。
// 従来の Evaluate は呼び出しごとにクラスタ全体を O(N^2) で再評価していたが、
//   - プレフィックス u[0:i] は周期をまたいでも不変（追記のみ）
//   - ピア状況 peerCtx はクラスタ開始時点で固定
//
// という性質を使い、新着 k 件だけを処理することで地震バースト中の負荷を抑える。
//
// 1 件追加あたりの処理は、信頼度・表示判定ともに running state を使って償却 O(1)
// （地域率の走査はたかだか地域数ぶんの定数）で、クラスタ全体では O(N) になる。
type IncrementalEvaluator struct {
	started epsp.EPSPTime
	ctx     peerCtx

	// running な userquake 件数（現在処理済みのプレフィックスに対応）
	countArea   map[epsp.AreaCode]int // mp=1
	countPref   map[epsp.AreaCode]int // mp=10
	countRegion map[epsp.AreaCode]int // mp=100

	processed int           // 処理済み Userquake 件数
	first     *time.Time    // u[0].Time（speed 計算用）
	last      epsp.EPSPTime // u[processed-1].Time（UpdatedAt と整合性チェック用）
	lastArea  epsp.AreaCode // u[processed-1].Area（整合性チェック用）

	// 全体信頼度: 各プレフィックスを評価した最大値
	maxAreaRate   float64    // 最大エリア発信率（件数追加で単調非減少なので running max で十分）
	maxConfidence Confidence // これまでの全プレフィックスでの最大信頼度

	// 表示判定（calcAreaConfidence 相当）: 表示対象エリアの集合と、その座標のバウンディングボックス
	display map[epsp.AreaCode]struct{}
	bboxSet bool
	left    int
	right   int
	top     int
	bottom  int
}

// NewIncrementalEvaluator はクラスタ開始時刻とピア状況を固定した評価器を作る。
func NewIncrementalEvaluator(started epsp.EPSPTime, peers epsp.Areapeers) *IncrementalEvaluator {
	ie := &IncrementalEvaluator{started: started, ctx: newPeerCtx(peers)}
	ie.reset()
	return ie
}

func (ie *IncrementalEvaluator) reset() {
	ie.countArea = map[epsp.AreaCode]int{}
	ie.countPref = map[epsp.AreaCode]int{}
	ie.countRegion = map[epsp.AreaCode]int{}
	ie.processed = 0
	ie.first = nil
	ie.last = epsp.EPSPTime{}
	ie.lastArea = 0
	ie.maxAreaRate = 0
	ie.maxConfidence = 0
	ie.display = map[epsp.AreaCode]struct{}{}
	ie.bboxSet = false
	ie.left, ie.right, ie.top, ie.bottom = 0, 0, 0, 0
}

// Advance は us のうち未処理の末尾分だけを取り込む。
// すでに処理済みのプレフィックスが us と矛盾していた場合（順序の入れ替わりや縮小）は、
// 安全側に倒して最初から作り直す。
func (ie *IncrementalEvaluator) Advance(us []epsp.Userquake) {
	if ie.staleAgainst(us) {
		ie.reset()
	}
	for ie.processed < len(us) {
		ie.step(us[ie.processed])
	}
}

// staleAgainst は処理済みプレフィックスと us の先頭・末尾・件数を O(1) で照合する。
// append-only が崩れている場合のみ true を返す。
func (ie *IncrementalEvaluator) staleAgainst(us []epsp.Userquake) bool {
	if ie.processed == 0 {
		return false
	}
	if len(us) < ie.processed {
		return true
	}
	if ie.first == nil || us[0].Time.Time == nil || !us[0].Time.Time.Equal(*ie.first) {
		return true
	}
	last := us[ie.processed-1]
	if last.Area != ie.lastArea {
		return true
	}
	return ie.last.Time == nil || last.Time.Time == nil || !last.Time.Time.Equal(*ie.last.Time)
}

// step は Userquake を 1 件取り込み、running state を更新する。
func (ie *IncrementalEvaluator) step(u epsp.Userquake) {
	area := u.Area
	ie.countArea[area]++
	ie.countPref[area/10*10]++
	ie.countRegion[area/100*100]++

	if ie.processed == 0 {
		ie.first = u.Time.Time
	}
	ie.processed++
	ie.last = u.Time
	ie.lastArea = area

	// 最大エリア発信率を更新（分母 ctx.area は固定なので比は単調非減少）
	if area/100 != 9 {
		if peers, ok := ie.ctx.area[area]; ok {
			if r := float64(ie.countArea[area]) / float64(peers); r > ie.maxAreaRate {
				ie.maxAreaRate = r
			}
		}
	}

	// 先頭 2 件はかならず表示対象（シード）。信頼度評価・表示判定は 3 件目以降。
	if ie.processed < 3 {
		ie.addDisplay(area)
		return
	}

	// 全体信頼度: このプレフィックスを評価して最大値を更新
	if c := ie.prefixConfidence(); c > ie.maxConfidence {
		ie.maxConfidence = c
	}

	// 表示判定（座標による）: 現在の表示集合のバウンディングボックス内なら最新エリアを追加
	if _, shown := ie.display[area]; !shown {
		if pos, ok := areaPositions[area]; ok && ie.bboxSet {
			if pos.x >= ie.left-allowXRange && pos.x <= ie.right+allowXRange &&
				pos.y >= ie.top-allowYRange && pos.y <= ie.bottom+allowYRange {
				ie.addDisplay(area)
			}
		}
	}

	// 表示判定（発信数による）
	if _, shown := ie.display[area]; !shown {
		if peers, ok := ie.ctx.area[area]; ok {
			count := ie.countArea[area]
			ratio := float64(count) / float64(peers)
			if (count >= 3 && ratio >= 0.5) || (count >= 5 && ratio >= 0.1) {
				ie.addDisplay(area)
			}
		}
	}
}

// addDisplay はエリアを表示集合に加え、バウンディングボックスを広げる。
// 座標未登録のエリアは (0,0) として扱う（原実装の bbox 計算と一致させるため）。
func (ie *IncrementalEvaluator) addDisplay(area epsp.AreaCode) {
	if _, ok := ie.display[area]; ok {
		return
	}
	ie.display[area] = struct{}{}

	p := areaPositions[area] // 未登録なら zero value {0, 0}
	if !ie.bboxSet {
		ie.left, ie.right, ie.top, ie.bottom = p.x, p.x, p.y, p.y
		ie.bboxSet = true
		return
	}
	if p.x < ie.left {
		ie.left = p.x
	}
	if p.x > ie.right {
		ie.right = p.x
	}
	if p.y < ie.top {
		ie.top = p.y
	}
	if p.y > ie.bottom {
		ie.bottom = p.y
	}
}

// prefixConfidence は現在のプレフィックス（processed 件）の信頼度を計算する。
func (ie *IncrementalEvaluator) prefixConfidence() Confidence {
	n := float64(ie.processed)
	speed := n / ie.last.Time.Sub(*ie.first).Seconds()
	rate := n / float64(ie.ctx.total)
	areaRate := ie.maxAreaRate
	regionRate := ie.maxRegionRate()
	return confidenceFromRates(speed, rate, areaRate, regionRate)
}

// maxRegionRate は現在のプレフィックスでの最大地域発信率（地域件数 / 総件数）を返す。
// 分母（総件数）が件数追加で増えるため running max にはできないが、地域数は定数なので毎回走査してよい。
func (ie *IncrementalEvaluator) maxRegionRate() float64 {
	rate := 0.0
	for region, count := range ie.countRegion {
		if region/100 == 9 {
			continue
		}
		if r := float64(count) / float64(ie.processed); r > rate {
			rate = r
		}
	}
	return rate
}

// Result は現在の処理済み状態から評価結果を組み立てる。
func (ie *IncrementalEvaluator) Result() Result {
	return Result{
		StartedAt:      ie.started,
		UpdatedAt:      ie.last,
		Count:          ie.processed,
		Confidence:     ie.maxConfidence,
		AreaConfidence: ie.areaConfidence(),
	}
}

// areaConfidence は表示対象エリアごとの信頼度・件数を組み立てる。
// 原実装ではプレフィックスごとに値を上書きしていたが、最終値は必ず全件（i=N）の計算結果になるため、
// ここでは確定した全件カウントで一度だけ算出する。
func (ie *IncrementalEvaluator) areaConfidence() map[epsp.AreaCode]AreaResult {
	result := map[epsp.AreaCode]AreaResult{}
	if ie.processed < 3 {
		return result
	}

	for area, count := range ie.countArea {
		if _, shown := ie.display[area]; shown {
			conf := Confidence(0)
			if _, ok := ie.ctx.area[area]; ok {
				conf = ie.areaConfidenceValue(area)
			}
			result[area] = AreaResult{Confidence: conf, Count: count}
		} else {
			result[area] = AreaResult{Confidence: -1, Count: count}
		}
	}
	return result
}

// areaConfidenceValue は表示対象かつピアの存在するエリアの信頼度を、全件カウントから計算する。
func (ie *IncrementalEvaluator) areaConfidenceValue(area epsp.AreaCode) Confidence {
	count := ie.countArea[area]

	pc := float64(count) / float64(ie.ctx.area[area]) * 100
	if float64(count)/float64(ie.ctx.total) < 0.01 {
		pc *= float64(count) / float64(ie.ctx.total) * 100
	} else {
		pc *= 1.2
	}

	pc *= float64(ie.countPref[area/10*10])/float64(ie.ctx.pref[area/10*10])*5 + 1
	pc *= float64(ie.countRegion[area/100*100])/float64(ie.ctx.region[area/100*100])*5 + 1
	pc = minF(pc, 100)
	if pc < 0 {
		pc = 0
	}

	return Confidence(pc / 100)
}

// confidenceFromRates は speed / rate / areaRate / regionRate から信頼度区分を決める。
func confidenceFromRates(speed, rate, areaRate, regionRate float64) Confidence {
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

func minF(values ...float64) float64 {
	min := values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return min
}
