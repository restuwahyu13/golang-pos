package services

import (
	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type serviceProduct struct {
	product entities.EntityProduct
}

func NewServiceProduct(product entities.EntityProduct) *serviceProduct {
	return &serviceProduct{product: product}
}

/**
* ==========================================
* Service Create New Merchant Teritory
*===========================================
 */

func (s *serviceProduct) EntityCreate(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError) {
	var product schemes.SchemeProduct
	product.Name = input.Name
	product.Image = input.Image
	product.SKU = input.SKU
	product.Price = input.Price
	product.OutletID = input.OutletID
	product.SupplierID = input.SupplierID

	res, err := s.product.EntityCreate(&product)
	return res, err
}

/**
* ==========================================
* Service Results All Merchant Teritory
*===========================================
 */

func (s *serviceProduct) EntityResults() (*[]models.ModelProduct, schemes.SchemeDatabaseError) {
	res, err := s.product.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Merchant By ID Teritory
*===========================================
 */

func (s *serviceProduct) EntityResult(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError) {
	var product schemes.SchemeProduct
	product.ID = input.ID

	res, err := s.product.EntityResult(&product)
	return res, err
}

/**
* ==========================================
* Service Delete Merchant By ID Teritory
*===========================================
 */

func (s *serviceProduct) EntityDelete(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError) {
	var product schemes.SchemeProduct
	product.ID = input.ID

	res, err := s.product.EntityDelete(&product)
	return res, err
}

/**
* ==========================================
* Service Update Merchant By ID Teritory
*===========================================
 */

func (s *serviceProduct) EntityUpdate(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError) {
	var product schemes.SchemeProduct
	product.Name = input.Name
	product.Image = input.Image
	product.SKU = input.SKU
	product.Price = input.Price
	product.OutletID = input.OutletID
	product.SupplierID = input.SupplierID

	res, err := s.product.EntityUpdate(&product)
	return res, err
}

/**
* ===============================================
* Service Result Product By Outlet ID Teritory
*================================================
 */

func (s *serviceProduct) EntityProductByOutlet(input *schemes.SchemeProduct) (*[]models.ModelProduct, schemes.SchemeDatabaseError) {
	var product schemes.SchemeProduct
	product.ID = input.ID

	res, err := s.product.EntityProductByOutlet(&product)
	return res, err
}
