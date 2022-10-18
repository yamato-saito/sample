package main

import "fmt"

func main() {
	//var n int = 100
	//const m int = 200

	for i := 0; i <= 100; i = i + 1 {
		fmt.Printf("%d-", i)
		if i%2 == 1 { //奇数
			fmt.Printf("奇数\n")
		} else { //偶数
			fmt.Printf("偶数\n")
		}
	}
	//fmt.Print("値段>")
	//fmt.Printf("%d円%d円引き\n", n, m)
}
