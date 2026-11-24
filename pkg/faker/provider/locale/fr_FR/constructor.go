// Package fr_FR provides localized fake data for France.
package fr_FR

import "github.com/ensoria/gofake/pkg/faker/provider"

// New creates the fr_FR locale.
func New() *provider.Localized {
	return &provider.Localized{
		People:       createPeople(),
		Addresses:    createAddresses(),
		Companies:    createCompanies(),
		PhoneNumbers: createPhoneNumbers(),
	}
}
