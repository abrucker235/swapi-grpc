package main

import (
	context "context"
	pb "swapi-grpc/swapi"
)

//Species is the GRPC implementation
func (s *Server) Species(ctx context.Context, input *pb.SpeciesRequest) (*pb.SpeciesReply, error) {
	var outputSpecies []*pb.Species
	if input.Id != 0 {
		for _, specie := range Species {
			if input.Id == specie.Id {
				outputSpecies = append(outputSpecies, specie)
				return &pb.SpeciesReply{Species: outputSpecies}, nil
			}
		}
	}

	if input.Name != "" {
		for _, specie := range Species {
			if input.Name == specie.Name {
				outputSpecies = append(outputSpecies, specie)
				return &pb.SpeciesReply{Species: outputSpecies}, nil
			}
		}
	}

	return &pb.SpeciesReply{Species: Species}, nil
}
