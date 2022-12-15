package services

import (
	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type serviceMerchant struct {
	merchant entities.EntityMerchant
}

func NewServiceMerchant(merchant entities.EntityMerchant) *serviceMerchant {
	return &serviceMerchant{merchant: merchant}
}

/**
* ==========================================
* Service Create New Merchant Teritory
*===========================================
 */

func (s *serviceMerchant) EntityCreate(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant schemes.SchemeMerchant
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

func (s *serviceMerchant) EntityResults() (*[]models.ModelMerchant, schemes.SchemeDatabaseError) {
	res, err := s.merchant.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityResult(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant schemes.SchemeMerchant
	merchant.ID = input.ID

	res, err := s.merchant.EntityResult(&merchant)
	return res, err
}

/**
* ==========================================
* Service Delete Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityDelete(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant schemes.SchemeMerchant
	merchant.ID = input.ID

	res, err := s.merchant.EntityDelete(&merchant)
	return res, err
}

/**
* ==========================================
* Service Update Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityUpdate(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant schemes.SchemeMerchant
	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.SupplierID = input.SupplierID

	res, err := s.merchant.EntityUpdate(&merchant)
	return res, err
}
