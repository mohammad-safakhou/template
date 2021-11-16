package utils

type StandardHttpResponse struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Code    int         `json:"code"`
}
type StandardHttpErrorResponse struct {
	error
	Message string `json:"message"`
	Code    int    `json:"code"`
}
