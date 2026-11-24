package en_IN

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

// Address layout: https://www.upu.int/UPU/media/upu/PostalEntitiesFiles/addressingUnit/indEn.pdf
func createAddresses() *provider.Addresses {
	return &provider.Addresses{
		Postcodes: []string{"%#####"},
		States: []string{
			"Andhra Pradesh", "Assam", "Bihar", "Chhattisgarh", "Gujarat", "Haryana", "Himachal Pradesh",
			"Jharkhand", "Karnataka", "Kerala", "Madhya Pradesh", "Maharashtra", "Odisha", "Punjab",
			"Rajasthan", "Tamil Nadu", "Telangana", "Uttar Pradesh", "Uttarakhand", "West Bengal",
		},
		CityNames: []string{
			"Ahmedabad", "Bengaluru", "Bhopal", "Bhubaneswar", "Chennai", "Guwahati", "Hyderabad",
			"Indore", "Jaipur", "Kochi", "Kolkata", "Lucknow", "Mumbai", "Mysuru", "Nagpur", "Patna",
			"Pune", "Surat",
		},
		CityFormats: []string{"{{.}}"},
		CreateCity:  localefmt.City,
		StreetNames: []string{
			"Ashoka Road", "Banyan Lane", "Canal Road", "College Road", "Garden Road", "Gulmohar Lane",
			"Hill View Road", "Jasmine Lane", "Lake Road", "Lotus Road", "Mango Grove Road",
			"Market Road", "Neem Lane", "Park Street", "Station Road", "Temple Street",
		},
		StreetFormats:        []string{"{{.}}"},
		CreateStreet:         localefmt.Street,
		BuildingNumbers:      []string{"%", "%#", "%##"},
		StreetAddressFormats: []string{"{{.Number}} {{.Street}}"},
		CreateStreetAddress:  localefmt.StreetAddress,
		AddressFormats:       []string{"{{.StreetAddress}}\n{{.City}} {{.State}} {{.Postcode}}"},
		CreateAddress:        localefmt.AddressWithState,
	}
}
