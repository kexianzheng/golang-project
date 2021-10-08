package main

import "fmt"

func main() {
	for i := 2; i <= 100; i++ {
		flag := true //jilujshifouweisushu
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false //
				break
			}
		}
		if flag { //==true
			fmt.Println(i)
		}
	}
}
