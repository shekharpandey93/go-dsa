package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)
// capital
// lower
// number
// speical char
// min 6 max 16
func main()  {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@$!%*?&"
	randomPW := getRandomString(chars, 1)
	val := validatePW(randomPW)
	if val {
		fmt.Printf("Give Password %s is valid", randomPW)
	}
	
}

func getRandomString(chars string, pwLength int) string  {
	b := make([]byte, pwLength)
	//fmt.Println("11111",rand.New())
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = chars[seededRand.Intn(len(chars))]
	}
	fmt.Println(string(b))
	return string(b)
}

func validatePW(randomPW string) bool {
	r, err := regexp.Compile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]`)
	fmt.Println("errrr",err)
	return r.MatchString(randomPW)
}
