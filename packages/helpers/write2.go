package helpers

import "fmt"

// write is an unexported function (lowercase).
// It's visible only within the helpers package.
func write(text string) {
	fmt.Println(text)
}
