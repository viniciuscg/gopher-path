package main

import (
	"fmt"
)

func main() {
	implicityType := "This type of varible as a implicitytype"			
	fmt.Println(implicityType)

	variable1, variable2 := "This is the first variable", "This is the second variable"
	fmt.Println(variable1, variable2)

	//This is how go makes easy to use temp variable
	variable1, variable2 = variable2, variable1 
	fmt.Println(variable1, variable2)

	var typedVariable string = "This is how typed variable looks like"
	fmt.Println(typedVariable)

	//the same thing but for typed variables.
	var (
		typedVariable1 string = "This is the first typed variable"
		typedVariable2 string = "This is the second typed variable"
	)
	
	typedVariable1, typedVariable2 = typedVariable2, typedVariable1 
	fmt.Println(variable1, variable2)
}