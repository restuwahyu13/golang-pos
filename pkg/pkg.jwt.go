package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"

	"github.com/restuwahyu13/golang-pos/schemas"
)

func Sign(configs *schemas.JWtMetaRequest) (string, error) {

	expiredAt := time.Now().Add(time.Duration(time.Minute) * configs.Options.ExpiredAt).Unix()

	claims := jwt.MapClaims{}
	claims["jwt"] = configs.Data
	claims["exp"] = (24 * 60) * expiredAt
	claims["audience"] = configs.Options.Audience
	claims["authorization"] = true

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(configs.SecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyToken(accessToken, SecretPublicKey string) (*jwt.Token, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretPublicKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}
