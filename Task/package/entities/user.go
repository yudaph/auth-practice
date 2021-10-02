package entities

import (
	"net/http"
	"task/errs"

	"github.com/segmentio/ksuid"
)

type User struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Address  string `db:"address"`
	Password string `db:"password"`
	Status 	 string `db:"status"`
}

type UserRegister struct {
	Email 				 string `json:"email"`
	Address 			 string `json:"address"`
	Password 			 string `json:"password"`
	PasswordConfirmation string `json:"password-confirmation"`
}

type UserDto struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

func (ur *UserRegister) CreateID() string {
	return ksuid.New().String()
}

func (u *User) ToUserDto() *UserDto {
	return &UserDto{
		ID: u.ID,
		Email: u.Email,
		Address: u.Address,
	}
}

func (ur *UserRegister) Validate() *errs.AppError{
	if ur.Email == "" {
		return errs.NewAppError("Email cannot be empty" , http.StatusBadRequest)
	}
	if ur.Address == "" {
		return errs.NewAppError("Address cannot be empty" , http.StatusBadRequest)
	}
	if ur.Password == "" {
		return errs.NewAppError("Password cannot be empty" , http.StatusBadRequest)
	}
	if ur.PasswordConfirmation == "" {
		return errs.NewAppError("Password Confirmation cannot be empty" , http.StatusBadRequest)
	}
	if ur.Password != ur.PasswordConfirmation {
		return errs.NewAppError("Password and Password Confirmation did not match" , http.StatusBadRequest)
	}
	return nil
}
