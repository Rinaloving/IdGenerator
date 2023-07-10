package main

import (
	"C"
	"fmt"
)

func main() {
	// UnRegister()
	var res int32 = RegisterOne(C.CString("127.0.0.1:6379"), "", 1, "", 63, 30, 15)
	//Validate(2)
	// var res int = Add(2, 5)
	fmt.Printf("res: %d\n", res)
	fmt.Println("hello go")
}

// type Books struct {
// 	Name  string
// 	Price int
// }

func TestAdd() {
	fmt.Println("print hello go")
}
