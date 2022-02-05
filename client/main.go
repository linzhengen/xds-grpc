package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/linzhengen/xds-grpc/proto/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
)

var (
	target = flag.String("target", "xds:///localhost:50051", "uri of the Greeter Server, e.g. 'xds:///helloworld-service:8080'")
	name   = flag.String("name", "world", "name you wished to be greeted by the server")
)

func main() {
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial(*target, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("grpc.Dial(%s) failed: %v", *target, err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c := helloworld.NewGreeterClient(conn)
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
