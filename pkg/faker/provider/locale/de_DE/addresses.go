package de_DE

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

// Address layout: https://www.upu.int/UPU/media/upu/PostalEntitiesFiles/addressingUnit/deuEn.pdf
func createAddresses() *provider.Addresses {
	return &provider.Addresses{
		Postcodes: []string{"#####"},
		States: []string{
			"Baden-Württemberg", "Bayern", "Berlin", "Brandenburg", "Bremen", "Hamburg", "Hessen",
			"Mecklenburg-Vorpommern", "Niedersachsen", "Nordrhein-Westfalen", "Rheinland-Pfalz",
			"Saarland", "Sachsen", "Sachsen-Anhalt", "Schleswig-Holstein", "Thüringen",
		},
		CityNames: []string{
			"Aachen", "Augsburg", "Berlin", "Bremen", "Dresden", "Düsseldorf", "Erfurt",
			"Freiburg im Breisgau", "Hamburg", "Hannover", "Köln", "Leipzig", "Lübeck", "Mainz",
			"München", "Nürnberg", "Potsdam", "Stuttgart",
		},
		CityFormats: []string{"{{.}}"},
		CreateCity:  localefmt.City,
		StreetNames: []string{
			"Ahornweg", "Am Stadtpark", "Bergstraße", "Birkenallee", "Brückenstraße", "Gartenweg",
			"Hafenstraße", "Im Wiesengrund", "Kastanienallee", "Kirchplatz", "Lindenstraße",
			"Marktstraße", "Mühlenweg", "Rosenweg", "Schloßstraße", "Wiesenring",
		},
		StreetFormats:        []string{"{{.}}"},
		CreateStreet:         localefmt.Street,
		BuildingNumbers:      []string{"%", "%#", "%##"},
		StreetAddressFormats: []string{"{{.Street}} {{.Number}}"},
		CreateStreetAddress:  localefmt.StreetAddress,
		AddressFormats:       []string{"{{.StreetAddress}}\n{{.Postcode}} {{.City}}"},
		CreateAddress:        localefmt.Address,
	}
}
