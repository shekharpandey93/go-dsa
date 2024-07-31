package main

//[30][30]board;
func main()  {
	board := [30] int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	dataMap := make(map[int]bool)

	for jdx := 0; jdx < len(board) ; jdx = jdx*jdx {
		for idx := 0; idx < len(board) ; idx {
			if _,ok:=dataMap[idx]; !ok {
				dataMap[idx] = false
			} else {
				dataMap[idx] = !dataMap[idx]
			}
		}
	}
	//for key := range dataMap {
	//	fmt.Println("dataMap",dataMap[key])
	//}
}






