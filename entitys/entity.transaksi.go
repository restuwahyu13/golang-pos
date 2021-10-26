package entitys

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type EntityTransaction interface {
	EntityCreate(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelTransaction, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError)
}
