package tilastokeskus

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	layout := "2006-01-02T15:04:05"
	nt, err := time.Parse(layout, s)
	*t = Time{Time: nt}
	return err
}

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02T15:04:05"
	s := t.Format(layout)
	return e.EncodeElement(s, start)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format("2006-01-02T15:04:05"))
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard format
	t.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try custom format
	t.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}
