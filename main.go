package main

import (
	"log"
	pb "user-service/proto/user"

	"github.com/micro/go-micro"
)

func main() {
	log.Println("User Service start..")

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.

	srv.Init()

	// Will comment this out now to save having to run this locally
	// publisher := micro.NewPublisher("user.created", srv.Client())

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})
	//pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println("The server failed to run....", err)
	}
	log.Println("User Service going down..")
}
