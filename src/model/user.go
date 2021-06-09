package model

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	UserNumber string `json:"user_number"`
	Email      string `json:"email"`
	ImageUrl   string `json:"image_url"`
}
