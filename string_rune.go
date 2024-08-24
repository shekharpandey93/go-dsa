package main

import (
	"log"
)

func main() {
	sample := "aabb$c"
	log.Printf(" input ===> %s", sample)
	samplerune := []rune(sample)

	for idx := 0; idx < len(samplerune)/2; idx++ {
		samplerune[idx], samplerune[len(samplerune)-1-idx] = samplerune[len(samplerune)-1-idx], samplerune[idx]
	}
	log.Printf("output ==> %s", string(samplerune))
	sampleMap := make(map[rune]bool)
	outPutrune := make([]rune, len(sampleMap))

	for _, val := range samplerune {
		if !sampleMap[val] {
			sampleMap[val] = true
			outPutrune = append(outPutrune, val)
		}

	}
	log.Printf("output ==> %s", string(outPutrune))

}
