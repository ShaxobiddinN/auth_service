package main

import (
	"fmt"
	"log"
	"net"

	"blogpost/auth_service/config"
	"blogpost/auth_service/protogen/blogpost"
	"blogpost/auth_service/services/auth"

	"blogpost/auth_service/storage"
	"blogpost/auth_service/storage/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	cfg := config.Load()

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	var err error
	var stg storage.StorageI

	stg, err = postgres.InitDb(psqlConnString)
	if err != nil {
		panic(err)
	}

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	authService := auth.NewAuthService(stg)
	blogpost.RegisterAuthServiceServer(srv, authService)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
