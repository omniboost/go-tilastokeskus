package tilastokeskus

import (
	"encoding/xml"
)

type RequestEnvelope struct {
	XMLName xml.Name
	// Namespaces []xml.Attr `xml:"-"`

	NS     []xml.Attr `xml:"-"`
	Header SOAPHeader `xml:"soap:Header"`
	Body   any        `xml:"soap:Body"`
}

func NewRequestEnvelope() RequestEnvelope {
	return RequestEnvelope{
		NS: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:stat"}, Value: "http://www.tilastokeskus.fi/StatTrans"},
			{Name: xml.Name{Space: "", Local: "xmlns:soap"}, Value: "http://www.w3.org/2003/05/soap-envelope"},
		},
		// Header: NewSOAPHeader(auth),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header SOAPHeader `xml:"Header"`
	Body   any        `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "soap:Envelope"}

	start.Attr = append(start.Attr, env.NS...)

	type alias RequestEnvelope
	a := alias(env)
	return e.EncodeElement(a, start)
}

type SOAPHeader interface{}

// func (h SOAPHeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	return omitempty.MarshalXML(h, e, start)
// }

// func (h SOAPHeader) IsEmpty() bool {
// 	return zero.IsZero(h)
// }

type AuthDetails struct {
	User        string
	Password    string
	CrsProperty string
	// VendorID    string
	// Version     string
	// Token       string
}

// func NewSOAPHeader(auth AuthDetails) SOAPHeader {
// 	return SOAPHeader{
// 		Authentication: Authentication{
// 			XmlnsNamespace: "Amadeus.Hospitality.HybridCloudEngine",
// 			User:           auth.User,
// 			Password:       auth.Password,
// 			CrsProperty:    auth.CrsProperty,
// 		},
// 	}
// }

type Authentication struct {
	XMLName        xml.Name `xml:"Authentication"`
	XmlnsNamespace string   `xml:"xmlns,attr"`
	User           string   `xml:"User"`
	Password       string   `xml:"Password"`
	CrsProperty    string   `xml:"CrsProperty"`
}
