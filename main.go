package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))
type Usuaurio struct {
	Username string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		template, err := template.ParseFiles("templates/index.html")
		if err != nil {
			panic(err)
		}
		user := Usuaurio{"Didier"}
		template.Execute(writer, user)
	})
	fmt.Println("Servidor corriendo en el puerto: 3000 ")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
