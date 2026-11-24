package fr_FR

import "github.com/ensoria/gofake/pkg/faker/provider"

// Fictional phone number ranges: https://www.arcep.fr/uploads/tx_gsavis/25-2215.pdf#page=82
func createPhoneNumbers() *provider.PhoneNumbers {
	return &provider.PhoneNumbers{Formats: []string{"01 99 00 ## ##", "06 39 98 ## ##"}}
}
