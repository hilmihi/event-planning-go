package service

import (
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/graph/model"
	"sirclo/api/repository"
)

type ServiceParticipant interface {
	ServiceParticipantCreate(input model.NewParticipant) (model.Participant, error)
	ServiceParticipantsGet(int) ([]*model.Participant, error)
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

	// cek jika participant sudah pernah join event
	existingParticipant, errGet := s.repository1.GetParticipants(input.IDEvent)
	if errGet != nil {
		return model.Participant{}, errGet
	}

	for _, v := range existingParticipant {
		if v == participant {
			return model.Participant{}, fmt.Errorf("Failed to Join, You Already Joined This Event!")
		}
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

func (su *serviceParticipant) ServiceParticipantsGet(id_event int) ([]*model.Participant, error) {
	comments, err := su.repository1.GetParticipant(id_event)
	if err != nil {
		return nil, err
	}

	var res []*model.Participant
	for _, val := range comments {
		name := val.Name
		email := val.Email
		photo := val.Photo
		comment := model.Participant{
			ID:      val.Id,
			IDEvent: val.Id_event,
			IDUser:  val.Id_user,
			Name:    &name,
			Email:   &email,
			Photo:   &photo,
		}
		res = append(res, &comment)
	}
	return res, nil
}
