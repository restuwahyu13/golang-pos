package services

import (
	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type serviceTransaction struct {
	transaction entities.EntityTransaction
}

func NewServiceTransaction(transaction entities.EntityTransaction) *serviceTransaction {
	return &serviceTransaction{transaction: transaction}
}

/**
* ==========================================
* Service Create New Transaction Teritory
*===========================================
 */

func (s *serviceTransaction) EntityCreate(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError) {
	var transaction schemes.SchemeTransaction
	transaction.CustomerID = input.CustomerID
	transaction.OutletID = input.OutletID
	transaction.Products = input.Products

	res, err := s.transaction.EntityCreate(&transaction)
	return res, err
}

/**
* ============================================
* Service Results All Transaction Teritory
*=============================================
 */

func (s *serviceTransaction) EntityResults() (*[]models.ModelTransaction, schemes.SchemeDatabaseError) {
	res, err := s.transaction.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Transaction By ID Teritory
*===========================================
 */

func (s *serviceTransaction) EntityResult(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError) {
	var transaction schemes.SchemeTransaction
	transaction.ID = input.ID

	res, err := s.transaction.EntityResult(&transaction)
	return res, err
}

/**
* ==========================================
* Service Delete Transaction By ID Teritory
*===========================================
 */

func (s *serviceTransaction) EntityDelete(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError) {
	var transaction schemes.SchemeTransaction
	transaction.ID = input.ID

	res, err := s.transaction.EntityDelete(&transaction)
	return res, err
}

/**
* ==========================================
* Service Update Transaction By ID Teritory
*===========================================
 */

func (s *serviceTransaction) EntityUpdate(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError) {
	var transaction schemes.SchemeTransaction
	transaction.CustomerID = input.CustomerID
	transaction.OutletID = input.OutletID
	transaction.Products = input.Products

	res, err := s.transaction.EntityUpdate(&transaction)
	return res, err
}
