package users

import (
	"net/http"
	"reflect"
	"strings"
	"task/errs"
	"task/package/entities"
	"task/zapLog"

	"github.com/jmoiron/sqlx"
)

type repositoryMysql struct {
	client *sqlx.DB
}

//NewRepo is the single instance repo that is being created.
func NewRepoMySQL(client *sqlx.DB) Repository {
	return &repositoryMysql{
		client: client,
	}
}

func (rm *repositoryMysql) FindBy(where *map[string]interface{}) (*[]entities.User, *errs.AppError) {
	items := make([]entities.User, 0)
	whereSQL := ""
	value := make([]interface{}, 0)
	for k, v := range *where {
		if whereSQL != "" {
			whereSQL += " AND "
		} else {
			whereSQL += "WHERE "
		}
		temp := ""
		if reflect.TypeOf(v).String() == "string" {
			s := strings.Split(v.(string), " ")
			if len(s) > 1 {
				f := strings.ToLower(s[0])
				if f == "between" || f == "like" {
					temp = k + " " + v.(string)
				}
			}
		}
		if temp == "" {
			temp = k + "=?"
			value = append(value, v)
		}

		whereSQL += temp
	}

	findSQL := "select * from users " + whereSQL
	//zapLog.Info(findSQL)
	err := rm.client.Select(&items, findSQL, value...)
	if err != nil {
		zapLog.Error("Error SQL " + err.Error())
		return nil, errs.NewAppError("Error", http.StatusInternalServerError)
	}
	return &items, nil
}

func (rm *repositoryMysql) FindOne(where *map[string]interface{}) (*entities.User, *errs.AppError) {
	items, err := rm.FindBy(where)
	if err != nil {
		return nil, errs.NewAppError("Error", http.StatusInternalServerError)
	}
	res := *items
	if len(res) <= 0 {
		return nil, errs.NewAppError("User not found", http.StatusNotFound)
	}

	return &res[0], nil
}

func (rm *repositoryMysql) Create(user *entities.User) (*entities.User, *errs.AppError) {
	insertSQL := "insert into users(`id`,`email`,`address`,`password`) values(?,?,?,?)"
	_, err := rm.client.Exec(insertSQL, user.ID, user.Email, user.Address, user.Password)
	if err != nil {
		zapLog.Error("Error SQL " + err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, errs.NewAppError("Email Already Exist", http.StatusBadRequest)
		}
		return nil, errs.NewAppError("Internal Server Error", http.StatusInternalServerError)
	}
	return user, nil
}

func (rm *repositoryMysql) Update(set *map[string]interface{}, id *string) *errs.AppError {
	setSQL := ""
	value := make([]interface{}, 0)
	for k, v := range *set {
		if setSQL != "" {
			setSQL += ", "
		}

		value = append(value, v)
		setSQL += k + "=?"
	}
	value = append(value, *id)

	updateSQL := "UPDATE users set " + setSQL + " where id = ?"
	_, err := rm.client.Exec(updateSQL, value...)
	if err != nil {
		zapLog.Error("Error SQL " + err.Error())
		return errs.NewAppError("Error", http.StatusInternalServerError)
	}
	return nil
}