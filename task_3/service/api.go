package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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
type InputTask3 [][]interface{}

func NewApiServer(wg *sync.WaitGroup) *Api {
	return &Api{wg}
}

func (s *Api) GetAndCheckTask(taskName string) {
	defer s.wg.Done()

	//I really disrespect that variable declaration inside if or case in golang has own scope, WHY???
	if taskName == TASK1 {
		r, _ := http.Get(TASK_HOST + taskName)
		var input Input
		json.NewDecoder(r.Body).Decode(&input)
		var results []int
		task1 := model.Task1{}
		for _, A := range input {
			results = append(results, task1.Solution(A[0]))
		}
		requestData := model.ApiRequestData{
			UserName: USERNAME,
			Task:     taskName,
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
		fmt.Print(apiRequestAnswer)
	}
	if taskName == TASK2 {
		r, _ := http.Get(TASK_HOST + taskName)
		var input InputTask3
		json.NewDecoder(r.Body).Decode(&input)
		var results [][]int
		task2 := model.Task2{}
		for _, data := range input {
			rawArray := reflect.ValueOf(data[0])
			rawK := reflect.ValueOf(data[1])

			K := int(rawK.Float())
			array := make([]int, rawArray.Len())

			for i := 0; i < rawArray.Len(); i++ {
				array[i] = int(rawArray.Index(i).Interface().(float64))
			}
			results = append(results, task2.Solution(array, K))
		}
		requestData := model.ApiRequestData{
			UserName: USERNAME,
			Task:     taskName,
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
		fmt.Print(apiRequestAnswer)
	}
	if taskName == TASK3 {
		r, _ := http.Get(TASK_HOST + taskName)
		var input Input
		json.NewDecoder(r.Body).Decode(&input)
		var results []int
		task3 := model.Task3{}
		for _, A := range input {
			results = append(results, task3.Solution(A[0]))
		}
		requestData := model.ApiRequestData{
			UserName: USERNAME,
			Task:     taskName,
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
		fmt.Print(apiRequestAnswer)
	}
	if taskName == TASK4 {
		r, _ := http.Get(TASK_HOST + taskName)
		var input Input
		json.NewDecoder(r.Body).Decode(&input)
		var results []int
		task4 := model.Task4{}
		for _, A := range input {
			results = append(results, task4.Solution(A[0]))
		}
		requestData := model.ApiRequestData{
			UserName: USERNAME,
			Task:     taskName,
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
		fmt.Print(apiRequestAnswer)
	}

}

func (s *Api) CheckAllTasks() {
	s.wg.Add(4)
	go s.GetAndCheckTask(TASK1)
	go s.GetAndCheckTask(TASK2)
	go s.GetAndCheckTask(TASK3)
	go s.GetAndCheckTask(TASK4)
}
