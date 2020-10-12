package main

import (
	"fmt"
	"github.com/binhk1004/mydict"
)

func main()  {
	//account := accounts.NewAccount("bin hyun")
	//account.Deposit(10)
	//fmt.Println(account.String())
	dictionary := mydict.Dictionary{}
	//definition, err := dictionary.Search("second")
	//if err != nil{
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(definition)
	//}
	baseWord := "hello"
	//definition := "Greeting"
	//err := dictionary.Add(word, definition)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//hello, _ := dictionary.Search(word)
	//fmt.Println("found: ", word, "definition: ", hello)
	//err2 := dictionary.Add(word, definition)
	//if err2 != nil{
	//	fmt.Println(err2)
	//}
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	//err := dictionary.Update(baseWord, "Second")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
