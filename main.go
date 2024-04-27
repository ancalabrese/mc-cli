package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ancalabrese/mc-cli/mc/auth"
	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/ancalabrese/mc-cli/utils"
	"github.com/hashicorp/go-hclog"
)

func main() {
	ctx := context.Background()

	loggerOptions := &hclog.LoggerOptions{
		Name:  "mc-cli",
		Level: hclog.Debug,
	}
	log := hclog.New(loggerOptions)

	c := config.NewConfig()
	c.Authentication.Load()

	err := auth.Login(ctx, c, log)
	utils.Check(err)

	err = c.Authentication.Write()
	utils.Check(err)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	_ = <-sigChan
}
