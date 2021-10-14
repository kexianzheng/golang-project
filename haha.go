package main

import "fmt"

func main() {
	res2, res4 := rectangle(6, 4)
	fmt.Println("zhouzhang", res2, "mianji", res4)
}
func rectangle(len, wid float64) (peri float64, area float64) {
	peri = (len + wid) * 2
	area = len * wid
	return
}
