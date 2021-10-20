package app

import (
	"os"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/lines"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func Start() {

	app := cli.NewApp()

	_ = setUpApp(app)

	_ = app.Run(os.Args)

}

func setUpApp(app *cli.App) error {
	app.Name = "SCADA db loader"
	app.Usage = "This app is intended for loading slave db data to master db for SCADA"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		{
			Name:    "bdashboard",
			Aliases: []string{"bd"},
			Usage:   "run `bdashboard` for all 28 branches, get data from respective dbs and store to local",
			Action: func(c *cli.Context) error {
				color.HiGreen("* Batch execution `started` for BDashboard")
				lines.RunConcurAllBDashboardBranches(28)
				return nil
			},
		},
		{
			Name:    "qweekly",
			Aliases: []string{"qw"},
			Usage:   "run `qweekly` for all 28 branches, get data from respective dbs and store to local",
			Action: func(c *cli.Context) error {
				color.HiGreen("* Batch execution `started` for QWeekly")
				lines.RunConcurAllQWeeklyBranches(28)
				return nil
			},
		},
		{
			Name:    "qsearch",
			Aliases: []string{"qs"},
			Usage:   "run `qsearch` for all 28 branches, get data from respective dbs and store to local",
			Action: func(c *cli.Context) error {
				color.HiGreen("* Batch execution not concurrent `started` for QSearch")
				lines.RunAllQSearchBranches(28)
				return nil
			},
		},
		{
			Name:    "setup",
			Aliases: []string{"s"},
			Usage:   "run `setup` to create all necessary tables on local (Master) DB.",
			Action: func(c *cli.Context) error {
				color.HiGreen("* setting up or creating necessary tables on local (Master) DB.")
				if err := local.SetUpTables(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "qdashboard",
			Aliases: []string{"qd"},
			Usage:   "run `qdashboard` for all 28 branches, get data from respective dbs and store to local",
			Action: func(c *cli.Context) error {
				color.HiGreen("* Batch execution `started` for QDashboard")
				lines.RunConcurAllQDashboardBranches(28)
				return nil
			},
		},
		{
			Name:    "qreport",
			Aliases: []string{"qr"},
			Usage:   "run `qreport` for all 28 branches, get data from respective dbs and store to local",
			Action: func(c *cli.Context) error {
				color.HiGreen("* Batch execution `started` for QReport")
				lines.RunConcurAllQReportBranches(28)
				return nil
			},
		},
		{
			Name:    "all",
			Aliases: []string{"all"},
			Usage:   "run all tasks, get data from respective dbs and store to local",
			Action: func(c *cli.Context) error {
				color.HiGreen("* Batch execution `started` for All tasks")
				RunAll()
				color.HiGreen("* Batch execution `completed` for All tasks")
				return nil
			},
		},
		{
			Name:    "start",
			Aliases: []string{"str"},
			Usage:   "run all tasks, get data from respective dbs and store to local every 24 hours",
			Action: func(c *cli.Context) error {
				for {
					color.HiGreen("* Batch execution `started` for All tasks for every 24 hours")
					RunAll()
					color.HiGreen("* Batch execution `completed` for All tasks for every 24 hours")
					time.Sleep(time.Hour * 24)
				}
				return nil
			},
		},
	}

	return nil
}

func RunAll() {
	lines.RunConcurAllBDashboardBranches(28)
	lines.RunConcurAllQWeeklyBranches(28)
	lines.RunAllQSearchBranches(28)
	lines.RunConcurAllQDashboardBranches(28)
	lines.RunConcurAllQReportBranches(28)
}