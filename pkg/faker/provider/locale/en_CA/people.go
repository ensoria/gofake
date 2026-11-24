package en_CA

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createPeople() *provider.People {
	const nameFormat = "{{.FirstName}} {{.LastName}}"
	return &provider.People{
		FirstNameMales: []string{
			"Aaron", "Adam", "Alexander", "Arjun", "Benjamin", "Cameron", "Daniel", "Ethan", "Gabriel",
			"Isaac", "Jacob", "Liam", "Lucas", "Noah", "Owen", "Ryan", "Samuel", "William",
		},
		FirstNameFemales: []string{
			"Abigail", "Amelia", "Avery", "Charlotte", "Chloe", "Emily", "Emma", "Grace", "Hannah",
			"Isabelle", "Jasmine", "Maya", "Nora", "Olivia", "Priya", "Sarah", "Sophie", "Zoe",
		},
		LastNames: []string{
			"Anderson", "Bouchard", "Campbell", "Chan", "Chen", "Clarke", "Fraser", "Gagnon", "Gill",
			"Grant", "Harris", "Li", "MacDonald", "Martin", "Nguyen", "Patel", "Robinson", "Singh",
			"Tremblay", "Wilson",
		},
		MaleNameFormats:   []string{nameFormat},
		CreateNameMale:    localefmt.MaleName,
		FemaleNameFormats: []string{nameFormat},
		CreateNameFemale:  localefmt.FemaleName,
	}
}
