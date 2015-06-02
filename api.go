package main

import (
	"bufio"
	"encoding/json"
	"github.com/ddliu/go-httpclient"
	"github.com/howeyc/gopass"
	"github.com/jmoiron/jsonq"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const baseApiEndPoint = "https://deploygate.com"

type App struct {
	name        string
	owner       string
	packageName string
	revision    int
	url         string
}

func Upload(filePath string, userName string, message string, isDisableNotify bool) (bool, App) {
	if !isLogin() {
		result := Login("", "")
		if !result {
			return false, App{}
		}
	}

	// progress
	print("upload start: ")
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			print(">>")
		}
	}()
	defer println("")
	defer ticker.Stop()

	disableNotify := "no"
	if isDisableNotify {
		disableNotify = "yes"
	}
	uploadData := map[string]string{
		"@file":   filePath,
		"message": message,
		"disable_notify": disableNotify,
	}

	name, _ := getSessions()
	if userName != "" {
		name = userName
	}

	json := postRequest("/api/users/"+name+"/apps", uploadData)
	error, message := checkError(json)
	if error {
		println(message)
		return false, App{}
	}

	appName, _ := json.String("results", "name")
	packageName, _ := json.String("results", "package_name")
	path, _ := json.String("results", "path")
	ownerName, _ := json.String("results", "user", "name")
	revision, _ := json.Int("results", "revision")

	app := App{
		name:        appName,
		packageName: packageName,
		url:         baseApiEndPoint + path,
		owner:       ownerName,
		revision:    revision,
	}
	return true, app
}

func Login(email string, password string) bool {
	if email == "" || password == "" {
		email, password = scanEmailAndPassword()
	}

	loginData := map[string]string{
		"email":    email,
		"password": password,
	}
	json := postRequest("/api/sessions", loginData)

	error, message := checkError(json)
	if error {
		println(message)
		return false
	}

	name, _ := json.String("results", "name")
	token, _ := json.String("results", "api_token")

	settings := `{"name":"` + name + `","token":"` + token + `"}`
	writeSettingFile(settings)

	return true
}

func Logout() {
	name, _ := getSessions()
	settings := `{"name":"` + name + `","token":""}`
	writeSettingFile(settings)
	println("Logout Success")
}

/******************
  private methods
*******************/

func scanEmailAndPassword() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	print("Email: ")
	scanner.Scan()
	email := scanner.Text()

	print("Password: ")
	pass := gopass.GetPasswd()
	return email, string(pass)
}

func isLogin() bool {
	name, token := getSessions()
	return name != "" && token != ""
}

func checkError(json *jsonq.JsonQuery) (bool, string) {
	error, _ := json.Bool("error")
	message := ""

	if error {
		message, _ = json.String("message")
	}
	return error, message
}

func getSessions() (string, string) {
	settingFile := getSettingFilePath()
	fileByte, err := ioutil.ReadFile(settingFile)
	if err != nil {
		return "", ""
	}

	json := stringToJsonq(string(fileByte))
	name, err := json.String("name")
	if err != nil {
		log.Fatal(err)
	}

	token, err := json.String("token")
	if err != nil {
		log.Fatal(err)
	}

	return name, token
}

func getSettingFilePath() string {
	return os.Getenv("HOME") + "/.dgate"
}

func writeSettingFile(settings string) {
	settingFile := getSettingFilePath()
	ioutil.WriteFile(settingFile, []byte(settings), 0644)
}

func stringToJsonq(jsonString string) *jsonq.JsonQuery {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonString))
	dec.Decode(&data)
	json := jsonq.NewQuery(data)
	return json
}

func responseToJsonq(res *httpclient.Response) *jsonq.JsonQuery {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonString := string(body)
	json := stringToJsonq(jsonString)

	return json
}

func postRequest(path string, params map[string]string) *jsonq.JsonQuery {
	url := baseApiEndPoint + path

	_, token := getSessions()
	setUpClient(token)
	res, err := httpclient.Post(url, params)
	if err != nil {
		log.Fatal(err)
	}

	json := responseToJsonq(res)
	return json
}

func setUpClient(apiToken string) {
	userAgent := "dgate/" + Version()
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: userAgent,
		"AUTHORIZATION":          apiToken,
	})
}
