package GorreiosHttp

import (
	"encoding/xml"
	"errors"
	"github.com/franela/goreq"
	"io/ioutil"
	"time"
	)

type SoapRequest struct {
	address           string
	XMLName           xml.Name `xml:"x:Envelope"`
	SoapenvNamespace1 string   `xml:"xmlns:x,attr"`
	SoapenvNamespace2 string   `xml:"xmlns:cli,attr"`
	Header            RequestHeader
	Body              RequestBody
}

type RequestHeader struct {
	XMLName xml.Name `xml:"x:Header"`
}

type RequestBody struct {
	XMLName xml.Name `xml:"x:Body"`
	Content RequestContent
}

type RequestContent interface {
}

func SoapRequestFactory() (*SoapRequest, error) {
	request := new(SoapRequest)
	request.SoapenvNamespace1 = "http://schemas.xmlsoap.org/soap/envelope/"
	request.SoapenvNamespace2 = "http://cliente.bean.master.sigep.bsb.correios.com.br/"

	return request, nil
}

type RequestContentCEP struct {
	XMLName xml.Name `xml:"cli:consultaCEP"`
	Cep     string   `xml:"cep,omitempty"`
}

func (this *SoapRequest) SetRequest(serviceAddress string, contentType string, content string) error {

	this.address = serviceAddress

	switch contentType {
	case "ConsultaCEP":
		requestContent := new(RequestContentCEP)
		requestContent.Cep = content
		this.Body.Content = requestContent
		break
	default:
		return errors.New("Unrecognized Request Type: " + contentType)
	}
	return nil
}

func (this *SoapRequest) Do() ([]byte, error) {
	formattedXml, err := xml.Marshal(this)
	if err != nil {
		return nil, err
	}

	httpResponse, err := goreq.Request{
		Method:      "POST",
		Uri:         this.address,
		ContentType: "text/xml;charset=UTF8",
		Body:        formattedXml,
		Timeout:     30 * time.Second,
		//ShowDebug: 	 true,
	}.Do()
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, errors.New("Unable to retrieve status")
	}

	byteBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return byteBody, nil
}
