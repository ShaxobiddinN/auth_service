package storage

import (
	"blogpost/auth_service/protogen/blogpost"
)

type StorageI interface {
	AddUser(id string, entity *blogpost.CreateUserRequest) error
	GetUserById(id string) (*blogpost.User, error)
	GetUserList(offset, limit int, search string) (resp *blogpost.GetUserListResponse, err error)
	UpdateUser(entity *blogpost.UpdateUserRequest) error
	RemoveUser(id string) error
}
