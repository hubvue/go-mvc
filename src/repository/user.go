package repository

import (
	"database/sql"
	"fmt"
	"server-test/src/database"
	"server-test/src/model"
)

type UserRepo interface {
	GetUserById(id int) *model.User
}

type userRepo struct {
	mysql *sql.DB
}

func NewUserRepo() *userRepo {
	return &userRepo{
		mysql: database.NewMysql(),
	}
}

func (repo *userRepo) GetUserById(id int) *model.User {
	var user model.User
	row := repo.mysql.QueryRow(" select * from user where id = ?", id)
	err := row.Scan(&user.Id, &user.Name, &user.UserNumber, &user.Email, &user.ImageUrl)
	if err != nil {
		fmt.Println("query fail", err)
		return nil
	}
	return &user
}