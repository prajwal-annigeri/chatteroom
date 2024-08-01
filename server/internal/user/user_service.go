package user

import (
	"context"
	"fmt"
	"os"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	c, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(c, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil
}

type JWTClaims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	c, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.GetUserByEmail(c, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}
	fmt.Printf("Username %s, Password %s", u.Username, u.Password)
	if err := util.CheckPassword(req.Password, u.Password); err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	signedString, err := token.SignedString([]byte(os.Getenv("secretkey")))
	if err != nil {
		return &LoginUserRes{}, nil
	}

	return &LoginUserRes{
		accessToken: signedString,
		Username: u.Username,
		ID: strconv.Itoa(int(u.ID)),
	}, nil
}
