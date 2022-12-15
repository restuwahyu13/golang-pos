package entities

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type EntitySupplier interface {
	EntityCreate(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelSupplier, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError)
}
