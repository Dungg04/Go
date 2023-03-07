package sort

import (
	"fmt"
	"sort"
)

func main() {
	var slide = []int{}
	fmt.Printf("Nhap n = ")
	var n int
	fmt.Scanf("%d", &n)
	var input int 
	for i:=0; i<n; i++ {
		fmt.Printf("slide[%d] = ", i)
		fmt.Scanf("%d", &input)
		slide = append(slide, input)
	}
	bubbleSoft(slide)
	fmt.Println(slide)
}

func bubbleSoft(slide sort.IntSlice) {
	for j:=0; j<len(slide)-1; j++ {
		for i:=len(slide)-2; i>=j; i-- {
			if slide[i] < slide[i+1] {
				temp := slide[i]
				slide[i] = slide[i+1]
				slide[i+1] = temp
			}
		}
	}
}