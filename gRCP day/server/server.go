package main

import (
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
	"os"
	"src/api/proto/transmitter"
	"time"
)

type server struct {
	transmitter.UnimplementedTransmitterServer
}

func (s *server) GetSampleStream(req *transmitter.Request, stream transmitter.Transmitter_GetSampleStreamServer) error {
	mean := -10 + rand.Float64()*(10 - -10)
	std := 0.3 + rand.Float64()*(1.5-0.3)
	log.Printf("mean: %v, std: %v\n", mean, std)
	session := uuid.New().String()

	for {
		normSample := rand.NormFloat64()*std + mean
		t := time.Now().UTC()
		timeStamp := timestamppb.New(t)
		err := stream.Send(&transmitter.Response{
			SessionId: session,
			Frequency: normSample,
			TimeStamp: timeStamp,
		})
		time.Sleep(1000 * time.Millisecond)
		if err != nil {
			return err
		}
	}
}

func setLogs() *os.File {
	logFile, err := os.OpenFile("values.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("error opening file %v", err)
	}
	log.SetOutput(logFile)
	return logFile
}

func main() {
	//defer setLogs().Close()

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	transmitter.RegisterTransmitterServer(s, &server{})
	log.Printf("server listening aat %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
