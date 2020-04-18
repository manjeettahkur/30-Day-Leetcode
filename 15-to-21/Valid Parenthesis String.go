package _5_to_21

// Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:
//
// Any left parenthesis '(' must have a corresponding right parenthesis ')'.
// Any right parenthesis ')' must have a corresponding left parenthesis '('.
// Left parenthesis '(' must go before the corresponding right parenthesis ')'.
// '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
// An empty string is also valid.
//
// Example 1:
//
// Input: "()"
// Output: True
//
// Example 2:
//
// Input: "(*)"
// Output: True
//
// Example 3:
//
// Input: "(*))"
// Output: True

// Approach 1

func CheckValidString1(s string) bool {
	var low int
	var high int
	for _, letter := range s {
		if letter == '(' {
			low += 1
			high += 1
		} else if letter == ')' {
			high -= 1
			low -= 1
		} else {
			low -= 1
			high += 1
		}

		if high < 0 {
			return false
		}

		if low < 0 {
			low = 0
		}
	}
	return low == 0
}

// Approach 2

func checkValidString2(s string) bool {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	length := len(s)

	if s == "" {
		return true
	}

	for i := 0; i < length; i++ {

		if s[i] == '(' {
			s1 = append(s1, i)

		} else if s[i] == ')' {
			if len(s1) > 0 {
				s1 = s1[:len(s1)-1]
			} else if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			} else {
				return false
			}
		} else if s[i] == '*' {
			s2 = append(s2, i)
		}
	}

	for len(s1) > 0 && len(s2) > 0 {
		val := s2 [len(s2)-1]
		s2 = s2[:len(s2)-1]
		if s1[len(s1)-1] < val {
			s1 = s1[:len(s1)-1]
		}
	}

	if len(s1) < 1 {
		return true
	}
	return false

}
