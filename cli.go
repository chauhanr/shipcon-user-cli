package main

import (
	"github.com/micro/go-micro"
	pb "github.com/chauhanr/shipcon-user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"context"
	"log"
	"os"
)

func main(){
	srv := micro.NewService(
		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)

	srv.Init()
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	name := "Ritesh Chauhan"
	email := "ritesh.chauhan@hcl.com"
	password := "test123"
	company := "hcl"

	r, err := client.Create(context.TODO(), &pb.User{
		Name: name,
		Email: email,
		Password: password,
		Company: company,
	})

	if err != nil{
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)
	getAll, err := client.GetAll(context.Background(), &pb.Request{})
    if err != nil{
		log.Fatalf("Could not list users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email: email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Printf("Your access token is : %s\n", authResponse.Token)
	os.Exit(0)
}