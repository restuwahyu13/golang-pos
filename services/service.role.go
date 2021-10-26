package services

import (
	"github.com/restuwahyu13/golang-pos/entitys"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type serviceRole struct {
	role entitys.EntityRole
}

func NewServiceRole(role entitys.EntityRole) *serviceRole {
	return &serviceRole{role: role}
}

/**
* ==========================================
* Service Create New Role Teritory
*===========================================
 */

func (s *serviceRole) EntityCreate(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError) {
	var role schemas.SchemaRole
	role.RoleName = input.RoleName
	role.RoleAccess = input.RoleAccess

	res, err := s.role.EntityCreate(&role)
	return res, err
}

/**
* ==========================================
* Service Results All Role Teritory
*===========================================
 */

func (s *serviceRole) EntityResults() (*[]models.ModelRole, schemas.SchemaDatabaseError) {
	res, err := s.role.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Role By ID Teritory
*===========================================
 */

func (s *serviceRole) EntityResult(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError) {
	var role schemas.SchemaRole
	role.ID = input.ID

	res, err := s.role.EntityResult(&role)
	return res, err
}

/**
* ==========================================
* Service Delete Role By ID Teritory
*===========================================
 */

func (s *serviceRole) EntityDelete(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError) {
	var role schemas.SchemaRole
	role.ID = input.ID

	res, err := s.role.EntityDelete(&role)
	return res, err
}

/**
* ==========================================
* Service Update Role By ID Teritory
*===========================================
 */

func (s *serviceRole) EntityUpdate(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError) {
	var role schemas.SchemaRole
	role.RoleName = input.RoleName
	role.RoleAccess = input.RoleAccess

	res, err := s.role.EntityUpdate(&role)
	return res, err
}
