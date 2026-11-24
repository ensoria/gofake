package company_test

import (
	"bytes"
	"log"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	fakerlog "github.com/ensoria/gofake/pkg/faker/common/log"
	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/company"
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

var _ = Describe("Company", func() {

	localized := en_US.New()
	coreRand := core.NewRand(util.RandSeed())
	comp := company.New(coreRand, localized)

	Describe("Company", func() {
		It("CompanyName should return a company name", func() {
			r := comp.CompanyName()
			Expect(r).To(BeElementOf(en_US.CompanyNames))
		})

		It("CompanySuffix should return a company suffix", func() {
			r := comp.CompanySuffix()
			Expect(r).To(BeElementOf(en_US.CompanySuffixes))
		})

		It("Name should return a company name", func() {
			r := comp.Name()
			testutil.Output("Company.Name", r)
		})
	})

	Describe("JobTitle", func() {
		It("jobTitleName should return a job title name", func() {
			r := comp.JobTitleName()
			Expect(r).To(BeElementOf(en_US.JobTitleNames))
		})

		It("JobTitle should return a job title", func() {
			r := comp.JobTitle()
			testutil.Output("Company.JobTitle", r)
		})
	})

	Describe("EIN", func() {
		It("EINPrefix should return a EIN prefix", func() {
			r := comp.EINPrefix()
			Expect(r).To(BeElementOf(en_US.EINPrefixes))
		})
		It("generates a complete U.S. EIN", func() {
			Expect(comp.EIN()).To(MatchRegexp(`^[0-9]{2}-[0-9]{7}$`))
		})

		It("uses the EIN prefixes supplied by a custom locale", func() {
			custom := &provider.Localized{Companies: &provider.Companies{EINPrefixes: []string{"42"}}}
			Expect(company.New(coreRand, custom).EIN()).To(MatchRegexp(`^42-[0-9]{7}$`))
		})

		DescribeTable("returns an empty EIN for a locale without U.S. EIN data",
			func(newLocale func() *provider.Localized) {
				Expect(company.New(coreRand, newLocale()).EIN()).To(BeEmpty())
			},
			Entry("de_DE", de_DE.New),
			Entry("en_GB", en_GB.New),
			Entry("fr_FR", fr_FR.New),
			Entry("en_CA", en_CA.New),
			Entry("en_IN", en_IN.New),
		)

		It("returns an empty string for an empty prefix list", func() {
			custom := &provider.Localized{Companies: &provider.Companies{EINPrefixes: []string{}}}
			Expect(company.New(coreRand, custom).EIN()).To(BeEmpty())
		})

		It("returns an empty EIN when a custom locale omits company data", func() {
			Expect(company.New(coreRand, &provider.Localized{}).EIN()).To(BeEmpty())
		})

		It("logs one unavailable-locale warning from EIN", func() {
			var output bytes.Buffer
			fakerlog.SetLogger(log.New(&output, "", 0))
			DeferCleanup(func() { fakerlog.SetLogger(nil) })

			company.New(coreRand, ja_JP.New()).EIN()
			Expect(strings.Count(output.String(), "\n")).To(Equal(1))
			Expect(output.String()).To(ContainSubstring("This method is not available for the current locale."))
			Expect(output.String()).To(ContainSubstring("(*Company).EIN]"))
		})

		It("returns an empty string without advancing the random source for ja_JP", func() {
			const seed = 42
			source := core.NewRand(util.SeededRand(seed))
			control := core.NewRand(util.SeededRand(seed))
			Expect(company.New(source, ja_JP.New()).EIN()).To(BeEmpty())
			Expect(source.Num.Int()).To(Equal(control.Num.Int()))
		})
	})

	jaJP := ja_JP.New()
	compJaJP := company.New(coreRand, jaJP)
	Describe("ja_JP Company", func() {
		It("CompanyPrefix should return a company prefix", func() {
			r := compJaJP.CompanyPrefix()
			Expect(r).To(BeElementOf(ja_JP.CompanyPrefixes))
		})

		It("Name should return a company name", func() {
			r := compJaJP.Name()
			testutil.Output("Company.Name", r)
		})
	})
})
