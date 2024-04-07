package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) CreateLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntrty()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{
			Result: "Failed to write log",
		}
		return res, err
	}

	res := &logs.LogResponse{Result: "Logged successfully"}
	return res, nil
}

func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	logs.RegisterLogServiceServer(server, &LogServer{Models: app.Models})

	log.Printf("gRPC Server started on port %s", gRpcPort)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
