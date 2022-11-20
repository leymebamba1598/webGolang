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
	fmt.Fprintf(w, "Se obtienen todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"]) //string
	//user, err := models.GetUser(userId)
	user, err := models.GetUserSlice(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	output, _ := json.Marshal(&user)
	fmt.Fprintf(w, string(output))
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
