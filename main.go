package main

import "flag"

var (
	deviceId string
)

func main() {
	flag.StringVar(&deviceId, "d", "", "")
}
