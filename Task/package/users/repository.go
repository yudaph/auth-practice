package users

import (
	"task/errs"
	"task/package/entities"
)

//Repository interface allows us to access the CRUD Operations in mysqlhere.
type Repository interface {
	FindBy(*map[string]interface{}) (*[]entities.User, *errs.AppError)
	FindOne(*map[string]interface{}) (*entities.User, *errs.AppError)
	Create(*entities.User) (*entities.User, *errs.AppError)
	Update(*map[string]interface{},*string) *errs.AppError
}