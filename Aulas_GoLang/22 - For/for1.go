package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60}

	for key, value := range slice {
		total := 0
		total += value
		fmt.Println(key, value)
	}

	fmt.Println("Forma simplicada de utilizar o For")

	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}

}
