package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)

		//Leer headers
		accesToken := r.Header.Get("acces_token")
		if len(accesToken) != 0 {
			fmt.Println(accesToken)
		}
		//agregar encabezados
		r.Header.Set("nombre", "didier")
		fmt.Println(r.Header)

		fmt.Fprintf(w, "Hola mundo")

	})

	//Iniciar un servidor en go
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
