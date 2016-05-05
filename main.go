package main

import (
	"encoding/xml"
	"fmt"
	"github.com/hcsouza/Gorreios/GorreiosHttp"
)

func main() {

	request, err := GorreiosHttp.SoapRequestFactory()

	if err != nil {
		fmt.Println(err)
	}

	request.SetRequest("https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente?wsdl", "ConsultaCEP", "01508000")

	byteBody, err := request.Do()
	if err != nil {
		fmt.Println(err)
	}

	var respCEP GorreiosHttp.SoapGetCepResponse

	err = xml.Unmarshal(byteBody, &respCEP)
	if err != nil {
		fmt.Println(err)
	}

	getCEP := respCEP.Body.CepResponse.Return_

	fmt.Println(
		fmt.Sprintf("teste:  %s", getCEP.Bairro),
	)
}
