package main

import (
	"fmt"
	"net/url"
)

func createUrl() string {
	u, err := url.Parse("/params")
	if err != nil {
		panic(err)
	}
	u.Host = "localhost:3000"
	u.Scheme = "http"

	query := u.Query() //Nos regresa un mapa
	query.Add("nombre", "valor")

	u.RawQuery = query.Encode()
	return u.String()
}
func main() {
	url := createUrl()
	fmt.Println("La url final es", url)
}
