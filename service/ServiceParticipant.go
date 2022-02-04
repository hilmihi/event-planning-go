package service

import (
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/graph/model"
	"sirclo/api/repository"
)

type ServiceParticipant interface {
	ServiceParticipantCreate(input model.NewParticipant) (model.Participant, error)
}

type serviceParticipant struct {
	repository1 repository.RepositoryParticipant
}

func NewParticipantService(repository1 repository.RepositoryParticipant) *serviceParticipant {
	return &serviceParticipant{repository1}
}

func (s *serviceParticipant) ServiceParticipantCreate(input model.NewParticipant) (model.Participant, error) {
	participant := entities.Participant{
		Id_event: input.IDEvent,
		Id_user:  input.IDUser,
	}

	CreateParticipant, err := s.repository1.CreateParticipant(participant)
	if err != nil {
		fmt.Println(err)
		return model.Participant{}, err
	}

	output := model.Participant{
		ID:      CreateParticipant.Id,
		IDEvent: CreateParticipant.Id_event,
		IDUser:  CreateParticipant.Id_user,
	}

	return output, nil
}
