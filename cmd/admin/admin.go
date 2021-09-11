package main

import (
	"flag"
	"log"

	"github.com/google/uuid"
	"github.com/mes1234/golock/adapters"
	"github.com/mes1234/golock/auth"
	"github.com/mes1234/golock/persistance"
)

func main() {

	username := flag.String("username", "", "name of user")
	password := flag.String("password", "", "password for user")

	flag.Parse()

	collection := persistance.NewClientRepository()

	client := adapters.Client{
		ClientName: *username,
	}
	collection.Retrieve(&client)

	if client.ClientId != uuid.Nil {
		log.Printf("Username : %v already defined", *username)
	} else {
		client.ClientId = uuid.New()
		salted := *password + client.ClientId.String()
		hash, _ := auth.HashPassword(salted)
		client.Password = hash
		log.Printf("attempt to insert\nUsername : %v\nPassword : %v \nId : %v", *username, *password, client.ClientId)
		collection.Insert(&client)
	}

}
