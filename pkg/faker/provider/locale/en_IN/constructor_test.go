package en_IN_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_IN"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localetest"
)

var _ = Describe("en_IN.New", func() {
	It("produces repeatable localized data", func() {
		localetest.ExpectRepeatable(en_IN.New)
	})

	It("uses local address, postcode and domestic phone formats", func() {
		const (
			locality      = `[\p{L}][\p{L} '’.-]*`
			streetAddress = `[1-9][0-9]* ` + locality
			postcode      = `[1-9][0-9]{5}`
		)
		localetest.ExpectFormats(en_IN.New, &localetest.Formats{
			StreetAddress: "^" + streetAddress + "$",
			FullAddress:   "^" + streetAddress + "\n" + locality + " " + locality + " " + postcode + "$",
			Postcode:      "^" + postcode + "$",
			Phone:         `^[6-9]([0-9]{9}|[0-9]{4} [0-9]{5}|[0-9]{3} [0-9]{3} [0-9]{3})$`,
			PostcodeMasks: []string{"%#####"},
			PhoneMasks:    []string{"6#########", "7#########", "8#########", "9#########", "6#### #####", "7#### #####", "8#### #####", "9#### #####", "6### ### ###", "7### ### ###", "8### ### ###", "9### ### ###"},
		})
	})

	It("places the address components in postal order", func() {
		data := en_IN.New()
		data.Addresses.BuildingNumbers = []string{"17"}
		data.Addresses.StreetNames = []string{"Lotus Road"}
		data.Addresses.CityNames = []string{"Pune"}
		data.Addresses.Postcodes = []string{"411001"}
		data.Addresses.States = []string{"Maharashtra"}
		address := faker.CreateWithLocaleAndSeed(data, 42).Address()
		Expect(address.Address()).To(Equal("17 Lotus Road\nPune Maharashtra 411001"))
	})

	It("provides state names", func() {
		address := faker.CreateWithLocaleAndSeed(en_IN.New(), 42).Address()
		Expect(address.State()).NotTo(BeEmpty())
	})
})
