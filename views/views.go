package views

type Response struct {
	Code     int         `json:"code"`
	Response interface{} `json:"body"`
}
