package main

import (
	"fmt"
	"log"
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

type User struct {
	name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola: "+this.name)
}

func main() {
	didier := &User{name: "Didier"}

	mux := http.NewServeMux()
	mux.Handle("/", didier)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error", err)
	}
}
