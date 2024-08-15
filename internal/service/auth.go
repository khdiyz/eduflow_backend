package service

import (
	"database/sql"
	"eduflow/config"
	"eduflow/internal/constants"
	"eduflow/internal/model"
	"eduflow/internal/repository"
	"eduflow/pkg/helper"
	"eduflow/pkg/logger"
	"eduflow/pkg/response"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewAuthService(repo repository.Repository, log logger.Logger) *AuthService {
	return &AuthService{
		repo: repo,
		log:  log,
	}
}

type jwtCustomClaim struct {
	jwt.StandardClaims
	UserId   int64  `json:"user_id"`
	RoleName string `json:"role"`
	Type     string `json:"type"`
}

func (s *AuthService) CreateToken(user model.User, tokenType string, expiresAt time.Time) (*model.Token, error) {
	claims := &jwtCustomClaim{
		UserId:   user.Id,
		RoleName: user.RoleName,
		Type:     tokenType,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	return &model.Token{
		User:      user,
		Token:     token,
		Type:      tokenType,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *AuthService) GenerateTokens(user model.User) (*model.Token, *model.Token, error) {
	accessExpiresAt := time.Now().Add(time.Duration(config.GetConfig().JWTAccessExpirationHours) * time.Hour)
	refreshExpiresAt := time.Now().Add(time.Duration(config.GetConfig().JWTRefreshExpirationDays) * time.Hour * 24)

	accessToken, err := s.CreateToken(user, constants.TokenTypeAccess, accessExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := s.CreateToken(user, constants.TokenTypeRefresh, refreshExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ParseToken(token string) (*jwtCustomClaim, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*jwtCustomClaim)
	if !ok {
		return nil, errors.New("token claims are not of type *jwtCustomClaim")
	}

	return claims, nil
}

func (s *AuthService) Login(input model.LoginRequest) (*model.Token, *model.Token, error) {
	user, err := s.repo.User.GetByUsername(input.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, response.ServiceError(errors.New("wrong username or password"), codes.Unauthenticated)
		}
		return nil, nil, response.ServiceError(err, codes.Internal)
	}

	hashPassword, err := helper.GenerateHash(input.Password)
	if err != nil {
		return nil, nil, response.ServiceError(err, codes.Internal)
	}

	if user.Password != hashPassword {
		return nil, nil, response.ServiceError(errors.New("wrong username or password"), codes.Unauthenticated)
	}

	return s.GenerateTokens(user)
}
