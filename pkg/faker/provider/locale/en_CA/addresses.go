package en_CA

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

// Address layout: https://www.canadapost-postescanada.ca/cpc/en/support/articles/addressing-guidelines/important-information.page
func createAddresses() *provider.Addresses {
	return &provider.Addresses{
		Postcodes: []string{"M%B #A#", "V#B #A#"},
		States: []string{
			"Alberta", "British Columbia", "Manitoba", "New Brunswick", "Newfoundland and Labrador",
			"Northwest Territories", "Nova Scotia", "Nunavut", "Ontario", "Prince Edward Island",
			"Quebec", "Saskatchewan", "Yukon",
		},
		StateAbbrs: []string{
			"AB", "BC", "MB", "NB", "NL", "NT", "NS", "NU", "ON", "PE", "QC", "SK", "YT",
		},
		CityNames: []string{
			"Calgary", "Charlottetown", "Edmonton", "Fredericton", "Halifax", "Hamilton", "Iqaluit",
			"Ottawa", "Regina", "Saskatoon", "St. John's", "Toronto", "Vancouver", "Victoria",
			"Whitehorse", "Winnipeg", "Yellowknife",
		},
		CityFormats: []string{"{{.}}"},
		CreateCity:  localefmt.City,
		StreetNames: []string{
			"Aspen Drive", "Birch Avenue", "Cedar Crescent", "Crescent Road", "Elm Street",
			"Harbour Road", "Lakeview Drive", "Maple Avenue", "Meadow Lane", "Mill Street",
			"Pine Ridge Road", "Prairie Trail", "River Road", "Spruce Court", "Station Street",
			"Willow Way",
		},
		StreetFormats:        []string{"{{.}}"},
		CreateStreet:         localefmt.Street,
		BuildingNumbers:      []string{"%", "%#", "%##"},
		StreetAddressFormats: []string{"{{.Number}} {{.Street}}"},
		CreateStreetAddress:  localefmt.StreetAddress,
		AddressFormats:       []string{"{{.StreetAddress}}\n{{.City}} {{.State}}  {{.Postcode}}"},
		CreateAddress:        localefmt.AddressWithStateAbbr,
	}
}
