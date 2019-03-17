package main

import (
	context "context"
	pb "swapi-grpc/swapi"
)

//Starships is the GRPC implementation
func (s *Server) Starships(ctx context.Context, input *pb.StarshipsRequest) (*pb.StarshipsReply, error) {
	var outputStarships []*pb.Starship
	if input.Id > 0 {
		for _, starship := range Starships {
			if input.Id == starship.Id {
				outputStarships = append(outputStarships, starship)
				return &pb.StarshipsReply{Starships: outputStarships}, nil
			}
		}
	}

	if input.Name != "" {
		for _, starship := range Starships {
			if input.Name == starship.Name {
				outputStarships = append(outputStarships, starship)
				return &pb.StarshipsReply{Starships: outputStarships}, nil
			}
		}
	}

	if input.Model != "" {
		for _, starship := range Starships {
			if input.Model == starship.Model {
				outputStarships = append(outputStarships, starship)
				return &pb.StarshipsReply{Starships: outputStarships}, nil
			}
		}
	}
	return &pb.StarshipsReply{Starships: Starships}, nil
}
