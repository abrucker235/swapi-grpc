package main

import (
	context "context"
	pb "swapi-grpc/swapi"
)

//Films is the GRPC implementation
func (s *Server) Films(ctx context.Context, input *pb.FilmsRequest) (*pb.FilmsReply, error) {
	var outputFilms []*pb.Film
	if input.Id != 0 {
		for _, film := range Films {
			if film.Id == input.Id {
				outputFilms = append(outputFilms, film)
				return &pb.FilmsReply{Films: outputFilms}, nil
			}
		}
	}

	if input.Title != "" {
		for _, film := range Films {
			if film.Title == input.Title {
				outputFilms = append(outputFilms, film)
				return &pb.FilmsReply{Films: outputFilms}, nil
			}
		}
	}

	return &pb.FilmsReply{Films: Films}, nil
}
