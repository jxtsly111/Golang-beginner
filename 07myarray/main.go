package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golangs")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "peach"

	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Fruit list is: ", len(fruitlist))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)

}
