package epsp

import (
	"testing"
)

func TestRecord(t *testing.T) {
	jsons := []string{
		`{"_id":{"$oid":"5eeae82e02add6471ccbe922"},"code":551,"earthquake":{"time":"2020/06/18 13:01:00","hypocenter":{"name":"宮城県南部","latitude":37.9,"longitude":140.9,"depth":80,"magnitude":3.3},"maxScale":20,"domesticTsunami":"None","foreignTsunami":"Unknown"},"expire":null,"issue":{"source":"気象庁","time":"2020/06/18 13:04:00","type":"DetailScale","correct":"None"},"points":[{"pref":"福島県","addr":"福島伊達市霊山町","scale":20,"isArea":false},{"pref":"福島県","addr":"田村市船引町","scale":10,"isArea":false}],"time":"2020/06/18 13:06:06.837","user-agent":"quake_checker/analyzer 20180128","ver":"20180128"}`,
		`{"_id":{"$oid":"5d090bbe02add629bb6fbd3a"},"areas":[],"cancelled":true,"code":552,"created_at":"2019/06/19 01:05:18.344","expire":null,"issue":{"source":"気象庁","time":"2019/06/19 01:02:00","type":"Focus"},"time":"2019/06/19 01:04:15.957","user-agent":"tsunami_checker/analyzer 20180128","ver":"20180128"}`,
		`{"_id":{"$oid":"5d08e68a02add621e84c6949"},"areas":[{"name":"山形県","grade":"Watch","immediate":true},{"name":"新潟県上中下越","grade":"Watch","immediate":true},{"name":"佐渡","grade":"Watch","immediate":true},{"name":"石川県能登","grade":"Watch","immediate":false}],"cancelled":false,"code":552,"created_at":"2019/06/18 22:26:34.475","expire":null,"issue":{"source":"気象庁","time":"2019/06/18 22:24:00","type":"Focus"},"time":"2019/06/18 22:25:28.306","user-agent":"tsunami_checker/analyzer 20180128","ver":"20180128"}`,
		`{"_id":{"$oid":"5ec7069302add60b2f68980f"},"areas":[{"id":10,"peer":49},{"id":15,"peer":5},{"id":25,"peer":1},{"id":30,"peer":2},{"id":35,"peer":2},{"id":40,"peer":1},{"id":45,"peer":1},{"id":50,"peer":2},{"id":55,"peer":7},{"id":60,"peer":1},{"id":65,"peer":8},{"id":70,"peer":4},{"id":75,"peer":1},{"id":100,"peer":14},{"id":105,"peer":16},{"id":110,"peer":5},{"id":111,"peer":1},{"id":115,"peer":16},{"id":120,"peer":21},{"id":125,"peer":27},{"id":130,"peer":13},{"id":135,"peer":7},{"id":140,"peer":1},{"id":141,"peer":2},{"id":142,"peer":11},{"id":143,"peer":4},{"id":150,"peer":16},{"id":151,"peer":5},{"id":152,"peer":5},{"id":200,"peer":16},{"id":205,"peer":20},{"id":210,"peer":7},{"id":215,"peer":21},{"id":220,"peer":2},{"id":225,"peer":25},{"id":230,"peer":18},{"id":231,"peer":113},{"id":232,"peer":2},{"id":240,"peer":21},{"id":241,"peer":69},{"id":242,"peer":7},{"id":250,"peer":345},{"id":270,"peer":152},{"id":275,"peer":31},{"id":300,"peer":3},{"id":301,"peer":4},{"id":302,"peer":12},{"id":305,"peer":2},{"id":310,"peer":3},{"id":315,"peer":3},{"id":325,"peer":8},{"id":330,"peer":6},{"id":340,"peer":6},{"id":345,"peer":6},{"id":350,"peer":8},{"id":351,"peer":12},{"id":355,"peer":5},{"id":405,"peer":20},{"id":410,"peer":4},{"id":411,"peer":24},{"id":415,"peer":23},{"id":416,"peer":13},{"id":420,"peer":18},{"id":425,"peer":71},{"id":430,"peer":11},{"id":435,"peer":2},{"id":440,"peer":2},{"id":445,"peer":4},{"id":455,"peer":20},{"id":460,"peer":48},{"id":465,"peer":26},{"id":470,"peer":2},{"id":475,"peer":35},{"id":480,"peer":11},{"id":490,"peer":6},{"id":495,"peer":1},{"id":500,"peer":1},{"id":505,"peer":4},{"id":510,"peer":4},{"id":515,"peer":2},{"id":520,"peer":1},{"id":525,"peer":12},{"id":530,"peer":1},{"id":535,"peer":13},{"id":541,"peer":4},{"id":550,"peer":8},{"id":560,"peer":15},{"id":570,"peer":4},{"id":575,"peer":4},{"id":576,"peer":1},{"id":580,"peer":1},{"id":581,"peer":4},{"id":600,"peer":12},{"id":601,"peer":8},{"id":602,"peer":4},{"id":605,"peer":7},{"id":610,"peer":1},{"id":615,"peer":3},{"id":625,"peer":1},{"id":641,"peer":4},{"id":646,"peer":1},{"id":650,"peer":1},{"id":651,"peer":3},{"id":665,"peer":5},{"id":670,"peer":5},{"id":675,"peer":3},{"id":701,"peer":2},{"id":710,"peer":1},{"id":900,"peer":58},{"id":901,"peer":5},{"id":905,"peer":2}],"code":555,"created_at":"2020/05/22 07:54:11.135","expire":"2020/05/22 07:56:28","hop":8,"time":"2020/05/22 07:54:11.114","uid":"2020/05/22 07:56:28","ver":"20150406"}`,
		`{"_id":{"$oid":"5eeaf6c802add60b3268bd4f"},"area":200,"code":561,"created_at":"2020/06/18 14:08:24.777","expire":"2020/06/18 14:09:24","hop":1,"time":"2020/06/18 14:08:24.772","uid":"9999920200618140824767","ver":"20150406"}`,
	}

	for _, json := range jsons {
		r, err := Parse(json)
		if err != nil {
			t.Errorf("record (%v) parse error: %v", json, err)
		}
		t.Logf("record: %v", r)
	}
}

func TestUserquake(t *testing.T) {
	json := `{"_id":{"$oid":"5eeaf6c802add60b3268bd4f"},"area":200,"code":561,"created_at":"2020/06/18 14:08:24.777","expire":"2020/06/18 14:09:24","hop":1,"time":"2020/06/18 14:08:24.772","uid":"9999920200618140824767","ver":"20150406"}`
	r, err := Parse(json)
	if err != nil {
		t.Errorf("record (%v) parse error: %v", json, err)
	}

	u := r.Userquake
	if u == nil {
		t.Errorf("record (%v) parse type is not userquake", json)
	}

	if u.Code != 561 || u.Area != 200 || u.CreatedAt != "2020/06/18 14:08:24.777" || u.Time.UnixNano() != 1592456904772000000 {
		t.Errorf("userquake (%v) values are wrong", u)
	}

	t.Logf("userquake: %v", u)
}

func TestAreapeers(t *testing.T) {
	json := `{"_id":{"$oid":"5ec7069302add60b2f68980f"},"areas":[{"id":10,"peer":49},{"id":15,"peer":5},{"id":25,"peer":1},{"id":30,"peer":2},{"id":35,"peer":2},{"id":40,"peer":1},{"id":45,"peer":1},{"id":50,"peer":2},{"id":55,"peer":7},{"id":60,"peer":1},{"id":65,"peer":8},{"id":70,"peer":4},{"id":75,"peer":1},{"id":100,"peer":14},{"id":105,"peer":16},{"id":110,"peer":5},{"id":111,"peer":1},{"id":115,"peer":16},{"id":120,"peer":21},{"id":125,"peer":27},{"id":130,"peer":13},{"id":135,"peer":7},{"id":140,"peer":1},{"id":141,"peer":2},{"id":142,"peer":11},{"id":143,"peer":4},{"id":150,"peer":16},{"id":151,"peer":5},{"id":152,"peer":5},{"id":200,"peer":16},{"id":205,"peer":20},{"id":210,"peer":7},{"id":215,"peer":21},{"id":220,"peer":2},{"id":225,"peer":25},{"id":230,"peer":18},{"id":231,"peer":113},{"id":232,"peer":2},{"id":240,"peer":21},{"id":241,"peer":69},{"id":242,"peer":7},{"id":250,"peer":345},{"id":270,"peer":152},{"id":275,"peer":31},{"id":300,"peer":3},{"id":301,"peer":4},{"id":302,"peer":12},{"id":305,"peer":2},{"id":310,"peer":3},{"id":315,"peer":3},{"id":325,"peer":8},{"id":330,"peer":6},{"id":340,"peer":6},{"id":345,"peer":6},{"id":350,"peer":8},{"id":351,"peer":12},{"id":355,"peer":5},{"id":405,"peer":20},{"id":410,"peer":4},{"id":411,"peer":24},{"id":415,"peer":23},{"id":416,"peer":13},{"id":420,"peer":18},{"id":425,"peer":71},{"id":430,"peer":11},{"id":435,"peer":2},{"id":440,"peer":2},{"id":445,"peer":4},{"id":455,"peer":20},{"id":460,"peer":48},{"id":465,"peer":26},{"id":470,"peer":2},{"id":475,"peer":35},{"id":480,"peer":11},{"id":490,"peer":6},{"id":495,"peer":1},{"id":500,"peer":1},{"id":505,"peer":4},{"id":510,"peer":4},{"id":515,"peer":2},{"id":520,"peer":1},{"id":525,"peer":12},{"id":530,"peer":1},{"id":535,"peer":13},{"id":541,"peer":4},{"id":550,"peer":8},{"id":560,"peer":15},{"id":570,"peer":4},{"id":575,"peer":4},{"id":576,"peer":1},{"id":580,"peer":1},{"id":581,"peer":4},{"id":600,"peer":12},{"id":601,"peer":8},{"id":602,"peer":4},{"id":605,"peer":7},{"id":610,"peer":1},{"id":615,"peer":3},{"id":625,"peer":1},{"id":641,"peer":4},{"id":646,"peer":1},{"id":650,"peer":1},{"id":651,"peer":3},{"id":665,"peer":5},{"id":670,"peer":5},{"id":675,"peer":3},{"id":701,"peer":2},{"id":710,"peer":1},{"id":900,"peer":58},{"id":901,"peer":5},{"id":905,"peer":2}],"code":555,"created_at":"2020/05/22 07:54:11.135","expire":"2020/05/22 07:56:28","hop":8,"time":"2020/05/22 07:54:11.114","uid":"2020/05/22 07:56:28","ver":"20150406"}`
	r, err := Parse(json)
	if err != nil {
		t.Errorf("record (%v) parse error: %v", json, err)
	}

	a := r.Areapeers
	if a == nil {
		t.Errorf("record (%v) parse type is not areapeers", json)
	}

	if a.Code != 555 || a.CreatedAt != "2020/05/22 07:54:11.135" || a.Time.UnixNano() != 1590101651114000000 {
		t.Errorf("areapeers (%v) values are wrong", a)
	}

	if a.Areas[0].Id != 10 || a.Areas[0].Peer != 49 {
		t.Errorf("areapeers.areas[0] (%v) values are wrong", a.Areas[0])
	}

	t.Logf("areapeers: %v", a)
}
