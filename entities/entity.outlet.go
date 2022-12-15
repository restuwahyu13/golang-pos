package entities

import (
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type EntityOutlet interface {
	EntityCreate(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelOutlet, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError)
}
