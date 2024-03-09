package usecase

import (
	"context"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/database/entity"
	pb "github.com/YukiOnishi1129/auto-stock-trading-system/user-service/grpc"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/infrastructure/mysql/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserUsecaseInterface interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
}

type UserUsecase struct {
	ur *repository.UserRepository
}

func NewUserUsecase(ur *repository.UserRepository) *UserUsecase {
	return &UserUsecase{ur}
}

// GetUsers is a usecase to get all users
func (uu *UserUsecase) GetUsers(ctx context.Context) (*pb.UsersResponse, error) {
	users, rErr := uu.ur.GetUsers(ctx)
	if rErr != nil {
		return nil, rErr
	}

	resUsers := make([]*pb.User, 0)
	for _, user := range users {
		resUser := &pb.User{
			Id:        user.ID,
			Email:     user.Email.String,
			Name:      user.Name.String,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		}
		if !user.DeletedAt.Valid {
			resUser.DeletedAt = timestamppb.New(user.DeletedAt.Time)
		}
		resUsers = append(resUsers, resUser)

	}
	return &pb.UsersResponse{
		Users: resUsers,
	}, nil
}
