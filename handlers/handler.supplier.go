package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/restuwahyu13/golang-pos/entitys"
	"github.com/restuwahyu13/golang-pos/helpers"
	"github.com/restuwahyu13/golang-pos/pkg"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type handleSupplier struct {
	supplier entitys.EntitySupplier
}

func NewHandlerSupplier(supplier entitys.EntitySupplier) *handleSupplier {
	return &handleSupplier{supplier: supplier}
}

/**
* ======================================
* Handler Ping Status Supplier Teritory
*=======================================
 */

func (h *handleSupplier) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Supplier", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Supplier Teritory
*======================================
 */

func (h *handleSupplier) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaSupplier
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorSupplier(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.supplier.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Supplier name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Supplier failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Supplier successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Supplier Teritory
*=======================================
 */

func (h *handleSupplier) HandlerResults(ctx *gin.Context) {
	res, error := h.supplier.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Supplier data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Supplier data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Supplier By ID Teritory
*=======================================
 */

func (h *handleSupplier) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaSupplier
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorSupplier(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.supplier.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Supplier data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Supplier data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Supplier By ID Teritory
*=======================================
 */

func (h *handleSupplier) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaSupplier
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorSupplier(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.supplier.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Supplier data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Supplier data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Supplier data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Supplier By ID Teritory
*=======================================
 */

func (h *handleSupplier) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaSupplier
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorSupplier(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.supplier.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Supplier data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Supplier data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Supplier data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Supplier
*=======================================
 */

func ValidatorSupplier(ctx *gin.Context, input schemas.SchemaSupplier, Type string) (interface{}, int) {
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
					Tag:     "gte",
					Field:   "Phone",
					Message: "Phone number must be 12 character",
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
					Tag:     "gte",
					Field:   "Phone",
					Message: "Phone number must be 12 character",
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
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
