package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main()  {
	//names := []string{"bin", "nua", "wouju"}
	//names = append(names, "jina")
	//bin := map[string] string {"name":"bin", "age":"32"}
	//for key, value := range bin {
	//	fmt.Println(key, value)
	//}
	//fmt.Println(names)
	//fmt.Println(bin)
	favFood := [] string {"수육", "치킨", "피자", "햄버거"}
	bin := person{"bin", 32, favFood}
	fmt.Println(bin)
}