package command

import (
	"github.com/lyloou/pugo/app/builder"
	"github.com/lyloou/pugo/app/server"
	"github.com/lyloou/pugo/app/sync"
	"github.com/urfave/cli"
)

var (
	// Doc is command of 'doc'
	Doc = cli.Command{
		Name:  "doc",
		Usage: "run documentation server",
		Flags: []cli.Flag{
			addrFlag,
			debugFlag,
			noServerDocFlag,
		},
		Before: Before,
		Action: docServ,
	}
)

func docServ(c *cli.Context) error {
	if !c.Bool("no-server") {
		builder.After(func(ctx *builder.Context) {
			if s == nil {
				s = server.New(ctx.DstDir())
				go s.Run(c.String("addr"))
			}
			if ctx.Source != nil && ctx.Source.Meta != nil {
				s.SetPrefix(ctx.Source.Meta.Path)
			}
		})
	}
	buildContext := newContext(c, false)
	buildContext.From = "doc/source"
	buildContext.To = "doc/dest"
	buildContext.ThemeName = "doc/theme"
	buildContext.Tree.Dest = buildContext.DstDir()
	buildContext.Sync = sync.NewSyncer(buildContext.DstDir())
	build(buildContext, !c.Bool("no-server"))
	return nil
}
