package main

import (
	"fmt"
	"log"
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

type customHandelers func(w http.ResponseWriter, r *http.Request)

type MuxFacilito struct {
	rutasFaciitas map[string]customHandelers //Handlers
}

func (this *MuxFacilito) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := this.rutasFaciitas[r.URL.Path]
	fn(w, r)
}
func (this *MuxFacilito) AddMux(ruta string, fn customHandelers) {
	this.rutasFaciitas[ruta] = fn
}
func main() {
	newMapa := make(map[string]customHandelers)
	mux := &MuxFacilito{rutasFaciitas: newMapa}

	mux.AddMux("/Hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola desde funcion anonima")
	})
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
