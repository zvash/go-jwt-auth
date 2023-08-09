package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zvash/go-jwt-auth/internal/config"
	"github.com/zvash/go-jwt-auth/internal/grpc/auth"
	"github.com/zvash/go-jwt-auth/internal/grpc/handlers"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	err := setupAndListen()
	if err != nil {
		log.Fatalf("grpc server failed with err: %v", err)
		return
	}
}

func setupAndListen() error {
	grpcConfig := config.GRPCConfig{}
	err := godotenv.Load()
	if err != nil {
		return err
	}
	portString := os.Getenv("GRPC_PORT")
	if portString == "" {
		portString = "9000"
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return err
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	grpcConfig.SetDB(conn)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", portString))
	if err != nil {
		log.Fatalf("failed to listen on port %s", portString)
		return err
	}
	grpcServer := grpc.NewServer()
	grpcConfig.SetServer(grpcServer)

	handlers.SetGRPCConfig(&grpcConfig)
	s := handlers.Server{}
	auth.RegisterAuthServiceServer(grpcServer, &s)

	log.Printf("Starting grpc server on port: %v and address: %v", portString, listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server on port %s with error: %v", portString, err)
	}
	return nil
}
