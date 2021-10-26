package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/pkg"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}

/**
* ==========================================
* Repository Register Auth Teritory
*===========================================
 */

func (r *repositoryUser) EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role
	user.Active = true

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_register_01",
		}
		return &user, <-err
	}

	addNewUser := db.Debug().Create(&user).Commit()

	if addNewUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

/**
* ==========================================
* Repository Login Auth Teritory
*===========================================
 */

func (r *repositoryUser) EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser
	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_login_01",
		}
		return &user, <-err
	}

	checkPasswordMatch := pkg.ComparePassword(user.Password, input.Password)

	if checkPasswordMatch != nil {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_login_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}
