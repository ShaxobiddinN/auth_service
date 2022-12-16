package auth

import (
	"blogpost/auth_service/protogen/blogpost"
	"blogpost/auth_service/util"
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login ...
func (s *authService) Login(ctx context.Context, req *blogpost.LoginRequest) (*blogpost.TokenResponse, error) {
	log.Println("Login...")

	errAuth := errors.New("Username or password wrong")
	//step 1:find user by username
	user, err := s.stg.GetUserByUsername(req.Username)
	if err != nil {
		log.Println(err.Error())
		return nil, status.Errorf(codes.Unauthenticated, errAuth.Error())
	}

	//step 2:find user by username
	match, err := util.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "util.ComparePassword: %s", err.Error())
	}
	if !match {
		return nil, status.Errorf(codes.Unauthenticated, errAuth.Error())
	}

	m := map[string]interface{}{
		"user_id":  user.Id,
		"username": user.Username,
	}
	tokenStr, err := util.GenerateJWT(m, time.Minute*10, s.cfg.SecretKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "util.GenerateJWT: %s", err.Error())
	}

	return &blogpost.TokenResponse{
		Token: tokenStr,
	}, nil
}

// HasAccess ...
func (s *authService) HasAccess(ctx context.Context, req *blogpost.TokenRequest) (*blogpost.HasAccessResponse, error) {
	log.Println("HasAccess...")

	//step 1: parse token
	result, err := util.ParseClaims(req.Token, s.cfg.SecretKey)
	if err != nil {
		log.Println(status.Errorf(codes.Unauthenticated, "util.ParseClaims: %s", err.Error()))
		return &blogpost.HasAccessResponse{
			User:      nil,
			HasAccess: false,
		}, nil
	}

	log.Println(result.Username)

	//step 2: get user by id
	user, err := s.stg.GetUserById(result.UserID)
	if err != nil {
		log.Println(status.Errorf(codes.Unauthenticated, "s.stg.GetUserById: %s", err.Error()))
		return &blogpost.HasAccessResponse{
			User:      nil,
			HasAccess: false,
		}, nil
	}

	return &blogpost.HasAccessResponse{
		User:      user,
		HasAccess: true,
	}, nil

}
