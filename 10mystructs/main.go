package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang

	sly := User{"Sly", "jxtsly@gmail.com", true, 16}
	fmt.Println(sly)
	fmt.Printf("sly details are: %+v\n", sly)
	fmt.Printf("Name is %v and email is %v.", sly.Name, sly.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
