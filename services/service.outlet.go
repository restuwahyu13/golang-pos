package services

import (
	"github.com/restuwahyu13/golang-pos/entitys"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type serviceOutlite struct {
	outlet entitys.EntityOutlet
}

func NewServiceOutlet(outlet entitys.EntityOutlet) *serviceOutlite {
	return &serviceOutlite{outlet: outlet}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceOutlite) EntityCreate(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError) {
	var outlet schemas.SchemaOutlet
	outlet.Name = input.Name
	outlet.Phone = input.Phone
	outlet.Address = input.Address
	outlet.MerchatID = input.MerchatID

	res, err := s.outlet.EntityCreate(&outlet)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet Teritory
*===========================================
 */

func (s *serviceOutlite) EntityResults() (*[]models.ModelOutlet, schemas.SchemaDatabaseError) {
	res, err := s.outlet.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Teritory
*===========================================
 */

func (s *serviceOutlite) EntityResult(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError) {
	var outlet schemas.SchemaOutlet
	outlet.ID = input.ID

	res, err := s.outlet.EntityResult(&outlet)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceOutlite) EntityDelete(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError) {
	var outlet schemas.SchemaOutlet
	outlet.ID = input.ID

	res, err := s.outlet.EntityDelete(&outlet)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceOutlite) EntityUpdate(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError) {
	var outlet schemas.SchemaOutlet
	outlet.Name = input.Name
	outlet.Phone = input.Phone
	outlet.Address = input.Address
	outlet.MerchatID = input.MerchatID

	res, err := s.outlet.EntityUpdate(&outlet)
	return res, err
}
