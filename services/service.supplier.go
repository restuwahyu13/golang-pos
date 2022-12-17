package services

import (
	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type serviceSupplier struct {
	supplier entities.EntitySupplier
}

func NewServiceSupplier(supplier entities.EntitySupplier) *serviceSupplier {
	return &serviceSupplier{supplier: supplier}
}

/**
* ==========================================
* Service Create New Supplier Teritory
*===========================================
 */

func (s *serviceSupplier) EntityCreate(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError) {
	var supplier schemes.SchemeSupplier
	supplier.Name = input.Name
	supplier.Phone = input.Phone
	supplier.Address = input.Address

	res, err := s.supplier.EntityCreate(&supplier)
	return res, err
}

/**
* ==========================================
* Service Results All Supplier Teritory
*===========================================
 */

func (s *serviceSupplier) EntityResults() (*[]models.ModelSupplier, schemes.SchemeDatabaseError) {
	res, err := s.supplier.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityResult(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError) {
	var supplier schemes.SchemeSupplier
	supplier.ID = input.ID

	res, err := s.supplier.EntityResult(&supplier)
	return res, err
}

/**
* ==========================================
* Service Delete Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityDelete(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError) {
	var supplier schemes.SchemeSupplier
	supplier.ID = input.ID

	res, err := s.supplier.EntityDelete(&supplier)
	return res, err
}

/**
* ==========================================
* Service Update Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityUpdate(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError) {
	var supplier schemes.SchemeSupplier
	supplier.Name = input.Name
	supplier.Phone = input.Phone
	supplier.Address = input.Address

	res, err := s.supplier.EntityUpdate(&supplier)
	return res, err
}
