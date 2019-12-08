package main

import (
	"os"
	"sort"

	"github.com/cikupin/feature-flag-example/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "feature-flag",
		Usage: "feature flag example using flagr",
		Commands: []*cli.Command{
			cmd.GenerateFeatureToggle,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
