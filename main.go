package main

import (
	"encoding/xml"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/hcsouza/Gorreios/GorreiosHttp"
	"net/http"
)

func main() {

	mserver := martini.Classic()

	mserver.Get("/", func() string {
		return "http://gorreios.com.br/cep/:id"
	})

	mserver.Get("/cep/:id", func(params martini.Params, writer http.ResponseWriter) string {
		writer.Header().Set("Content-Type", "application/json")
		return searchCep(params["id"])
	})

	mserver.Run()

}

func searchCep(id string) string {

	// uriCorreios := String("uriCorreios", "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente", "service CEP URI")
	// cep := flag.String("cep", "01508000", "cep number")
	// flag.Parse()
	uriCorreios := "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente"

	GorreiosRequest, err := GorreiosHttp.SoapRequestFactory()

	if err != nil {
		fmt.Println(err)
	}

	GorreiosRequest.SetRequest(uriCorreios, "ConsultaCEP", id)

	byteBody, err := GorreiosRequest.Do()
	if err != nil {
		fmt.Println(err)
	}

	var respCEP GorreiosHttp.SoapGetCepResponse

	err = xml.Unmarshal(byteBody, &respCEP)
	if err != nil {
		fmt.Println(err)
	}

	getCEP := respCEP.Body.CepResponse.Return

	return getCEP.End

}
