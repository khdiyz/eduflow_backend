package service

import (
	"database/sql"
	"eduflow/config"
	"eduflow/internal/models"
	"eduflow/internal/repository"
	"eduflow/pkg/helper"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	repo *repository.Repository
	cfg  *config.Config
}

func NewAuthService(repo *repository.Repository, cfg *config.Config) *AuthService {
	return &AuthService{
		repo: repo,
		cfg:  cfg,
	}
}

type jwtCustomClaim struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id"`
	RoleId uuid.UUID `json:"role_id"`
	Type   string    `json:"type"`
}

func (s *AuthService) CreateToken(user models.User, tokenType string, expiresAt time.Time) (*models.Token, error) {
	claims := &jwtCustomClaim{
		UserId: user.Id,
		RoleId: user.RoleId,
		Type:   tokenType,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return &models.Token{
		Token:     token,
		Type:      tokenType,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *AuthService) GenerateTokens(user models.User) (*models.Token, *models.Token, error) {
	accessExpiresAt := time.Now().Add(time.Duration(s.cfg.JWTAccessExpirationHours) * time.Hour)
	refreshExpiresAt := time.Now().Add(time.Duration(s.cfg.JWTRefreshExpirationDays) * time.Hour * 24)

	accessToken, err := s.CreateToken(user, config.TokenTypeAccess, accessExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := s.CreateToken(user, config.TokenTypeRefresh, refreshExpiresAt)
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

func (s *AuthService) Login(request models.LoginRequest) (*models.Token, *models.Token, error) {
	user, err := s.repo.User.GetByUsername(request.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, serviceError(errors.New("user not found"), codes.Unauthenticated)
		}
		return nil, nil, serviceError(err, codes.Internal)
	}

	hashPassword, err := helper.GenerateHash(request.Password)
	if err != nil {
		return nil, nil, serviceError(err, codes.Internal)
	}

	if user.Password != hashPassword {
		return nil, nil, serviceError(errors.New("incorrect username or password"), codes.Unauthenticated)
	}

	return s.GenerateTokens(user)
}
