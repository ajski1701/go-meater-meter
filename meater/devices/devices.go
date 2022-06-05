package devices

import (
	"encoding/json"
	models "go-meater-meter/meater/models"
	"io/ioutil"
	"net/http"
)

func GetDevices(token string) []models.Devices {
	var bearer = "Bearer " + token

	url := "https://public-api.cloud.meater.com/v1/devices"
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
