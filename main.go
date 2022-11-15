package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.RawQuery)//Ver parametros enviados por url

		fmt.Println(r.URL.Query()) //Retorna un mapa
		//Obtenemos los datos del mapa

		name := r.URL.Query().Get("name")
		if len(name) != 0 {
			fmt.Println("Nombre", name)
		}
		fmt.Fprintf(w, "Hola mundo")

	})
	//Iniciar un servidor en go
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
