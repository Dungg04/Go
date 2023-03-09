func maxArea(height []int) int {
    start := 0
    end := len(height) - 1
    maxArea := 0
    for start < end {
        currArea := end - start
        if height[start] < height[end] {
            currArea *= height[start]
            start++
        } else {
            currArea *= height[end]
            end--
        }
        if maxArea < currArea {
            maxArea = currArea
        }
    }
    
    return maxArea
}


func PredictTheWinner(nums []int) bool {
    cases := make(map[string]int)
    score := getScore(nums, 0, len(nums) - 1, cases)
    return score >= sum(nums) - score
}

func getScore(nums []int, i, j int, cases map[string]int) int {
    if i > j { return 0 }
    if i == j { return nums[i] }
    key := strconv.Itoa(i) + "," + strconv.Itoa(j)
    if val, ok := cases[key]; ok { return val }
    leftLeft := getScore(nums, i + 2, j, cases)
    leftRight := getScore(nums, i + 1, j - 1, cases)
    rightRight := getScore(nums, i, j - 2, cases)
    takeLeft := nums[i] + min(leftLeft, leftRight)
    takeRight := nums[j] + min(leftRight, rightRight)
    final := max(takeLeft, takeRight)
    cases[key] = final
    return final
}

func sum(a []int) int {
    res := 0
    for _, val := range a { res += val }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func nextGreaterElements(nums []int) []int {
    n := len(nums)
    res, stack := make([]int, n), []int{}

    for i := 0; i < n * 2; i++ {
        num := nums[i % n]

        for len(stack) > 0 && num > nums[stack[len(stack) - 1]] {
            index := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            res[index] = num
        }

        if i < n {
            stack = append(stack, i)
        }
    }

    for i := 0; i < len(stack); i++ {
        index := stack[i]
        res[index] = -1
    }

    return res
}