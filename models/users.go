package models

import (
	"github.com/wooogi123/gin-mysql-boilerplate/db"
	"github.com/wooogi123/gin-mysql-boilerplate/forms"
)

type User struct {
	Id           int    `json:"id"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Group        string `json:"group"`
	Star         int    `json:"star"`
	Finish_count int    `json:"finish_count"`
	Point        int    `json:"point"`
}

func (u User) SignIn(userPayload forms.UserSignIn) (user User, err error) {
	db := db.GetDB()
	user = User{
		Email:    userPayload.Email,
		Password: userPayload.Password,
	}

	row := db.QueryRow(`
        SELECT Users.id, Users.nickname, Users.group, Users.star, Users.finish_count, Users.point
        FROM Users
        WHERE Users.email = ? AND Users.password=  ?
    `, user.Email, user.Password)
	err = row.Scan(&user.Id, &user.Nickname, &user.Group, &user.Star, &user.Finish_count, &user.Point)
	if err != nil {
		return
	}
	return
}

func (u User) SignUp(userPayload forms.UserSignUp) (user User, err error) {
	db := db.GetDB()
	user = User{
		Nickname: userPayload.Nickname,
		Email:    userPayload.Email,
		Password: userPayload.Password,
		Group:    userPayload.Group,
	}

	stmt, err := db.Prepare(`
        INSERT INTO Users(Users.nickname, Users.email, Users.password, Users.group)
        VALUES(?, ?, ?, ?)
    `)
	if err != nil {
		return
	}
	rs, err := stmt.Exec(user.Nickname, user.Email, user.Password, user.Group)
	if err != nil {
		return
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return
	}
	defer stmt.Close()

	row := db.QueryRow(`
        SELECT Users.id, Users.nickname, Users.group, Users.star, Users.finish_count, Users.point
        FROM Users
        WHERE Users.id = ?
    `, id)
	err = row.Scan(&user.Id, &user.Nickname, &user.Group, &user.Star, &user.Finish_count, &user.Point)
	if err != nil {
		return
	}
	return
}

func (u User) Update(userPayload forms.UserUpdate) (user User, err error) {
	db := db.GetDB()
	user = User{
		Nickname:     userPayload.Nickname,
		Email:        userPayload.Email,
		Password:     userPayload.Password,
		Group:        userPayload.Group,
		Star:         userPayload.Star,
		Finish_count: userPayload.Finish_count,
		Point:        userPayload.Point,
	}

	row := db.QueryRow(`
        SELECT Users.id
        FROM Users
        WHERE Users.email = ? AND Users.password = ?
    `, user.Email, user.Password)
	err = row.Scan(&user.Id)
	if err != nil {
		return
	}

	stmt, err := db.Prepare(`
        UPDATE Users
        SET Users.nickname = ?, Users.password = ?, Users.group = ?, Users.star = ?, Users.finish_count = ?, Users.point = ?, Users.updated_at = CURRENT_TIMESTAMP
        WHERE Users.id = ?
    `)
	if err != nil {
		return
	}
	_, err = stmt.Exec(user.Nickname, user.Password, user.Group, user.Star, user.Finish_count, user.Point, user.Id)
	if err != nil {
		return
	}
	defer stmt.Close()
	return
}
