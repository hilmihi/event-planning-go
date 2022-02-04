package service

import (
	"sirclo/api/entities"
	"sirclo/api/graph/model"
	"sirclo/api/repository"
)

type ServiceEvent interface {
	ServiceEventsGet() ([]entities.Event, error)
	ServiceEventGet(id int) (model.EventDetail, error)
	ServiceEventCreate(input entities.Event) (model.Event, error)
	ServiceEventUpdate(id int, input model.NewEvent) error
	ServiceEventDelete(id int) (entities.Event, error)
}

type serviceEvent struct {
	repo repository.RepositoryEvent
}

func NewEventService(repo repository.RepositoryEvent) *serviceEvent {
	return &serviceEvent{repo: repo}
}

// get all event
func (se *serviceEvent) ServiceEventsGet() ([]entities.Event, error) {
	events, err := se.repo.GetEvents()
	if err != nil {
		return events, err
	}
	return events, nil
}

// get event by id
func (se *serviceEvent) ServiceEventGet(id int) (model.EventDetail, error) {
	event, err := se.repo.GetEvent(id)
	if err != nil {
		return model.EventDetail{}, err
	}

	output := model.EventDetail{
		ID:         event.Id,
		IDUser:     event.Id_user,
		IDCategory: event.Id_category,
		Title:      event.Title,
		StartDate:  event.Start_date,
		EndDate:    event.End_date,
		Location:   event.Location,
		Details:    event.Details,
		Photo:      &event.Photo,
		// Comments: ,
		// Participant: ,
	}

	return output, nil
}

// create event
func (se *serviceEvent) ServiceEventCreate(input entities.Event) (model.Event, error) {
	createEvent, err := se.repo.CreateEvent(input)
	if err != nil {
		return model.Event{}, err
	}

	output := model.Event{
		ID:         createEvent.Id,
		IDUser:     createEvent.Id_user,
		IDCategory: createEvent.Id_category,
		Title:      createEvent.Title,
		StartDate:  createEvent.Start_date,
		EndDate:    createEvent.End_date,
		Location:   createEvent.Location,
		Details:    createEvent.Details,
		Photo:      &createEvent.Photo,
	}
	return output, nil
}

// update event
func (se *serviceEvent) ServiceEventUpdate(id int, input model.NewEvent) error {
	event, err := se.repo.GetEvent(id)
	if err != nil {
		return err
	}

	event.Id_user = input.IDUser
	event.Id_category = input.IDCategory
	event.Title = input.Title
	event.Start_date = input.StartDate
	event.End_date = input.EndDate
	event.Location = input.Location
	event.Details = input.Details
	if input.Photo != nil {
		event.Photo = *input.Photo
	}

	_, err = se.repo.UpdateEvent(event)
	if err != nil {
		return err
	}
	return nil
}

// delete event
func (se *serviceEvent) ServiceEventDelete(id int) (entities.Event, error) {
	event, err := se.repo.GetEvent(id)
	if err != nil {
		return entities.Event{}, err
	}

	deleteEvent, err := se.repo.DeleteEvent(event)
	if err != nil {
		return deleteEvent, err
	}
	return deleteEvent, nil
}
