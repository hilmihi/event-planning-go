package repository

import (
	"database/sql"
	"sirclo/api/entities"
	"strings"
)

type RepositoryEvent interface {
	GetEvents(int, int) ([]entities.Event, error)
	GetEventsByCategory(int, int, int) ([]entities.Event, error)
	SearchEvents(string) ([]entities.Event, error)
	GetMyEvents(int) ([]entities.Event, error)
	GetEventsHistory(int) ([]entities.Event, error)
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
func (re *Repository_Event) GetEvents(limit, offset int) ([]entities.Event, error) {
	var events []entities.Event
	result, err := re.db.Query(`select e.id, e.id_user, e.id_category, e.title, e.start_date, e.end_date, e.location, e.details, e.photo, u.name
								from event e 
								join users u on u.id = e.id_user and u.deleted_at is null
								WHERE e.deleted_at IS NULL
								LIMIT ?, ?`, offset, limit)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event entities.Event
		err = result.Scan(&event.Id, &event.Id_user, &event.Id_category, &event.Title, &event.Start_date, &event.End_date, &event.Location, &event.Details, &event.Photo, &event.HostedBy)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func (re *Repository_Event) GetEventsByCategory(limit, offset, id_category int) ([]entities.Event, error) {
	var events []entities.Event
	result, err := re.db.Query(`select e.id, e.id_user, e.id_category, e.title, e.start_date, e.end_date, e.location, e.details, e.photo, u.name
								from event e 
								join users u on u.id = e.id_user and u.deleted_at is null
								WHERE e.deleted_at IS NULL AND e.id_category = ?
								LIMIT ?, ?`, id_category, offset, limit)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event entities.Event
		err = result.Scan(&event.Id, &event.Id_user, &event.Id_category, &event.Title, &event.Start_date, &event.End_date, &event.Location, &event.Details, &event.Photo, &event.HostedBy)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func (re *Repository_Event) SearchEvents(title string) ([]entities.Event, error) {
	var events []entities.Event
	result, err := re.db.Query(`select e.id, e.id_user, e.id_category, e.title, e.start_date, e.end_date, e.location, e.details, e.photo, u.name
								from event e 
								join users u on u.id = e.id_user and u.deleted_at is null
								WHERE e.deleted_at IS NULL AND LOWER(e.title) LIKE ?`, "%"+strings.ToLower(title)+"%")
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event entities.Event
		err = result.Scan(&event.Id, &event.Id_user, &event.Id_category, &event.Title, &event.Start_date, &event.End_date, &event.Location, &event.Details, &event.Photo, &event.HostedBy)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

// get my events
func (re *Repository_Event) GetMyEvents(id_user int) ([]entities.Event, error) {
	var events []entities.Event
	result, err := re.db.Query(`select e.id, e.id_user, e.id_category, e.title, e.start_date, e.end_date, e.location, e.details, e.photo, u.name
								from event e 
								join users u on u.id = e.id_user and u.deleted_at is null
								WHERE e.deleted_at IS NULL AND e.id_user = ?`, id_user)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event entities.Event
		err = result.Scan(&event.Id, &event.Id_user, &event.Id_category, &event.Title, &event.Start_date, &event.End_date, &event.Location, &event.Details, &event.Photo, &event.HostedBy)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

// get events history
func (re *Repository_Event) GetEventsHistory(id_user int) ([]entities.Event, error) {
	var events []entities.Event
	result, err := re.db.Query(`select e.id, e.id_user, e.id_category, e.title, e.start_date, e.end_date, e.location, e.details, e.photo, u.name
								from event e 
								join users u on u.id = e.id_user and u.deleted_at is null
								join participant p on p.id_event = e.id
								WHERE e.deleted_at IS NULL AND p.id_user = ?`, id_user)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event entities.Event
		err = result.Scan(&event.Id, &event.Id_user, &event.Id_category, &event.Title, &event.Start_date, &event.End_date, &event.Location, &event.Details, &event.Photo, &event.HostedBy)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

// create event
func (re *Repository_Event) CreateEvent(event entities.Event) (entities.Event, error) {
	query := `INSERT INTO event (id_user, id_category, title, start_date, end_date, location, details, photo, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, now())`

	statement, err := re.db.Prepare(query)
	if err != nil {
		return event, err
	}
	defer statement.Close()

	_, err = statement.Exec(event.Id_user, event.Id_category, event.Title, event.Start_date, event.End_date, event.Location, event.Details, event.Photo)
	if err != nil {
		return event, err
	}
	return event, nil
}

// get event by id
func (re *Repository_Event) GetEvent(id int) (entities.Event, error) {
	var event entities.Event
	result := re.db.QueryRow(`select e.id, e.id_user, e.id_category, e.title, e.start_date, e.end_date, e.location, e.details, e.photo, u.name
								from event e 
								join users u on u.id = e.id_user and u.deleted_at is null
								WHERE e.deleted_at IS NULL AND e.id=? `, id)

	err := result.Scan(&event.Id, &event.Id_user, &event.Id_category, &event.Title, &event.Start_date, &event.End_date, &event.Location, &event.Details, &event.Photo, &event.HostedBy)
	if err != nil {
		return event, err
	}

	return event, nil
}

// Update Event
func (re *Repository_Event) UpdateEvent(event entities.Event) (entities.Event, error) {
	query := `UPDATE event SET id = ?, id_user = ?, id_category = ?, title = ?, start_date = ?, end_date = ?, location = ?, details = ?, photo = ?, updated_at = now() WHERE id = ?`

	statement, err := re.db.Prepare(query)
	if err != nil {
		return event, err
	}
	defer statement.Close()

	_, err = statement.Exec(event.Id, event.Id_user, event.Id_category, event.Title, event.Start_date, event.End_date, event.Location, event.Details, event.Photo, event.Id)
	if err != nil {
		return event, err
	}
	return event, nil
}

// delete event
func (re *Repository_Event) DeleteEvent(event entities.Event) (entities.Event, error) {
	query := `UPDATE event SET deleted_at = now() WHERE id = ?`

	statement, err := re.db.Prepare(query)
	if err != nil {
		return event, err
	}
	defer statement.Close()

	_, err = statement.Exec(event.Id)
	if err != nil {
		return event, err
	}

	return event, nil
}
