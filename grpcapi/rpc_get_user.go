package grpcapi

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/blueai2022/appsubmission/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			notFoundErr := fmt.Errorf("username %s not found", req.Username)
			return nil, status.Errorf(codes.NotFound, "user not found: %s", notFoundErr)
		}
		return nil, status.Errorf(codes.Internal, "failed to find user: %s", err)
	}

	rsp := &pb.GetUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}
