package helper

import (
	"sirclo/api/entities"
)

type RequestUserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RequestUserCreate struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Birth_date   string `json:"birth_date" form:"birth_date"`
	Phone_number int    `json:"phone_number" form:"phone_number"`
	Photo        string `json:"photo" form:"photo"`
	Gender       string `json:"gender" form:"gender"`
	Address      string `json:"address" form:"address"`
}

type RequestUserUpdate struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RequestEventCreate struct {
	Id_user     entities.User     `json:"id_uesr" form:"id_user"`
	Id_category entities.Category `json:"id_category" form:"id_category"`
	Title       string            `json:"title" form:"title"`
	Start_date  string            `json:"start_date" form:"start_date"`
	End_date    string            `json:"end_date" form:"end_date"`
	Location    string            `json:"location" form:"location"`
	Details     string            `json:"details" form:"details"`
	Photo       string            `json:"photo" form:"photo"`
}

type RequestProductUpdate struct {
	Id_user     entities.User     `json:"id_uesr" form:"id_user"`
	Id_category entities.Category `json:"id_category" form:"id_category"`
	Title       string            `json:"title" form:"title"`
	Start_date  string            `json:"start_date" form:"start_date"`
	End_date    string            `json:"end_date" form:"end_date"`
	Location    string            `json:"location" form:"location"`
	Details     string            `json:"details" form:"details"`
	Photo       string            `json:"photo" form:"photo"`
}
