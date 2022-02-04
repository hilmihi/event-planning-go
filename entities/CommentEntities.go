package entities

import "time"

type Comment struct {
	Id         int       `json:"id" form:"id"`
	Id_event   int       `json:"id_event" form:"id_event"`
	Id_user    int       `json:"id_user" form:"id_user"`
	Comment    string    `json:"comment" form:"comment"`
	Created_at time.Time `json:"created_at" form:"created_at"`
	Name       string    `json:"name" form:"name"`
	Email      string    `json:"email" form:"email"`
	Photo      string    `json:"photo" form:"photo"`
}
