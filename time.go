package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Day struct {
	Date time.Time     `json:"date"`
	Val  time.Duration `json:"val"`
}

type Interval struct {
	Days []*Day `json:"days"`
}

func DecodeInterval(i *Interval, b []byte) error {
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(i); err != nil {
		return err
	}
	return nil
}

func (x *Interval) Encode() ([]byte, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(x); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (x *Interval) LastDay() *Day {
	if x == nil || len(x.Days) == 0 {
		return nil
	}
	return x.Days[len(x.Days)-1]
}

func NewInterval() *Interval {
	i := new(Interval)
	i.Days = append(i.Days, &Day{Date: time.Now()})
	return i
}

func (x *Interval) Print() {
	if x == nil || len(x.Days) < 1 {
		fmt.Println("deepwork: no recorded days yet, run `deepwork` to start a timer!")
	}
	for _, d := range x.Days {
		fmt.Println(
			d.Date.Year(),
			d.Date.Month(),
			d.Date.Weekday(),
			d.Val.Round(time.Second),
		)
	}
}
