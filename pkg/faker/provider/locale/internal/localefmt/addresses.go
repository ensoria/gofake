package localefmt

func City(generator any) any {
	return generator.(interface{ CityName() string }).CityName()
}

func Street(generator any) any {
	return generator.(interface{ StreetName() string }).StreetName()
}

func StreetAddress(generator any) any {
	a := generator.(interface {
		Street() string
		BuildingNumber() string
	})
	return &struct{ Street, Number string }{Street: a.Street(), Number: a.BuildingNumber()}
}

type address struct {
	StreetAddress string
	City          string
	Postcode      string
	State         string
}

func addressParts(generator any) *address {
	a := generator.(interface {
		StreetAddress() string
		City() string
		Postcode() string
	})
	return &address{StreetAddress: a.StreetAddress(), City: a.City(), Postcode: a.Postcode()}
}

func Address(generator any) any {
	return addressParts(generator)
}

func AddressWithState(generator any) any {
	data := addressParts(generator)
	data.State = generator.(interface{ State() string }).State()
	return data
}

func AddressWithStateAbbr(generator any) any {
	data := addressParts(generator)
	data.State = generator.(interface{ StateAbbr() string }).StateAbbr()
	return data
}
