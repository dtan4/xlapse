package main

import (
	"fmt"
	"os"
)

func main() {
	if err := realMain(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func realMain(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("object path is missing")
	}
	path := args[1]

	fmt.Println(path)

	return nil
}
