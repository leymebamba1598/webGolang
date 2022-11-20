package main

import (
	"log"
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

type User struct {
	name string
}

func main() {
	redirect := http.RedirectHandler("https://www.facebook.com", http.StatusMovedPermanently)
	notFound := http.NotFoundHandler()

	mux := http.NewServeMux()

	mux.Handle("/redirect", redirect)
	mux.Handle("/not", notFound)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error", err)
	}
}
