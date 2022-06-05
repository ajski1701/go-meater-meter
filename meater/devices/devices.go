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
	url = "http://127.0.0.1:8080/device_json"
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

	// for _, device := range result.Data.Devices {
	// 	localOutputMap := map[string]string{}

	// 	localOutputMap["device_id"] = device.Id
	// 	localOutputMap["temperature_ambient"] = device.Temperature.Ambient
	// 	localOutputMap["temperature_internal"] = device.Temperature.Internal
	// 	localOutputMap["cook_id"] = device.Cook.Id
	// 	localOutputMap["cook_name"] = device.Cook.Name
	// 	localOutputMap["cook_state"] = device.Cook.State
	// 	localOutputMap["cook_target"] = device.Cook.Temperature.Target
	// 	localOutputMap["cook_peak"] = device.Cook.Temperature.Peak
	// 	localOutputMap["cook_time_elapsed"] = device.Cook.Time.Elapsed
	// 	localOutputMap["cook_time_remaining"] = device.Cook.Time.Remaining
	// 	localOutputMap["updated_at"] = device.Temperature.Internal
	// 	outputArray = append(outputArray, localOutputMap)
	// }
	return result.Data.Devices
}
