// Package localetest checks the output of locale providers through the public Faker API.
package localetest

import (
	"regexp"
	"strings"
	"unicode/utf8"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
	"github.com/ensoria/gofake/pkg/faker/provider"
)

func sample(f *faker.Faker) []string {
	return []string{
		f.Person().MaleName(), f.Person().FemaleName(), f.Person().Name(),
		f.Address().City(), f.Address().Street(), f.Address().StreetAddress(),
		f.Address().Address(), f.Address().Postcode(),
		f.Company().Name(), f.Company().JobTitle(), f.PhoneNumber().PhoneNumber(),
	}
}

// ExpectRepeatable verifies the supported methods and their shared seed behavior.
func ExpectRepeatable(newLocale func() *provider.Localized) {
	GinkgoHelper()
	const seed = 42
	data := newLocale()
	f := faker.CreateWithLocaleAndSeed(data, seed)
	other := faker.CreateWithLocaleAndSeed(newLocale(), seed)
	for range 4 {
		values := sample(f)
		Expect(values).To(Equal(sample(other)))
		for _, value := range values {
			Expect(strings.TrimSpace(value)).NotTo(BeEmpty())
			Expect(utf8.ValidString(value)).To(BeTrue())
		}
	}
	company := f.Company().Name()
	Expect(data.Companies.CompanySuffixes).To(ContainElement(Satisfy(func(form string) bool {
		return strings.HasSuffix(company, " "+form)
	})), "company %q must include a legal form", company)
	Expect(f.Company().JobTitle()).To(BeElementOf(data.Companies.JobTitleNames))
}

type Formats struct {
	StreetAddress string
	FullAddress   string
	Postcode      string
	Phone         string
	PostcodeMasks []string
	PhoneMasks    []string
}

// ExpectFormats checks every configured mask and samples rendered output under fixed seeds.
func ExpectFormats(newLocale func() *provider.Localized, formats *Formats) {
	GinkgoHelper()
	data := newLocale()
	Expect(data.Addresses.Postcodes).To(ConsistOf(formats.PostcodeMasks))
	Expect(data.PhoneNumbers.Formats).To(ConsistOf(formats.PhoneMasks))
	const seeds = 32
	var postcodes, phones []string
	for seed := range int64(seeds) {
		f := faker.CreateWithLocaleAndSeed(data, seed)
		Expect(f.Address().StreetAddress()).To(MatchRegexp(formats.StreetAddress), "seed %d", seed)
		Expect(f.Address().Address()).To(MatchRegexp(formats.FullAddress), "seed %d", seed)
		for range 4 {
			postcode := f.Address().Postcode()
			phone := f.PhoneNumber().PhoneNumber()
			Expect(postcode).To(MatchRegexp(formats.Postcode), "seed %d", seed)
			Expect(phone).To(MatchRegexp(formats.Phone), "seed %d", seed)
			postcodes = append(postcodes, postcode)
			phones = append(phones, phone)
		}
	}
	expectAllMasksReached(formats.PostcodeMasks, postcodes)
	expectAllMasksReached(formats.PhoneMasks, phones)
}

func expectAllMasksReached(masks, samples []string) {
	GinkgoHelper()
	digits := strings.NewReplacer("#", "[0-9]", "%", "[1-9]")
	for _, mask := range masks {
		pattern := "^" + digits.Replace(regexp.QuoteMeta(mask)) + "$"
		Expect(samples).To(ContainElement(MatchRegexp(pattern)), "expected an output matching %q", mask)
	}
}
