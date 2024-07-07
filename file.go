package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	fileName           = "deepwork.json"
	whereErrorOccurred = "creating initial file: "
)

func filePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("filepath: getting home dir")
	}
	return fmt.Sprintf("%s/%s", homeDir, fileName), nil
}

// openFile will open the config file and return it.
// Caller is responsible to close it.
func openFile() (*os.File, error) {
	path, err := filePath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	return f, nil
}

func createInitialFile(off bool) error {
	if off {
		fmt.Println("create file off, doing nothing")
		return nil
	}

	path, err := filePath()
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stdout, "creating initial file in %s", filePath)
			cF, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("%s, %w", fmt.Sprint(whereErrorOccurred+"creating file"), err)
			}
			defer func() {
				if err := cF.Close(); err != nil {
					fmt.Fprintf(os.Stdout, "%s, %s", fmt.Sprint(whereErrorOccurred+"closing file"), err.Error())
				}
			}()
		}
		return nil
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(whereErrorOccurred + "closing file")
		}
	}()

	return nil
}
