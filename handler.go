package main

import (
	"log"

	micro "github.com/micro/go-micro"
	// _ "github.com/micro/go-plugins/broker/nats"
	auth "github.com/rickynyairo/plaeve-auth/proto/auth"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

const (
	topic = "user.created"
)

// Handler implements auth.AuthHandler to allow signup, login
type Handler struct {
	repo         Repository
	tokenService Authable
	Publisher    micro.Publisher
}

// Create allows creation of a new user account
func (srv *Handler) Create(ctx context.Context, req *auth.User, res *auth.Response) error {
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	if err := srv.Publisher.Publish(ctx, req); err != nil {
		return err
	}
	return nil
}

// Login authenticates a user using email and password
func (srv *Handler) Login(ctx context.Context, req *auth.User, res *auth.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := srv.repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

// Get returns a user
func (srv *Handler) Get(ctx context.Context, req *auth.User, res *auth.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

// GetAll returns all users
func (srv *Handler) GetAll(ctx context.Context, req *auth.Request, res *auth.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

// ValidateToken validates atoken
func (srv *Handler) ValidateToken(ctx context.Context, req *auth.Token, res *auth.Token) error {
	return nil
}
