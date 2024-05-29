package main

import (
	"flag"
	"log"

	jaegerClient "github.com/ibiscum/Go-for-DevOps/chapter11/ops/internal/jaeger/client"
	"github.com/ibiscum/Go-for-DevOps/chapter11/ops/internal/prom"
	"github.com/ibiscum/Go-for-DevOps/chapter11/ops/internal/server"
	"github.com/ibiscum/Go-for-DevOps/chapter11/petstore/client"
)

var (
	addr         = flag.String("addr", "0.0.0.0:7000", "The address to run the service on.")
	jaegerAddr   = flag.String("jaegerAddr", "127.0.0.1:16685", "The address of the Jaeger query service.")
	promAddr     = flag.String("promAddr", "127.0.0.1:9000", "The address of the Prometheus service.")
	petstoreAddr = flag.String("petstoreAddr", "127.0.0.1:6742", "The address of the Petstore.")
)

func main() {
	flag.Parse()

	j, err := jaegerClient.New(*jaegerAddr)
	if err != nil {
		log.Fatalf("could not connect to Jaeger: %s", err)
	}

	p, err := prom.New("http://" + *promAddr)
	if err != nil {
		log.Fatalf("could not connect to Prometheus: %s", err)
	}

	ps, err := client.New(*petstoreAddr)
	if err != nil {
		log.Fatalf("could not connecto the Petstore: %s", err)
	}
	clients := server.Clients{
		Jaeger:   j,
		Prom:     p,
		Petstore: ps,
	}
	serv, err := server.New(*addr, clients)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("serving starting on: ", *addr)
	err = serv.Start()
	if err != nil {
		log.Fatal(err)
	}

}
