package tilastokeskus_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	tilastokeskus "github.com/omniboost/go-tilastokeskus"
)

var (
	client *tilastokeskus.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	username := os.Getenv("TSK_USERNAME")
	password := os.Getenv("TSK_PASSWORD")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	client = tilastokeskus.NewClient(nil)
	if debug != "" {
		client.SetDebug(true)
	}

	client.SetUsername(username)
	client.SetPassword(password)

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	m.Run()
}
