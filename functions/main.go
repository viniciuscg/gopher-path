package main

import "fmt"

func main() {
	sum := func(x, y int) int {
		return x + y
	}

	fmt.Println(sum(10, 20))

	result := sum(20, 20)
	fmt.Println(result)

	text := printText("This is a function that do and return something")

	resultSub, resultSum := maths(10, 20)
	fmt.Println(resultSub, resultSum)

	resultSubstract, _ := maths(10, 20)
	fmt.Println(resultSubstract)
}

func printText(text string) string {
	fmt.Println(text)
	return text
}

func maths(x, y int) (int, int) {
	var sub int = x - y
	var sum int = x + y
	return sub, sum
}
