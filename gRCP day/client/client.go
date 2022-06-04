package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/montanaflynn/stats"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	tr "src/api/proto/transmitter"
	"sync"
)

var pool = sync.Pool{
	// New creates an object when the pool has nothing available to return.
	// New must return an interface{} to make it flexible. You have to cast
	// your type after getting it.
	New: func() any {
		// Pools often contain things like *bytes.Buffer, which are
		// temporary and re-usable.
		return []float64{}
	},
}

func receiveMessages(c tr.TransmitterClient, k float64) {
	var samples []float64
	resStream, err := c.GetSampleStream(context.Background(), &tr.Request{})
	if err != nil {
		log.Fatalf("Error while calling GetSampleStream RPC: %v", err)
	}

	flag := 1
	for i := 0; i < 50; i++ {
		samples = pool.Get().([]float64)
		msg, err := resStream.Recv()
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		if flag == 1 {
			fmt.Printf("session: %s\n", msg.SessionId)
			flag = 0
		}
		samples = append(samples, msg.Frequency)
		pool.Put(samples)
	}
	samples = pool.Get().([]float64)
	mean, err := stats.Mean(samples)
	if err != nil {
		log.Fatalf("while calculating mean: %v", err)
	}
	std, err := stats.StandardDeviation(samples)
	pool.Put(samples)
	if err != nil {
		log.Fatalf("while calculating std: %v", err)
	}
	fmt.Printf("mean: %f\nstd: %f\n", mean, std)
	n := 0
	an := 0
	notAn := 0
	for {
		msg, err := resStream.Recv()
		n++
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		if msg.Frequency < mean-k*std || msg.Frequency > mean+k*std {
			an++
			log.Printf("an anomaly: %f\t%f\n", msg.Frequency, float64(an)/float64(n))
		} else {
			notAn++
			log.Printf("not an anomaly: %f\t%f\n", msg.Frequency, float64(notAn)/float64(n))
		}
	}
}

func main() {
	kPtr := flag.Float64("k", 0.1, "number of samples")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	c := tr.NewTransmitterClient(conn)

	receiveMessages(c, *kPtr)
}
