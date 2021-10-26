package repositorys

import (
	"net/http"
	"time"

	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type repositoryTransaction struct {
	db *gorm.DB
}

func NewRepositoryTransaction(db *gorm.DB) *repositoryTransaction {
	return &repositoryTransaction{db: db}
}

/**
* ==========================================
* Repository Create New Transaction Teritory
*===========================================
 */

func (r *repositoryTransaction) EntityCreate(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction models.ModelTransaction
	transaction.CustomerID = input.CustomerID
	transaction.OutletID = input.OutletID
	transaction.Products = input.Products
	transaction.PurchaseDate = time.Now().Local()

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&transaction)

	addtranTransaction := db.Debug().Create(&transaction).Commit()

	if addtranTransaction.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_01",
		}
		return &transaction, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &transaction, <-err
}

/**
* ==========================================
* Repository Results All Transaction Teritory
*===========================================
 */

func (r *repositoryTransaction) EntityResults() (*[]models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction []models.ModelTransaction

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&transaction)

	checkTransaction := db.Debug().Order("created_at DESC").Find(&transaction)

	if checkTransaction.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &transaction, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &transaction, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryTransaction) EntityResult(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction models.ModelTransaction
	transaction.CustomerID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&transaction)

	checkTransactionId := db.Debug().First(&transaction, "customer_id = ?", input.ID)

	if checkTransactionId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &transaction, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &transaction, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryTransaction) EntityDelete(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction models.ModelTransaction
	transaction.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&transaction)

	checkTransactionId := db.Debug().First(&transaction)

	if checkTransactionId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &transaction, <-err
	}

	deleteTransaction := db.Debug().Delete(&transaction)

	if deleteTransaction.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &transaction, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &transaction, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryTransaction) EntityUpdate(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction models.ModelTransaction
	transaction.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&transaction)

	checkTransactionId := db.Debug().First(&transaction, "customer_id = ?", input.ID)

	if checkTransactionId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &transaction, <-err
	}

	transaction.CustomerID = input.CustomerID
	transaction.OutletID = input.OutletID
	transaction.Products = input.Products

	updateTransaction := db.Debug().Updates(&transaction)

	if updateTransaction.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &transaction, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &transaction, <-err
}
