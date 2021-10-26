package entitys

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type EntityOutlet interface {
	EntityCreate(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelOutlet, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError)
}
