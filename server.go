package main

import (
	"context"
	"fmt"
	grpc_calc "grpc_practice/grpc_calc"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	grpc_calc.CalcPerformerServer
}

func (s *Server) Calculate(ctx context.Context, req *grpc_calc.Operation) (*grpc_calc.OperationResult, error) {
	fmt.Println(req.First, req.Operation, req.Second)
	answer := req.First + req.Second
	return &grpc_calc.OperationResult{Result: answer, Error: "0"}, nil
}

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := grpc_calc.RegisterCalcPerformerHandlerServer(ctx, mux, &Server{})
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Mount("/", mux)

	if err := http.ListenAndServe(":8081", r); err != nil {
		panic(err)
	}
}

func runGrpc() {
	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpc_calc.RegisterCalcPerformerServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	go runRest()
	runGrpc()
}
