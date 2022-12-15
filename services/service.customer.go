package services

import (
	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type serviceCustomer struct {
	customer entities.EntityCustomer
}

func NewServiceCustomer(customer entities.EntityCustomer) *serviceCustomer {
	return &serviceCustomer{customer: customer}
}

/**
* ==========================================
* Service Create New Customer Teritory
*===========================================
 */

func (s *serviceCustomer) EntityCreate(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer schemes.SchemeCustomer
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

func (s *serviceCustomer) EntityResults() (*[]models.ModelCustomer, schemes.SchemeDatabaseError) {
	res, err := s.customer.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityResult(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer schemes.SchemeCustomer
	customer.ID = input.ID

	res, err := s.customer.EntityResult(&customer)
	return res, err
}

/**
* ==========================================
* Service Delete Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityDelete(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer schemes.SchemeCustomer
	customer.ID = input.ID

	res, err := s.customer.EntityDelete(&customer)
	return res, err
}

/**
* ==========================================
* Service Update Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityUpdate(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer schemes.SchemeCustomer
	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address

	res, err := s.customer.EntityUpdate(&customer)
	return res, err
}
