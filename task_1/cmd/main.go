package main

import (
	"fmt"
	"task_1/dummy"
)

func main() {
	fmt.Println("Test")
	d := dummy.New(100, 1000, 200, 300, false)
	dummy.Generate(d)
	// u := model.NewUser("Vasya", "89162463456", "vasok@mail.ru")
	// fmt.Println(u)

}
