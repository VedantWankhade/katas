package katas

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	p  person
	id int
}

// embed person
type employee struct {
	empId int
	person
}

type reader interface {
	read() string
}

type writer interface {
	write(s string)
}

type readwriter interface {
	reader
	writer
}

type io struct {
}

func (d io) read() string {
	fmt.Println("reading")
	return "read"
}

func (d io) write(s string) {
	fmt.Println("writing")
}

func struct_interface_embedding() {
	s1 := student{
		id: 55,
		p: person{
			name: "x",
			age:  22,
		},
	}

	fmt.Println(s1.id)
	fmt.Println(s1.p.name)
	// fmt.Println(s1.name) // no workie

	e1 := employee{
		empId: 6666,
		person: person{
			name: "xsd",
			age:  22,
		},
	}

	fmt.Println(e1.empId)
	// fmt.Println(e1.p.name) // no workie
	fmt.Println(e1.name)

	e2 := employee{
		23232,
		person{
			"asfg",
			22,
		},
	}
	fmt.Printf("%+v\n", e2)

	//- ---
	socket := io{}
	useSocket(socket)
}

func useSocket(d io) {
	d.read()
	d.write("test")
}
