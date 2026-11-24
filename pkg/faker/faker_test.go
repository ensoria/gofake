package faker_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/address"
	"github.com/ensoria/gofake/pkg/faker/generator/barcode"
	"github.com/ensoria/gofake/pkg/faker/generator/color"
	"github.com/ensoria/gofake/pkg/faker/generator/company"
	"github.com/ensoria/gofake/pkg/faker/generator/file"
	"github.com/ensoria/gofake/pkg/faker/generator/image"
	"github.com/ensoria/gofake/pkg/faker/generator/internet"
	"github.com/ensoria/gofake/pkg/faker/generator/lorem"
	"github.com/ensoria/gofake/pkg/faker/generator/medical"
	"github.com/ensoria/gofake/pkg/faker/generator/payment"
	"github.com/ensoria/gofake/pkg/faker/generator/person"
	"github.com/ensoria/gofake/pkg/faker/generator/phonenumber"
	"github.com/ensoria/gofake/pkg/faker/generator/useragent"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/ja_JP"
)

func expectSameGenerator[T any](name string, get func() *T) {
	GinkgoHelper()
	first := get()
	Expect(first).NotTo(BeNil(), name)
	Expect(get()).To(BeIdenticalTo(first), name)
}

func expectStableGenerators(f *faker.Faker) {
	GinkgoHelper()
	expectSameGenerator[core.RandBool]("Bool", f.Bool)
	expectSameGenerator[core.RandNum]("Num", f.Num)
	expectSameGenerator[core.RandStr]("Str", f.Str)
	expectSameGenerator[core.RandTime]("Time", f.Time)
	expectSameGenerator[core.RandSlice]("Slice", f.Slice)
	expectSameGenerator[core.RandMap]("Map", f.Map)
	expectSameGenerator[person.Person]("Person", f.Person)
	expectSameGenerator[color.Color]("Color", f.Color)
	expectSameGenerator[address.Address]("Address", f.Address)
	expectSameGenerator[barcode.Barcode]("Barcode", f.Barcode)
	expectSameGenerator[company.Company]("Company", f.Company)
	expectSameGenerator[file.File]("File", f.File)
	expectSameGenerator[image.Image]("Image", f.Image)
	expectSameGenerator[internet.Internet]("Internet", f.Internet)
	expectSameGenerator[lorem.Lorem]("Lorem", f.Lorem)
	expectSameGenerator[medical.Medical]("Medical", f.Medical)
	expectSameGenerator[payment.Payment]("Payment", f.Payment)
	expectSameGenerator[phonenumber.PhoneNumber]("PhoneNumber", f.PhoneNumber)
	expectSameGenerator[useragent.UserAgent]("UserAgent", f.UserAgent)
}

var _ = DescribeTable("Faker constructors expose stable generators",
	func(create func() *faker.Faker) {
		expectStableGenerators(create())
	},
	Entry("Create", faker.Create),
	Entry("CreateWithLocale", func() *faker.Faker { return faker.CreateWithLocale(ja_JP.New()) }),
	Entry("CreateWithSeed", func() *faker.Faker { return faker.CreateWithSeed(testSeed) }),
	Entry("CreateWithLocaleAndSeed", func() *faker.Faker { return faker.CreateWithLocaleAndSeed(ja_JP.New(), testSeed) }),
)
