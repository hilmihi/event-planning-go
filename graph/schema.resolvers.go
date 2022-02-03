package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/graph/generated"
	"sirclo/api/graph/model"
	"sirclo/api/helper"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	new := entities.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		Birth_date:   *input.BirthDate,
		Phone_number: *input.PhoneNumber,
		Photo:        *input.Photo,
		Gender:       *input.Gender,
		Address:      *input.Address,
	}

	resp, err := r.userService.ServiceUserCreate(new)

	return &resp, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUser, id int) (*model.ResponseMessage, error) {
	auth_user := ctx.Value("EchoContextKey").(int)
	if auth_user == 0 {
		return &model.ResponseMessage{Code: 400, Message: "Not Authorized"}, fmt.Errorf("Not Authorized")
	}

	user := entities.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		Birth_date:   *input.BirthDate,
		Phone_number: *input.PhoneNumber,
		Photo:        *input.Photo,
		Gender:       *input.Gender,
		Address:      *input.Address,
	}

	if id != auth_user {
		return &model.ResponseMessage{Code: 400, Message: "Not Allowed"}, fmt.Errorf("Not Authorized")
	}

	_, err := r.userService.ServiceUserUpdate(int(auth_user), user)
	if err != nil {
		fmt.Println(err)
		return &model.ResponseMessage{Code: 500, Message: "Internal Server Error"}, err
	}

	return &model.ResponseMessage{Code: 200, Message: "Succesfull Operation"}, err
}

func (r *mutationResolver) DeleteUserByID(ctx context.Context, id int) (*model.ResponseMessage, error) {
	auth_user := ctx.Value("EchoContextKey").(int)
	if auth_user == 0 {
		return &model.ResponseMessage{Code: 400, Message: "Not Authorized"}, fmt.Errorf("Not Authorized")
	}

	if id != auth_user {
		return &model.ResponseMessage{Code: 400, Message: "Not Allowed"}, fmt.Errorf("Not Authorized")
	}

	_, err := r.userService.ServiceUserDelete(int(auth_user))
	if err != nil {
		fmt.Println(err)
		return &model.ResponseMessage{Code: 500, Message: "Internal Server Error"}, err
	}

	return &model.ResponseMessage{Code: 200, Message: "Succesfull Operation"}, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	responseData, err := r.userService.ServiceUsersGet()

	if err != nil {
		return nil, err
	}

	userResponseData := []*model.User{}

	for _, v := range responseData {
		userResponseData = append(userResponseData, &model.User{ID: v.Id, Name: v.Name, Email: v.Email, Password: v.Password})
	}

	return userResponseData, nil
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.ResponseLogin, error) {
	input := helper.RequestUserLogin{
		Email:    email,
		Password: password,
	}

	token, err := r.userService.ServiceUserLoginGraph(input)
	if err != nil {
		fmt.Println("login: ", err)
		return &model.ResponseLogin{Code: 400, Token: "Failed Login!"}, err
	}

	return &model.ResponseLogin{Code: 200, Token: token}, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
