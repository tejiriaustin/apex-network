package service

import "github.com/tejiriaustin/apex-network/env"

type (
	Service struct {
		UserService UserServiceInterface
		config      env.Env
	}
)

type IService interface {
}

func NewService(env env.Env) IService {
	return &Service{
		UserService: NewUserService(env),
		config:      env,
	}
}
