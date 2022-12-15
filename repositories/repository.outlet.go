package repositories

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type repositoryOutlite struct {
	db *gorm.DB
}

func NewRepositoryOutlet(db *gorm.DB) *repositoryOutlite {
	return &repositoryOutlite{db: db}
}

/**
* ==========================================
* Repository Create New Outlet Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityCreate(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet models.ModelOutlet
	phone, _ := strconv.ParseUint(input.Phone, 10, 64)
	outlet.Name = input.Name
	outlet.Phone = phone
	outlet.Address = input.Address
	outlet.MerchantID = input.MerchatID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	checkOutletName := db.Debug().First(&outlet, "name = ?", outlet.Name)

	if checkOutletName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &outlet, <-err
	}

	checkOutletPhone := db.Debug().First(&outlet, "phone = ?", outlet.Phone)

	if checkOutletPhone.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_02",
		}
		return &outlet, <-err
	}

	addoutlet := db.Debug().Create(&outlet).Commit()

	if addoutlet.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &outlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}

/**
* ==========================================
* Repository Results All Outlet Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityResults() (*[]models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet []models.ModelOutlet

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	checkOutletName := db.Debug().Order("created_at DESC").Find(&outlet)

	if checkOutletName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &outlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityResult(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet models.ModelOutlet
	outlet.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	checkOutletName := db.Debug().First(&outlet)

	if checkOutletName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &outlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityDelete(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet models.ModelOutlet
	outlet.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	checkOutletName := db.Debug().First(&outlet)

	if checkOutletName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &outlet, <-err
	}

	deleteoutlet := db.Debug().Delete(&outlet)

	if deleteoutlet.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &outlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityUpdate(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError) {
	var outlet models.ModelOutlet
	outlet.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	checkOutletName := db.Debug().First(&outlet)

	if checkOutletName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &outlet, <-err
	}

	phone, _ := strconv.ParseUint(input.Phone, 10, 64)
	outlet.Name = input.Name
	outlet.Phone = phone
	outlet.Address = input.Address
	outlet.MerchantID = input.MerchatID

	updateoutlet := db.Debug().Updates(&outlet)

	if updateoutlet.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &outlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}

/**
* ================================================
* Repository Result Outlet By Merchant Teritory
*=================================================
 */

func (r *repositoryOutlite) EntitySaleProduct(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError) {
	var transaction models.ModelTransaction
	transaction.CustomerID = input.ID
	transaction.OutletID = input.OutletID
	transaction.Products = input.Products

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&transaction)

	addTransaction := db.Debug().Create(&transaction).Commit()

	if addTransaction.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_result_01",
		}
		return &transaction, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &transaction, <-err
}
