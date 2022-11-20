package main

import (
	"fmt"
	"log"
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mux DefaultServerMux")
}
func Hola2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mux propio")
}
func main() {
	http.HandleFunc("/", Hola) //DefaultServerMux

	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola2)
	mux.HandleFunc("/dos", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hola mundo desde mux propio dos")
	})

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux, //Si es nill utilizamos el DefaultServerMux
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error", err)
	}
}
