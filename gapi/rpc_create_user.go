package gapi

import (
	"context"
	"time"

	db "github.com/OmSingh2003/vaultguard-api/db/sqlc"
	"github.com/OmSingh2003/vaultguard-api/pb"
	"github.com/OmSingh2003/vaultguard-api/util"
	"github.com/OmSingh2003/vaultguard-api/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	// NOTE: TODO - use DB transaction
	taskPayload := &worker.PayloadSendVerifyEmail{
		Username: user.Username,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue("critical"),
	}
	err = server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to distribute task: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

