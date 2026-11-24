package en_GB

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

// Address layout: https://www.upu.int/UPU/media/upu/PostalEntitiesFiles/addressingUnit/gbrEn.pdf
func createAddresses() *provider.Addresses {
	return &provider.Addresses{
		Postcodes: []string{"SW#A #AA", "M## #BB"},
		CityNames: []string{
			"Aberdeen", "Bath", "Belfast", "Birmingham", "Bristol", "Cambridge", "Cardiff", "Derby",
			"Dundee", "Edinburgh", "Glasgow", "Leeds", "Liverpool", "London", "Manchester",
			"Newcastle upon Tyne", "Swansea", "York",
		},
		CityFormats: []string{"{{.}}"},
		CreateCity:  localefmt.City,
		StreetNames: []string{
			"Abbey Road", "Alder Close", "Beech Avenue", "Bridge Street", "Castle Lane", "Cedar Crescent",
			"Church Lane", "Elm Grove", "High Street", "Kingfisher Way", "Market Place", "Mill Road",
			"Orchard Lane", "Station Road", "Victoria Terrace", "Willow Drive",
		},
		StreetFormats:        []string{"{{.}}"},
		CreateStreet:         localefmt.Street,
		BuildingNumbers:      []string{"%", "%#", "%##"},
		StreetAddressFormats: []string{"{{.Number}} {{.Street}}"},
		CreateStreetAddress:  localefmt.StreetAddress,
		AddressFormats:       []string{"{{.StreetAddress}}\n{{.City}}\n{{.Postcode}}"},
		CreateAddress:        localefmt.Address,
	}
}
