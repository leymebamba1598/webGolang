package handlers

import (
	"GoWebTotal/rest/models"
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
)

//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtienen todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{Id: 2, UserName: "Didier", Password: "23432"}

	output, _ := yaml.Marshal(&user)
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
