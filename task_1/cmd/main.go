package main

import (
	"fmt"
	"task_1/dummy"
)

func main() {
	var d dummy.Dummy
	fmt.Println("Test")
	dummy.New()
	d := dummy.New(100, 1000, 200, 300, false)

}
