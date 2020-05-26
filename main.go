package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	err := p.SetTwitterHandler("@jam_wils")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twiter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.Country())
	println(p.FullName())
}
