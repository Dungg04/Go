package main

import (
	"fmt"
)

type data interface {
	getData() []string
}

type sql struct {}

func (db sql) getData() []string {
	return []string{"db1", "db2", "db3"}
}

type repo struct {
	dt data
}

func (r repo) readData() []string {
	return r.dt.getData()
}

func main() {
	s := sql{}
	r := repo{dt: s}
	fmt.Println(r.readData())
}