// Package localefmt supplies the data referenced by shared locale templates.
package localefmt

type fullName struct {
	FirstName string
	LastName  string
}

func MaleName(generator any) any {
	p := generator.(interface {
		FirstNameMale() string
		LastName() string
	})
	return &fullName{FirstName: p.FirstNameMale(), LastName: p.LastName()}
}

func FemaleName(generator any) any {
	p := generator.(interface {
		FirstNameFemale() string
		LastName() string
	})
	return &fullName{FirstName: p.FirstNameFemale(), LastName: p.LastName()}
}
