package main

import (
	"sync"
	"task_3/task_3/service"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	SolutionsServer := service.NewSolutionServer(&wg)
	go SolutionsServer.Serve()
	Api := service.NewApiServer(&wg)
	wg.Add(1)
	Api.GetTask()
	wg.Wait()
}
