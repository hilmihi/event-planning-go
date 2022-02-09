package service

import (
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/graph/model"
	"sirclo/api/repository"
)

type ServiceEvent interface {
	ServiceEventsGet(int, int) ([]entities.Event, error)
	ServiceEventsByCategoryGet(int, int, int) ([]entities.Event, error)
	ServiceEventsPaginationGet(int, int) (model.Pagination, error)
	ServiceSearctEventsGet(string) ([]entities.Event, error)
	ServiceMyEventsGet(int) ([]entities.Event, error)
	ServiceEventsHistoryGet(int) ([]entities.Event, error)
	ServiceEventGet(id int) (model.EventDetail, error)
	ServiceEventCreate(input entities.Event) (model.Event, error)
	ServiceEventUpdate(id int, input model.NewEvent) (entities.Event, error)
	ServiceEventDelete(id int) (entities.Event, error)
}

type serviceEvent struct {
	repo            repository.RepositoryEvent
	repoComment     repository.RepositoryComment
	repoParticipant repository.RepositoryParticipant
}

func NewEventService(repo repository.RepositoryEvent, repoComment repository.RepositoryComment, repoParticipant repository.RepositoryParticipant) *serviceEvent {
	return &serviceEvent{repo: repo, repoComment: repoComment, repoParticipant: repoParticipant}
}

// get all event
func (se *serviceEvent) ServiceEventsGet(limit, offset int) ([]entities.Event, error) {
	events, err := se.repo.GetEvents(limit, offset)
	if err != nil {
		return events, err
	}
	return events, nil
}

// get events by category
func (se *serviceEvent) ServiceEventsByCategoryGet(limit, offset, id_category int) ([]entities.Event, error) {
	events, err := se.repo.GetEventsByCategory(limit, offset, id_category)
	if err != nil {
		return events, err
	}
	return events, nil
}

// pagination
func (se *serviceEvent) ServiceEventsPaginationGet(limit, offset int) (model.Pagination, error) {
	events, err := se.repo.GetEvents(limit, offset)
	if err != nil {
		return model.Pagination{}, err
	}

	total_page := (len(events) / limit) + 1

	eventResponseData := []*model.Event{}

	for _, v := range events {
		eventResponseData = append(eventResponseData, &model.Event{ID: v.Id, IDUser: &v.Id_user, IDCategory: v.Id_category, Title: v.Title, StartDate: v.Start_date, EndDate: v.End_date, Location: v.Location, Details: v.Details, Photo: &v.Photo})
	}

	pagination := model.Pagination{
		TotalPage: total_page,
		Data:      eventResponseData,
	}

	return pagination, nil
}

// search event by title
func (se *serviceEvent) ServiceSearctEventsGet(title string) ([]entities.Event, error) {
	events, err := se.repo.SearchEvents(title)
	if err != nil {
		return events, err
	}
	return events, nil
}

// get my event
func (se *serviceEvent) ServiceMyEventsGet(id_user int) ([]entities.Event, error) {
	events, err := se.repo.GetMyEvents(id_user)
	if err != nil {
		return events, err
	}
	return events, nil
}

// get history event
func (se *serviceEvent) ServiceEventsHistoryGet(id_user int) ([]entities.Event, error) {
	events, err := se.repo.GetMyEvents(id_user)
	if err != nil {
		return events, err
	}
	return events, nil
}

// get event by id
func (se *serviceEvent) ServiceEventGet(id int) (model.EventDetail, error) {
	event, err := se.repo.GetEvent(id)
	if err != nil {
		fmt.Println("ServiceEventGet-GetEvent: ", err)
		return model.EventDetail{}, err
	}

	comments, err := se.repoComment.GetComments(id)
	if err != nil {
		fmt.Println("ServiceEventGet-GetComments: ", err)
		return model.EventDetail{}, err
	}

	participants, err := se.repoParticipant.GetParticipant(id)
	if err != nil {
		fmt.Println("ServiceEventGet-GetParticipant: ", err)
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
	}

	for _, val := range comments {
		tempName := val.Name
		tempComment := val.Comment
		tempEmail := val.Email
		tempPhoto := val.Photo
		comment := model.Comment{
			ID:      val.Id,
			IDEvent: val.Id_event,
			IDUser:  val.Id_user,
			Comment: &tempComment,
			Name:    &tempName,
			Email:   &tempEmail,
			Photo:   &tempPhoto,
		}

		output.Comments = append(output.Comments, &comment)
	}

	for _, val := range participants {
		tempName := val.Name
		tempEmail := val.Email
		tempPhoto := val.Photo
		participant := model.Participant{
			ID:      val.Id,
			IDEvent: val.Id_event,
			IDUser:  val.Id_user,
			Name:    &tempName,
			Email:   &tempEmail,
			Photo:   &tempPhoto,
		}

		output.Participant = append(output.Participant, &participant)
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
		IDUser:     &createEvent.Id_user,
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
func (se *serviceEvent) ServiceEventUpdate(id int, input model.NewEvent) (entities.Event, error) {
	event, err := se.repo.GetEvent(id)
	if err != nil {
		return entities.Event{}, err
	}

	event.Id_category = input.IDCategory
	event.Title = input.Title
	event.Start_date = input.StartDate
	event.End_date = input.EndDate
	event.Location = input.Location
	event.Details = input.Details
	if input.Photo != nil {
		event.Photo = *input.Photo
	}

	updateEvent, errUpdate := se.repo.UpdateEvent(event)
	if errUpdate != nil {
		return updateEvent, err
	}
	return updateEvent, nil
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
