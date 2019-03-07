package main

import (
	"log"
	"time"
	pb "user-service/proto/user"

	"github.com/dgrijalva/jwt-go"
)

var (
	// this key is used in hashing our token
	key = []byte("mySecretKey")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

//Decode is
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	//Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	//Validate the toketn
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//Encode is
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "shipping.user",
		},
	}

	//Create a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Println("Token Created")
	//Sign the token with key and return
	return token.SignedString(key)
}
