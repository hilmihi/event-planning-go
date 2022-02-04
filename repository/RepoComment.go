package repository

import (
	"database/sql"
	"fmt"
	"sirclo/api/entities"
)

type RepositoryComment interface {
	GetComments(int) ([]entities.Comment, error)
	CreateComment(user entities.Comment) (entities.Comment, error)
}

type Repository_Comment struct {
	db *sql.DB
}

func NewRepositoryComment(db *sql.DB) *Repository_Comment {
	return &Repository_Comment{db}
}

//get users
func (r *Repository_Comment) GetComments(id_event int) ([]entities.Comment, error) {
	var comments []entities.Comment
	results, err := r.db.Query(`select c.id, c.id_event, c.id_user, c.comment, c.created_at, u.name, u.email, u.photo
								from comment c
								join users u on u.id = c.id_user
								where c.deleted_at is null order by created_at asc`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var comment entities.Comment

		err = results.Scan(&comment.Id, &comment.Id_event, &comment.Id_user, &comment.Comment, &comment.Created_at, &comment.Name, &comment.Email, &comment.Photo)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		comments = append(comments, comment)
	}
	return comments, nil
}

//create user
func (r *Repository_Comment) CreateComment(comment entities.Comment) (entities.Comment, error) {
	query := `INSERT INTO comment (id_event, id_user, comment, created_at, updated_at) VALUES (?, ?, ?, now(), now())`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return comment, err
	}

	defer statement.Close()

	_, err = statement.Exec(comment.Id_event, comment.Id_user, comment.Comment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
