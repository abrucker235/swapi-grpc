package main

import (
	context "context"
	pb "swapi-grpc/swapi"
)

//Planets is the GRPC implementation
func (s *Server) Planets(ctx context.Context, input *pb.PlanetsRequest) (*pb.PlanetsReply, error) {
	var outputPlanets []*pb.Planet
	if input.Id != 0 {
		for _, planet := range Planets {
			if input.Id == planet.Id {
				outputPlanets = append(outputPlanets, planet)
				return &pb.PlanetsReply{Planets: outputPlanets}, nil
			}
		}
	}

	if input.Name != "" {
		for _, planet := range Planets {
			if input.Name == planet.Name {
				outputPlanets = append(outputPlanets, planet)
				return &pb.PlanetsReply{Planets: outputPlanets}, nil
			}
		}
	}

	return &pb.PlanetsReply{Planets: Planets}, nil
}
