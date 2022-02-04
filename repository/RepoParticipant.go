package repository

import (
	"database/sql"
	"fmt"
	"sirclo/api/entities"
)

type RepositoryParticipant interface {
	GetComments(int) ([]entities.Participant, error)
	CreateParticipant(user entities.Participant) (entities.Participant, error)
}

type Repository_Participant struct {
	db *sql.DB
}

func NewRepositoryParticipant(db *sql.DB) *Repository_Participant {
	return &Repository_Participant{db}
}

//get users
func (r *Repository_Participant) GetComments(id_event int) ([]entities.Participant, error) {
	var comments []entities.Participant
	results, err := r.db.Query("select id, id_event, id_user, comment, created_at from comment where deleted_at is null order by created_at asc")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var comment entities.Participant

		err = results.Scan(&comment.Id, &comment.Id_event, &comment.Id_user)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		comments = append(comments, comment)
	}
	return comments, nil
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
