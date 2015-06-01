package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dgate"
	app.Usage = "A command-line interface for DeployGate"

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
			Name:      "push",
			Aliases:     []string{"p"},
			Usage:     "push apps to deploygate",
			Flags:       []cli.Flag{paramOwner, paramMessage},
			Action: func(c *cli.Context) {
				owner   := c.String("owner")
				message := c.String("message")

				println("push file path:", c.Args().First())
				if len(owner) > 0 {
					println("owner:", owner)
				}
				if len(message) > 0 {
					println("message:", message)
				}
			},
		},
		{
			Name:      "login",
			Usage:     "login to deploygate",
			Flags:       []cli.Flag{paramEmail, paramPassword},
			Action: func(c *cli.Context) {
				email := c.String("email")
				password := c.String("password")
				println("login", "email:", email, "password:", password)
			},
		},
		{
			Name:      "logout",
			Usage:     "logout to deploygate",
			Action: func(c *cli.Context) {
				println("logout")
			},
		},
	}

	app.Run(os.Args)
}
