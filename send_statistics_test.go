package tilastokeskus_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSendStatisticsRequest(t *testing.T) {
	req := client.NewSendStatisticsRequest()
	// req.RequestBody().HcsListServices.PropertyCode = "UTR"
	// req.RequestBody().HcsListServices.Language = "NL"
	// data := tilastokeskus.StatisticsData{
	// 	ReportDate:     time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
	// 	EstablishmentID: "12345",
	// 	Statistics: []tilastokeskus.StatisticsEntry{
	// 		{
	// 			CountryCode: "FI",
	// 			Arrivals:    50,
	// 			NightsSpent: 70,
	// 			Purpose: tilastokeskus.PurposeStats{
	// 				Leisure:  40,
	// 				Business: 30,
	// 				Other:    0,
	// 			},
	// 			Accommodation: tilastokeskus.AccommodationStats{
	// 				Room:    70,
	// 				Caravan: 0,
	// 				Tent:    0,
	// 			},
	// 		},
	// 		{
	// 			CountryCode: "SE",
	// 			Arrivals:    8,
	// 			NightsSpent: 10,
	// 			Purpose: tilastokeskus.PurposeStats{
	// 				Leisure:  6,
	// 				Business: 4,
	// 				Other:    0,
	// 			},
	// 			Accommodation: tilastokeskus.AccommodationStats{
	// 				Room:    10,
	// 				Caravan: 0,
	// 				Tent:    0,
	// 			},
	// 		},
	// 	},
	// }

	// req.SetData(data)
	// resp, err := req.Do()
	// if err != nil {
	// 	t.Error(err)
	// }

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
