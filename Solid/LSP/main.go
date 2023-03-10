package main

import (
	"fmt"
)

type transport interface {
	getName() string
}


type kind struct {
	name string
}

func (k kind) getName() string {
	return k.name
}

type cut struct {
	kind
	age int
	sounding string
}

type dog struct {
	kind 
	age int
}

type printer struct {}

func (printer) transportName(p transport) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	k := kind{name: "TT"}

	c := cut {
		kind: kind{name: "Tom"},
		age: 2,
		sounding: "mewmew",
	}

	d := dog {
		kind: kind{name: "Sonwy"},
		age: 3,
	}

	p := printer{}
	p.transportName(k)
	p.transportName(c)
	p.transportName(d)
}