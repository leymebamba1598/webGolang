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
	//user, err := models.GetUser(userId)
	response := models.GetDefaultResponse(w)
	user, err := models.GetUserSlice(userId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.NotFound("usuario no encontrado")
	} else {
		response.Data = user
		response.Message = "Usuario traido correctamente"
	}

	response.Send()
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
