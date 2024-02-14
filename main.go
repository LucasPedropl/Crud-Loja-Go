package main

import (
	"loja/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
