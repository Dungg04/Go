package main

import (
	"fmt"
)

type dog struct {
	name string
}

func (d dog) sounding() string {
	temp := d.name + " sounding gougou"
	return temp
	// fmt.Println(d.name, "sounding gougou")
}

type cut struct {
	name string
}

func (c cut) sounding() string {
	// fmt.Println(c.name, "sounding mewmew")
	temp := c.name + " sounding mewmew"
	return temp
}

type species interface {
	sounding() string
}

type outPrinter struct{}

func (o outPrinter) toText(s species) string {
	return s.sounding()
}

func main() {
	d := dog{name: "Sonwy"}
	d.sounding()

	c := cut{name: "Tom"}
	c.sounding()

	o := outPrinter{}
	fmt.Println(o.toText(d))
	fmt.Println(o.toText(c))
}
