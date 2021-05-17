package service

import (
	"context"
	pb "go-micro-demo-project/api/user/service/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() pb.UserServer {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) RegisterOrLogin(ctx context.Context, req *pb.RegisterOrLoginRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateReply, error) {
	return &pb.AuthenticateReply{}, nil
}
