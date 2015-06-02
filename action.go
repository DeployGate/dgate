package main

import (
	"github.com/codegangsta/cli"
	"github.com/skratchdot/open-golang/open"
)

func LoginAction(context *cli.Context) {
	email := context.String("email")
	password := context.String("password")

	result := Login(email, password)
	if result {
		welcomeMessage := `Welcome to DeployGate!
Let's upload the app to DeployGate!`
		println(welcomeMessage)
	}
}

func LogoutAction(context *cli.Context) {
	Logout()
}

func PushAction(context *cli.Context) {
	filePath := context.Args().First()
	userName := context.String("user")
	message := context.String("message")
	isOpen := context.Bool("open")

	result, App := Upload(filePath, userName, message)
	if result {
		println("Push app file successful!")
		println("Name :    ", App.name)
		println("Owner :   ", App.owner)
		println("Package : ", App.packageName)
		println("Revision :", App.revision)
		println("URL :     ", App.url)
	}

	if isOpen {
		open.Run(App.url)
	}
}
