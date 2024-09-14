package main

import "fmt"

const Login string = "Divi" //public

func main() {
	var username string = "Divi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.474747474
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.474747474
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	//default variables and alias
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	//implicit type
	var implicitVar = "Divi"
	fmt.Println(implicitVar)

	//no var style
	implicitVar2 := 5000
	fmt.Println(implicitVar2)
}
