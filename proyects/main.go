package main

import (
	mux2 "GoWebTotal/proyects/mux"
	"fmt"
	"log"
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

func main() {

	mux := mux2.CreateMux()

	mux.AddMux("/Hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola desde funcion anonima")
	})
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
