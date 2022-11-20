package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[int]User)
var listUser []User

func SetDefaultUser() {
	user := User{Id: 1, UserName: "didier", Password: "4343"}
	user2 := User{Id: 5, UserName: "Carlos", Password: "453"}
	//Usando mapas
	users[user.Id] = user
	//usando slice
	listUser = append(listUser, user, user2)
	fmt.Println(listUser)
}
func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("El usuario no se encunetra dentro de mapa")
}
func GetUserSlice(userId int) (User, error) {
	for _, user := range listUser {
		if user.Id == userId {
			return user, nil
		}
	}
	return User{}, errors.New("El usuario no se encunetra dentro de mapa")
}
