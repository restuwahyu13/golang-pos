package entities

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type EntityUser interface {
	EntityRegister(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
	EntityLogin(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
}
