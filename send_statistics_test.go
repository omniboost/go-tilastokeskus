package tilastokeskus_test

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"testing"

	"github.com/omniboost/go-tilastokeskus"
	"golang.org/x/net/html/charset"
)

func TestSendStatisticsRequest(t *testing.T) {
	b, err := os.ReadFile("docs/majoitustilasto.xml")
	if err != nil {
		t.Fatal(err)
	}

	majoitustilasto := tilastokeskus.Majoitustilasto{}
	decoder := xml.NewDecoder(bytes.NewReader(b))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&majoitustilasto)
	if err != nil {
		t.Fatal(err)
	}

	req := client.NewSendStatisticsRequest()
	req.RequestBody().DataTransfer.Data.Majoitustilasto = majoitustilasto

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
