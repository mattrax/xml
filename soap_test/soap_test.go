package soap_test

import (
	"testing"

	"github.com/mattrax/xml"
)

type Envelope struct {
	XMLName    xml.Name `xml:"s:Envelope"`
	NamespaceA string   `xml:"xmlns:a,attr"`
	NamespaceS string   `xml:"xmlns:s,attr"`
	Header     struct {
		Action struct {
			MustUnderstand string `xml:"s:mustUnderstand,attr,omitempty"`
			Value          string `xml:",innerxml"`
		} `xml:"a:Action"`
	} `xml:"s:Header"`
	Body struct {
		Request struct {
			Data string
		}
	} `xml:"s:Body"`
}

var soapRaw = `<?xml version="1.0"?><s:Envelope xmlns:a="http://www.w3.org/2005/08/addressing" xmlns:s="http://www.w3.org/2003/05/soap-envelope"><s:Header><a:Action s:mustUnderstand="1">THIS_IS_THE_ACTION</a:Action></s:Header><s:Body><Request><Data>THIS_IS_THE_DATA</Data></Request></s:Body></s:Envelope>`

func TestSOAP(t *testing.T) {
	var e Envelope
	if err := xml.Unmarshal([]byte(soapRaw), &e); err != nil {
		t.Fatal(err)
	}

	if e.XMLName.Local != "s:Envelope" {
		t.Errorf("got %s, want %s", e.XMLName.Local, "s:Envelope")
	}

	if e.Header.Action.Value != "THIS_IS_THE_ACTION" {
		t.Errorf("got %s, want %s", e.Header.Action.Value, "THIS_IS_THE_ACTION")
	}

	if e.Body.Request.Data != "THIS_IS_THE_DATA" {
		t.Errorf("got %s, want %s", e.Body.Request.Data, "THIS_IS_THE_DATA")
	}

	e.Header.Action.MustUnderstand = "1"
	e.NamespaceA = "http://www.w3.org/2005/08/addressing"
	e.NamespaceS = "http://www.w3.org/2003/05/soap-envelope"

	if outSoapRaw, err := xml.Marshal(&e); err != nil {
		t.Fatal(err)
	} else if outSoap := `<?xml version="1.0"?>` + string(outSoapRaw); outSoap != soapRaw {
		t.Errorf("got '%s', want '%s'", outSoap, soapRaw)
	}
}
