package repository

import (
	"database/sql"
	"fmt"
	"sirclo/api/entities"
)

type RepositoryParticipant interface {
	GetParticipants(int) ([]entities.Participant, error)
	GetParticipant(int) ([]entities.Participant, error)
	CreateParticipant(user entities.Participant) (entities.Participant, error)
}

type Repository_Participant struct {
	db *sql.DB
}

func NewRepositoryParticipant(db *sql.DB) *Repository_Participant {
	return &Repository_Participant{db}
}

//get users
func (r *Repository_Participant) GetParticipant(id_event int) ([]entities.Participant, error) {
	var participants []entities.Participant
	results, err := r.db.Query(`select p.id, p.id_user, p.id_event, u.name, u.email, u.photo
								from participant p
								join users u on u.id = p.id_user AND u.deleted_at is null
								where p.deleted_at is null and p.id_event = ? order by p.created_at asc`, id_event)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var participant entities.Participant

		err = results.Scan(&participant.Id, &participant.Id_user, &participant.Id_event, &participant.Name, &participant.Email, &participant.Photo)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		participants = append(participants, participant)
	}
	return participants, nil
}

//create user
func (r *Repository_Participant) CreateParticipant(participant entities.Participant) (entities.Participant, error) {
	query := `INSERT INTO participant (id_event, id_user, created_at, updated_at) VALUES (?, ?, now(), now())`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return participant, err
	}

	defer statement.Close()

	_, err = statement.Exec(participant.Id_event, participant.Id_user)
	if err != nil {
		return participant, err
	}

	return participant, nil
}

// get id_user participant
func (r *Repository_Participant) GetParticipants(idEvent int) ([]entities.Participant, error) {
	var participants []entities.Participant
	results, err := r.db.Query(`select id_event, id_user from participant where deleted_at is null AND id_event = ?`, idEvent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var participant entities.Participant

		err = results.Scan(&participant.Id_event, &participant.Id_user)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		participants = append(participants, participant)
	}
	return participants, nil
}
