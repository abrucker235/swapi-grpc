package main

import (
	context "context"
	pb "swapi-grpc/swapi"
)

//Vehicles is the GRPC implementation
func (s *Server) Vehicles(cxt context.Context, input *pb.VehiclesRequest) (*pb.VehiclesReply, error) {
	var outputVehicles []*pb.Vehicle
	if input.Id > 0 {
		for _, vehilce := range Vehicles {
			if input.Id == vehilce.Id {
				outputVehicles = append(outputVehicles, vehilce)
				return &pb.VehiclesReply{Vehicles: outputVehicles}, nil
			}
		}
	}

	if input.Name != "" {
		for _, vehilce := range Vehicles {
			if input.Name == vehilce.Name {
				outputVehicles = append(outputVehicles, vehilce)
				return &pb.VehiclesReply{Vehicles: outputVehicles}, nil
			}
		}
	}

	if input.Model != "" {
		for _, vehicle := range Vehicles {
			if input.Model == vehicle.Model {
				outputVehicles = append(outputVehicles, vehicle)
				return &pb.VehiclesReply{Vehicles: outputVehicles}, nil
			}
		}
	}

	return &pb.VehiclesReply{Vehicles: Vehicles}, nil
}
