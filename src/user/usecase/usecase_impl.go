package usecase

import (
	"fmt"

	"github.com/agungdwiprasetyo/agungdpcms/helper"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/token"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/serializer"
)

type userUc struct {
	repo  *repository.Repository
	token token.Token
}

// NewUserUsecase construct user usecase
func NewUserUsecase(repo *repository.Repository, token token.Token) User {
	return &userUc{
		repo:  repo,
		token: token,
	}
}

func (uc *userUc) Login(username, password string) (res shared.Result) {
	res = <-uc.repo.User.FindByUsername(username)
	if res.Error != nil {
		return shared.Result{Error: fmt.Errorf("invalid username/password")}
	}

	userData := res.Data.(*domain.User)
	password = helper.ComputeHmac256(password)
	if userData.Password != password {
		return shared.Result{Error: fmt.Errorf("invalid username/password")}
	}

	res = <-uc.repo.Role.FindByID(userData.RoleID)
	if res.Error != nil {
		return shared.Result{Error: fmt.Errorf("invalid username/password")}
	}

	userData.Role = res.Data.(*domain.Role)

	claims := token.NewClaim(userData)
	t, err := uc.token.Generate(claims)
	if err != nil {
		res.Error = err
		return
	}
	userData.Token = t

	res.Data = &serializer.UserSchema{User: userData}
	return
}
