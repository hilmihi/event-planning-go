package entities

type Participant struct {
	Id       int `json:"id" form:"id"`
	Id_event int `json:"id_event" form:"id_event"`
	Id_user  int `json:"id_user" form:"id_user"`
}
