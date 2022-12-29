package service

import (
	"fmt"
	"net/http"
	"sync"
)

type SolutionServer struct {
	handler    *http.ServeMux
	wg         *sync.WaitGroup
	apiHandler *Api
}

func NewSolutionServer(wg *sync.WaitGroup) *SolutionServer {
	return &SolutionServer{http.NewServeMux(), wg, NewApiServer(wg)}
}

func (s *SolutionServer) solveAllTasks(w http.ResponseWriter, r *http.Request) {
	s.apiHandler.CheckAllTasks()
}

func (s *SolutionServer) solveTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	s.apiHandler.GetAndCheckTask(r.RequestURI)
}

func (s *SolutionServer) Serve() {
	defer s.wg.Done()
	s.handler.HandleFunc("/tasks", s.solveAllTasks)
	s.handler.HandleFunc("/task/", s.solveTask)
	http.ListenAndServe(":8080", s.handler)
}
