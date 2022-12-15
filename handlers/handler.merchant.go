package handlers

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/helpers"
	"github.com/restuwahyu13/golang-pos/pkg"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type handlerMerchant struct {
	merchant entities.EntityMerchant
}

func NewHandlerMerchant(merchant entities.EntityMerchant) *handlerMerchant {
	return &handlerMerchant{merchant: merchant}
}

/**
* ======================================
* Handler Ping Status Merchant Teritory
*=======================================
 */

func (h *handlerMerchant) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Merchant", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Merchant Teritory
*======================================
 */

func (h *handlerMerchant) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeMerchant
	file, _ := ctx.FormFile("logo")
	body.Logo = file.Filename
	body.Name = ctx.PostForm("name")
	body.Phone = ctx.PostForm("phone")
	body.Address = ctx.PostForm("address")
	body.SupplierID = ctx.PostForm("supplier_id")

	err := ctx.SaveUploadedFile(file, path.Join("images/"+file.Filename))
	fmt.Println(err)

	if err != nil {
		helpers.APIResponse(ctx, "File upload error", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorMerchant(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.merchant.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Merchant name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Merchant phone number already taken", error.Code, nil)
		return
	}

	if error.Type == "error_create_03" {
		helpers.APIResponse(ctx, "Create new merchant failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new merchant successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Merchant Teritory
*=======================================
 */

func (h *handlerMerchant) HandlerResults(ctx *gin.Context) {
	res, error := h.merchant.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Merchant data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Merchant data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Merchant By ID Teritory
*=======================================
 */

func (h *handlerMerchant) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeMerchant
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorMerchant(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.merchant.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Merchant data not found for this id %s ", body.ID), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Merchant data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Merchant By ID Teritory
*=======================================
 */

func (h *handlerMerchant) HandlerDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	var body schemes.SchemeMerchant
	body.ID = id

	errors, code := ValidatorMerchant(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.merchant.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Merchant data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete merchant data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete merchant data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Merchant By ID Teritory
*=======================================
 */

func (h *handlerMerchant) HandlerUpdate(ctx *gin.Context) {
	var body schemes.SchemeMerchant
	id := ctx.Param("id")
	file, _ := ctx.FormFile("logo")
	body.ID = id
	body.Logo = file.Filename
	body.Name = ctx.PostForm("name")
	body.Phone = ctx.PostForm("phone")
	body.Address = ctx.PostForm("address")
	body.SupplierID = ctx.PostForm("supplier_id")

	err := ctx.SaveUploadedFile(file, path.Join("images/"+file.Filename))
	fmt.Println(err)

	if err != nil {
		helpers.APIResponse(ctx, "File upload error", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorMerchant(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.merchant.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Merchant data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update merchant data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update merchant data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Merchant
*=======================================
 */

func ValidatorMerchant(ctx *gin.Context, input schemes.SchemeMerchant, Type string) (interface{}, int) {
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
					Field:   "Phone",
					Message: "Phone is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Phone",
					Message: "Phone must be number",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Address",
					Message: "Address is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "max",
					Field:   "Address",
					Message: "Address maximal 1000 character",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Suppliers",
					Message: "Suppliers is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Suppliers",
					Message: "Suppliers is required on body",
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
					Field:   "Phone",
					Message: "Phone is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Phone",
					Message: "Phone must be number",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Address",
					Message: "Address is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "max",
					Field:   "Address",
					Message: "Address maximal 1000 character",
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
