package main

import "fmt"

func main() {
	fmt.Println("Hello!!")
	fmt.Print("Nhap a = ")
	var a int;
	fmt.Scanf("%d", &a)
	fmt.Print("Nhap b = ")
	var b int;
	fmt.Scanf("%d", &b)

	if a<b {
		fmt.Println(b, "-", a, "=", total(a,b))
	}
 	
	if a>=b {
		fmt.Println(a, "-", b, "=", total(a,b))
	}

}

func total(a, b int) int {
	if a<b {
		return b-a;
	}
	return a - b;
}