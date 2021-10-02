package users

import (
	"task/errs"
	"task/package/entities"
)

type Service interface {
	InsertUser(register *entities.UserRegister) (*entities.UserDto, *errs.AppError)
	GetUsers(where *map[string]interface{}) (*[]entities.UserDto, *errs.AppError)
	GetUser(where *map[string]interface{}) (*entities.UserDto, *errs.AppError)
	UpdateUser(set *map[string]interface{}, id *string) (*errs.AppError)
	RemoveUser(id *string) (*errs.AppError)
}

type service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertUser(register *entities.UserRegister) (*entities.UserDto, *errs.AppError) {
	
	if err:=register.Validate(); err!=nil{
		return nil, err
	}

	user := &entities.User{
		ID: register.CreateID(),
		Email: register.Email,
		Address: register.Address,
		Password: entities.EncryptPassword(&register.Password),
		Status: "on",
	}

	if _, err := s.repository.Create(user); err!=nil{
		return nil, err
	}

	return user.ToUserDto(), nil
}
func (s *service) GetUsers(where *map[string]interface{}) (*[]entities.UserDto, *errs.AppError) {
	users, err := s.repository.FindBy(where)
	if err!=nil{
		return nil, err
	}

	usersDto := make([]entities.UserDto, 0)
	for _,v :=range *users {
		usersDto = append(usersDto, *v.ToUserDto())
	}

	return &usersDto, nil
}

func (s *service) GetUser(where *map[string]interface{}) (*entities.UserDto, *errs.AppError) {
	user, err := s.repository.FindOne(where)
	if err!=nil{
		return nil, err
	}

	return user.ToUserDto(), nil
}

func (s *service) UpdateUser(set *map[string]interface{}, id *string) (*errs.AppError) {
	if _, err := s.repository.FindOne(&map[string]interface{}{
		"id": *id,
	}); err != nil {
		return err
	}

	//prevent update other than email & address
	for i := range *set{
		if i!="email" && i!="address" {
			delete(*set, i)
		}
	}

	if err := s.repository.Update(set, id); err!=nil{
		return err
	}

	return nil
}

func (s *service) RemoveUser(id *string) (*errs.AppError) {
	if _, err := s.repository.FindOne(&map[string]interface{}{
		"id": *id,
	}); err != nil {
		return err
	}

	if err := s.repository.Update(&map[string]interface{}{"status": "off"}, id); err!=nil{
		return err
	}

	return nil
}