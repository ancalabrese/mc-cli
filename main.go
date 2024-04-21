package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ancalabrese/mc-cli/mc/auth"
	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/ancalabrese/mc-cli/utils"
)

func main() {
	ctx := context.Background()

	c := config.NewConfig()
	c.Authentication.Load()

	err := auth.Login(ctx, c)
	utils.Check(err)

	err = c.Authentication.Write()
	utils.Check(err)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	_ = <-sigChan
}
