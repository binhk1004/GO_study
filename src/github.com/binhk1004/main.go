package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (lenght int, uppercase string){
	defer fmt.Println("끝났습니다.")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers{
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func checkAge(age int) bool {
	switch age {
	case 10 :
		return false
	case 18 :
		return true
	}
	return false
}
func main()  {
	//fmt.Println(multiply(2, 2))
	//totalLenght, upperName := lenAndUpper("bin hyun")
	//fmt.Println(totalLenght, upperName)
	//repeatMe("bin", "jun", "nua", "woju", "jina")
	//result := superAdd(1,2,3,4,5,6,7,8,9)
	//fmt.Println(result)
	fmt.Println(canIDrink(16))
	fmt.Println(checkAge(18))
}