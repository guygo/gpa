package main

import "fmt"

type location struct {
	latitude  int
	longitude int
}
type Address struct {
	street string `constraint:"Primary Key"`
	city   string `constraint:"NOT NULL"`
	loc    []location
}
type person struct {
	name    string
	age     int
	address Address
}

func (p *person) addAdress(street string, city string) {
	p.address.city = city
	p.address.street = street
}

func main() {

	// p := person{"guy", 4, Address{}}
	// var rep Repo
	// rep.init()
	// fmt.Print(rep.createTable(p))

	// rep.insertTable(p)
	// rep.createTable(Address{"hacramel", "tel", []location{{1, 2}}})
	// rep.createTable(location{3, 4})
	// for k, v := range rep.tableList {
	// 	fmt.Printf("t:%s command:%s\n", k, v.sqlCommand)
	// }
	// fmt.Printf("%+v", rep.createRealations())
	// var p person
	// p.addAdress("vinkler", "PT")
	// fmt.Printf("%+v", p)
	var p *person = &person{}
	fmt.Printf("%T", p)
}
