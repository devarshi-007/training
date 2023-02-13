package webService

import (
	"errors"
)

var (
	users  []User
	newtId int = 1
)

type OurCustomClass struct {
	IsAllGood bool
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Done      OurCustomClass
}

func GetUsers() []User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = newtId
	newtId++
	users = append(users, u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, v := range users {
		currentUser := v
		if currentUser.ID == id {
			return currentUser, nil
		}
	}

	return User{}, errors.New("User not found")
}

func UpdateUserById(usr User) error {
	for _, v := range users {
		if v.ID == usr.ID {
			v = usr
			return errors.New("user update sucessfully")
		}
	}
	return errors.New("user not found")
}

func DeleteUserById(usr User) error {
	for i, v := range users {
		if v.ID == usr.ID {
			if len(users) < i+1 {
				users = users[:i]
			} else {
				users = append(users[:i], users[i+1:]...)
			}

			return errors.New("user update sucessfully")
		}
	}
	return errors.New("user not found")
}
