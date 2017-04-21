package myservice

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type GetWeather struct {
	XMLName xml.Name `xml:"http://www.webserviceX.NET GetWeather"`

	CityName    string `xml:"CityName,omitempty"`
	CountryName string `xml:"CountryName,omitempty"`
}

type GetWeatherResponse struct {
	XMLName xml.Name `xml:"http://www.webserviceX.NET GetWeatherResponse"`

	GetWeatherResult string `xml:"GetWeatherResult,omitempty"`
}

type GetCitiesByCountry struct {
	XMLName xml.Name `xml:"http://www.webserviceX.NET GetCitiesByCountry"`

	CountryName string `xml:"CountryName,omitempty"`
}

type GetCitiesByCountryResponse struct {
	XMLName xml.Name `xml:"http://www.webserviceX.NET GetCitiesByCountryResponse"`

	GetCitiesByCountryResult string `xml:"GetCitiesByCountryResult,omitempty"`
}

type GlobalWeatherSoap struct {
	client *SOAPClient
}

func NewGlobalWeatherSoap(url string, tls bool, auth *BasicAuth) *GlobalWeatherSoap {
	if url == "" {
		url = "http://www.webservicex.com/globalweather.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &GlobalWeatherSoap{
		client: client,
	}
}

/* Get weather report for all major cities around the world. */
func (service *GlobalWeatherSoap) GetWeather(request *GetWeather) (*GetWeatherResponse, error) {
	response := new(GetWeatherResponse)
	err := service.client.Call("http://www.webserviceX.NET/GetWeather", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Get all major cities by country name(full / part). */
func (service *GlobalWeatherSoap) GetCitiesByCountry(request *GetCitiesByCountry) (*GetCitiesByCountryResponse, error) {
	response := new(GetCitiesByCountryResponse)
	err := service.client.Call("http://www.webserviceX.NET/GetCitiesByCountry", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GlobalWeatherHttpGet struct {
	client *SOAPClient
}

func NewGlobalWeatherHttpGet(url string, tls bool, auth *BasicAuth) *GlobalWeatherHttpGet {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &GlobalWeatherHttpGet{
		client: client,
	}
}

/* Get weather report for all major cities around the world. */
func (service *GlobalWeatherHttpGet) GetWeather(request *String) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Get all major cities by country name(full / part). */
func (service *GlobalWeatherHttpGet) GetCitiesByCountry(request *String) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GlobalWeatherHttpPost struct {
	client *SOAPClient
}

func NewGlobalWeatherHttpPost(url string, tls bool, auth *BasicAuth) *GlobalWeatherHttpPost {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &GlobalWeatherHttpPost{
		client: client,
	}
}

/* Get weather report for all major cities around the world. */
func (service *GlobalWeatherHttpPost) GetWeather(request *String) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Get all major cities by country name(full / part). */
func (service *GlobalWeatherHttpPost) GetCitiesByCountry(request *String) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
