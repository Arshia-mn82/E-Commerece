package service

import (
	"E-Commerce/internal/dto"
	"E-Commerce/internal/model"
	"E-Commerce/internal/repository"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	Users     *repository.UserRepository
	jwtSecret []byte
}

func NewAuthService(users *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{Users: users, jwtSecret: []byte(jwtSecret)}
}

func (as *AuthService) Signup(req dto.SignUpRequest) (string, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return "", errors.New("name, email, password are required")
	}

	_, err := as.Users.FindByEmail(req.Email)
	if err == nil {
		return "", errors.New("email already in use")
	}

	if err != nil && !repository.IsNotFound(err) {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	u := &model.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         "user",
	}
	if err := as.Users.Create(u); err != nil {
		return "", err
	}

	return as.issueToken(u.ID, u.Role)

}

func (as *AuthService) Login(req dto.LoginRequest) (string, error) {
	if req.Email == "" || req.Password == "" {
		return "", errors.New("email and password are required")
	}

	u, err := as.Users.FindByEmail(req.Email)
	if err != nil {
		if repository.IsNotFound(err) {
			return "", errors.New("invalud credentials")
		}
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return as.issueToken(u.ID, u.Role)
}

func (as *AuthService) issueToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(as.jwtSecret)
}
