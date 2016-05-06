package main

import (
	// "encoding/xml"
	// "fmt"
	"github.com/go-martini/martini"
	"github.com/hcsouza/Gorreios/GorreiosService"
	"net/http"
)

func main() {

	mserver := martini.Classic()

	mserver.Get("/", func() string {
		return "http://gorreios.com.br/cep/:id"
	})

	mserver.Get("/cep/:id", func(params martini.Params, writer http.ResponseWriter) string {
		writer.Header().Set("Content-Type", "application/json")
		return GorreiosService.searchCep(params["id"])
	})

	mserver.Run()
}
