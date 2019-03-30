package main

import (
  "log"
   "net"
   "fmt"
   "google.golang.org/grpc"
   "github.com/beelarr/grpc-go-course/greet/greetpb"
   )

type server struct{}

func main() {
  fmt.Println("Hello World")
  lis, err := net.Listen("tcp", " 0.0.0.0:50051"
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  s:=grpc.NewServer()
  greetpb.RegistewrGreetSErviceServer(s, &server{})

  if err := s.Server(lis); err != nil {
    log.Fatalf("failed to server: %v", err)
  }
}