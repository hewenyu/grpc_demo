package greeter_client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	pb "github.com/hewenyu/grpc_demo/helloworld"
	"google.golang.org/grpc"
)

func connect(name string) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}

// 性能测试
func benchmark(b *testing.B, name string) {

	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		_name := fmt.Sprintf("%v_%v", name, i)

		go func() {
			connect(_name)
			wg.Done()
		}()
	}

	wg.Wait()

}

func benchmark_connect(b *testing.B, name string) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	for i := 0; i < b.N; i++ {

		_name := fmt.Sprintf("%v_%v", name, i)

		c := pb.NewGreeterClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: _name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())

	}

}

func BenchmarkMore(b *testing.B) { benchmark(b, defaultName) }

func BenchmarkDeConn(b *testing.B) { benchmark_connect(b, defaultName) }
