package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola mundo")
	})

	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Numero dos")
	})
	//Iniciar un servidor en go
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
