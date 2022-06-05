package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	models "go-meater-meter/meater/models"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/ini.v1"
)

func GetAuth(cfg *ini.File) string {
	email := cfg.Section("api-authentication").Key("email").String()
	password := cfg.Section("api-authentication").Key("password").String()
	values := map[string]string{"email": email, "password": password}

	jsonValue, _ := json.Marshal(values)
	//https://github.com/apption-labs/meater-cloud-public-rest-api#api-endpoints
	resp, err := http.Post("https://public-api.cloud.meater.com/v1/login", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(time.Now().Format(time.RFC3339), "Authentication Response status:", resp.Status)

	if resp.Status != "200 OK" {
		fmt.Println("Authentication failed. Exiting.")
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var result models.Authentication
	json.Unmarshal([]byte(body), &result)
	sessionToken := result.Data.Token
	return sessionToken
}
