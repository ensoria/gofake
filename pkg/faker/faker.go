package faker

import (
	"github.com/ensoria/gofake/pkg/faker/common/util"
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
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/global"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_US"
)

// Faker is the main entry point that aggregates all fake data generators.
//
// すべてのフェイクデータジェネレーターを集約するメインのエントリーポイント構造体。
type Faker struct {
	rand        *core.Rand
	person      *person.Person
	color       *color.Color
	address     *address.Address
	barcode     *barcode.Barcode
	company     *company.Company
	file        *file.File
	image       *image.Image
	internet    *internet.Internet
	lorem       *lorem.Lorem
	medical     *medical.Medical
	payment     *payment.Payment
	phoneNumber *phonenumber.PhoneNumber
	userAgent   *useragent.UserAgent
}

// REF: https://fakerphp.github.io/

// Create creates a new Faker instance with the default en_US locale.
//
// デフォルトのen_USロケールで新しいFakerインスタンスを作成する。
func Create() *Faker {
	localized := en_US.New()
	return CreateWithLocale(localized)
}

// CreateWithLocale creates a new Faker instance with the specified locale data.
//
// 指定されたロケールデータで新しいFakerインスタンスを作成する。
func CreateWithLocale(localized *provider.Localized) *Faker {
	return newFaker(core.NewRand(util.RandSeed()), localized)
}

// CreateWithSeed creates a new Faker instance with the default en_US locale and
// a fixed seed, producing deterministic output (useful for tests / golden files).
//
// デフォルトのen_USロケールと固定シードで新しいFakerインスタンスを作成する。
// 出力は決定的になる（テストやゴールデンファイル向け）。
func CreateWithSeed(seed int64) *Faker {
	return CreateWithLocaleAndSeed(en_US.New(), seed)
}

// CreateWithLocaleAndSeed creates a new Faker instance with the specified locale
// data and a fixed seed, producing deterministic output.
//
// 指定されたロケールデータと固定シードで新しいFakerインスタンスを作成する。
func CreateWithLocaleAndSeed(localized *provider.Localized, seed int64) *Faker {
	return newFaker(core.NewRand(util.SeededRand(seed)), localized)
}

// newFaker wires all generators onto the given core.Rand and locale.
//
// 与えられた core.Rand とロケールで全ジェネレーターを組み立てる内部ヘルパー。
func newFaker(coreRand *core.Rand, localized *provider.Localized) *Faker {
	global := global.New()
	return &Faker{
		rand:        coreRand,
		barcode:     barcode.New(coreRand),
		color:       color.New(coreRand, global),
		person:      person.New(coreRand, localized),
		address:     address.New(coreRand, localized),
		company:     company.New(coreRand, localized),
		file:        file.New(coreRand, global),
		image:       image.New(coreRand, global),
		internet:    internet.New(coreRand, global),
		lorem:       lorem.New(coreRand, global),
		medical:     medical.New(coreRand, global),
		payment:     payment.New(coreRand, global),
		phoneNumber: phonenumber.New(coreRand, global, localized),
		userAgent:   useragent.New(coreRand, global),
	}
}

// Bool returns the generator for booleans.
func (f *Faker) Bool() *core.RandBool {
	return f.rand.Bool
}

// Num returns the generator for numbers.
func (f *Faker) Num() *core.RandNum {
	return f.rand.Num
}

// Str returns the generator for strings.
func (f *Faker) Str() *core.RandStr {
	return f.rand.Str
}

// Time returns the generator for times and durations.
func (f *Faker) Time() *core.RandTime {
	return f.rand.Time
}

// Slice returns the generator for sampling slices.
func (f *Faker) Slice() *core.RandSlice {
	return f.rand.Slice
}

// Map returns the generator for sampling maps.
func (f *Faker) Map() *core.RandMap {
	return f.rand.Map
}

// Person returns the generator for personal data.
func (f *Faker) Person() *person.Person {
	return f.person
}

// Color returns the generator for colors.
func (f *Faker) Color() *color.Color {
	return f.color
}

// Address returns the generator for addresses.
func (f *Faker) Address() *address.Address {
	return f.address
}

// Barcode returns the generator for barcodes.
func (f *Faker) Barcode() *barcode.Barcode {
	return f.barcode
}

// Company returns the generator for company data.
func (f *Faker) Company() *company.Company {
	return f.company
}

// File returns the generator for file data.
func (f *Faker) File() *file.File {
	return f.file
}

// Image returns the generator for images.
func (f *Faker) Image() *image.Image {
	return f.image
}

// Internet returns the generator for internet data.
func (f *Faker) Internet() *internet.Internet {
	return f.internet
}

// Lorem returns the generator for placeholder text.
func (f *Faker) Lorem() *lorem.Lorem {
	return f.lorem
}

// Medical returns the generator for medical data.
func (f *Faker) Medical() *medical.Medical {
	return f.medical
}

// Payment returns the generator for payment data.
func (f *Faker) Payment() *payment.Payment {
	return f.payment
}

// PhoneNumber returns the generator for phone numbers.
func (f *Faker) PhoneNumber() *phonenumber.PhoneNumber {
	return f.phoneNumber
}

// UserAgent returns the generator for user agent strings.
func (f *Faker) UserAgent() *useragent.UserAgent {
	return f.userAgent
}
