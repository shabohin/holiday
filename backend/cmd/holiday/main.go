package main

import (
	"github.com/shabohin/holiday.git/internal/containers"
	"github.com/urfave/cli/v2"
	"os"
)

const version = "0.1.0"

const configPath = "configs/config.toml"

func main() {
	app := &cli.App{
		Name:    "holiday",
		Usage:   "service",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from `FILE`",
				EnvVars:     []string{"HOLIDAY_CONFIG_PATH"},
				TakesFile:   true,
				Value:       configPath,
				DefaultText: configPath,
				HasBeenSet:  false,
			},
		},
		Action: runApp,
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

// runApp - run app
func runApp(context *cli.Context) error {
	app := containers.NewHoliday(configPath)
	err := app.Start(context.Context)
	if err != nil {
		return err
	}
	return nil
}
