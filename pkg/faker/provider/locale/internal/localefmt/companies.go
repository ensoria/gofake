package localefmt

func Company(generator any) any {
	c := generator.(interface {
		CompanyName() string
		CompanySuffix() string
	})
	return &struct{ Name, Suffix string }{Name: c.CompanyName(), Suffix: c.CompanySuffix()}
}

func JobTitle(generator any) any {
	return generator.(interface{ JobTitleName() string }).JobTitleName()
}
