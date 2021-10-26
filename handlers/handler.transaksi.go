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

type handleTransaction struct {
	transaction entitys.EntityTransaction
}

func NewHandlerTransaction(transaction entitys.EntityTransaction) *handleTransaction {
	return &handleTransaction{transaction: transaction}
}

/**
* =========================================
* Handler Ping Status Transaction Teritory
*==========================================
 */

func (h *handleTransaction) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Transaction", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Transaction Teritory
*======================================
 */

func (h *handleTransaction) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaTransaction
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorTransaction(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.transaction.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Transaction name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new transaction failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Transaction successfully", http.StatusCreated, nil)
}

/**
* =========================================
* Handler Results All Transaction Teritory
*==========================================
 */

func (h *handleTransaction) HandlerResults(ctx *gin.Context) {
	res, error := h.transaction.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Transaction data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Transaction data already to use", http.StatusOK, res)
}

/**
* ==========================================
* Handler Result Transaction By ID Teritory
*===========================================
 */

func (h *handleTransaction) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaTransaction
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorTransaction(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.transaction.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Transaction data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Transaction data already to use", http.StatusOK, res)
}

/**
* ==========================================
* Handler Delete Transaction By ID Teritory
*===========================================
 */

func (h *handleTransaction) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaTransaction
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorTransaction(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.transaction.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Transaction data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Transaction data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Transaction data for this id %s success", id), http.StatusOK, res)
}

/**
* ==========================================
* Handler Update Transaction By ID Teritory
*===========================================
 */

func (h *handleTransaction) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaTransaction
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorTransaction(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.transaction.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Transaction data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Transaction data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Transaction data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ==========================================
*  All Validator User Input For Transaction
*===========================================
 */

func ValidatorTransaction(ctx *gin.Context, input schemas.SchemaTransaction, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "CustomerID",
					Message: "CustomerID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "CustomerID",
					Message: "CustomerID must be uuid",
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
					Field:   "Products",
					Message: "Products is required on body",
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
					Field:   "CustomerID",
					Message: "CustomerID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "CustomerID",
					Message: "CustomerID must be uuid",
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
					Field:   "Products",
					Message: "Products is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
