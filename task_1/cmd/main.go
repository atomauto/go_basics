package main

import (
	"fmt"
	"task_1/dummy"
)

func main() {
	fmt.Println("Test")
	d := dummy.New(100, 1000, 200, 300, false)
	d = dummy.Generate()

}
