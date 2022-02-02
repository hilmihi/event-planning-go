package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sirclo/api/graph/generated"
	"sirclo/api/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUser, id int) (*model.ResponseMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUserByID(ctx context.Context, id int) (*model.ResponseMessage, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
