package serializer

import "github.com/agungdwiprasetyo/agungdpcms/src/user/domain"

type UserSchema struct {
	User *domain.User
}

func (r *UserSchema) ID() int32 {
	return int32(r.User.ID)
}

func (r *UserSchema) Username() string {
	return r.User.Username
}

func (r *UserSchema) Password() string {
	return r.User.Password
}

func (r *UserSchema) Name() string {
	return r.User.Name
}

func (r *UserSchema) Token() string {
	return r.User.Token
}
