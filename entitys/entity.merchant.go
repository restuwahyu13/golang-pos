package entitys

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type EntityMerchant interface {
	EntityCreate(input *schemas.SchemaMerchant) (*models.ModelMerchant, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaMerchant) (*models.ModelMerchant, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelMerchant, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaMerchant) (*models.ModelMerchant, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaMerchant) (*models.ModelMerchant, schemas.SchemaDatabaseError)
}
