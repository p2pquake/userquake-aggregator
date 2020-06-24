package epsp

import (
	"encoding/json"
	"time"
)

const earthquake = 551
const tsunami = 552
const areapeers = 555
const userquake = 561

type AreaCode int

type EPSPTime struct {
	*time.Time
}

type Record struct {
	Code      int
	CreatedAt string `json:"created_at"`
	Time      EPSPTime
	Userquake *Userquake
	Areapeers *Areapeers
}

type Userquake struct {
	Code      int
	CreatedAt string `json:"created_at"`
	Time      EPSPTime
	Area      AreaCode
}

type Areapeers struct {
	Code      int
	CreatedAt string `json:"created_at"`
	Time      EPSPTime
	Areas     []Areapeer
}

type Areapeer struct {
	Id   AreaCode
	Peer int
}

func (et *EPSPTime) UnmarshalJSON(data []byte) error {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t, err := time.ParseInLocation("\"2006/01/02 15:04:05.999\"", string(data), loc)
	*et = EPSPTime{&t}
	return err
}

func Parse(body string) (Record, error) {
	r := Record{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return r, err
	}

	if r.Code == userquake {
		u := Userquake{}
		err = json.Unmarshal([]byte(body), &u)

		if err == nil {
			r.Userquake = &u
		}
	}

	if r.Code == areapeers {
		a := Areapeers{}
		err = json.Unmarshal([]byte(body), &a)

		if err == nil {
			r.Areapeers = &a
		}
	}

	return r, nil
}
