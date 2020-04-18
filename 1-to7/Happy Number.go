package __to7

//Write an algorithm to determine if a number n is "happy".
//
//A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
//
//Return True if n is a happy number, and False if not.

func isHappy(n int) bool {

	hash := make(map[int]bool)
	sum, tmp, x := 0, 0, n
	hash[n] = true
	for ; ;{
		sum, x = 0, n
		for ; x > 0 ; {
			tmp = x%10
			sum += tmp*tmp
			x = (x - tmp)/10
		}
		_ , exists := hash[sum]
		if exists {
			break
		}
		n = sum
		hash[n] = true
	}
	if 1 == n {
		return true
	}
	return false
}