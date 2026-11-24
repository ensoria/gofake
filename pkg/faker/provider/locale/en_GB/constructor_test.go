package en_GB_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_GB"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localetest"
)

var _ = Describe("en_GB.New", func() {
	It("produces repeatable localized data", func() {
		localetest.ExpectRepeatable(en_GB.New)
	})

	It("uses local address, postcode and domestic phone formats", func() {
		const (
			locality      = `[\p{L}][\p{L} '’.-]*`
			streetAddress = `[1-9][0-9]* ` + locality
			postcode      = `(SW[0-9]A [0-9]AA|M[0-9]{2} [0-9]BB)`
		)
		localetest.ExpectFormats(en_GB.New, &localetest.Formats{
			StreetAddress: "^" + streetAddress + "$",
			FullAddress:   "^" + streetAddress + "\n" + locality + "\n" + postcode + "$",
			Postcode:      "^" + postcode + "$",
			Phone:         `^(020 7946 0|07700 900)[0-9]{3}$`,
			PostcodeMasks: []string{"SW#A #AA", "M## #BB"},
			PhoneMasks:    []string{"020 7946 0###", "07700 900###"},
		})
	})

	It("places the address components in postal order", func() {
		data := en_GB.New()
		data.Addresses.BuildingNumbers = []string{"17"}
		data.Addresses.StreetNames = []string{"Orchard Lane"}
		data.Addresses.CityNames = []string{"Bristol"}
		data.Addresses.Postcodes = []string{"M12 3BB"}
		address := faker.CreateWithLocaleAndSeed(data, 42).Address()
		Expect(address.Address()).To(Equal("17 Orchard Lane\nBristol\nM12 3BB"))
	})
})
