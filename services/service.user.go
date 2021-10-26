package services

import (
	"github.com/restuwahyu13/golang-pos/entitys"
	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type serviceUser struct {
	user entitys.EntityUser
}

func NewServiceUser(user entitys.EntityUser) *serviceUser {
	return &serviceUser{user: user}
}

func (s *serviceUser) EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaUser
	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password
	schema.Role = input.Role

	res, err := s.user.EntityRegister(&schema)
	return res, err
}

func (s *serviceUser) EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaUser
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.user.EntityLogin(&schema)
	return res, err
}
