package fr_FR

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createPeople() *provider.People {
	const nameFormat = "{{.FirstName}} {{.LastName}}"
	return &provider.People{
		FirstNameMales: []string{
			"Adrien", "Antoine", "Bastien", "Benoît", "Clément", "Émile", "Étienne", "François",
			"Gabriel", "Hugo", "Jules", "Louis", "Lucien", "Mathieu", "Nicolas", "Olivier", "Raphaël",
			"Théo",
		},
		FirstNameFemales: []string{
			"Adèle", "Agathe", "Alice", "Amélie", "Anaïs", "Camille", "Chloé", "Claire", "Élodie", "Inès",
			"Jeanne", "Julie", "Léa", "Lucie", "Maëlle", "Manon", "Margaux", "Zoé",
		},
		LastNames: []string{
			"André", "Bernard", "Bertrand", "Chevalier", "Dubois", "Dupont", "Durand", "Faure",
			"Fontaine", "François", "Garnier", "Lefèvre", "Leroy", "Martin", "Mercier", "Moreau",
			"Perrin", "Petit", "Richard", "Roux",
		},
		MaleNameFormats:   []string{nameFormat},
		CreateNameMale:    localefmt.MaleName,
		FemaleNameFormats: []string{nameFormat},
		CreateNameFemale:  localefmt.FemaleName,
	}
}
