package en_IN_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/phonenumber"
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/global"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/en_IN"
)

var _ = Describe("en_IN phone numbers", func() {
	It("generates all combinations of leading digit and digit grouping", func() {
		const (
			seed                 = 42
			sampleCount          = 256
			expectedCombinations = 12
			domesticPhonePattern = `^[6-9]([0-9]{9}|[0-9]{4} [0-9]{5}|[0-9]{3} [0-9]{3} [0-9]{3})$`
		)

		localized := &provider.Localized{PhoneNumbers: en_IN.CreatePhoneNumbers()}
		numbers := phonenumber.New(core.NewRand(util.SeededRand(seed)), global.New(), localized)
		seen := make(map[string]struct{}, expectedCombinations)

		for range sampleCount {
			number := numbers.PhoneNumber()
			Expect(number).To(MatchRegexp(domesticPhonePattern))

			spaceCount := strings.Count(number, " ")
			combination := fmt.Sprintf("%c/%d spaces", number[0], spaceCount)
			seen[combination] = struct{}{}
		}

		Expect(seen).To(HaveLen(expectedCombinations))
	})
})
