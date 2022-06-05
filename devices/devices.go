package devices

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajski1701/go-meater-meter/config"

	models "github.com/ajski1701/go-meater-meter/models"
)

func GetDevices(token string) []models.Devices {
	url_debug := config.LoadConfigIni().Section("debug").Key("device_api_url").String()
	var url string
	var bearer = "Bearer " + token

	if url_debug == "" {
		url = "https://public-api.cloud.meater.com/v1/devices"
	} else {
		url = url_debug
	}
	//url = "http://127.0.0.1:8080/device_json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result models.Device
	json.Unmarshal([]byte(body), &result)

	return result.Data.Devices
}
