package main

import (
	"GoWebTotal/rest/handlers"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))
func main() {
	mux := mux2.NewRouter()
	mux.HandleFunc("/api/v1/users/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POS")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods(http.MethodDelete)

	log.Println("El servidor esta escuchando el puerto: 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
