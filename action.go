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
	filePath, userName := "", ""
	if len(context.Args()) >= 2 {
		userName = context.Args().Get(0)
		filePath = context.Args().Get(1)
	} else {
		filePath = context.Args().First()
	}
	message := context.String("message")
	isOpen := context.Bool("open")
	isDisableNotify := context.Bool("disable-notify")
	isPublic := context.Bool("public")

	result, App := Upload(filePath, userName, message, isDisableNotify, isPublic)
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
