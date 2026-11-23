package en_IN

import "github.com/ensoria/gofake/pkg/faker/provider"

// CreatePhoneNumbers provides synthetic ten-digit phone formats for en_IN.
func CreatePhoneNumbers() *provider.PhoneNumbers {
	return &provider.PhoneNumbers{
		Formats: []string{
			"6#########", "7#########", "8#########", "9#########",
			"6#### #####", "7#### #####", "8#### #####", "9#### #####",
			"6### ### ###", "7### ### ###", "8### ### ###", "9### ### ###",
		},
	}
}
