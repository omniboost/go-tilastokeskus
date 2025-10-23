package tilastokeskus

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Format("2006-01-02")
}

func (t *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	layout := "2006-01-02"
	nt, err := time.Parse(layout, s)
	if err == nil {
		*t = Date{Time: nt}
	}

	layout = "2006-01-02T15:04:05"
	nt, err = time.Parse(layout, s)
	if err == nil {
		*t = Date{Time: nt}
	}

	return err
}

func (t Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02"
	s := t.Format(layout)
	return e.EncodeElement(s, start)
}

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	// Try multiple formats
	layouts := []string{
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02",
	}

	for _, layout := range layouts {
		parsed, err := time.Parse(layout, s)
		if err == nil {
			*t = DateTime{Time: parsed}
			return nil
		}
	}

	return fmt.Errorf("could not parse date: %s", s)
}

func (t DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02T15:04:05"
	s := t.Format(layout)
	return e.EncodeElement(s, start)
}
