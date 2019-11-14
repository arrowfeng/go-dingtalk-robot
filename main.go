/*
 * File : main.go.go
 * Author : arrowfeng
 * Date : 2019/11/13
 */
package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)


const url = "https://oapi.dingtalk.com/robot/send"
var access_token  string

func main() {
	app := cli.NewApp()
	app.Name = "dingtalk tools"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "arrowfeng",
			Email: "husterzdf@gmail.com",
		},
	}

	app.Copyright = "(c) 2019 Halo"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token",
			Usage:       "global test",
			EnvVar:      "DINGTALK_TOKEN",
			Required:    true,
			Destination: &access_token,
		},

	}

	app.Commands = []cli.Command {
		sendCommand,
		textCommand,
		linkCommand,
		markDownCommand,
		actionCardCommand,
		feedCardCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
