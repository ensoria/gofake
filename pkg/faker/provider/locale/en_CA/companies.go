package en_CA

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createCompanies() *provider.Companies {
	return &provider.Companies{
		CompanyNames: []string{
			"Aspen Ridge Surveying", "Birch Lake Design", "Cedar Bay Logistics",
			"Copper Spruce Engineering", "Elm Prairie Foods", "Harbour Pine Textiles",
			"Lake Trail Publishing", "Maple Shore Ceramics", "North Marsh Energy", "River Bend Software",
			"Spruce Grove Consulting", "Willow Creek Forestry",
		},
		CompanySuffixes: []string{"Inc.", "Ltd.", "Corp."},
		CompanyFormats:  []string{"{{.Name}} {{.Suffix}}"},
		CreateCompany:   localefmt.Company,
		JobTitleNames: []string{
			"Accountant", "Architect", "Civil engineer", "Community planner", "Data analyst",
			"Electrician", "Environmental technician", "Forestry technician", "Graphic designer",
			"Land surveyor", "Librarian", "Occupational therapist", "Paramedic", "Registered nurse",
			"Software developer", "Water treatment operator",
		},
		JobTitleFormats: []string{"{{.}}"},
		CreateJobTitle:  localefmt.JobTitle,
	}
}
