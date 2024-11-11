package main

import "fmt"

// using const you always start with the Capital letter to indicate it
const LoginToken string = "login"

func main(){
	var username string ="olowoleru"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type %T \n", smallVal)

	var smallFloat float32 = 255.45545542234
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type %T \n", smallFloat)

	var bigFloat float64 = 255.45545542234
	fmt.Println(bigFloat)
	fmt.Printf("variable is of type %T \n", bigFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type %T \n", anotherVariable)

	// implicit type
	
	var website = "http://www.example.com"
	fmt.Println(website)

	//no var style
	numberOfUser := 3000000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)

	

};