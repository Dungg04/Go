package struct 

import ("fmt")

type Person struct {
	id int
	name string
	age int
}

func main() {
	var p Person
	fmt.Printf("Nhap id: ")
	fmt.Scanf("%d", &p.id)
	fmt.Printf("Nhap ten: ")
	fmt.Scanf("%s", &p.name)
	fmt.Printf("Nhap tuoi: ")
	fmt.Scanf("%d", &p.age)

	fmt.Println("id: ", p.id)
	fmt.Println("name: ", p.name)
	fmt.Println("age: ", p.age)
}