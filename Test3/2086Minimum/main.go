package main

import "fmt"

func main() {
	fmt.Printf("Nap chuoi: ")
	var s string
	fmt.Scanf("%s", s)

	count = minimumBuckets(s)
}

func minimumBuckets(hamsters string) int {
    count := 0
    n = len(hamsters)
     for i:=0; i<n; i++ {
         if hamsters[i]=='H' {
            if (i+1)<n && hamsters[i+1]=='.' {
                count++
                i+=3
            } else if (i-1)>=0 && hamsters[i-1]=='.' {
                count++
                i+=1
            } else {
                return -1
            }
         } else {
             i++
         }

         
	}
	return count
}