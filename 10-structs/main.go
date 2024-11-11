package main

import "fmt"

type User struct{
	Name string
	Email string
	Status bool
	Age int
}


func main(){
	fmt.Println("Structs in golang")

	olowoleru := User{"olowoler","olowo@email.com,", true, 16}
	fmt.Println(olowoleru)
	fmt.Printf("hitesh details are: %+v\n", olowoleru)
	fmt.Printf("Name is %v and email is %v.", olowoleru.Email, olowoleru.Name)
} 

