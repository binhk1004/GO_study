package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

func main()  {
	fmt.Println(multiply(2, 2))
	totalLenght, upperName := lenAndUpper("bin hyun")
	fmt.Println(totalLenght, upperName)
}