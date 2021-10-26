package services

import (
	"github.com/restuwahyu13/golang-pos/entitys"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type serviceMerchant struct {
	merchant entitys.EntityMerchant
}

func NewServiceMerchant(merchant entitys.EntityMerchant) *serviceMerchant {
	return &serviceMerchant{merchant: merchant}
}

/**
* ==========================================
* Service Create New Merchant Teritory
*===========================================
 */

func (s *serviceMerchant) EntityCreate(input *schemas.SchemaMerchant) (*models.ModelMerchant, schemas.SchemaDatabaseError) {
	var merchant schemas.SchemaMerchant
	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.SupplierID = input.SupplierID

	res, err := s.merchant.EntityCreate(&merchant)
	return res, err
}

/**
* ==========================================
* Service Results All MerchantTeritory
*===========================================
 */

func (s *serviceMerchant) EntityResults() (*[]models.ModelMerchant, schemas.SchemaDatabaseError) {
	res, err := s.merchant.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityResult(input *schemas.SchemaMerchant) (*models.ModelMerchant, schemas.SchemaDatabaseError) {
	var merchant schemas.SchemaMerchant
	merchant.ID = input.ID

	res, err := s.merchant.EntityResult(&merchant)
	return res, err
}

/**
* ==========================================
* Service Delete Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityDelete(input *schemas.SchemaMerchant) (*models.ModelMerchant, schemas.SchemaDatabaseError) {
	var merchant schemas.SchemaMerchant
	merchant.ID = input.ID

	res, err := s.merchant.EntityDelete(&merchant)
	return res, err
}

/**
* ==========================================
* Service Update Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityUpdate(input *schemas.SchemaMerchant) (*models.ModelMerchant, schemas.SchemaDatabaseError) {
	var merchant schemas.SchemaMerchant
	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.SupplierID = input.SupplierID

	res, err := s.merchant.EntityUpdate(&merchant)
	return res, err
}
