package repository

import (
	"database/sql"
	"sirclo/api/entities"
)

type RepositoryEvent interface {
	GetEvents() ([]entities.Event, error)
	CreateEvent(event entities.Event) (entities.Event, error)
	GetEvent(id int) (entities.Event, error)
	UpdateEvent(event entities.Event) (entities.Event, error)
	DeleteEvent(event entities.Event) (entities.Event, error)
}

type Repository_Event struct {
	db *sql.DB
}

func NewRepositoryEvent(db *sql.DB) *Repository_Event {
	return &Repository_Event{db: db}
}

// get events
func (re *Repository_Event) GetEvents() ([]entities.Event, error) {
	var events []entities.Event
	result, err := re.db.Query(`select id, id_user, id_category, title, start_date, end_date, location, details, photo, from event WHERE deleted_ad IS NULL`)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event entities.Event
		err = result.Scan(&event.Id, &event.Id_user, &event.Id_category, &event.Title, &event.Start_date, &event.End_date, &event.Location, &event.Details, &event.Photo)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}
