package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"

	"github.com/labstack/echo"
	"github.com/sunmi-OS/gocore/api"
	"github.com/urfave/cli"
)

type EchoApi struct {
}

var eApi EchoApi

func (a *EchoApi) echoStart(c *cli.Context) error {
	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Route => handler
	e.Any("/", func(c echo.Context) error {

		request := api.NewRequest(c)
		response := api.NewResponse(c)

		err := request.InitRawJson()
		if err != nil {
			return response.RetError(err, 400)
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		return response.RetSuccess("sccess")
	})

	e.Any("/gorm2", func(c echo.Context) error {

		request := api.NewRequest(c)
		response := api.NewResponse(c)

		err := request.InitRawJson()
		if err != nil {
			return response.RetError(err, 400)
		}

		return response.RetSuccess("sccess")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

func main() {

	runtime.GOMAXPROCS(1)

	app := cli.NewApp()
	app.Name = "IOT-seanbox"
	app.Usage = "IOT-seanbox"
	app.Email = "wenzhenxi@sunmi.com"
	app.Version = "1.0.0"
	app.Usage = "IOT-seanbox"
	app.Action = eApi.echoStart

	// 指定对于的命令
	app.Commands = []cli.Command{
		{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "api",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开启API-DEMO",
					Action: eApi.echoStart,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)

}
