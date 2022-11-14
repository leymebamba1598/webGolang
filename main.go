package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("El metodo es ", r.Method)
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hola mundo el metodo es GET")
		case "POST":
			fmt.Fprintf(w, "Hola el metodo es POST")
		case "PUT":
			fmt.Fprintf(w, "Hola el metodo es PUT")
		case "DELETE":
			fmt.Fprintf(w, "Hola el metodo es DELETE")
		default:
			http.Error(w, "Metodo no valido", http.StatusBadRequest)
		}
	})
	//Iniciar un servidor en go
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
