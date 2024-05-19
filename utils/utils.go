package utils

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		os.Exit(1)
	}
}

func FlushStdin() {
	var discard string
	for {
		n, _ := fmt.Scanln(&discard)
		if n == 0 {
			break
		}
	}
}
