package epsp

import "encoding/json"

const earthquake = 551
const tsunami = 552
const areapeers = 555
const userquake = 561

type Record struct {
	Code      int
	CreatedAt string `json:"created_at"`
	Time      string
}

type Userquake struct {
	Code      int
	CreatedAt string `json:"created_at"`
	Time      string
	Area      int
}

type Areapeers struct {
	Code      int
	CreatedAt string `json:"created_at"`
	Time      string
	Areas     []Areapeer
}

type Areapeer struct {
	Id   int
	Peer int
}

func ParseRecord(body string) (Record, error) {
	r := Record{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func Parse(body string) (interface{}, error) {
	r, err := ParseRecord(body)
	if err != nil {
		return r, err
	}

	if r.Code == userquake {
		u := Userquake{}
		err = json.Unmarshal([]byte(body), &u)
		return u, err
	}

	if r.Code == areapeers {
		a := Areapeers{}
		err = json.Unmarshal([]byte(body), &a)
		return a, err
	}

	return r, nil
}
