package services

import (
	"github.com/restuwahyu13/golang-pos/entities"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemes"
)

type serviceRole struct {
	role entities.EntityRole
}

func NewServiceRole(role entities.EntityRole) *serviceRole {
	return &serviceRole{role: role}
}

/**
* ==========================================
* Service Create New Role Teritory
*===========================================
 */

func (s *serviceRole) EntityCreate(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError) {
	var role schemes.SchemeRole
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

func (s *serviceRole) EntityResults() (*[]models.ModelRole, schemes.SchemeDatabaseError) {
	res, err := s.role.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Role By ID Teritory
*===========================================
 */

func (s *serviceRole) EntityResult(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError) {
	var role schemes.SchemeRole
	role.ID = input.ID

	res, err := s.role.EntityResult(&role)
	return res, err
}

/**
* ==========================================
* Service Delete Role By ID Teritory
*===========================================
 */

func (s *serviceRole) EntityDelete(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError) {
	var role schemes.SchemeRole
	role.ID = input.ID

	res, err := s.role.EntityDelete(&role)
	return res, err
}

/**
* ==========================================
* Service Update Role By ID Teritory
*===========================================
 */

func (s *serviceRole) EntityUpdate(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError) {
	var role schemes.SchemeRole
	role.RoleName = input.RoleName
	role.RoleAccess = input.RoleAccess

	res, err := s.role.EntityUpdate(&role)
	return res, err
}
