package main

import (
	"fmt"
	"os"
)

const (
	fileName           = "deepwork.json"
	whereErrorOccurred = "creating initial file: "
)

func createInitialFile(off bool) error {
	if off {
		fmt.Println("create file off, doing nothing")
		return nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(whereErrorOccurred + "getting user home dir")
	}
	filePath := fmt.Sprintf("%s/%s", homeDir, fileName)

	f, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stdout, "creating initial file in %s", filePath)
			cF, err := os.Create(filePath)
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
