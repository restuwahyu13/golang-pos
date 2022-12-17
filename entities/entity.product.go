package entities

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type EntityProduct interface {
	EntityCreate(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelProduct, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError)
	EntityProductByOutlet(input *schemes.SchemeProduct) (*[]models.ModelProduct, schemes.SchemeDatabaseError)
}
