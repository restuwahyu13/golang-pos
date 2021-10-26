package schemas

type SchemaDatabaseError struct {
	Type string
	Code int
}

type SchemaErrorResponse struct {
	StatusCode int         `json:"code"`
	Error      interface{} `json:"error"`
}

type SchemaUnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
