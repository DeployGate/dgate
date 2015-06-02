package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "dgate"
	app.Usage = "A command-line interface for DeployGate"
	app.Version = Version()

	paramOwner := cli.StringFlag{
		Name:  "owner, o",
		Value: "",
		Usage: "app update owner",
	}
	paramMessage := cli.StringFlag{
		Name:  "message, m",
		Value: "",
		Usage: "app update message",
	}
	paramEmail := cli.StringFlag{
		Name:  "email, e",
		Value: "",
		Usage: "login email to DeployGate",
	}
	paramPassword := cli.StringFlag{
		Name:  "password, p",
		Value: "",
		Usage: "login password to DeployGate",
	}

	app.Commands = []cli.Command{
		{
			Name:    "push",
			Aliases: []string{"p"},
			Usage:   "push apps to deploygate",
			Flags:   []cli.Flag{paramOwner, paramMessage},
			Action: func(c *cli.Context) {
				PushAction(c)
			},
		},
		{
			Name:  "login",
			Usage: "login to deploygate",
			Flags: []cli.Flag{paramEmail, paramPassword},
			Action: func(c *cli.Context) {
				LoginAction(c)
			},
		},
		{
			Name:  "logout",
			Usage: "logout to deploygate",
			Action: func(c *cli.Context) {
				LogoutAction(c)
			},
		},
	}

	app.Run(os.Args)
}
