package en_GB

import "github.com/ensoria/gofake/pkg/faker/provider"

// Fictional phone number ranges: https://www.ofcom.org.uk/phones-and-broadband/phone-numbers/numbers-for-drama
func createPhoneNumbers() *provider.PhoneNumbers {
	return &provider.PhoneNumbers{Formats: []string{"020 7946 0###", "07700 900###"}}
}
