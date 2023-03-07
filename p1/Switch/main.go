package main

import ("fmt")

func main() {
	fmt.Printf("Nhap thang: ")
	var month int
	fmt.Scanf("%d", &month)

	switch month {
		case 1:
			fmt.Println("Januaty")
		case 2:
			fmt.Println("February")
		case 3:
			fmt.Println("March")
		case 4:
			fmt.Println("April")
		case 5:
			fmt.Println("May")
		case 6:
			fmt.Println("June")
	}
}