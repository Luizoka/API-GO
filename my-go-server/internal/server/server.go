package server

import (
    "fmt"
    "net/http"
)

type Server struct {
    port int
}

func NewServer(port int) *Server {
    return &Server{port: port}
}

func (s *Server) Start() {
    http.HandleFunc("/", s.handleRequest)
    addr := fmt.Sprintf(":%d", s.port)
    fmt.Printf("Server is listening on port %d\n", s.port)
    if err := http.ListenAndServe(addr, nil); err != nil {
        fmt.Printf("Error starting server on port %d: %v\n", s.port, err)
    }
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've reached the server on port %d!", s.port)
}