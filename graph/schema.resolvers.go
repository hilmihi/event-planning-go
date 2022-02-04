package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/graph/generated"
	"sirclo/api/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	new := &entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	resp, err := r.userService.ServiceUserCreate(*new)

	if err != nil {
		fmt.Println("graph controller create user:", err)
		return &model.User{}, err
	}

	return &resp, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUser, id int) (*model.ResponseMessage, error) {
	auth_user, bol := ctx.Value("EchoContextKey").(int)
	if bol != true {
		return &model.ResponseMessage{Code: 400, Message: "Not Authorized"}, fmt.Errorf("Not Authorized")
	}

	if id != auth_user {
		return &model.ResponseMessage{Code: 400, Message: "Not Allowed"}, fmt.Errorf("Not Authorized")
	}

	err := r.userService.ServiceUserUpdate(int(auth_user), input)
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

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, input model.NewEvent, id int) (*model.ResponseMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEventByID(ctx context.Context, id int) (*model.ResponseMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	auth_user := ctx.Value("EchoContextKey").(int)
	if auth_user == 0 {
		return &model.Comment{}, fmt.Errorf("Not Authorized")
	}

	resp, err := r.commentService.ServiceCommentCreate(input)

	if err != nil {
		fmt.Println("graph controller create Comment:", err)
		return &model.Comment{}, err
	}

	return &resp, err
}

func (r *mutationResolver) DeleteCommentByID(ctx context.Context, id int) (*model.ResponseMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateParticipant(ctx context.Context, input model.NewParticipant) (*model.Participant, error) {
	auth_user := ctx.Value("EchoContextKey").(int)
	if auth_user == 0 {
		return &model.Participant{}, fmt.Errorf("Not Authorized")
	}

	resp, err := r.participantService.ServiceParticipantCreate(input)

	if err != nil {
		fmt.Println("graph controller create Comment:", err)
		return &model.Participant{}, err
	}

	return &resp, err
}

func (r *mutationResolver) UpdateParticipant(ctx context.Context, input model.NewParticipant, id int) (*model.ResponseMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteParticipantByID(ctx context.Context, id int) (*model.ResponseMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.ResponseLogin, error) {
	token, err := r.userService.ServiceUserLoginGraph(email, password)
	if err != nil {
		fmt.Println("login: ", err)
		return &model.ResponseLogin{Code: 400, Token: "Failed Login!"}, err
	}

	return &model.ResponseLogin{Code: 200, Token: token}, err
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

func (r *queryResolver) UsersByID(ctx context.Context, id *int) (*model.User, error) {
	responseData, err := r.userService.ServiceUserGet(*id)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EventsByID(ctx context.Context, id *int) (*model.EventDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Category(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comments(ctx context.Context, idEvent int) ([]*model.Comment, error) {
	auth_user := ctx.Value("EchoContextKey").(int)
	if auth_user == 0 {
		return nil, fmt.Errorf("Not Authorized")
	}

	responseData, err := r.commentService.ServiceCommentsGet(idEvent)

	if err != nil {
		return nil, err
	}

	return responseData, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	responseData, err := r.userService.ServiceUserGet(id)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
