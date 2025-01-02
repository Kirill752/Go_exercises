package main

func removeElement(nums []int, val int) int {
	var k int = 0
	for i, v := range nums {
		if v == val {
			nums[i] = -1 // так как 0 <= nums[i] <= 50
		} else {
			k++
		}
	}
	// fmt.Println(nums)
	for L := 0; L < k; L++ {
		if nums[L] == -1 {
			for i := L; i < len(nums); i++ {
				if nums[i] != -1 {
					swap(nums, L, i)
					// fmt.Println(nums)
					break
				}
			}
		}
	}
	// fmt.Println(nums)
	return k
}
