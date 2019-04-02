package main

import (
	"net/http"
	"time"
	"github.com/sweetbrain/sample-api/common"
)

var tmpUsers = map[string]User{}

type User struct {
	ID 			string		`json:"id"`
	Name 		string 		`json:"name"`
	Password 	string 		`json:"password"`
	Description	string 		`json:"description"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
}

type Users []User

func RegistUser(user User) (User, error) {
	date := time.Now()
	user.CreatedAt = date
	user.UpdatedAt = date

	id, err := uuid.NewV4()
	if err != nil {
		return user, NewError(http.StatusInternalServerError, err.Error())
	}
	user.ID = id.String()

	tmpUsers[user.ID] = user

	return user, nil
}

func ReadUser(id string) (User, error) {
	user, exist := tmpUsers[id]

	if !exist {
		return user, NewError(http.StatusNotFound, "not found user")
	}
	return user, nil
}

func ListUser() (Users, error) {
	users := Users{}

	for _, user := range tmpUsers {
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(newUser User) (User, error) {
	if _, err := ReadUser(newUser.ID); err != nil {
		return newUser, err
	}

	date := time.Now()
	newUser.CreatedAt = date
	newUser.UpdatedAt = date

	tmpUsers[newUser.ID] =newUser
	return newUser, nil
}

func DeleteUser(id string)error {
	if _, err := ReadUser(id); err != nil {
		return err
	}

	delete(tmpUsers, id)
	return nil
}