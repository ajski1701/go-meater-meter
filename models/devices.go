package models

type Device struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Data       struct {
		Devices []Devices
	}
}

type Devices struct {
	Id          string
	Temperature struct {
		Internal int
		Ambient  int
	}
	Cook struct {
		Id          string
		Name        string
		State       string
		Temperature struct {
			Target int
			Peak   int
		}
		Time struct {
			Elapsed   int
			Remaining int
		}
	}
	Updated_At int `json:"updated_at"`
}
