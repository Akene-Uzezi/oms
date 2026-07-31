package main

import pb "oms-common/api"

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}
