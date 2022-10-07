package grpcapi

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/blueai2022/appsubmission/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RenewAccessToken(ctx context.Context, req *pb.RenewAccessTokenRequest) (*pb.RenewAccessTokenResponse, error) {
	refreshPayload, err := server.tokenMaker.VerifyToken(req.GetRefreshToken())
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "session not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get session info: %s", err)
	}

	if session.IsBlocked {
		err = errors.New("blocked session")
		return nil, unauthenticatedError(err)
	}

	if session.Username != refreshPayload.Username {
		err = errors.New("incorrect session username")
		return nil, unauthenticatedError(err)
	}

	if session.RefreshToken != req.RefreshToken {
		err = errors.New("mismatched session refresh token")
		return nil, unauthenticatedError(err)
	}

	if time.Now().After(session.ExpiresAt) {
		err = errors.New("session expired")
		return nil, unauthenticatedError(err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err)
	}

	rsp := &pb.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiredAt: timestamppb.New(accessPayload.ExpiredAt),
	}

	return rsp, nil
}
