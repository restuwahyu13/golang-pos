package services

import (
	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type serviceOutlite struct {
	outlet entities.EntityOutlet
}

func NewServiceOutlet(outlet entities.EntityOutlet) *serviceOutlite {
	return &serviceOutlite{outlet: outlet}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceOutlite) EntityCreate(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet schemes.SchemeOutlet
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

func (s *serviceOutlite) EntityResults() (*[]models.ModelOutlet, schemes.SchemeDatabaseError) {
	res, err := s.outlet.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Teritory
*===========================================
 */

func (s *serviceOutlite) EntityResult(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet schemes.SchemeOutlet
	outlet.ID = input.ID

	res, err := s.outlet.EntityResult(&outlet)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceOutlite) EntityDelete(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet schemes.SchemeOutlet
	outlet.ID = input.ID

	res, err := s.outlet.EntityDelete(&outlet)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceOutlite) EntityUpdate(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet schemes.SchemeOutlet
	outlet.Name = input.Name
	outlet.Phone = input.Phone
	outlet.Address = input.Address
	outlet.MerchatID = input.MerchatID

	res, err := s.outlet.EntityUpdate(&outlet)
	return res, err
}
