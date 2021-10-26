package repositorys

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type repositorySupplier struct {
	db *gorm.DB
}

func NewRepositorySupplier(db *gorm.DB) *repositorySupplier {
	return &repositorySupplier{db: db}
}

/**
* ==========================================
* Repository Create New Supplier Teritory
*===========================================
 */

func (r *repositorySupplier) EntityCreate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier models.ModelSupplier
	phone, _ := strconv.ParseUint(input.Phone, 10, 64)
	supplier.Name = input.Name
	supplier.Phone = phone
	supplier.Address = input.Address

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierName := db.Debug().First(&supplier, "name = ?", supplier.Name)

	if checkSupplierName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &supplier, <-err
	}

	checkSupplierPhone := db.Debug().First(&supplier, "phone = ?", supplier.Phone)

	if checkSupplierPhone.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_02",
		}
		return &supplier, <-err
	}

	addSupplier := db.Debug().Create(&supplier).Commit()

	if addSupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Results All Supplier Teritory
*===========================================
 */

func (r *repositorySupplier) EntityResults() (*[]models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier []models.ModelSupplier

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplier := db.Debug().Order("created_at DESC").Find(&supplier)

	if checkSupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityResult(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier models.ModelSupplier
	supplier.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierId := db.Debug().First(&supplier)

	if checkSupplierId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityDelete(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier models.ModelSupplier
	supplier.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierId := db.Debug().First(&supplier)

	if checkSupplierId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
	}

	deleteSupplier := db.Debug().Delete(&supplier)

	if deleteSupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityUpdate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier models.ModelSupplier
	supplier.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierId := db.Debug().First(&supplier)

	if checkSupplierId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &supplier, <-err
	}

	phone, _ := strconv.ParseUint(input.Phone, 10, 64)
	supplier.Name = input.Name
	supplier.Phone = phone
	supplier.Address = input.Address

	updateSupplier := db.Debug().Updates(&supplier)

	if updateSupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}
