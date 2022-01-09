package main

func main() {
	var a = 200
	var b = &a
	println("value store in a is = ", a)
	println("address of a is = ", &a)
	println("address storein pointer is = ", b)

	println("value store in a(*b) before changing = ", *b)
	*b = 111
	println("value store in a(*b) after changing = ", a)
}
