package GorreiosHttp

import (
	"encoding/xml"
)

type SoapGenericResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	XSINamespace  string   `xml:"xmlns:xsi,attr"`
	XSDNamespace  string   `xml:"xmlns:xsd,attr"`
	SoapNamespace string   `xml:"xmlns:soap,attr"`
}

type SoapGetCepResponse struct {
	SoapGenericResponse
	Body SoapCepBodyResponse
}

type SoapCepBodyResponse struct {
	XMLName     xml.Name `xml:"Body"`
	CepResponse ConsultaCepResponse
}

type ConsultaCepResponse struct {
	XMLName xml.Name  `xml:"http://cliente.bean.master.sigep.bsb.correios.com.br/ consultaCEPResponse"`
	Return_ *Endereco `xml:"return,omitempty"`
}

type Endereco struct {
	XMLName      xml.Name `xml:"http://cliente.bean.master.sigep.bsb.correios.com.br/ enderecoERP"`
	Bairro       string   `xml:"bairro,omitempty"`
	Cep          string   `xml:"cep,omitempty"`
	Cidade       string   `xml:"cidade,omitempty"`
	Complemento  string   `xml:"complemento,omitempty"`
	Complemento2 string   `xml:"complemento2,omitempty"`
	End          string   `xml:"end,omitempty"`
	Id           int64    `xml:"id,omitempty"`
	Uf           string   `xml:"uf,omitempty"`
}
