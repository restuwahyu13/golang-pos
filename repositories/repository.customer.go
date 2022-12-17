package repositories

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"

	models "github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type repositoryCustomer struct {
	db *gorm.DB
}

func NewRepositoryCustomer(db *gorm.DB) *repositoryCustomer {
	return &repositoryCustomer{db: db}
}

/**
* ==========================================
* Repository Create New Csutomer Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityCreate(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer models.ModelCustomer
	phone, _ := strconv.ParseUint(input.Phone, 10, 64)
	customer.Name = input.Name
	customer.Phone = phone
	customer.Address = input.Address

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer, "name = ?", customer.Name)

	if checkCustomerName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &customer, <-err
	}

	addCustomer := db.Debug().Create(&customer).Commit()

	if addCustomer.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}

/**
* ==========================================
* Repository Results All Csutomer Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityResults() (*[]models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer []models.ModelCustomer

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomer := db.Debug().Order("created_at DESC").Find(&customer)

	if checkCustomer.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityResult(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer models.ModelCustomer
	customer.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomerId := db.Debug().First(&customer)

	if checkCustomerId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityDelete(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer models.ModelCustomer
	customer.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomerId := db.Debug().First(&customer)

	if checkCustomerId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &customer, <-err
	}

	deleteCustomer := db.Debug().Delete(&customer)

	if deleteCustomer.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityUpdate(input *schemes.SchemeCustomer) (*models.ModelCustomer, schemes.SchemeDatabaseError) {
	var customer models.ModelCustomer
	customer.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomerId := db.Debug().First(customer.ID)

	if checkCustomerId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &customer, <-err
	}

	phone, _ := strconv.ParseUint(input.Phone, 10, 64)
	customer.Name = input.Name
	customer.Phone = phone
	customer.Address = input.Address

	updateCustomer := db.Debug().Updates(&customer)

	if updateCustomer.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}
