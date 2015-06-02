package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	initHelp()

	app := cli.NewApp()
	app.Name = "dgate"
	app.Usage = "A command-line interface for DeployGate"
	app.Version = Version()

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
	paramOpen := cli.BoolFlag{
		Name:  "open, o",
		Usage: "open with browser (Mac OS only)",
	}
	paramDisableNotify := cli.BoolFlag{
		Name:  "disable-notify",
		Usage: "disable notify via email (iOS app only)",
	}
	paramPublic := cli.BoolFlag{
		Name:  "public",
		Usage: "set public visibility(new app upload only)",
	}

	app.Commands = []cli.Command{
		{
			Name:    "push",
			Aliases: []string{"p"},
			Usage:   "push apps to deploygate",
			Flags:   []cli.Flag{paramMessage, paramOpen, paramDisableNotify, paramPublic},
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

func initHelp() {
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
   {{.Version}}{{if len .Authors}}

AUTHOR(S):
   {{range .Authors}}{{ . }}{{end}}{{end}}

COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
EXAMPLE:
    login to deploygate
    $ dgate login -e <your email> -p <your password>

    Push/Update app to your own
    $ dgate push <app_file_path>

    Push/Update app to inviter who invited your as developer
    $ dgate push <owner_name> <app_file_path>

    Push/Update app to group with message and open it in browser after push
    $ dgate push <group_name> <app_file_path> -m 'develop build' -o

    Change account or logout
    $ dgate logout
`
}
