package entities

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type EntityCustomer interface {
	EntityCreate(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelCustomer, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError)
}
