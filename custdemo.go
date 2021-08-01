package main

import (
	"fmt"
)

type Address struct {
	hno     int
	village string
	state   string
}
type Customer struct {
	custname string
	salary   int
	Address
}

func main() {
	initialize()

}
func objreference(cust *Customer) {
	fmt.Println("customer name:", (*cust).custname, cust.salary, cust.Address)
}
func valueref(cust Customer) {
	fmt.Println("customer data", cust.custname)
}
func initialize() {
	cust1 := Customer{}
	cust1.custname = "Maheshbabu"
	cust1.salary = 25000
	cust1.hno = 27
	cust1.village = "grpt"
	cust1.state = "Telangana"
	fmt.Println(cust1)
	fmt.Println()
	add := Address{
		hno:     20,
		village: "grpt",
		state:   "telangana",
	}
	cust2 := Customer{
		custname: "Hemanth",
		salary:   26000,
		Address:  add,
	}
	fmt.Println(cust2)
	objreference(&cust2)
	valueref(cust2)
	//create a map
	//var datainfo map[int] Customer
	data := make(map[int]Customer)
	data[100] = cust1
	data[102] = cust1
	data[101] = cust2
	fmt.Println(data)
	//read
	var add1 = data[100]
	fmt.Println("customer info based on key 100: ", add1)
	//delete
	delete(data, 100)
	fmt.Println("deleted key is 100:", data)
	//update
	data[101] = cust1
	fmt.Println("updated key 101:", data)

}
