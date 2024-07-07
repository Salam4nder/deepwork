package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Day struct {
	Date time.Time
	Val  time.Duration
}

type Interval struct {
	Days []*Day
}

func DecodeInterval(b []byte) (*Interval, error) {
	var i *Interval
	if err := gob.NewDecoder(bytes.NewReader(b)).Decode(i); err != nil {
		return nil, err
	}
	return i, nil
}

func (x *Interval) Encode() ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(x); err != nil {
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

func (x *Interval) NewIfEmpty() *Interval {
	if x == nil || len(x.Days) == 0 {
		x = new(Interval)
		x.Days = append(x.Days, &Day{})
		return x
	}
	return x
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
			d.Val.Round(time.Hour),
			" in hours",
		)
	}
}
