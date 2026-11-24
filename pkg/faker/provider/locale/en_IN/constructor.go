// Package en_IN provides localized fake data for India in English.
package en_IN

import "github.com/ensoria/gofake/pkg/faker/provider"

// New creates the en_IN locale.
func New() *provider.Localized {
	return &provider.Localized{
		People:       createPeople(),
		Addresses:    createAddresses(),
		Companies:    createCompanies(),
		PhoneNumbers: CreatePhoneNumbers(),
	}
}
