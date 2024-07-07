package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer stop()

	var (
		dFlag = flag.Bool("d", false, "Print the day overview")
		wFlag = flag.Bool("w", false, "Print the week overview")
		mFlag = flag.Bool("m", false, "Print the month overview")
	)
	flag.Parse()

	switch {
	case *dFlag:
	case *wFlag:
	case *mFlag:
	default:
		startTimer(ctx)
	}
}

func startTimer(ctx context.Context) error {
	createInitialFile(true)
	fmt.Println("deepwork: timer starting, good luck!")
	start := time.Now()

	select {
	case <-ctx.Done():
		file, err := openFile()
		if err != nil {
			return err
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println("start timer: closing file")
			}
		}()

		b, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("start timer: reading file, %w", err)
		}

		var year *Year
		if err := json.Unmarshal(b, year); err != nil {
			return fmt.Errorf("start timer: unmarshaling, %w", err)
		}

		var w *Week
		if len(year.Months) >= 12 {
			// TODO(kg): One full year has passed, decide what to do.
			// For now only support one year of tracking.
			oneYearPlaceHolder()
		}
		m := year.LastMonth()
		if m == nil {
			year.Months = append(year.Months, &Month{})
			year.Months[0].Weeks = []*Week{}
			year.Months[0].Weeks = append(year.Months[0].Weeks, w)
		} else {
			// Is it time for a new month?
			if len(m.Weeks) >= 4 {
			}
		}

		elapsed := time.Now().Sub(start)
		rounded := elapsed.Round(elapsed)
		switch time.Now().Weekday() {
		case time.Monday:
			if interval.W.Monday != 0 {
				interval.W.Monday += rounded
				break
			}
			w.Monday = rounded
		case time.Tuesday:
			w.Tuesday = rounded
		case time.Wednesday:
			w.Wednesday = rounded
		case time.Thursday:
			w.Thursday = rounded
		case time.Friday:
			w.Friday = rounded
		case time.Saturday:
			w.Staruday = rounded
		case time.Sunday:
			w.Sunday = rounded
		default:
			return errors.New("invalid weekday")
		}
		fmt.Println(elapsed.Round(elapsed))
		return nil
	}
}

func oneYearPlaceHolder() {
	panic("not implemented")
}
