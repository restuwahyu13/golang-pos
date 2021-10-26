package entitys

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type EntityProduct interface {
	EntityCreate(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelProduct, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
	EntityProductByOutlet(input *schemas.SchemaProduct) (*[]models.ModelProduct, schemas.SchemaDatabaseError)
}
