package main
import "fmt"
func twoSum(nums []int, target int) []int {
    map1 := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if val, ok := map1[diff]; ok {
			return []int{i, val}
		}
		map1[diff] = i
	}
	return []int{-1, -1}
}

func main(){
    nums := [4]int{2, 7, 11, 15}
	target := 9
	twoSum(nums[:], target)
}