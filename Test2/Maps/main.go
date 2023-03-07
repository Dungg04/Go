package main

import ("fmt")

func main() {
	var m1 = map[int] string {1:"A", 2:"B", 3:"C", 4:"D", 5:"E"}

	fmt.Println(m1)

	//create map using make() func
	var m2 =  make(map[int]string)
	fmt.Println(m2)
	m2[1] = "S"
	m2[2] = "G"
	m2[3] = "C"
	fmt.Println(m2)

	//check if key exists
	val, ok := m1[2]
	fmt.Println(val, ok)
	//check false
	valF, f := m1[7]
	fmt.Println(valF, f)

	//update
	m2[3] = "O"
	fmt.Println(m2)

	//delete
	delete(m1, 3)
	fmt.Println(m1)

}