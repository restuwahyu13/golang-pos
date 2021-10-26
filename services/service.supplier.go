package services

import (
	"github.com/restuwahyu13/golang-pos/entitys"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type serviceSupplier struct {
	supplier entitys.EntitySupplier
}

func NewServiceSupplier(supplier entitys.EntitySupplier) *serviceSupplier {
	return &serviceSupplier{supplier: supplier}
}

/**
* ==========================================
* Service Create New Supplier Teritory
*===========================================
 */

func (s *serviceSupplier) EntityCreate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
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

func (s *serviceSupplier) EntityResults() (*[]models.ModelSupplier, schemas.SchemaDatabaseError) {
	res, err := s.supplier.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityResult(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.ID = input.ID

	res, err := s.supplier.EntityResult(&supplier)
	return res, err
}

/**
* ==========================================
* Service Delete Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityDelete(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.ID = input.ID

	res, err := s.supplier.EntityDelete(&supplier)
	return res, err
}

/**
* ==========================================
* Service Update Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityUpdate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.Name = input.Name
	supplier.Phone = input.Phone
	supplier.Address = input.Address

	res, err := s.supplier.EntityUpdate(&supplier)
	return res, err
}
