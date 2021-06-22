package main

import (
	"log"
	"os"

	"github.com/dictybase-playground/analytics-backup/internal/app"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "analytics-backup"
	app.Usage = "cli for backing up google analytics data"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-format",
			Usage: "format of the logging out, either of json or text",
			Value: "json",
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "log level for the application",
			Value: "error",
		},
	}
	app.Commands = []cli.Command{RegCmd()}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error in running command %s", err)
	}
}

func requiredFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:     "credentials",
			Usage:    "filename for service account credentials json",
			Value:    "key.json",
			Required: true,
		},
		cli.StringFlag{
			Name:     "view-id",
			Usage:    "google analytics view id",
			Required: true,
		},
		cli.StringFlag{
			Name:     "start-date",
			Usage:    "start date (YYYY-mm-dd)",
			Required: true,
		},
		cli.StringFlag{
			Name:     "end-date",
			Usage:    "end date (YYYY-mm-dd)",
			Required: true,
		},
		cli.StringFlag{
			Name:     "metrics",
			Usage:    "metrics to include (separated by comma)",
			Value:    "ga:sessions,ga:users",
			Required: true,
		},
		cli.StringFlag{
			Name:     "dimensions",
			Usage:    "dimensions to include (separated by comma)",
			Value:    "ga:date",
			Required: true,
		},
	}
}

func RegCmd() cli.Command {
	return cli.Command{
		Name:   "reports",
		Usage:  "gets google analytics report and converts to csv",
		Action: app.GetAnalyticReports,
		Flags:  requiredFlags(),
	}
}
