package entitys

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type EntityUser interface {
	EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError)
	EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError)
}
