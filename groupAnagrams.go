package main

func main()  {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	groupAnagrams(strs)
}

func groupAnagrams(strs []string) [][]string {
	//mapOne := map[[26]int][]string{}
	//for _, str := range strs {
	//	tempKey := [26]int{}
	//	for i := 0; i < len(str); i++ {
	//		tempKey[str[i]-'a']++
	//	}
	//	mapOne[tempKey] = append(mapOne[tempKey], str)
	//}
	//fmt.Println("mpmpmpmp",mapOne)
	//res := [][]string{}
	//for _, val := range mapOne {
	//	res = append(res, val)
	//}
	//return res

	groups := map[[26]int][]string{}
	for _, str := range strs {
		var occ [26]int   // chars a-z mapped to occurance count
		for _, char := range str {
			occ[int(char) - int('a')]++
		}
		// update or create groups
		if _, ok := groups[occ]; ok {
			groups[occ] = append(groups[occ], str)
		} else {
			groups[occ] = []string{ str }
		}
	}
	result := [][]string{}
	for k, _ := range groups {
		result = append(result, groups[k])
	}
	return result
}
