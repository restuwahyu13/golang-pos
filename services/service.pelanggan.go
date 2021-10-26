package services

import (
	"github.com/restuwahyu13/golang-pos/entitys"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type serviceCustomer struct {
	customer entitys.EntityCustomer
}

func NewServiceCustomer(customer entitys.EntityCustomer) *serviceCustomer {
	return &serviceCustomer{customer: customer}
}

/**
* ==========================================
* Service Create New Customer Teritory
*===========================================
 */

func (s *serviceCustomer) EntityCreate(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer schemas.SchemaCustomer
	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address

	res, err := s.customer.EntityCreate(&customer)
	return res, err
}

/**
* ==========================================
* Service Results All Customer Teritory
*===========================================
 */

func (s *serviceCustomer) EntityResults() (*[]models.ModelCustomer, schemas.SchemaDatabaseError) {
	res, err := s.customer.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityResult(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer schemas.SchemaCustomer
	customer.ID = input.ID

	res, err := s.customer.EntityResult(&customer)
	return res, err
}

/**
* ==========================================
* Service Delete Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityDelete(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer schemas.SchemaCustomer
	customer.ID = input.ID

	res, err := s.customer.EntityDelete(&customer)
	return res, err
}

/**
* ==========================================
* Service Update Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityUpdate(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer schemas.SchemaCustomer
	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address

	res, err := s.customer.EntityUpdate(&customer)
	return res, err
}
