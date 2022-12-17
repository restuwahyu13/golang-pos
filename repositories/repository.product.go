package repositories

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db: db}
}

/**
* ==========================================
* Repository Create New Product Teritory
*===========================================
 */

func (r *repositoryProduct) EntityCreate(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError) {
	var product models.ModelProduct
	product.Name = input.Name
	product.Image = input.Image
	product.SKU = input.SKU
	product.Price = input.Price
	product.OutletID = input.OutletID
	product.SupplierID = input.SupplierID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	checkproductName := db.Debug().First("name = ?", product.Name)

	if checkproductName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &product, <-err
	}

	addProduct := db.Debug().Create(&product).Commit()

	if addProduct.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}

/**
* ==========================================
* Repository Results All Product Teritory
*===========================================
 */

func (r *repositoryProduct) EntityResults() (*[]models.ModelProduct, schemes.SchemeDatabaseError) {
	var product []models.ModelProduct

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	checkProduct := db.Debug().Order("created_at DESC").Find(&product)

	if checkProduct.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryProduct) EntityResult(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError) {
	var product models.ModelProduct
	product.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	checkProductId := db.Debug().First(&product)

	if checkProductId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryProduct) EntityDelete(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError) {
	var product models.ModelProduct
	product.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	checkProductId := db.Debug().First(&product)

	if checkProductId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &product, <-err
	}

	deleteProduct := db.Debug().Delete(&product)

	if deleteProduct.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryProduct) EntityUpdate(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError) {
	var product models.ModelProduct
	product.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	checkProductId := db.Debug().First(&product)

	if checkProductId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &product, <-err
	}

	product.Name = input.Name
	product.Image = input.Image
	product.SKU = input.SKU
	product.Price = input.Price
	product.OutletID = input.OutletID
	product.SupplierID = input.SupplierID

	updateProduct := db.Debug().Updates(&product)

	if updateProduct.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}

/**
* ================================================
* Repository Result Porduct By Merchant ID Teritory
*=================================================
 */

func (r *repositoryProduct) EntityProductByMerchant(input *schemes.SchemeProduct) (*[]models.ModelProduct, schemes.SchemeDatabaseError) {
	var product []models.ModelProduct

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	getMerchant := db.Debug().Find(&product, "merchant_id = ?", input.ID)

	if getMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}

/**
* ================================================
* Repository Result Porduct By Outlet ID Teritory
*=================================================
 */

func (r *repositoryProduct) EntityProductByOutlet(input *schemes.SchemeProduct) (*[]models.ModelProduct, schemes.SchemeDatabaseError) {
	var product []models.ModelProduct

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	getOutlet := db.Debug().Find(&product, "outlet_id = ?", input.ID)

	if getOutlet.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}
