package fr_FR

import (
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/locale/internal/localefmt"
)

func createCompanies() *provider.Companies {
	return &provider.Companies{
		CompanyNames: []string{
			"Atelier des Ormes", "Bois des Collines", "Céramique du Val", "Éditions du Verger",
			"Énergie des Rives", "Filature des Prés", "Jardins du Levant", "Les Ateliers du Canal",
			"Papeterie des Saules", "Services du Bocage", "Textiles des Dunes",
			"Transports de la Clairière",
		},
		CompanySuffixes: []string{"SARL", "SAS", "SA", "EURL"},
		CompanyFormats:  []string{"{{.Name}} {{.Suffix}}"},
		CreateCompany:   localefmt.Company,
		JobTitleNames: []string{
			"Architecte", "Bibliothécaire", "Chargée de projet", "Comptable", "Conductrice de travaux",
			"Développeur logiciel", "Électricienne", "Enseignant", "Géomètre", "Graphiste", "Infirmière",
			"Ingénieur civil", "Mécanicienne", "Pharmacien", "Technicienne de laboratoire", "Urbaniste",
		},
		JobTitleFormats: []string{"{{.}}"},
		CreateJobTitle:  localefmt.JobTitle,
	}
}
