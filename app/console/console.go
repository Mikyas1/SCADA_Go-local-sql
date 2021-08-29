package console

//import (
//	"fmt"
//	"github.com/urfave/cli"
//)
//
//app := cli.NewApp()
//app.Name = "SCADA db loader"
//app.Usage = "This app is intended for loading slave db to master db for SCADA"
//app.Version = "0.1"
//
//var language string
//
//app.Flags = []cli.Flag {
//cli.StringFlag{
//Name: "lang, l",
//Value: "english",
//Usage: "language for the greeting",
//Destination: &language,
//},
//}
//
//cli.VersionFlag = cli.BoolFlag{
//Name: "print-version, V",
//Usage: "print only the version",
//}
//
//app.Action = func(c *cli.Context) error {
//	name := "Nefertiti"
//	if c.NArg() > 0 {
//		name = c.Args()[0]
//	}
//	if language == "spanish" {
//		fmt.Println("Hola", name)
//	} else {
//		fmt.Println("Hello", name)
//	}
//	return nil
//}
//
//app.Commands = []cli.Command{
//{
//Name:    "complete",
//Aliases: []string{"c"},
//Usage:   "complete a task on the list",
//Action: func(c *cli.Context) error {
//color.Red("Prints text in cyan.")
//return nil
//},
//},
//}
//
//app.Run(os.Args)
