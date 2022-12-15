package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/helpers"
	"github.com/restuwahyu13/golang-pos/pkg"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type handleCustomer struct {
	customer entities.EntityCustomer
}

func NewHandlerCustomer(customer entities.EntityCustomer) *handleCustomer {
	return &handleCustomer{customer: customer}
}

/**
* ======================================
* Handler Ping Status Customer Teritory
*=======================================
 */

func (h *handleCustomer) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Customer", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Customer Teritory
*======================================
 */

func (h *handleCustomer) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeCustomer
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorCustomer(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.customer.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Customer name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Customer failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Customer successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Customer Teritory
*=======================================
 */

func (h *handleCustomer) HandlerResults(ctx *gin.Context) {
	res, error := h.customer.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Customer data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Customer data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Customer By ID Teritory
*=======================================
 */

func (h *handleCustomer) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeCustomer
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorCustomer(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.customer.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Customer data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Customer data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Customer By ID Teritory
*=======================================
 */

func (h *handleCustomer) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeCustomer
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorCustomer(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.customer.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Customer data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Customer data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Customer data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Customer By ID Teritory
*=======================================
 */

func (h *handleCustomer) HandlerUpdate(ctx *gin.Context) {
	var body schemes.SchemeCustomer
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorCustomer(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.customer.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Customer data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Customer data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Customer data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Customer
*=======================================
 */

func ValidatorCustomer(ctx *gin.Context, input schemes.SchemeCustomer, Type string) (interface{}, int) {
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
