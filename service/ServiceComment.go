package service

import (
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/graph/model"
	"sirclo/api/repository"
)

type ServiceComment interface {
	ServiceCommentsGet(int) ([]*model.Comment, error)
	ServiceCommentCreate(input model.NewComment) (model.Comment, error)
}

type serviceComment struct {
	repository1 repository.RepositoryComment
}

func NewCommentService(repository1 repository.RepositoryComment) *serviceComment {
	return &serviceComment{repository1}
}

func (su *serviceComment) ServiceCommentsGet(id_event int) ([]*model.Comment, error) {
	comments, err := su.repository1.GetComments(id_event)
	if err != nil {
		return nil, err
	}

	var res []*model.Comment
	for _, val := range comments {
		com := val.Comment
		date := val.Created_at.String()
		name := val.Name
		email := val.Email
		photo := val.Photo
		comment := model.Comment{
			ID:        val.Id,
			IDEvent:   val.Id_event,
			IDUser:    val.Id_user,
			Comment:   &com,
			CreatedAt: &date,
			Name:      &name,
			Email:     &email,
			Photo:     &photo,
		}
		res = append(res, &comment)
	}
	return res, nil
}
func (s *serviceComment) ServiceCommentCreate(input model.NewComment) (model.Comment, error) {
	comment := entities.Comment{
		Id_event: input.IDEvent,
		Id_user:  input.IDUser,
	}

	if input.Comment == nil {
		return model.Comment{}, fmt.Errorf("Need value of comment!")
	}

	comment.Comment = *input.Comment

	createComment, err := s.repository1.CreateComment(comment)
	if err != nil {
		fmt.Println(err)
		return model.Comment{}, err
	}

	output := model.Comment{
		ID:      createComment.Id,
		IDEvent: createComment.Id_event,
		IDUser:  createComment.Id_user,
		Comment: &createComment.Comment,
	}

	return output, nil
}
