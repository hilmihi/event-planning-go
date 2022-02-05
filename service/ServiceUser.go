package service

import (
	"fmt"
	addmiddleware "sirclo/api/addMiddleware"
	"sirclo/api/entities"
	"sirclo/api/graph/model"
	"sirclo/api/helper"
	"sirclo/api/repository"
)

type ServiceUser interface {
	ServiceUserLogin(input helper.RequestUserLogin) (entities.User, error)
	ServiceUserLoginGraph(string, string) (int, string, error)
	ServiceUsersGet() ([]entities.User, error)
	ServiceUserGet(id int) (model.User, error)
	ServiceUserCreate(input entities.User) (model.User, error)
	ServiceUserUpdate(id int, input model.NewUser) error
	ServiceUserDelete(id int) (entities.User, error)
}

type serviceUser struct {
	repository1 repository.RepositoryUser
}

func NewUserService(repository1 repository.RepositoryUser) *serviceUser {
	return &serviceUser{repository1}
}

func (su *serviceUser) ServiceUserLogin(input helper.RequestUserLogin) (entities.User, error) {
	email := input.Email
	password := input.Password

	var user entities.User
	user, err := su.repository1.FindByEmail(email)
	if err != nil {
		return user, err
	}

	match, err := helper.CheckPasswordHash(password, user.Password)
	if err != nil {
		return user, err
	}

	if !match {
		return user, fmt.Errorf("Email atau Password Anda Salah!")
	}

	return user, nil
}

func (su *serviceUser) ServiceUserLoginGraph(email, password string) (int, string, error) {
	var user entities.User
	user, err := su.repository1.FindByEmail(email)
	if err != nil {
		return 0, "", err
	}

	match, err := helper.CheckPasswordHash(password, user.Password)
	if err != nil {
		return 0, "", err
	}

	if !match {
		return 0, "", fmt.Errorf("Email atau Password Anda Salah!")
	}

	token, err := addmiddleware.GenerateToken(user.Id)

	return user.Id, token, nil
}

func (su *serviceUser) ServiceUsersGet() ([]entities.User, error) {
	users, err := su.repository1.GetUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *serviceUser) ServiceUserGet(id int) (model.User, error) {
	user, err := s.repository1.GetUser(id)
	if err != nil {
		return model.User{}, err
	}

	output := model.User{
		ID:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		BirthDate:   &user.Birth_date,
		PhoneNumber: &user.Phone_number,
		Photo:       &user.Phone_number,
		Gender:      &user.Gender,
		Address:     &user.Address,
	}

	return output, nil
}

func (s *serviceUser) ServiceUserCreate(input entities.User) (model.User, error) {
	var err error
	user, _ := s.repository1.FindByEmail(input.Email)
	if user.Email == input.Email {
		return model.User{}, fmt.Errorf("Email sudah terdaftar")
	}

	input.Password, err = helper.HashPassword(input.Password)

	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}

	createUser, err := s.repository1.CreateUser(input)
	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}

	output := model.User{
		ID:          createUser.Id,
		Name:        createUser.Name,
		Email:       createUser.Email,
		Password:    createUser.Password,
		BirthDate:   &createUser.Birth_date,
		PhoneNumber: &createUser.Phone_number,
		Photo:       &createUser.Phone_number,
		Gender:      &createUser.Gender,
		Address:     &createUser.Address,
	}

	return output, nil
}

func (s *serviceUser) ServiceUserUpdate(id int, input model.NewUser) error {
	user, err := s.repository1.GetUser(id)
	if err != nil {
		return err
	}

	user.Name = input.Name
	user.Password = input.Password
	user.Email = input.Email
	if input.BirthDate != nil {
		user.Birth_date = *input.BirthDate
	}
	if input.PhoneNumber != nil {
		user.Phone_number = *input.PhoneNumber
	}
	if input.Photo != nil {
		user.Photo = *input.Photo
	}
	if input.Gender != nil {
		user.Gender = *input.Gender
	}
	if input.Address != nil {
		user.Address = *input.Address
	}

	_, err = s.repository1.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceUser) ServiceUserDelete(id int) (entities.User, error) {
	userID, err := s.repository1.GetUser(id)
	if err != nil {
		return entities.User{}, err
	}

	deleteUser, err := s.repository1.DeleteUser(userID)
	if err != nil {
		return deleteUser, err
	} else {
		return deleteUser, nil
	}
}
