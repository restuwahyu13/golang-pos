package schemes

type SchemeDatabaseError struct {
	Type string
	Code int
}

type SchemeErrorResponse struct {
	StatusCode int         `json:"code"`
	Error      interface{} `json:"error"`
}

type SchemeUnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
