package chapters

import "fmt"

type Person struct{
	FirstName string
	MiddleName string
	FamilyName string
}

func (p *Person) Equals(other Person) bool {
	return other.FirstName == p.FirstName && other.FamilyName == p.FamilyName && other.MiddleName == p.MiddleName
}

func NewPerson(firstName, lastName, middleName string) *Person {
	return &Person{
		FirstName: firstName,
		MiddleName: middleName,
		FamilyName: lastName,
	}
}

func Execute() {
	tohru :=  NewPerson("Tohru", "Yaginuma", "")
	shizuka :=  NewPerson("Tohru", "Yaginuma", "")

	fmt.Println(tohru.Equals(*shizuka))
}