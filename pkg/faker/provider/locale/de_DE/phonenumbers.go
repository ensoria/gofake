package de_DE

import "github.com/ensoria/gofake/pkg/faker/provider"

// Fictional phone number ranges: https://bundesnetzagentur.de/DE/Fachthemen/Telekommunikation/Nummerierung/DramaNumbers/artikel.html
func createPhoneNumbers() *provider.PhoneNumbers {
	return &provider.PhoneNumbers{Formats: []string{"030 23125###", "040 66969###"}}
}
