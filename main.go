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
		pFlag = flag.Bool("p", false, "Print the overview")
	)
	flag.Parse()

	switch {
	case *pFlag:
		file, err := openFile()
		if err != nil {
			fmt.Println("deepwork: failed to open file, you may need to create it by running `deepwork`")
			os.Exit(1)
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println("deepwork: failed closing file, " + err.Error())
			}
		}()

		b, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("deepwork: failed reading file, " + err.Error())
			os.Exit(1)
		}

		var i Interval
		if err = DecodeInterval(&i, b); err != nil {
			fmt.Println("deepwork: failed decoding file to Interval " + err.Error())
			os.Exit(1)
		}
		i.Print()
	default:
		if err := startTimer(ctx); err != nil {
			fmt.Println("deepwork: " + err.Error())
		}
	}
}

func startTimer(ctx context.Context) error {
	if err := createInitialFile(); err != nil {
		return fmt.Errorf("start timer: creating file, %w", err)
	}
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

		var i Interval
		if err = DecodeInterval(&i, b); err != nil {
			return fmt.Errorf("start timer: decoding interval, %w", err)
		}
		lastDay := i.LastDay()

		var isCurrentDay bool
		// Is it a new day?
		tY, tM, tD := time.Now().Date()
		sY, sM, sD := start.Date()
		y, m, d := lastDay.Date.Date()
		if y == tY && m == tM && d == tD {
			// It's not.
			isCurrentDay = true
		} else if y == sY && m == sM && d == sD {
			// It is, but we started the timer on the day before.
			// Let's count it as such.
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
		path, err := filePath()
		if err != nil {
			return fmt.Errorf("start timer: getting path, %w", err)
		}
		if err := os.WriteFile(path, encB, os.ModePerm); err != nil {
			return fmt.Errorf("start timer: writing file, %w", err)
		}

		return nil
	}
}
