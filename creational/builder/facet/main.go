package facet

import "fmt"

type Person struct {
	// Personal details
	name, address, pin string
	// Job details
	workAddress, company, position string
	salary                         int
}

type PersonBuilder struct {
	person *Person
}
type PersonAddressBuilder struct {
	PersonBuilder
}
type PersonJobBuilder struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (a *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*a}
}

func (a *PersonAddressBuilder) At(addr string) *PersonAddressBuilder {
	a.person.address = addr
	return a
}

func (a *PersonAddressBuilder) WithPostalCode(pin string) *PersonAddressBuilder {
	a.person.pin = pin
	return a
}

func (j *PersonJobBuilder) As(position string) *PersonJobBuilder {
	j.person.position = position
	return j
}

// For adds company to person
func (j *PersonJobBuilder) For(company string) *PersonJobBuilder {
	j.person.company = company
	return j
}

// In adds company address to person
func (j *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	j.person.workAddress = companyAddress
	return j
}

// WithSalary adds salary to person
func (j *PersonJobBuilder) WithSalary(salary int) *PersonJobBuilder {
	j.person.salary = salary
	return j
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func RunBuilderFacet() {
	pb := NewPersonBuilder()
	pb.Lives().
		At("Bangalore").
		WithPostalCode("1234").
		Works().
		As("LE").
		For("Last9").
		In("KTM").
		WithSalary(100)

	person := pb.Build()
	fmt.Println(person)
}
