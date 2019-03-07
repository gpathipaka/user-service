package main

import (
	"context"
	"log"

	pb "user-service/proto/user"
)

const topic = "user.created"

type service struct {
	repo         Repository
	tokenService Authable
}

func (s *service) Create(ctx context.Context, user *pb.User, res *pb.Response) error {
	log.Println("Creating User: ", user)
	//Generate hashed version of the password
	log.Println("Executing the Create method")
	if err := s.repo.Create(user); err != nil {
		return err
	}
	res.User = user
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
func (s *service) Auth(ctx context.Context, user *pb.User, t *pb.Token) error {
	user, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		return err
	}
	t.Token = "testing...."
	return nil
}
func (s *service) ValidateToken(ctx context.Context, t *pb.Token, tkn *pb.Token) error {
	return nil
}
