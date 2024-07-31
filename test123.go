package main

import "fmt"

func main() {
	//ch := make(chan bool)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go even1(ch, &wg)
	//go odd1(ch, &wg)
	//wg.Wait()

	input := []string{"nat", "top", "eat", "tea", "bat"}
	val := groupAnagrams(input)
	fmt.Println(val)
}

func groupAnagrams(strs []string) [][]string {
	mapData := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		if _, ok := mapData[strs[i]]; ok {
			map1 := make(map[byte]int)
			mapData[strs[i]] = append(mapData[strs[i]], strs[i])
		} else {
			mapData[strs[i]] = []string{strs[i]}
		}

	}

	for k, v := range mapData {
		mapData[k] = append(mapData[k], v...)
	}
	fmt.Println("mapData", mapData)

	return [][]string{}
}
