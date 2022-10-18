package main

func judge(n int) (x bool) {
	if n%2 == 0 {
		x = true
	} else {
		x = false
	}
	return
}

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if judge(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}
