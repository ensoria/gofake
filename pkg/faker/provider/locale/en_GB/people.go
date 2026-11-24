package en_GB

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createPeople() *provider.People {
	const nameFormat = "{{.FirstName}} {{.LastName}}"
	return &provider.People{
		FirstNameMales: []string{
			"Alfie", "Arthur", "Benjamin", "Callum", "Daniel", "Dylan", "Edward", "Finlay", "George",
			"Harry", "Henry", "Idris", "Isaac", "James", "Lewis", "Oliver", "Rhys", "Thomas",
		},
		FirstNameFemales: []string{
			"Alice", "Amelia", "Charlotte", "Daisy", "Eleanor", "Elsie", "Emily", "Freya", "Grace",
			"Harriet", "Isla", "Jasmine", "Maisie", "Megan", "Olivia", "Phoebe", "Ruby", "Sophie",
		},
		LastNames: []string{
			"Ahmed", "Bennett", "Clarke", "Davies", "Edwards", "Evans", "Fraser", "Griffiths", "Hughes",
			"Jones", "Kaur", "Khan", "MacLeod", "Morgan", "Patel", "Reid", "Taylor", "Walker", "Williams",
			"Wright",
		},
		MaleNameFormats:   []string{nameFormat},
		CreateNameMale:    localefmt.MaleName,
		FemaleNameFormats: []string{nameFormat},
		CreateNameFemale:  localefmt.FemaleName,
	}
}
