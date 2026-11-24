// Package de_DE provides localized fake data for Germany.
package de_DE

import "github.com/ensoria/gofake/pkg/faker/provider"

// New creates the de_DE locale.
func New() *provider.Localized {
	return &provider.Localized{
		People:       createPeople(),
		Addresses:    createAddresses(),
		Companies:    createCompanies(),
		PhoneNumbers: createPhoneNumbers(),
	}
}
