package tilastokeskus

import (
	"encoding/xml"
)

type DataTransfer struct {
	DataID   int    `xml:"stat:dataId"`
	UserID   string `xml:"stat:userId"`
	Password string `xml:"stat:password"`
	Flags    string `xml:"stat:flags"`
	Data     Data   `xml:"stat:data"`
}

type Data struct {
	Majoitustilasto Majoitustilasto `xml:"majoitustilasto"`
}

func (d Data) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// double encode xml
	b, err := xml.Marshal(d.Majoitustilasto)
	if err != nil {
		return err
	}

	return e.EncodeElement(b, start)
}

type DataTransferResponse struct {
	XMLName            xml.Name   `xml:"DataTransferResponse"`
	Xmlns              string     `xml:"xmlns,attr"`
	DataTransferResult BaseResult `xml:"DataTransferResult"`
}

type Majoitustilasto struct {
	XMLName xml.Name `xml:"majoitustilasto"`
	Xmlns   string   `xml:"xmlns,attr"`
	Liike   []Liike  `xml:"liike"`
}

type Liike struct {
	Tunnus    string      `xml:"tunnus,attr"`
	Ajankohta []Ajankohta `xml:"ajankohta"`
}

type Ajankohta struct {
	Vuosi         string        `xml:"vuosi,attr"`
	Kuukausi      string        `xml:"kuukausi,attr"`
	Kuukausitieto Kuukausitieto `xml:"kuukausitieto"`
	Yopymiset     []Yopymiset   `xml:"yopymiset"`
}

type Kuukausitieto struct {
	Aukiolopaivat  int `xml:"aukiolopaivat"`
	Huoneet        int `xml:"huoneet"`
	Vuoteet        int `xml:"vuoteet"`
	Kuukausikaytto int `xml:"kuukausikaytto"`
	Myynti         int `xml:"myynti"`
}

type Yopymiset struct {
	Avain           string          `xml:"avain,attr"`
	Saapuneet       int             `xml:"saapuneet"`
	Yopyneet        int             `xml:"yopyneet"`
	Matkantarkoitus Matkantarkoitus `xml:"matkantarkoitus"`
	Majoitusmuoto   Majoitusmuoto   `xml:"majoitusmuoto"`
	Maaerittely     *Maaerittely    `xml:"maaerittely,omitempty"`
}

type Maaerittely struct {
	Maa []Maa `xml:"maa"`
}

type Matkantarkoitus struct {
	VapaaAika int `xml:"vapaaaika"`
	Ammatti   int `xml:"ammatti"`
	Muut      int `xml:"muut"`
}

type Majoitusmuoto struct {
	Huone         int `xml:"huone"`
	Matkailuvaunu int `xml:"matkailuvaunu"`
	Teltta        int `xml:"teltta"`
}

type Maa struct {
	MaaKoodi  string `xml:"maakoodi,attr"`
	Saapuneet int    `xml:"saapuneet"`
	Yopyneet  int    `xml:"yopyneet"`
}

var baseResultErrors = map[string]string{
	"0":  "Operation successful",
	"5":  "Save failed",
	"10": "Authentication failed. Please check user ID and password",
	"11": "Problem with user information. Please contact Statistics Finland",
}
