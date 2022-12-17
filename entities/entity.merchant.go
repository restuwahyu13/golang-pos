package entities

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type EntityMerchant interface {
	EntityCreate(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelMerchant, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError)
}
