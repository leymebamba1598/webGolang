package main

import (
	"fmt"
	"log"
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

func Hola2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde funcion")
}
func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hola mundo")
	})
	http.HandleFunc("/hola2", Hola2) //DefaultServerMux

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: nil, //Si es nill utilizamos el DefaultServerMux
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error", err)
	}
}
