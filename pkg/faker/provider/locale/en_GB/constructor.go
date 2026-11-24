// Package en_GB provides localized fake data for the United Kingdom.
package en_GB

import "github.com/ensoria/gofake/pkg/faker/provider"

// New creates the en_GB locale.
func New() *provider.Localized {
	return &provider.Localized{
		People:       createPeople(),
		Addresses:    createAddresses(),
		Companies:    createCompanies(),
		PhoneNumbers: createPhoneNumbers(),
	}
}
