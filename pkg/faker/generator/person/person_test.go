package person_test

import (
	"bytes"
	"log"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	fakerlog "github.com/ensoria/gofake/pkg/faker/common/log"
	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/person"
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/de_DE"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_CA"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_GB"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_IN"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_US"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/fr_FR"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/ja_JP"
	"github.com/ensoria/gofake/pkg/faker/testutil"
)

const kanaTestSeed = 42

func expectKanaUnavailable(locale *provider.Localized, method string, generate func(*person.Person) string) {
	GinkgoHelper()
	source := core.NewRand(util.SeededRand(kanaTestSeed))
	control := core.NewRand(util.SeededRand(kanaTestSeed))
	var warnings bytes.Buffer
	fakerlog.SetLogger(log.New(&warnings, "", 0))
	defer fakerlog.SetLogger(nil)

	Expect(generate(person.New(source, locale))).To(BeEmpty())
	Expect(strings.Count(warnings.String(), "\n")).To(Equal(1))
	Expect(warnings.String()).To(ContainSubstring("This method is not available for the current locale."))
	Expect(warnings.String()).To(ContainSubstring("(*Person)." + method + "]"))
	Expect(source.Num.Int()).To(Equal(control.Num.Int()))
}

var _ = Describe("Person", func() {
	localized := en_US.New()
	coreRand := core.NewRand(util.RandSeed())
	p := person.New(coreRand, localized)

	DescribeTable("handles unavailable kana names without consuming randomness",
		func(newLocale func() *provider.Localized) {
			locale := newLocale()
			expectKanaUnavailable(locale, "MaleKanaName", (*person.Person).MaleKanaName)
			expectKanaUnavailable(locale, "FemaleKanaName", (*person.Person).FemaleKanaName)
			expectKanaUnavailable(locale, "KanaName", (*person.Person).KanaName)
		},
		Entry("en_US", en_US.New),
		Entry("de_DE", de_DE.New),
		Entry("en_CA", en_CA.New),
		Entry("en_GB", en_GB.New),
		Entry("en_IN", en_IN.New),
		Entry("fr_FR", fr_FR.New),
		Entry("missing people provider", func() *provider.Localized { return &provider.Localized{} }),
	)

	It("supports explicit requests when only male kana names are configured", func() {
		locale := ja_JP.New()
		locale.People.CreateKanaNameFemale = nil
		expectKanaUnavailable(locale, "FemaleKanaName", (*person.Person).FemaleKanaName)
		expectKanaUnavailable(locale, "KanaName", (*person.Person).KanaName)
		Expect(person.New(coreRand, locale).MaleKanaName()).NotTo(BeEmpty())
	})

	It("supports explicit requests when only female kana names are configured", func() {
		locale := ja_JP.New()
		locale.People.MaleNameFormats = nil
		expectKanaUnavailable(locale, "MaleKanaName", (*person.Person).MaleKanaName)
		expectKanaUnavailable(locale, "KanaName", (*person.Person).KanaName)
		Expect(person.New(coreRand, locale).FemaleKanaName()).NotTo(BeEmpty())
	})

	It("preserves the seeded ja_JP kana sequence", func() {
		names := person.New(core.NewRand(util.SeededRand(kanaTestSeed)), ja_JP.New())
		Expect(names.MaleKanaName()).To(Equal("サカモト ヒデキ"))
		Expect(names.FemaleKanaName()).To(Equal("ヤマモト ヨウコ"))
		Expect(names.KanaName()).To(Equal("ヤマグチ ハルカ"))
	})

	Describe("FirstNameMale", func() {
		It("should return a male first name", func() {
			r := p.FirstNameMale()
			Expect(r).To(BeElementOf(en_US.FirstNameMales))
		})
	})

	Describe("FirstNameFemale", func() {
		It("should return a female first name", func() {
			r := p.FirstNameFemale()
			Expect(r).To(BeElementOf(en_US.FirstNameFemales))
		})
	})

	Describe("FirstName", func() {
		It("should return a first name", func() {
			r := p.FirstName()

			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.FirstName", r)
		})
	})

	Describe("LastName", func() {
		It("should return a last name", func() {
			r := p.LastName()
			Expect(r).To(BeElementOf(en_US.LastNames))
		})
	})

	Describe("TitleMale", func() {
		It("should return a male title", func() {
			r := p.TitleMale()
			Expect(r).To(BeElementOf(en_US.TitleMales))
		})
	})

	Describe("TitleFemale", func() {
		It("should return a female title", func() {
			r := p.TitleFemale()
			Expect(r).To(BeElementOf(en_US.TitleFemales))
		})
	})

	Describe("Title", func() {
		It("should return a title", func() {
			r := p.Title()

			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.Title", r)
		})
	})

	Describe("Suffix", func() {
		It("should return a suffix", func() {
			r := p.Suffix()
			Expect(r).To(BeElementOf(en_US.Suffixes))
		})
	})

	Describe("MaleName", func() {
		It("should return a male name", func() {
			r := p.MaleName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.MaleName", r)
		})
	})

	Describe("FemaleName", func() {
		It("should return a female name", func() {
			r := p.FemaleName()

			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.FemaleName", r)
		})
	})

	Describe("Name", func() {
		It("should return a person's full name", func() {
			r := p.Name()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.Name", r)
		})
	})

	Describe("SSN", func() {
		It("should return a person's SSN", func() {
			r := p.SSN()
			Expect(r).To(MatchRegexp(`^\d{3}-\d{2}-\d{4}$`))
		})
	})

	jaJP := ja_JP.New()
	pJaJP := person.New(coreRand, jaJP)
	Describe("FirstKanaNameMale", func() {
		It("should return a male first kana name", func() {
			r := pJaJP.FirstKanaNameMale()
			Expect(r).To(BeElementOf(ja_JP.FirstKanaNameMales))
		})
	})

	Describe("FirstKanaNameFemale", func() {
		It("should return a female first kana name", func() {
			r := pJaJP.FirstKanaNameFemale()
			Expect(r).To(BeElementOf(ja_JP.FirstKanaNameFemales))
		})
	})

	Describe("FirstKanaName", func() {
		It("should return a first kana name", func() {
			r := pJaJP.FirstKanaName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.FirstKanaName", r)
		})
	})

	Describe("LastKanaName", func() {
		It("should return a last kana name", func() {
			r := pJaJP.LastKanaName()
			Expect(r).To(BeElementOf(ja_JP.LastKanaNames))
		})
	})

	Describe("MaleKanaName", func() {
		It("should return a male kana name", func() {
			r := pJaJP.MaleKanaName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.MaleKanaName", r)
		})
	})

	Describe("FemaleKanaName", func() {
		It("should return a female kana name", func() {
			r := pJaJP.FemaleKanaName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.FemaleKanaName", r)
		})
	})

	Describe("KanaName", func() {
		It("should return a kana name", func() {
			r := pJaJP.KanaName()
			Expect(len(r)).To(BeNumerically(">", 0))
			testutil.Output("Person.KanaName", r)
		})
	})
})
