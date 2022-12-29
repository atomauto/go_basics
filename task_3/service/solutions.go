package service

import (
	"fmt"
	"net/http"
	"sync"
)

type SolutionServer struct {
	handler *http.ServeMux
	wg      *sync.WaitGroup
}

func NewSolutionServer(wg *sync.WaitGroup) *SolutionServer {
	return &SolutionServer{http.NewServeMux(), wg}
}

func (*SolutionServer) solveAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SOLVING ALL TASK")
}

func (*SolutionServer) solveTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Solving Task")
}

func (s *SolutionServer) Serve() {
	defer s.wg.Done()
	s.handler.HandleFunc("/tasks", s.solveAllTasks)
	s.handler.HandleFunc("/task/", s.solveTask)
	http.ListenAndServe(":8080", s.handler)
}
