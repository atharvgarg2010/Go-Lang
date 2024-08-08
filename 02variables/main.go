package main

import "fmt"

func main() {
	// Variable Types in Go Lang
	var username string = "Atharv"
	fmt.Println(username)
	fmt.Printf("Type of Var: %T \n", username)

	var IsLogined bool = true
	fmt.Println(IsLogined)
	fmt.Printf("Type of Var: %T \n", IsLogined)

	var SmallVal uint8 = 200
	fmt.Println(SmallVal)
	fmt.Printf("Type of Var: %T \n", SmallVal)

	var SmallFloat float64 = 0.2000000000100001
	fmt.Println(SmallFloat)
	fmt.Printf("Type of Var: %T \n", SmallFloat)

	// Default Values Of Datatypes

	var otherVariable int
	fmt.Println(otherVariable)
	fmt.Printf("Type of Var: %T \n", otherVariable)
	// 0

	var AnotherVariable string
	fmt.Println(AnotherVariable)
	fmt.Printf("Type of Var: %T \n", AnotherVariable)

	// other Way of Variables

	var Website = "Atharv.com"
	fmt.Println(Website)

	// other Way of Variables

	Users := 0.11111
	fmt.Println(Users)

}
