package helper

import (
	"sirclo/api/entities"
)

type AuthFormat struct {
	Token string `json:"token"`
}

func FormatAuth(user entities.User, token string) AuthFormat {
	formatter := AuthFormat{
		Token: token,
	}
	return formatter
}

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseProduct2 struct {
	Id          int     `json:"id" form:"id"`
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price"`
	Quantity    int     `json:"quantity" form:"quantity"`
}

func FormatUser(user entities.User) UserFormatter {
	formatter := UserFormatter{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	return formatter
}
