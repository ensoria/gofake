package en_CA_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_CA"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localetest"
)

var _ = Describe("en_CA.New", func() {
	It("produces repeatable localized data", func() {
		localetest.ExpectRepeatable(en_CA.New)
	})

	It("uses local address, postcode and domestic phone formats", func() {
		const (
			locality      = `[\p{L}][\p{L} '’.-]*`
			streetAddress = `[1-9][0-9]* ` + locality
			postcode      = `(M[1-9]|V[0-9])B [0-9]A[0-9]`
		)
		localetest.ExpectFormats(en_CA.New, &localetest.Formats{
			StreetAddress: "^" + streetAddress + "$",
			FullAddress:   "^" + streetAddress + "\n" + locality + ` (AB|BC|MB|NB|NL|NS|NT|NU|ON|PE|QC|SK|YT)  ` + postcode + "$",
			Postcode:      "^" + postcode + "$",
			Phone:         `^(416|604)-555-01[0-9]{2}$`,
			PostcodeMasks: []string{"M%B #A#", "V#B #A#"},
			PhoneMasks:    []string{"416-555-01##", "604-555-01##"},
		})
	})

	It("places the address components in postal order", func() {
		data := en_CA.New()
		data.Addresses.BuildingNumbers = []string{"17"}
		data.Addresses.StreetNames = []string{"Maple Road"}
		data.Addresses.CityNames = []string{"Ottawa"}
		data.Addresses.Postcodes = []string{"M1B 2A3"}
		data.Addresses.StateAbbrs = []string{"ON"}
		address := faker.CreateWithLocaleAndSeed(data, 42).Address()
		Expect(address.Address()).To(Equal("17 Maple Road\nOttawa ON  M1B 2A3"))
	})

	It("provides province and territory names and abbreviations", func() {
		address := faker.CreateWithLocaleAndSeed(en_CA.New(), 42).Address()
		Expect(address.State()).NotTo(BeEmpty())
		Expect(address.StateAbbr()).To(MatchRegexp(`^(AB|BC|MB|NB|NL|NS|NT|NU|ON|PE|QC|SK|YT)$`))
	})
})
