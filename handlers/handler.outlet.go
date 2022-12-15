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

type handlerOutlet struct {
	outlet entities.EntityOutlet
}

func NewHandlerOutlet(outlet entities.EntityOutlet) *handlerOutlet {
	return &handlerOutlet{outlet: outlet}
}

/**
* ======================================
* Handler Ping Status Outlet Teritory
*=======================================
 */

func (h *handlerOutlet) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Outlet", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Outlet Teritory
*======================================
 */

func (h *handlerOutlet) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeOutlet
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorOutlet(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.outlet.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Outlet name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Outlet phone number already taken", error.Code, nil)
		return
	}

	if error.Type == "error_create_03" {
		helpers.APIResponse(ctx, "Create new Outlet failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Outlet successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Outlet Teritory
*=======================================
 */

func (h *handlerOutlet) HandlerResults(ctx *gin.Context) {
	res, error := h.outlet.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Outlet data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Outlet data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Outlet By ID Teritory
*=======================================
 */

func (h *handlerOutlet) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeOutlet
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorOutlet(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.outlet.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Outlet data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Outlet data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Outlet By ID Teritory
*=======================================
 */

func (h *handlerOutlet) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeOutlet
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorOutlet(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.outlet.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Outlet data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Outlet data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Outlet data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Outlet By ID Teritory
*=======================================
 */

func (h *handlerOutlet) HandlerUpdate(ctx *gin.Context) {
	var body schemes.SchemeOutlet
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorOutlet(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.outlet.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Outlet data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Outlet data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Outlet data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Outlet
*=======================================
 */

func ValidatorOutlet(ctx *gin.Context, input schemes.SchemeOutlet, Type string) (interface{}, int) {
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
					Field:   "MerchantID",
					Message: "MerchantID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "MerchantID value must be uuid",
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
					Field:   "MerchantID",
					Message: "MerchantID is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "MerchantID value must be uuid",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
