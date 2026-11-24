package faker_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/person"
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_US"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/ja_JP"
)

const (
	testSeed     = 42
	sampleCount  = 20
	numberLimit  = 1_000_000
	stringLength = 12
)

var _ = Describe("Seeded Faker", func() {
	DescribeTable("uses one random stream for core and locale generators",
		func(create func() *faker.Faker, locale func() *provider.Localized) {
			f := create()
			source := core.NewRand(rand.New(rand.NewSource(testSeed)))
			names := person.New(source, locale())

			expectStableGenerators(f)
			for range sampleCount {
				Expect(f.Num().IntBt(0, numberLimit)).To(Equal(source.Num.IntBt(0, numberLimit)))
				Expect(f.Str().AlphaFixedLength(stringLength)).To(Equal(source.Str.AlphaFixedLength(stringLength)))
				Expect(f.Person().Name()).To(Equal(names.Name()))
			}
		},
		Entry("default en_US", func() *faker.Faker { return faker.CreateWithSeed(testSeed) }, en_US.New),
		Entry("explicit ja_JP", func() *faker.Faker { return faker.CreateWithLocaleAndSeed(ja_JP.New(), testSeed) }, ja_JP.New),
	)

	It("produces the same sequence for the same seed", func() {
		a := faker.CreateWithSeed(testSeed)
		b := faker.CreateWithSeed(testSeed)

		for range sampleCount {
			Expect(a.Num().IntBt(0, numberLimit)).To(Equal(b.Num().IntBt(0, numberLimit)))
			Expect(a.Str().AlphaFixedLength(stringLength)).To(Equal(b.Str().AlphaFixedLength(stringLength)))
			Expect(a.Person().Name()).To(Equal(b.Person().Name()))
		}
	})

	It("produces different sequences for different seeds", func() {
		a := faker.CreateWithSeed(1)
		b := faker.CreateWithSeed(2)

		var first, second []int
		for range sampleCount {
			first = append(first, a.Num().IntBt(0, numberLimit))
			second = append(second, b.Num().IntBt(0, numberLimit))
		}
		Expect(first).NotTo(Equal(second))
	})
})
