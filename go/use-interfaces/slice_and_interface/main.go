package main

func main() {
	ns := []int{1, 2, 3, 4}

	// Compile error
	// cannot use ns (variable of type []int) as []interface{} value in variable declaration
	var vs []interface{} = ns
}
