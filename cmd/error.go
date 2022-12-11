package cmd

import (
	"fmt"
	"os"
)

func errExit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting: %s", err.Error())
		os.Exit(1)
	}
}
