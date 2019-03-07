package main

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"

	pb "user-service/proto/user"
)

const topic = "user.created"

type service struct {
	repo         Repository
	tokenService Authable
}

func (s *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	log.Println("Creating User: ", req)
	//Generate hashed version of the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error on Hashing the password: ", err)
		return err
	}
	req.Password = string(hashedPass)
	log.Println("Executing the Create method")
	if err := s.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}
func (s *service) Get(ctx context.Context, user *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(user.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}
func (s *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
func (s *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with: ", req.Email, req.Password)
	user, err := s.repo.GetByEmail(req.Email)
	log.Println("User after calling GetByEmail(): ", user)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Println("Error on Bcrypt Compare..", err)
		return err
	}

	token, err := s.tokenService.Encode(user)
	if err != nil {
		log.Println("Error Encoding user ", err)
		return err
	}
	res.Token = token
	return nil
}
func (s *service) ValidateToken(ctx context.Context, t *pb.Token, tkn *pb.Token) error {
	return nil
}
