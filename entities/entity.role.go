package entities

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type EntityRole interface {
	EntityCreate(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelRole, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError)
}
