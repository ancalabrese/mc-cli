package utils

import "os"

func Check(err error) {
	if err != nil {
		println("ERROR: %w", err)
		os.Exit(1)
	}
}
