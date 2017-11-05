package main

import (
	"os"
	"time"

	"github.com/lyloou/pugo/app/command"
	"github.com/lyloou/pugo/app/vars"
	"github.com/urfave/cli"
)

//go:generate go-bindata -o=app/asset/asset.go -pkg=asset source/... doc/source/... doc/theme/...
//go:generate gofmt -w -s .

var (
	commit string
)

func main() {
	vars.Commit = commit
	app := cli.NewApp()
	app.Name = vars.Name
	app.Usage = vars.Desc
	app.Version = vars.Version
	app.Compiled = time.Now()
	app.Commands = []cli.Command{
		command.Build,
		command.Server,
		command.New,
		command.Doc,
		command.Deploy,
		command.Version,
	}
	app.HideVersion = true
	app.Run(os.Args)
}
