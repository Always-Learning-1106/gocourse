package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

func (p Person) fullName() string {
	return fmt.Sprintf("%v %v", p.firstName, p.lastName)
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

type Address struct {
	city    string
	country string
}
type PhoneHomeCell struct {
	home string
	cell string
}

func main() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "Tampa",
			country: "USA",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "123456",
			cell: "789",
		},
	}
	p1 := Person{
		firstName: "Jane",
		lastName:  "Virgin",
		age:       25,
		address: Address{
			city:    "St. Pete",
			country: "Russia",
		},
	}
	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	// Anonymous structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoEmail@d.com",
	}
	fmt.Println(user)
	fmt.Println(p.fullName())
	fmt.Println(p1.fullName())
	fmt.Println(user.email)
	fmt.Println(p1.age)
	p1.incrementAgeByOne()
	fmt.Println(p1.age)
	fmt.Println(p1.address.city)
	fmt.Println(p.address.city)
}
