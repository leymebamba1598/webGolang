package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

//-----------------------------TODO METODOS STRUCT RESPONSE -----------------------------//

//Envia respuesta si el usuario no existe
func (this *Response) NotFound() {
	this.Status = http.StatusNotFound
	this.Message = "Resource not found"
}

func (this *Response) UnprocessableEntity() {
	this.Status = http.StatusUnprocessableEntity
	this.Message = "UnprocessableEntity"
}
func (this *Response) NotContent() {
	this.Status = http.StatusNoContent
	this.Message = "No contente"
}

//Envia la respuesta al cliente
func (this *Response) Send() {
	this.writer.Header().Set("Content-Type", this.contentType)
	this.writer.WriteHeader(this.Status)

	output, _ := json.Marshal(&this) // Convierte e json
	fmt.Fprintf(this.writer, string(output))

}

//-----------------------------TODO FUNCIONES -----------------------------//
//datos validos
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

// Cuando no se encuntra
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

// Cuando no se puede mapear el json
func SendUnprocesableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotContent()
	response.Send()
}
