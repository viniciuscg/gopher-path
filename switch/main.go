package main

import "fmt"

func main() {
	// --- BASIC SWITCH USAGE ---

	// Declare a variable to switch on
	day := "Monday"

	// The basic switch structure
	switch day {
	case "Monday":
		fmt.Println("Start of the week!") // This will be printed
	case "Friday":
		fmt.Println("End of the week!")
	default:
		fmt.Println("Midweek or weekend.")
	}

	// --- MULTIPLE CASE VALUES ---

	// You can match multiple values in a single case
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Weekday.") // This will be printed for "Monday"
	}

	// --- SWITCH WITHOUT A CONDITION ---

	// This form is useful as a cleaner version of chained if-else statements
	age := 20
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teen")
	default:
		fmt.Println("Adult") // Output: Adult
	}

	// --- SWITCH WITH SHORT VARIABLE DECLARATION ---

	// You can initialize and switch in one line
	switch score := 85; {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B") // Output: Grade: B
	default:
		fmt.Println("Grade: C or lower")
	}

	// --- FALLTHROUGH ---

	// In Go, switch does NOT fall through by default.
	// But you can force it with the 'fallthrough' keyword.
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two") // This will print due to fallthrough
	default:
		fmt.Println("Other")
	}

	// --- TYPE SWITCH ---

	// Type switch is used with interfaces to determine the actual type
	var x interface{} = "hello"

	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v) // Output: string: hello
	default:
		fmt.Println("unknown type")
	}

	// --- SWITCH IN LOOPS ---

	// You can use switch inside loops for input handling or decision trees
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		default:
			fmt.Println("Something else")
		}
	}
}
