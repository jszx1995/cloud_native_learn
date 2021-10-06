package main

import "fmt"


func main()  {
	err, result := DuplicateString("aaa")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func DuplicateString(input string) (err error, result string) {
	if input == "aaa"{
		err = fmt.Errorf("aaa is not allowed")
		return
	}
	result = input + input
	return
}