package en_GB

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createCompanies() *provider.Companies {
	return &provider.Companies{
		CompanyNames: []string{
			"Alder Fen Trading", "Amber Quay Surveying", "Beech Grove Publishing", "Bracken Vale Design",
			"Copper Lane Engineering", "Fern Brook Textiles", "Hazel Moor Foods", "Meadowgate Print",
			"North Beck Logistics", "Rowan Hill Ceramics", "Severn Vale Studios",
			"Willow Reach Consulting",
		},
		CompanySuffixes: []string{"Ltd", "plc", "LLP"},
		CompanyFormats:  []string{"{{.Name}} {{.Suffix}}"},
		CreateCompany:   localefmt.Company,
		JobTitleNames: []string{
			"Accountant", "Architect", "Building surveyor", "Civil engineer", "Community nurse",
			"Data analyst", "Electrician", "Graphic designer", "Librarian", "Mechanical engineer",
			"Occupational therapist", "Quantity surveyor", "Rail technician", "Software developer",
			"Town planner", "Veterinary surgeon",
		},
		JobTitleFormats: []string{"{{.}}"},
		CreateJobTitle:  localefmt.JobTitle,
	}
}
