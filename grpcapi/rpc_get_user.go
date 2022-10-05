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
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.Username != req.Username {
		return nil, status.Errorf(codes.PermissionDenied, "cannot access other user's info")
	}

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			verr := fmt.Errorf("username %s not found", req.Username)
			return nil, status.Errorf(codes.NotFound, "user not found: %s", verr)
		}
		return nil, status.Errorf(codes.Internal, "failed to find user: %s", err)
	}

	rsp := &pb.GetUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}
