// Package en_CA provides localized fake data for English-speaking Canada.
package en_CA

import "github.com/ensoria/gofake/pkg/faker/provider"

// New creates the en_CA locale.
func New() *provider.Localized {
	return &provider.Localized{
		People:       createPeople(),
		Addresses:    createAddresses(),
		Companies:    createCompanies(),
		PhoneNumbers: createPhoneNumbers(),
	}
}
