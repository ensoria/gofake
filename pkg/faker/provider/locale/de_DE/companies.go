package de_DE

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createCompanies() *provider.Companies {
	return &provider.Companies{
		CompanyNames: []string{
			"Ahorn Handel", "Bergland Technik", "Birkenwerk", "Eichenhain Logistik", "Elbwiesen Verlag",
			"Feldrain Energie", "Hafenlicht Medien", "Isarwerk", "Lindenhof Textil", "Nordufer Bau",
			"Rheinbogen Keramik", "Wiesengrund Gartenbau",
		},
		CompanySuffixes: []string{"GmbH", "AG", "KG", "OHG"},
		CompanyFormats:  []string{"{{.Name}} {{.Suffix}}"},
		CreateCompany:   localefmt.Company,
		JobTitleNames: []string{
			"Architektin", "Bibliothekar", "Buchhalterin", "Datenanalyst", "Elektrotechnikerin",
			"Erzieher", "Fachinformatikerin", "Gärtner", "Ingenieurin", "Lehrer", "Mediengestalterin",
			"Pflegefachkraft", "Projektleiter", "Übersetzerin", "Verfahrenstechniker", "Werkzeugmacherin",
		},
		JobTitleFormats: []string{"{{.}}"},
		CreateJobTitle:  localefmt.JobTitle,
	}
}
