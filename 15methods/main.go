package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang

	sly := User{"Sly", "jxtsly@gmail.com", true, 16}
	fmt.Println(sly)
	fmt.Printf("sly details are: %+v\n", sly)
	fmt.Printf("Name is %v and email is %v.\n", sly.Name, sly.Email)
	sly.GetStatus()
	sly.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
