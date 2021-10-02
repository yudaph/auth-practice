package auth

import (
	"net/http"
	"strings"
	"task/errs"
	"task/package/entities"
	"task/package/users"
)

type Service interface {
	Login(login *entities.Login) (*entities.LoginRes, *errs.AppError)
	ChangePassword(pw *entities.ChangePassword) (*errs.AppError)
	Auth(authorization *string) (ID *string, error *errs.AppError)
}

type service struct {
	repository users.Repository
}

//NewService is used to create a single instance of the service
func NewService(r users.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Login(login *entities.Login) (*entities.LoginRes, *errs.AppError){

	if err := login.Validate(); err!=nil {
		return nil, err
	}

	user, err := s.repository.FindOne(&map[string]interface{}{
		"email": login.Email,
		"password": entities.EncryptPassword(&login.Password),
		"status": "on",
	})
	if err!=nil {
		return nil, errs.NewAppError("Email or Password did not match", http.StatusBadRequest)
	}

	token, err := entities.CreateJWT(user.ToUserDto())
	if err!= nil {
		return nil, err
	}

	return &entities.LoginRes{
		Token: *token,
	},nil
}

func (s *service) Auth(authorization *string) (ID *string, error *errs.AppError){
	token := strings.Split(*authorization, " ")[1]
	if token == "" {
		return nil, errs.NewAppError("Please login", http.StatusUnauthorized)
	}

	userJwt, err := entities.ValidateJWT(&token)
	if err!=nil{
		return nil, err
	}

	if _, err:= s.repository.FindOne(&map[string]interface{}{
		"id": userJwt.ID,
		"status": "on",
	}); err!=nil{
		return nil, err
	}

	return &userJwt.ID, nil
}

func (s *service) ChangePassword(pw *entities.ChangePassword) (*errs.AppError) {
	
	if err := pw.Validate(); err!=nil{
		return err
	}

	//check user 
	_, err := s.repository.FindOne(&map[string]interface{}{
		"id": pw.ID,
		"password": entities.EncryptPassword(&pw.OldPassword),
		"status": "on",
	})
	if err!=nil {
		return err
	}

	if err = s.repository.Update(&map[string]interface{}{
		"password": entities.EncryptPassword(&pw.Password),
	}, &pw.ID); err!=nil{
		return err
	}
	
	return nil
}