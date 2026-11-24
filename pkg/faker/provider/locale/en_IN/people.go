package en_IN

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createPeople() *provider.People {
	const nameFormat = "{{.FirstName}} {{.LastName}}"
	return &provider.People{
		FirstNameMales: []string{
			"Aarav", "Aditya", "Arjun", "Dev", "Farhan", "Harish", "Imran", "Ishan", "Kabir", "Karan",
			"Manoj", "Nikhil", "Pranav", "Rahul", "Rohan", "Sameer", "Varun", "Vivek",
		},
		FirstNameFemales: []string{
			"Aditi", "Ananya", "Anjali", "Anusha", "Diya", "Farah", "Ishita", "Kavya", "Lakshmi", "Leela",
			"Meera", "Nandini", "Neha", "Priya", "Riya", "Sana", "Sneha", "Vidya",
		},
		LastNames: []string{
			"Bose", "Das", "Desai", "Gupta", "Iyer", "Joshi", "Kaur", "Khan", "Kulkarni", "Menon", "Nair",
			"Patel", "Pillai", "Rao", "Reddy", "Sen", "Shah", "Sharma", "Singh", "Verma",
		},
		MaleNameFormats:   []string{nameFormat},
		CreateNameMale:    localefmt.MaleName,
		FemaleNameFormats: []string{nameFormat},
		CreateNameFemale:  localefmt.FemaleName,
	}
}
