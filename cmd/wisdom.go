package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/arv28/wisdom-service/lib/api"
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
		cli.Command{
			Name:      "serve",
			Usage:     "run API server to dispense programming wisdom",
			UsageText: "wisdom serve -e ENV -p Port --host HOST --api-path API_PATH",
			Action:    RunServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "env,e",
					Value:  api.DefaultEnv,
					Usage:  "(dev | test | stage | prod)",
					EnvVar: "ENV",
				},
				cli.StringFlag{
					Name:   "host",
					Value:  api.DefaultHost,
					Usage:  "host to listen on",
					EnvVar: "HOST",
				},
				cli.IntFlag{
					Name:   "port,p",
					Value:  api.DefaultPort,
					Usage:  "port to listen on",
					EnvVar: "PORT",
				},
				cli.StringFlag{
					Name:   "api-path",
					Value:  api.DefaultAPIPath,
					Usage:  "uri path prefix for mounting api router",
					EnvVar: "API_PATH",
				},
			},
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

	rand.Seed(time.Now().UnixNano())
	d, err := wisdom.FromFile("quotes.json")
	if err != nil {
		fmt.Println("wisdom.FromFile failed")
		return
	}
	fmt.Println(d.Random())
	//q := wisdom.NewQuote("Go is growing", "Jim gordon")
	//fmt.Println(q)
	//return nil
}

func RunServer() error {
	fmt.Println("TODO: run the server")
	return nil
}
