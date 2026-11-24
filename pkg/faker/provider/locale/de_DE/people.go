package de_DE

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createPeople() *provider.People {
	const nameFormat = "{{.FirstName}} {{.LastName}}"
	return &provider.People{
		FirstNameMales: []string{
			"Alexander", "Andreas", "Ben", "Emil", "Elias", "Felix", "Florian", "Friedrich", "Jakob",
			"Jonas", "Jörg", "Lukas", "Matthias", "Moritz", "Niklas", "Paul", "Sebastian", "Tim",
		},
		FirstNameFemales: []string{
			"Amelie", "Anna", "Charlotte", "Clara", "Elisabeth", "Emilia", "Frieda", "Hannah", "Johanna",
			"Julia", "Klara", "Lena", "Lina", "Luise", "Marlene", "Nele", "Sophie", "Theresa",
		},
		LastNames: []string{
			"Bauer", "Becker", "Brandt", "Braun", "Fischer", "Fuchs", "Hartmann", "Hoffmann", "Keller",
			"König", "Krüger", "Lange", "Müller", "Richter", "Schäfer", "Schmidt", "Schneider",
			"Schröder", "Vogel", "Weber",
		},
		MaleNameFormats:   []string{nameFormat},
		CreateNameMale:    localefmt.MaleName,
		FemaleNameFormats: []string{nameFormat},
		CreateNameFemale:  localefmt.FemaleName,
	}
}
