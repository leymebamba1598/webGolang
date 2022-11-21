package models

import (
	"errors"
	"fmt"
	"math/rand"
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

//Obtener un usuario
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

type Users []User

//Obtener todos los usuarios
func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}
func GetUsersSlice() []User {
	return listUser
}

//Crear el usuario
func SaveUser(user User) User {
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}
func SaveUserSlice(user User) User {
	user.Id = rand.Intn(10)
	listUser = append(listUser, user)
	return user
}

//Actualizar usuario
func UpdateUser(user User, username, pass string) User {
	user.UserName = username
	user.Password = pass
	users[user.Id] = user
	return user
}
func UpdateUserSlice(user User, id int) User {
	for index, usuario := range listUser {
		if id == usuario.Id {
			user.Id = id
			listUser[index].UserName = user.UserName
			listUser[index].Password = user.Password

		}
	}

	return user
}
