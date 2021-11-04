package authentication

type Authentication struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Data       struct {
		Token  string
		userId string
	}
}
