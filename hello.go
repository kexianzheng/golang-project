package main

import "fmt"

func main() {
<<<<<<< HEAD
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
=======
	res := getSum()
	fmt.Println("1-10dehe: ", res)
}
func getSum() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
>>>>>>> qiusushu
}
