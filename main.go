package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor del header")
		w.Header().Add("Nombre2", "Valor del header dos")
		//fmt.Fprintf(w, "Hola mundo")

		//Redireccionar url
		http.Redirect(w, r, "/dos", http.StatusMovedPermanently)

	})

	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Numero dos")
		http.NotFound(w, r)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Este es un error", http.StatusNotFound)
	})

	//Iniciar un servidor en go
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
