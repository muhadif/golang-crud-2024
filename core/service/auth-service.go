package module

import (
	"context"
	"golang-crud-2024/config"
	"golang-crud-2024/core/entity"
	coreErr "golang-crud-2024/core/error"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/auth"
	"golang-crud-2024/pkg/fault"
	stringLib "golang-crud-2024/pkg/string"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	RefreshToken(ctx context.Context, req *entity.RefreshTokenRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterRequest) error
	Logout(ctx context.Context, req *entity.LogoutRequest) error
}

type authModule struct {
	userRepository repository.UserRepository
	cfg            config.Config
}

func NewAuthService(userRepository repository.UserRepository, cfg config.Config) AuthService {
	return &authModule{
		userRepository: userRepository,
		cfg:            cfg,
	}
}

func (a *authModule) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {
	checkUser, err := a.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if checkUser == nil {
		return nil, fault.ErrorDictionary(fault.HTTPPreconditionFailedError, coreErr.ErrLoginIncorrect)
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateJWTToken(checkUser, a.cfg)
	if err != nil {
		return nil, err
	}

	err = a.userRepository.UpsertRefreshToken(ctx, checkUser.Serial, token.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{AccessToken: token}, nil
}

func (a *authModule) Register(ctx context.Context, req *entity.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	checkUser, err := a.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if checkUser != nil {
		return fault.ErrorDictionary(fault.HTTPBadRequestError, coreErr.ErrEmailTaken)
	}

	user := &entity.User{
		Serial:       stringLib.GenerateSerial(entity.UserSerialPrefix, entity.UserSerialLength),
		Email:        req.Email,
		Role:         entity.UserRoleUser,
		Username:     req.Username,
		Password:     string(hashedPassword),
		AccessStatus: entity.UserStatusDisabled,
		Status:       entity.UserStatusDisabled,
	}
	err = a.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *authModule) RefreshToken(ctx context.Context, req *entity.RefreshTokenRequest) (*entity.LoginResponse, error) {
	if isValid := auth.ValidateRefreshToken(req.RefreshToken); !isValid {
		return nil, fault.ErrorDictionary(fault.HTTPUnauthorizedError, coreErr.ErrTokenNotValid)
	}

	userSerial, err := a.userRepository.GetUserSerialByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := a.userRepository.GetUserByUserSerial(ctx, userSerial)
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateJWTToken(user, a.cfg)
	if err != nil {
		return nil, err
	}

	err = a.userRepository.UpsertRefreshToken(ctx, user.Serial, token.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{AccessToken: token}, nil
}

func (a *authModule) Logout(ctx context.Context, req *entity.LogoutRequest) error {
	if err := a.userRepository.RemoveUserToken(ctx, req.UserSerial); err != nil {
		return err
	}
	return nil
}
