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
	nombre := vars["nombre"]
	edad := vars["edad"]

	fmt.Fprintf(w, "Los parametros son "+nombre+" "+edad)
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/didier/{nombre}/{edad:[0-9]+}", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}
