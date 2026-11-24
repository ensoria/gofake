package de_DE_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/de_DE"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localetest"
)

var _ = Describe("de_DE.New", func() {
	It("produces repeatable localized data", func() {
		localetest.ExpectRepeatable(de_DE.New)
	})

	It("uses local address, postcode and domestic phone formats", func() {
		const (
			locality      = `[\p{L}][\p{L} '’.-]*`
			streetAddress = locality + ` [1-9][0-9]*`
			postcode      = `[0-9]{5}`
		)
		localetest.ExpectFormats(de_DE.New, &localetest.Formats{
			StreetAddress: "^" + streetAddress + "$",
			FullAddress:   "^" + streetAddress + "\n" + postcode + " " + locality + "$",
			Postcode:      "^" + postcode + "$",
			Phone:         `^(030 23125|040 66969)[0-9]{3}$`,
			PostcodeMasks: []string{"#####"},
			PhoneMasks:    []string{"030 23125###", "040 66969###"},
		})
	})

	It("places the address components in postal order", func() {
		data := de_DE.New()
		data.Addresses.BuildingNumbers = []string{"17"}
		data.Addresses.StreetNames = []string{"Gartenstraße"}
		data.Addresses.CityNames = []string{"München"}
		data.Addresses.Postcodes = []string{"80331"}
		address := faker.CreateWithLocaleAndSeed(data, 42).Address()
		Expect(address.Address()).To(Equal("Gartenstraße 17\n80331 München"))
	})

	It("provides state names", func() {
		address := faker.CreateWithLocaleAndSeed(de_DE.New(), 42).Address()
		Expect(address.State()).NotTo(BeEmpty())
	})

	It("renders names with German characters from the supplied data", func() {
		names := faker.CreateWithLocaleAndSeed(de_DE.New(), 42).Person()
		var rendered []string
		for range 32 {
			rendered = append(rendered, names.Name())
		}
		Expect(rendered).To(ContainElement(MatchRegexp(`[äöüÄÖÜß]`)))
	})
})
