// ctrl+fn+f2 belgilangan 1ta sozni hammasini belgilab ozgartirish
package auth

import (
	blogpost "blogpost/auth_service/protogen/blogpost"
	"blogpost/auth_service/storage"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthService...
type authService struct {
	stg storage.StorageI
	blogpost.UnimplementedAuthServiceServer
}

func NewAuthService(stg storage.StorageI) *authService {
	return &authService{
		stg: stg,
	}
}

// Ping ...
func (s *authService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Pingg")
	return &blogpost.Pong{
		Message: "Ok",
	}, nil
}

// CreateAuth...
func (s *authService) CreateUser(ctx context.Context, req *blogpost.CreateUserRequest) (*blogpost.User, error) {
	id := uuid.New()
	//Todo - hash password

	err := s.stg.AddUser(id.String(), req)
	if err != nil {

		return nil, status.Errorf(codes.Internal, "s.stg.AddUser: %s", err.Error())
	}

	user, err := s.stg.GetUserById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserById: %s", err.Error())

	}

	return user, nil
}

// UpdateUser...
func (s *authService) UpdateUser(ctx context.Context, req *blogpost.UpdateUserRequest) (*blogpost.User, error) {
	//Todo - hash password

	err := s.stg.UpdateUser(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateUser: %s", err.Error())

	}

	user, err := s.stg.GetUserById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserById: %s", err.Error())
	}

	return user, nil
}

// DeleteUser...
func (s *authService) DeleteUser(ctx context.Context, req *blogpost.DeleteUserRequest) (*blogpost.User, error) {
	user, err := s.stg.GetUserById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserById: %s", err.Error())
	}

	err = s.stg.RemoveUser(user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteUser: %s", err.Error())
	}
	return user, nil

}

// GetUserList...
func (s *authService) GetUserList(ctx context.Context, req *blogpost.GetUserListRequest) (*blogpost.GetUserListResponse, error) {
	fmt.Println("----------GetUserList----------->")

	res, err := s.stg.GetUserList(int(req.Offset), int(req.Limit), string(req.Search))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserList: %s", err.Error())
	}

	return res, nil
}

// GetUserById...
func (s *authService) GetUserById(ctx context.Context, req *blogpost.GetUserByIdRequest) (*blogpost.User, error) {
	user, err := s.stg.GetUserById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserById: %s", err.Error())
	}

	return user, nil
}
