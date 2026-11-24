package fr_FR_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/fr_FR"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localetest"
)

var _ = Describe("fr_FR.New", func() {
	It("produces repeatable localized data", func() {
		localetest.ExpectRepeatable(fr_FR.New)
	})

	It("uses local address, postcode and domestic phone formats", func() {
		const (
			locality      = `[\p{L}][\p{L} '’.-]*`
			streetAddress = `[1-9][0-9]* ` + locality
			postcode      = `[0-9]{5}`
		)
		localetest.ExpectFormats(fr_FR.New, &localetest.Formats{
			StreetAddress: "^" + streetAddress + "$",
			FullAddress:   "^" + streetAddress + "\n" + postcode + " " + locality + "$",
			Postcode:      "^" + postcode + "$",
			Phone:         `^(01 99 00|06 39 98) [0-9]{2} [0-9]{2}$`,
			PostcodeMasks: []string{"#####"},
			PhoneMasks:    []string{"01 99 00 ## ##", "06 39 98 ## ##"},
		})
	})

	It("places the address components in postal order", func() {
		data := fr_FR.New()
		data.Addresses.BuildingNumbers = []string{"17"}
		data.Addresses.StreetNames = []string{"rue des Érables"}
		data.Addresses.CityNames = []string{"Orléans"}
		data.Addresses.Postcodes = []string{"45000"}
		address := faker.CreateWithLocaleAndSeed(data, 42).Address()
		Expect(address.Address()).To(Equal("17 rue des Érables\n45000 Orléans"))
	})

	It("renders names with French characters from the supplied data", func() {
		names := faker.CreateWithLocaleAndSeed(fr_FR.New(), 42).Person()
		var rendered []string
		for range 32 {
			rendered = append(rendered, names.Name())
		}
		Expect(rendered).To(ContainElement(MatchRegexp(`[àâçéèêëîïôùûüÿœ]`)))
	})
})
