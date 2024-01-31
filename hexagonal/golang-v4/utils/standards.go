package utils

type StandardHttpResponse struct {
	error
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

const (
	Ok                   = "OK"
	NotValidField        = "Not a valid %s"
	NotFound             = "Data not found!"
	NotValidData         = "Request is not Valid!"
	NoAccess             = "You do not have access!"
	ProblemInGettingData = "Problem in getting data!"
	ProblemInSystem      = "Problem in system!"
	CustomMessage        = "%s"
)
