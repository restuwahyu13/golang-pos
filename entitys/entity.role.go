package entitys

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type EntityRole interface {
	EntityCreate(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelRole, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError)
}
