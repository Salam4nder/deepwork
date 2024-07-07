package main

import "time"

type Week struct {
	Monday    time.Duration `json:"monday"`
	Tuesday   time.Duration `json:"tuesday"`
	Wednesday time.Duration `json:"wednesday"`
	Thursday  time.Duration `json:"thursday"`
	Friday    time.Duration `json:"friday"`
	Staruday  time.Duration `json:"saturday"`
	Sunday    time.Duration `json:"sunday"`
}

func (x *Week) OverwriteIfExist(w time.Weekday, d time.Duration) {
	switch w {
	case time.Monday:
		if x.Monday != 0 {
			x.Monday += d
			return
		}
		x.Monday = d
	case time.Tuesday:
		if x.Tuesday != 0 {
			x.Tuesday += d
			return
		}
		x.Tuesday = d
	case time.Wednesday:
		if x.Wednesday != 0 {
			x.Wednesday += d
			return
		}
		x.Wednesday = d
	case time.Thursday:
		if x.Thursday != 0 {
			x.Thursday += d
			return
		}
		x.Thursday = d
	case time.Friday:
		if x.Friday != 0 {
			x.Friday += d
			return
		}
		x.Friday = d
	case time.Saturday:
		if x.Staruday != 0 {
			x.Staruday += d
			return
		}
		x.Staruday = d
	case time.Sunday:
		if x.Staruday != 0 {
			x.Staruday += d
			return
		}
		x.Sunday = d
	default:
		return
	}
}

type Month struct {
	Weeks []*Week `json:"weeks"`
}

func (x *Month) LastWeek() *Week {
	if len(x.Weeks) == 0 {
		return nil
	}
	return x.Weeks[len(x.Weeks)-1]
}

type Year struct {
	Months []*Month `json:"months"`
}

func (x *Year) LastMonth() *Month {
	if len(x.Months) == 0 {
		return nil
	}
	return x.Months[len(x.Months)-1]
}
