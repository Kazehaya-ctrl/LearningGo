package main

import "fmt"

func add (a int, b int) int {
	return a + b;
}

func loop () {
	for i := 0; i <= 9; i++ {
		fmt.Println(i)
	}
}

func whileLoop () {
	age := 100
	for age > 0 {
		fmt.Println(1)
	}
}

func ifelseif (value *int) bool {
	if *value < 4 {
		return true
	} else {
		return false
	}
}

func pointer (var1 *int) {
	*var1 = *var1 + 1;
	fmt.Println(*var1)
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(2)
	fmt.Println(add(2, 4))
	loop()
	// whileLoop()
	var value int = 123;
	fmt.Println(ifelseif(&value))
	fmt.Println(value)
	pointer(&value);
	fmt.Println(value)
}