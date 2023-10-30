package model

type ApiResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  *ErrorData  `json:"error,omitempty"`
}

type ErrorData struct {
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Action  string `json:"action,omitempty"`
}
