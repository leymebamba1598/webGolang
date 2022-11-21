package handlers

import (
	"GoWebTotal/rest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsersSlice())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocesableEntity(w)
	} else {
		models.SendData(w, models.SaveUserSlice(user))
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return

	}
	userResponse := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userResponse); err != nil {
		models.SendUnprocesableEntity(w)
		return
	}
	//user = models.UpdateUser(user, userResponse.UserName, userResponse.Password)
	user = models.UpdateUserSlice(userResponse, user.Id)
	models.SendData(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina un usuario")
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUserSlice(userId); err != nil {
		return user, err
	} else {
		return user, nil
	}

}
