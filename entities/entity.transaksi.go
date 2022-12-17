package entities

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type EntityTransaction interface {
	EntityCreate(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelTransaction, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError)
}
