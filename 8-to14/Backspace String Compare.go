package __to14

//Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
//
//Note that after backspacing an empty text, the text will continue empty.
//
//Example 1:
//
//Input: S = "ab#c", T = "ad#c"
//Output: true
//Explanation: Both S and T become "ac".
//
//Example 2:
//
//Input: S = "ab##", T = "c#d#"
//Output: true
//Explanation: Both S and T become "".
//
//Example 3:
//
//Input: S = "a##c", T = "#a#c"
//Output: true
//Explanation: Both S and T become "c".
//
//Example 4:
//
//Input: S = "a#c", T = "b"
//Output: false
//Explanation: S becomes "c" while T becomes "b".
//

func backspaceCompare(S string, T string) bool {
	// make two stack

	s1 := make([]byte, 0)
	s2 := make([]byte, 0)

	for c1 := range S{
		if S[c1] != '#' {
			s1 = append(s1, S[c1])
		} else {
			if len(s1) > 0 {
				s1 = s1[:len(s1)-1]
			}
		}
	}

	for c2 := range T {
		if T[c2] != '#'{
			s2 = append(s2, T[c2])
		} else {
			if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			}
		}
	}
	return string(s1) == string(s2)
}
