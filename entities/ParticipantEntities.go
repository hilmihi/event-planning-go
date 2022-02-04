package entities

type Participant struct {
	Id       int    `json:"id" form:"id"`
	Id_event int    `json:"id_event" form:"id_event"`
	Id_user  int    `json:"id_user" form:"id_user"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Photo    string `json:"photo" form:"photo"`
}
