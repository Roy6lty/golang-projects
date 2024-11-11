package main

import "fmt"

type User struct {
	Name string
	Email string
	Status bool
	Age int

}

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang;No super or parent

	olowoleru := User{"Olowoleru", "olowoleru@mail.com", true, 16}
	fmt.Println(olowoleru)
	fmt.Printf("olowoleru details are: %+v\n", olowoleru)
	fmt.Printf("Name is %v and email is %v.", olowoleru.Name, olowoleru.Email)
	olowoleru.GetStatus()
	olowoleru.NewMail()
}

// a method is prt of a class 
// golang does not have classes but they have structs
//GetStatus is a method for user struct
func (u User) GetStatus(){
	fmt.Println("Is usee active: ", u.Status)

}

// getters mutates a copy of the struct object
func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)

}