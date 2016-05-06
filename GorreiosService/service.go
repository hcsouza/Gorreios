package GorreiosService

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/hcsouza/Gorreios/GorreiosHttp"
)

func searchCep(id string) string {

	uriCorreios := flag.String("uriCorreios", "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente", "service CEP URI")
	cep := flag.String("cep", "01508000", "cep number")
	flag.Parse()

	GorreiosRequest, err := GorreiosHttp.SoapRequestFactory()

	if err != nil {
		fmt.Println(err)
	}

	GorreiosRequest.SetRequest(*uriCorreios, "ConsultaCEP", *cep)

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
