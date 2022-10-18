package main

type MyInt int

func (m *MyInt) Inc() { *m += 1 }
func main() {
	var n MyInt = 0
	println(n)
	n.Inc()
	println(n)
}
