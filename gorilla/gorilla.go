package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

func YourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "Los parametros son "+id)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id:[0-9]+}", YourHandler).Methods("GET", "DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}
