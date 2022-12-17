package repositories

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type repositoryMerchant struct {
	db *gorm.DB
}

func NewRepositoryMerchant(db *gorm.DB) *repositoryMerchant {
	return &repositoryMerchant{db: db}
}

/**
* ==========================================
* Repository Create New Merchant Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityCreate(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant models.ModelMerchant
	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.SupplierID = input.SupplierID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchantName := db.Debug().First(&merchant, "name = ?", input.Name)

	if checkMerchantName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &merchant, <-err
	}

	checkMerchantPhone := db.Debug().First(&merchant, "phone = ?", input.Phone)

	if checkMerchantPhone.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_02",
		}
		return &merchant, <-err
	}

	addMerchant := db.Debug().Create(&merchant).Commit()

	if addMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}

/**
* ==========================================
* Repository Results All Merchant Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityResults() (*[]models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant []models.ModelMerchant

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchantName := db.Debug().Order("created_at DESC").Find(&merchant)

	if checkMerchantName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityResult(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant models.ModelMerchant
	merchant.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchantName := db.Debug().First(&merchant)

	if checkMerchantName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityDelete(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant models.ModelMerchant
	merchant.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchantName := db.Debug().First(&merchant)

	if checkMerchantName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &merchant, <-err
	}

	deleteMerchant := db.Debug().Delete(&merchant)

	if deleteMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityUpdate(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError) {
	var merchant models.ModelMerchant
	merchant.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchantName := db.Debug().First(&merchant)

	if checkMerchantName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &merchant, <-err
	}

	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.SupplierID = input.SupplierID

	updateMerchant := db.Debug().Updates(&merchant)

	if updateMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}
