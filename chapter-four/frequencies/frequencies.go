// Package set solves exercise 4.9 4-11.
package frequencies

// majorityElementUsingMap returns the element that appears more than n/2 times.
// Time: O(N)
// Space: O(N)
func majorityElementUsingMap(in []int) (int, bool) {
	frequencies := make(map[int]int)
	for _, num := range in {
		frequencies[num]++
	}

	for num, frequency := range frequencies {
		if frequency > len(in)/2 {
			return num, true
		}
	}

	return 0, false
}

// majorityElementUsingStack returns the element that appears more than n/2 times.
// Time: O(N)
// Space: O(N)
func majorityElementUsingStack(in []int) (int, bool) {
	nums := make([]int, len(in))
	copy(nums, in)

	var stack []int
	for len(nums) != 0 {
		if len(stack) != 0 && len(nums) != 0 {
			if stack[len(stack)-1] != nums[0] {
				stack = stack[0 : len(stack)-1]
				nums = nums[1:]
			} else {
				stack = append(stack, nums[0])
				nums = nums[1:]
			}
		} else if len(nums) != 0 {
			stack = append(stack, nums[0])
			nums = nums[1:]
		}
	}

	if len(stack) == 0 {
		return 0, false
	}

	var count int
	for _, num := range in {
		if stack[0] == num {
			count++
		}
	}

	if count > len(in)/2 {
		return stack[0], true
	}

	return 0, false
}
