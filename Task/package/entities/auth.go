package entities

import (
	"net/http"
	"task/errs"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type ChangePassword struct{
	ID					 string
	OldPassword 		 string `json:"old-password"`
	Password 			 string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func (l *Login) Validate() *errs.AppError {
	if l.Email == "" {
		return errs.NewAppError("Email cannot be empty", http.StatusBadRequest)
	}
	if l.Password == "" {
		return errs.NewAppError("Password cannot be empty", http.StatusBadRequest)
	}

	return nil
}

func (c *ChangePassword) Validate() *errs.AppError {
	if c.OldPassword == "" {
		return errs.NewAppError("Old Password cannot be empty", http.StatusBadRequest)
	}
	
	if c.Password == "" {
		return errs.NewAppError("Password cannot be empty", http.StatusBadRequest)
	}
	if c.PasswordConfirmation == "" {
		return errs.NewAppError("Password Confirmation cannot be empty", http.StatusBadRequest)
	}
	if c.Password != c.PasswordConfirmation {
		return errs.NewAppError("Password and Password Confirmation did not match", http.StatusBadRequest)
	}

	return nil
}