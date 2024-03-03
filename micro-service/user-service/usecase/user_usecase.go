package usecase

import (
	"context"
	"database/sql"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/database/entity"
	pb "github.com/YukiOnishi1129/auto-stock-trading-system/user-service/grpc"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/infrastructure/mysql/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type UserUsecaseInterface interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
}

type UserUsecase struct {
	db *sql.DB
	ur *repository.UserRepository
}

func NewUserUsecase(db *sql.DB, ur *repository.UserRepository) *UserUsecase {
	return &UserUsecase{db, ur}
}

// GetUsers is a usecase to get all users
func (uu *UserUsecase) GetUsers(ctx context.Context) (*pb.UsersResponse, error) {
	users, rErr := uu.ur.GetUsers(ctx)
	if rErr != nil {
		return nil, rErr
	}

	resUsers := make([]*pb.User, 0)
	for _, user := range users {
		resUsers = append(resUsers, &pb.User{
			Id:        strconv.Itoa(user.ID),
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}
	return &pb.UsersResponse{
		Users: resUsers,
	}, nil
}
