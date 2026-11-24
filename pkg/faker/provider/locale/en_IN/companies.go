package en_IN

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createCompanies() *provider.Companies {
	return &provider.Companies{
		CompanyNames: []string{
			"Amber Loom Textiles", "Banyan Grove Systems", "Copper Leaf Engineering",
			"Deccan Field Logistics", "Hillcrest Ceramics", "Jasmine Valley Foods", "Lotus Creek Design",
			"Mango Grove Trading", "Monsoon Ridge Energy", "Neem Leaf Publishing", "Riverbank Analytics",
			"Saffron Field Services",
		},
		CompanySuffixes: []string{"Private Limited", "Limited", "LLP"},
		CompanyFormats:  []string{"{{.Name}} {{.Suffix}}"},
		CreateCompany:   localefmt.Company,
		JobTitleNames: []string{
			"Accountant", "Architect", "Civil engineer", "Company secretary", "Data analyst",
			"Electrical engineer", "Graphic designer", "Laboratory technician", "Lecturer", "Nurse",
			"Operations manager", "Pharmacist", "Project coordinator", "Software developer",
			"Textile designer", "Urban planner",
		},
		JobTitleFormats: []string{"{{.}}"},
		CreateJobTitle:  localefmt.JobTitle,
	}
}
