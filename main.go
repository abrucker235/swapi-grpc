package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	pb "swapi-grpc/swapi"

	"google.golang.org/grpc"
)

//Server is a struct to implement SwapiServer interface
type Server struct{}

var Films []*pb.Film

var People []*pb.Person

var Planets []*pb.Planet

var Species []*pb.Species

func init() {
	readJSONFile("data/films.json", &Films)
	readJSONFile("data/people.json", &People)
	readJSONFile("data/planets.json", &Planets)
	readJSONFile("data/species.json", &Species)
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSwapiServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func readJSONFile(path string, v interface{}) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(bytes, v); err != nil {
		fmt.Println(err)
	}
}
