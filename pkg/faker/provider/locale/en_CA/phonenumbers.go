package en_CA

import "github.com/ensoria/gofake/pkg/faker/provider"

// Fictional phone number ranges: https://www.cnac.ca/other_codes/555/555_line_numbers.htm
func createPhoneNumbers() *provider.PhoneNumbers {
	return &provider.PhoneNumbers{Formats: []string{"416-555-01##", "604-555-01##"}}
}
