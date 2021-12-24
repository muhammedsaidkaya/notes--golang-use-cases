package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func processFile(path string) error {
	if path == "" {
		return errors.New("path is empty")
	}
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "unable to open a file")
	}
	if f.Name() != "expectedName" {
		return fmt.Errorf("unexpected filename: %s", f.Name())
	}
	return nil
}

func main() {
	err := processFile("./foo")
	if err != nil {
		switch errors.Cause(err).(type) {
		case *os.PathError:
			fmt.Printf("file not found: %v", err)
		default:
			fmt.Printf("unknown error: %v", err)
		}
		return
	}
}
