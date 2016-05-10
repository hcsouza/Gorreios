package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/hcsouza/Gorreios/GorreiosHttp"
	"net/http"
	"github.com/zieckey/goini"
)

var(
	ini = goini.New()
)

func init(){
	err := ini.ParseFile("gorreios.ini")
	if err != nil {
	    fmt.Printf("parse INI file gorreios.ini failed : %v\n", err.Error())
	    return
	}
}

func main() {
	mserver := martini.Classic()
	mserver.Get("/", func() string {
		return "http://gorreios.com.br/cep/:id"
	})

	mserver.Get("/cep/:id", func(params martini.Params, writer http.ResponseWriter) string {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		searched := searchCep(params["id"])
		if searched == "null" {
			 writer.WriteHeader(404)
			 return "{\"error\":\"resource_not_found\"}"
		}
		writer.WriteHeader(302)
		return searched
	})

	mserver.Run()
}

func searchCep(id string) string {
	uriCorreios, _ := ini.Get("uriCorreios")
	GorreiosRequest, err := GorreiosHttp.SoapRequestFactory()

	if err != nil {
		fmt.Println(err)
	}
	GorreiosRequest.CreateRequest(uriCorreios, "ConsultaCEP", id)

	byteBody, err := GorreiosRequest.Execute()
	if err != nil {
		fmt.Println(err)
	}
	var respCEP GorreiosHttp.SoapGetCepResponse

	err = xml.Unmarshal(byteBody, &respCEP)
	if err != nil {
		fmt.Println(err)
	}
	getCEP := respCEP.Body.CepResponse.Return
	end, _ := json.Marshal(getCEP)
	return string(end)
}
