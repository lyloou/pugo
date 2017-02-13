package cmd

import (
	"github.com/Unknwon/com"
	"github.com/go-xiaohei/pugo/app/helper/printer"
	"github.com/urfave/cli"
)

var commonFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "debug",
		Usage: "print all debug info",
	},
}

func isSiteAvailable() bool {
	printer.Logf("check meta.toml")
	if !com.IsFile("meta.toml") {
		printer.Error("Error: meta.toml is not found.")
		printer.Print("You need create a new site here with 'pugo new-site'")
		return false
	}
	return true
}
