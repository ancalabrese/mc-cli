package main

import (
	"os"
	"os/signal"

	"github.com/ancalabrese/mc-cli/mc/auth"
	"github.com/ancalabrese/mc-cli/mc/config"
)

func main() {
	c := config.NewConfig()
	c.Authentication.Load()
	auth.RequestAuthCode(c)
	err := c.Authentication.Write()
	if err != nil {
		panic(err)
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	_ = <-sigChan
}
