func maxArea(height []int) int {
    total := 0
	temp := 0
	t := 1
    for i:=0; i<len(height)-1; i++ {
        for j:=i+1; j<len(height); j++ {
			t = j-i
            if height[i] >= height[j] {
                temp = height[j]*t
            } else {
				temp = height[i]*t
			}

            if(total < temp) {
                total = temp
            }
        }
    }

    return total
}