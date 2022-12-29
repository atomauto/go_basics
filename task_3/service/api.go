package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"task_3/task_3/model"
)

const (
	TASK1      = "Проверка последовательности"
	TASK2      = "Циклическая ротация"
	TASK3      = "Чудные вхождения в массив"
	TASK4      = "Поиск отсутствующего элемента"
	TASK_HOST  = "https://kuvaev-ituniversity.vps.elewise.com/tasks/"
	CHECK_HOST = "https://kuvaev-ituniversity.vps.elewise.com/tasks/solution"
	USERNAME   = "atomauto"
)

type Api struct {
	wg *sync.WaitGroup
}
type Input [][][]int

func NewApiServer(wg *sync.WaitGroup) *Api {
	return &Api{wg}
}

func (s *Api) GetTask() {
	defer s.wg.Done()
	r, _ := http.Get(TASK_HOST + TASK1)
	var input Input
	json.NewDecoder(r.Body).Decode(&input)
	var results []int
	task1 := model.Task1{}
	for _, A := range input {
		results = append(results, task1.Solution(A[0]))
	}
	fmt.Print(results)

	requestData := model.ApiRequestData{
		UserName: USERNAME,
		Task:     TASK1,
		Results: &model.Results{
			Payload: input,
			Results: results,
		},
	}
	body, _ := json.Marshal(requestData)
	r, err := http.Post(CHECK_HOST, "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}

	var apiRequestAnswer model.ApiRequestAnswer
	json.NewDecoder(r.Body).Decode(&apiRequestAnswer)
	fmt.Println("Answer from server")
	fmt.Print(apiRequestAnswer)

}
