package main

import (
	"context"
	"reflect"
)

type GetUserInfoRequest struct {
	Id int
}

type GetUserInfoResponse struct {
	Id int
}

// UserService is example for dynamic agent
// 1. All the field must be func
// 2. All func have two args and two return values
// 3. Args must be context.Context and a request's pointer by user design
// 4. Return values muse be a response's pointer by user design and error
type UserService struct {
	GetUserInfo func(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error)
}

func (u *UserService) Name() string {
	return "UserService"
}

func main() {
	var user UserService
	fns := initStub(&user)

	for _, fn := range fns {
		ctx := reflect.ValueOf(context.Background())
		req := reflect.ValueOf(&GetUserInfoRequest{Id: 1})
		fn.Call([]reflect.Value{ctx, req})
	}
}
