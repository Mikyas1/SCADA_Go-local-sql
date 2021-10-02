package app

import (
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/lines"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"os"
)

func Start() {

	app := cli.NewApp()

	_ = setUpApp(app)

	_ = app.Run(os.Args)

	//lines.RunConcurAllQWeeklyBranches(28)
	//go lines.RunQWeeklyLine(3)
	// go lines.RunBDashboardLine(3)
	// go lines.RunQSearchLine(1)

	//fmt.Scanln()

	// if err := local.SetUpTables(); err != nil {
	// 	panic(err)
	// }

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
	}

	return nil
}
