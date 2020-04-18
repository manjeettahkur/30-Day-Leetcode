package __to7

//Given an array of strings, group anagrams together.
//
//Example:
//
//Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Output:
//[
//["ate","eat","tea"],
//["nat","tan"],
//["bat"]
//]
//
//Note:
//
//All inputs will be in lowercase.
//The order of your output does not matter.

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var arr [][]string
	hash := make(map[string][]string, len(strs))

	for i := 0; i < len(strs); i++ {
		s := strings.Split(strs[i], "")
		sort.Strings(s)
		rs := strings.Join(s, "")
		if val, ok := hash[rs]; ok {
			hash[rs] = append(val, strs[i])
		} else {
			var hashArr []string
			hashArr = append(hashArr, strs[i])
			hash[rs] = hashArr
		}
	}

	for _, v := range hash {
		arr = append(arr, v)
	}
	return arr
}
