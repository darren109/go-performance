package main

import (
	"os"
	"runtime"
	"sort"

	"github.com/labstack/echo"
	"github.com/sunmi-OS/go-performance/orm/model"
	"github.com/sunmi-OS/gocore/api"
	"github.com/sunmi-OS/gocore/gorm"
	"github.com/sunmi-OS/gocore/viper"
	"github.com/urfave/cli"
)

type EchoApi struct {
}

var eApi EchoApi

type Test struct {
	Test string `json:"test"`
}

func (a *EchoApi) echoStart(c *cli.Context) error {
	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Route => handler
	e.Any("/", func(c echo.Context) error {
		response := api.NewResponse(c)

		return response.RetSuccess(model.MachineHarder.GetList())
	})

	e.Any("/gorm2", func(c echo.Context) error {
		response := api.NewResponse(c)

		return response.RetSuccess(model.MachineHarder.GetList2())
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

	viper.NewConfig("config", "conf")
	gorm.NewDB("dbDefault")

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
