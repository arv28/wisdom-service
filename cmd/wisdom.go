package main

import (
	"fmt"
	"os"

	"github.com/arv28/wisdom-service/lib/wisdom"

	//"github.eagleview.com/engineering/pa-gosdk/errors"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "wisdom"
	app.Usage = "dispense some programming wisdom fortune cookies"
	app.UsageText = "wisdom dispense"
	app.Version = wisdom.Version

	app.Commands = []cli.Command{
		cli.Command{
			Name:      "dispense",
			Usage:     "dispense a single programming wisdom fortune cookie",
			UsageText: "wisdom dispense",
			Action:    DispenseWisdom,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		//errors.ReportError(err)
		//errors.FlushReports()
		os.Exit(1)
	}

}

func DispenseWisdom(c *cli.Context) {
	fmt.Println("TODO: dispense some wisdom")
	//return nil
}
