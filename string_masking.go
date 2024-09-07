package main

import (
	"fmt"
)

func main() {
	input := "00010110"
	mask := "XX0XXX11"
	op := applyMasking(input, mask)
	fmt.Println(op)
}

func applyMasking(input string, mask string) string {
	var result string
	for idx, val := range mask {
		if val == 'X' {
			result += string(input[idx])
			continue
		}
		result += string(mask[idx])
	}
	return result

	close()
}
