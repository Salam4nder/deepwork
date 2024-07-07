package main

import (
	"context"
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
		pFlag = flag.Bool("p", false, "Print the day overview")
	)
	flag.Parse()

	switch {
	case *pFlag:
	default:
		startTimer(ctx)
	}
}

func startTimer(ctx context.Context) error {
	createInitialFile()
	fmt.Println("deepwork: timer starting, good luck!")
	start := time.Now()

	select {
	case <-ctx.Done():
		var err error
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

		var i *Interval
		i, err = DecodeInterval(b)
		if err != nil {
			return fmt.Errorf("start timer: decoding interval, %w", err)
		}

		i = i.NewIfEmpty()
		lastDay := i.LastDay()

		var isCurrentDay bool
		// Is it a new day?
		tY, tM, tD := time.Now().Date()
		t, m, d := lastDay.Date.Date()
		if t == tY && m == tM && d == tD {
			// It's not.
			isCurrentDay = true
		} else {
			// It is.
			lastDay = new(Day)
			lastDay.Date = time.Now()
			i.Days = append(i.Days, lastDay)
		}

		elapsed := time.Now().Sub(start)

		if isCurrentDay {
			lastDay.Val += elapsed
		} else {
			lastDay.Val = elapsed
		}
		encB, err := i.Encode()
		if err != nil {
			return fmt.Errorf("start timer: encoding interval, %w", err)
		}

		if err = file.Truncate(0); err != nil {
			return fmt.Errorf("start timer: truncating file, %w", err)
		}
		_, err = file.Write(encB)
		if err != nil {
			return fmt.Errorf("start timer: writing file, %w", err)
		}

		return nil
	}
}
