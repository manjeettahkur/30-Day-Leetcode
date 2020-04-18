package __to14

//Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
//
//Example 1:
//
//Input: [0,1]
//Output: 2
//Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
//
//Example 2:
//
//Input: [0,1,0]
//Output: 2
//Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
//
//Note: The length of the given binary array will not exceed 50,000.


func findMaxLength(nums []int) int {

	hash := make(map[int]int, len(nums))
	hash[0] = -1
	var maxLength int
	var count int


	for i:=0; i< len(nums); i++{
		if nums[i] == 0 {
			count += -1
		} else {
			count += 1
		}

		if _, ok := hash[count]; ok {
			maxLength = max(maxLength, i - hash[count])

		} else {
			hash[count] = i
		}

	}
	return maxLength

}

//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}