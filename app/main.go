package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/asadzeynal/gymbro-api/gen/pb"
	"github.com/asadzeynal/gymbro-api/internal/controller/gapi"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConn, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	server := gapi.NewServer()
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterGymBroServer(grpcServer, server)

	grpcAddr := viper.GetString("server.address")
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting grpc server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
