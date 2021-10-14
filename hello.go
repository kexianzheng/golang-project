package main

import "fmt"

func main() {
	res1, res2 := rectangle(6, 4)
	fmt.Println("zhouchang: ", res1, "mianji:", res2)
}
func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}
