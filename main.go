package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.RawQuery) Ver Query completo
		//fmt.Println(r.URL.Query())  Ver mapa

		fmt.Println(r.URL)
		values := r.URL.Query()

		values.Del("otro") //elimina un parametro

		values.Add("name", "Didier")
		values.Add("curos", "Goweb")
		values.Add("edad", "24")
		r.URL.RawQuery = values.Encode()

		fmt.Println(r.URL)
		fmt.Fprintf(w, "Hola mundo")

	})
	//Iniciar un servidor en go
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
