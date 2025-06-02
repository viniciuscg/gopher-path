package helpers

import "fmt"

// Write prints the first string, then delegates the second string
// to an internal unexported function.
func Write(text1, text2 string) {
	fmt.Println(text1)
	write(text2) // call private function from same package
}