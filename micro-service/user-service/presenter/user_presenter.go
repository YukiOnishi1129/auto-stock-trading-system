package presenter

import (
	"context"
	"database/sql"
	pb "github.com/YukiOnishi1129/auto-stock-trading-system/user-service/grpc"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/usecase"
)

type userPresenter struct {
	db *sql.DB
	uu *usecase.UserUsecase
}

func NewUserPresenter(db *sql.DB, uu *usecase.UserUsecase) *userPresenter {
	return &userPresenter{db, uu}
}

func (p *userPresenter) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (p *userPresenter) GetUsers(ctx context.Context, req *pb.Empty) (*pb.UsersResponse, error) {
	users, rErr := p.uu.GetUsers(ctx)
	if rErr != nil {
		return nil, rErr
	}
	return users, nil
}

func (p *userPresenter) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (p *userPresenter) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (p *userPresenter) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}
