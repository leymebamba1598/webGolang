package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		template, err := template.New("Hola").Parse("Hola mundo")
		if err != nil {
			panic(err)
		}
		template.Execute(writer, nil)
	})
	fmt.Println("Servidor corriendo en el puerto: 3000 ")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
