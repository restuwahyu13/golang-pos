package schemas

import "time"

type JwtMetaOptions struct {
	Audience  string
	ExpiredAt time.Duration
}

type JWtMetaRequest struct {
	Data      map[string]interface{}
	SecretKey string
	Options   JwtMetaOptions
}
