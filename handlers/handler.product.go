package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/restuwahyu13/golang-pos/entitys"
	"github.com/restuwahyu13/golang-pos/helpers"
	"github.com/restuwahyu13/golang-pos/pkg"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type handleProduct struct {
	product entitys.EntityProduct
}

func NewHandlerProduct(product entitys.EntityProduct) *handleProduct {
	return &handleProduct{product: product}
}

/**
* ======================================
* Handler Ping Status Product Teritory
*=======================================
 */

func (h *handleProduct) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Product", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Product Teritory
*======================================
 */

func (h *handleProduct) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaProduct
	file, _ := ctx.FormFile("image")
	sku, _ := strconv.ParseUint(ctx.PostForm("sku"), 10, 64)
	price, _ := strconv.ParseUint(ctx.PostForm("price"), 10, 64)

	body.Image = file.Filename
	body.Name = ctx.PostForm("name")
	body.SKU = sku
	body.Price = price
	body.OutletID = ctx.PostForm("outlet_id")
	body.SupplierID = ctx.PostForm("supplier_id")

	err := ctx.SaveUploadedFile(file, path.Join("images/"+file.Filename))
	fmt.Println(err)

	if err != nil {
		helpers.APIResponse(ctx, "File upload error", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorProduct(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.product.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Product name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Add new Product failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Add new Product successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Product Teritory
*=======================================
 */

func (h *handleProduct) HandlerResults(ctx *gin.Context) {
	res, error := h.product.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Product data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Product data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Product By ID Teritory
*=======================================
 */

func (h *handleProduct) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaProduct
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorProduct(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.product.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Product data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Product data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Product By ID Teritory
*=======================================
 */

func (h *handleProduct) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaProduct
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorProduct(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.product.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Product data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Product data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Product data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Product By ID Teritory
*=======================================
 */

func (h *handleProduct) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaProduct
	file, _ := ctx.FormFile("logo")
	sku, _ := strconv.ParseUint(ctx.PostForm("sku"), 10, 64)
	price, _ := strconv.ParseUint(ctx.PostForm("price"), 10, 64)
	id := ctx.Param("id")

	body.ID = id
	body.Image = file.Filename
	body.Name = ctx.PostForm("name")
	body.SKU = sku
	body.Price = price
	body.OutletID = ctx.PostForm("outlet_id")
	body.SupplierID = ctx.PostForm("supplier_id")

	errors, code := ValidatorProduct(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.product.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Product data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Product data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Product data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ======================================
* Handler Result Product By ID Teritory
*=======================================
 */

func (h *handleProduct) HandlerResultByOutlet(ctx *gin.Context) {
	var body schemas.SchemaProduct
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorProduct(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.product.EntityProductByOutlet(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Product data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Product data already to use", http.StatusOK, res)
}

/**
* ======================================
*  All Validator User Input For Product
*=======================================
 */

func ValidatorProduct(ctx *gin.Context, input schemas.SchemaProduct, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Name",
					Message: "Name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "lowercase",
					Field:   "Name",
					Message: "Name must be lowercase",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "SKU",
					Message: "SKU is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "numeric",
					Field:   "Phone",
					Message: "SKU must be number",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Price",
					Message: "Price is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "price",
					Field:   "Price",
					Message: "Price must be number",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "MerchantID",
					Message: "MerchantID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "MerchantID must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "OutletID",
					Message: "OutletID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "OutletID",
					Message: "OutletID must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "SupplierID",
					Message: "SupplierID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "SupplierID",
					Message: "SupplierID must be uuid",
				},
			},
		}
	}

	if Type == "result" || Type == "delete" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "ID",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "ID",
					Message: "ID must be uuid",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "ID",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "ID",
					Message: "ID must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Name",
					Message: "Name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "lowercase",
					Field:   "Name",
					Message: "Name must be lowercase",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "SKU",
					Message: "SKU is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "numeric",
					Field:   "Phone",
					Message: "SKU must be number",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Price",
					Message: "Price is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "price",
					Field:   "Price",
					Message: "Price must be number",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "MerchantID",
					Message: "MerchantID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "MerchantID must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "OutletID",
					Message: "OutletID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "OutletID",
					Message: "OutletID must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "SupplierID",
					Message: "SupplierID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "SupplierID",
					Message: "SupplierID must be uuid",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
