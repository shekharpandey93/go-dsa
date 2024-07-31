package main

import "fmt"

func main()  {
	towerhanoi(2, "a", "b", "c")
}

func towerhanoi(n int, a,b,c string) {
	if n == 1 {
		fmt.Printf("Move 1 from %s to %s \n", a,c)
		return
	}
	towerhanoi(n-1,a, c, b )
	fmt.Printf("Move %d from %s to %s \n", n, a,c)
	towerhanoi(n-1,b, a, c )
}
