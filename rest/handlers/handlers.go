package handlers

import (
	"GoWebTotal/rest/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtienen todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"]) //string

	if user, err := models.GetUserSlice(userId); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se crea un usuario")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina un usuario")
}
