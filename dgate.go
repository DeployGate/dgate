package main

import (
	"github.com/codegangsta/cli"
	"os"
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
			Name:    "push",
			Aliases: []string{"p"},
			Usage:   "push apps to deploygate",
			Flags:   []cli.Flag{paramOwner, paramMessage},
			Action: func(c *cli.Context) {
				filePath := c.Args().First()
				owner := c.String("owner")
				message := c.String("message")

				result, App := Upload(filePath, owner, message)
				if result {
					println("Push app file successful!")
					println("Name :    ", App.name)
					println("Owner :   ", App.owner)
					println("Package : ", App.packageName)
					println("Revision :", App.revision)
					println("URL :     ", App.url)
				}
			},
		},
		{
			Name:  "login",
			Usage: "login to deploygate",
			Flags: []cli.Flag{paramEmail, paramPassword},
			Action: func(c *cli.Context) {
				email := c.String("email")
				password := c.String("password")

				result := Login(email, password)
				if result {
					welcomeMessage := `Welcome to DeployGate!
Let's upload the app to DeployGate!`
					println(welcomeMessage)
				}
			},
		},
		{
			Name:  "logout",
			Usage: "logout to deploygate",
			Action: func(c *cli.Context) {
				Logout()
			},
		},
	}

	app.Run(os.Args)
}
