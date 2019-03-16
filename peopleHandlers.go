package main

import (
	context "context"
	pb "swapi-grpc/swapi"
)

//People is the GRPC implementation
func (s *Server) People(ctx context.Context, input *pb.PeopleRequest) (*pb.PeopleReply, error) {
	var outputPeople []*pb.Person
	if input.Id != 0 {
		for _, person := range People {
			if input.Id == person.Id {
				outputPeople = append(outputPeople, person)
				return &pb.PeopleReply{People: outputPeople}, nil
			}
		}
	}

	if input.Name != "" {
		for _, person := range People {
			if input.Name == person.Name {
				outputPeople = append(outputPeople, person)
				return &pb.PeopleReply{People: outputPeople}, nil
			}
		}
	}

	return &pb.PeopleReply{People: People}, nil
}
