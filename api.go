package main

import (
	"encoding/json"
	"github.com/ddliu/go-httpclient"
	"github.com/jmoiron/jsonq"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const baseApiEndPoint = "https://deploygate.com/"

func DoLogin(email string, password string) {
	loginData := map[string]string{
		"email":    email,
		"password": password,
	}
	json := postRequest("api/sessions", loginData)

	error, message := checkError(json)
	if error {
		println(message)
		return
	}

	name := getResultsJsonStringValue(json, "name")
	token := getResultsJsonStringValue(json, "api_token")

	settings := `{"name":"` + name + `","token":"` + token + `"}`
	writeSettingFile(settings)

	welcomeMessage := `Welcome to DeployGate!
Let's upload the app to DeployGate!`

	println(welcomeMessage)
}

func IsLogin() bool {
	name, token := getSessions()
	return name != "" && token != ""
}

/******************
  private methods
*******************/

func checkError(json *jsonq.JsonQuery) (bool, string) {
	error, _ := json.Bool("error")
	message := ""

	if error {
		message, _ = json.String("message")
	}
	return error, message
}

func getResultsJsonStringValue(json *jsonq.JsonQuery, key string) string {
	value, err := json.String("results", key)
	if err != nil {
		log.Fatal(err)
	}

	return value
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
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "dgate/golang", // TODO: Versionを取得するようにする
		"AUTHORIZATION":          apiToken,
	})
}
