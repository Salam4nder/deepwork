package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	fileName           = "deepwork.txt"
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

func createInitialFile() error {
	path, err := filePath()
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("deepwork: creating initial file in %s\n", path)
			cF, err := os.Create(path)
			if err != nil {
				return err
			}
			defer func() {
				if err := cF.Close(); err != nil {
					fmt.Fprintf(os.Stdout, "%s, %s", fmt.Sprint(whereErrorOccurred+"closing file"), err.Error())
				}
			}()
			i := NewInterval()
			encB, err := i.Encode()
			if err != nil {
				return err
			}

			_, err = cF.Write(encB)
			if err != nil {
				return err
			}
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
