package mux

import (
	"net/http"
)

//Server mux -- Listado de rutas las cuales tienen asociadas acciones
//Handlers -- acciones asociadas a las rutas (responder al cliente (cuerpo, encabezado, status))

type customHandelers func(w http.ResponseWriter, r *http.Request)

type MuxFacilito struct {
	rutasFaciitas map[string]customHandelers //Handlers
}

func (this *MuxFacilito) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fn, ok := this.rutasFaciitas[r.URL.Path]; ok {
		fn(w, r)
	} else {
		http.NotFound(w, r)
	}

}
func (this *MuxFacilito) AddMux(ruta string, fn customHandelers) {
	this.rutasFaciitas[ruta] = fn
}
func CreateMux() *MuxFacilito {
	newMapa := make(map[string]customHandelers)
	return &MuxFacilito{newMapa}
}
