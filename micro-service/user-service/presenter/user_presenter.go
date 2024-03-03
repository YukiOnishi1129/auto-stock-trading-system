package presenter

import (
	"context"
	pb "github.com/YukiOnishi1129/auto-stock-trading-system/user-service/grpc"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/usecase"
)

type UserPresenter struct {
	uu *usecase.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewUserPresenter(uu *usecase.UserUsecase) *UserPresenter {
	return &UserPresenter{uu: uu}
}

func (p *UserPresenter) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (p *UserPresenter) GetUsers(ctx context.Context, req *pb.Empty) (*pb.UsersResponse, error) {
	users, rErr := p.uu.GetUsers(ctx)
	if rErr != nil {
		return nil, rErr
	}
	return users, nil
}

func (p *UserPresenter) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (p *UserPresenter) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (p *UserPresenter) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}
