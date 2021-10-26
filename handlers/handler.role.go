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

type handleRole struct {
	role entitys.EntityRole
}

func NewHandlerRole(role entitys.EntityRole) *handleRole {
	return &handleRole{role: role}
}

/**
* ======================================
* Handler Ping Status Role Teritory
*=======================================
 */

func (h *handleRole) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Role", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Role Teritory
*======================================
 */

func (h *handleRole) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaRole
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := Role(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.role.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Role name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Role failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Role successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Role Teritory
*=======================================
 */

func (h *handleRole) HandlerResults(ctx *gin.Context) {
	res, error := h.role.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Role data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Role data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Role By ID Teritory
*=======================================
 */

func (h *handleRole) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaRole
	id := ctx.Param("id")
	body.ID = id

	errors, code := Role(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.role.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Role data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Role data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Role By ID Teritory
*=======================================
 */

func (h *handleRole) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaRole
	id := ctx.Param("id")
	body.ID = id

	errors, code := Role(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.role.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Role data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Role data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Role data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Role By ID Teritory
*=======================================
 */

func (h *handleRole) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaRole
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := Role(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.role.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Role data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Role data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Role data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Role
*=======================================
 */

func Role(ctx *gin.Context, input schemas.SchemaRole, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "RoleName",
					Message: "RoleName is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "lowercase",
					Field:   "RoleName",
					Message: "RoleName must be lowercase",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "RoleAccess",
					Message: "RoleAccess is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "lowercase",
					Field:   "RoleAccess",
					Message: "RoleAccess must be lowercase",
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
					Field:   "RoleName",
					Message: "RoleName is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "lowercase",
					Field:   "RoleName",
					Message: "RoleName must be lowercase",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "RoleAccess",
					Message: "RoleAccess is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "lowercase",
					Field:   "RoleAccess",
					Message: "RoleAccess must be lowercase",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
