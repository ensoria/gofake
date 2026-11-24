package fr_FR

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

// Address layout: https://www.upu.int/UPU/media/upu/PostalEntitiesFiles/addressingUnit/fraEn.pdf
func createAddresses() *provider.Addresses {
	return &provider.Addresses{
		Postcodes: []string{"#####"},
		CityNames: []string{
			"Amiens", "Angers", "Besançon", "Bordeaux", "Brest", "Clermont-Ferrand", "Dijon", "Grenoble",
			"Lille", "Lyon", "Marseille", "Montpellier", "Nantes", "Nîmes", "Orléans", "Paris", "Rennes",
			"Toulouse",
		},
		CityFormats: []string{"{{.}}"},
		CreateCity:  localefmt.City,
		StreetNames: []string{
			"rue des Acacias", "chemin du Moulin", "rue des Érables", "allée des Tilleuls",
			"avenue de la Gare", "boulevard des Fleurs", "rue de la Liberté", "place du Marché",
			"rue des Ormes", "impasse des Prés", "avenue du Parc", "rue des Rosiers", "allée des Saules",
			"rue du Verger", "quai des Pêcheurs", "rue de l’Église",
		},
		StreetFormats:        []string{"{{.}}"},
		CreateStreet:         localefmt.Street,
		BuildingNumbers:      []string{"%", "%#", "%##"},
		StreetAddressFormats: []string{"{{.Number}} {{.Street}}"},
		CreateStreetAddress:  localefmt.StreetAddress,
		AddressFormats:       []string{"{{.StreetAddress}}\n{{.Postcode}} {{.City}}"},
		CreateAddress:        localefmt.Address,
	}
}
