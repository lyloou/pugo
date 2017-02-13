package main

import (
	"os"
	"time"

	"github.com/go-xiaohei/pugo/app/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "PuGo"
	app.Usage = "a Simple Static Site Generator"
	app.UsageText = app.Usage
	app.Commands = []cli.Command{
		cmd.Init,
		cmd.Build,
		cmd.Server,
		cmd.Asset,
	}
	app.Version = "0.11.0"
	app.Compiled = time.Now()
	// app.HideHelp = true
	app.Run(os.Args)
}
