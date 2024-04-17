package main

import (
	"github.com/ancalabrese/mc-cli/mc/config"
)

var (
	deviceId string
)

func main() {
	c := config.NewConfig()
	c.Authentication.Load()
	err := c.Authentication.Write()
	if err != nil {
		panic(err)
	}
	println(c.Authentication)
}
