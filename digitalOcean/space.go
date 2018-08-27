package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

func main() {
	accessKey := "BUIX2OH6JJK7DIR4UOBM"
	secKey := "BRV7l4Vz/+VAgwuZwKYn9kgb581xaChuQRrXkzYGtv0"
	endpoint := "0xdev.ams3.digitaloceanspaces.com"
	spaceName := "0xdev" // Space names must be globally unique
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Space.
	err = client.MakeBucket(spaceName, "us-east-1")
	if err != nil {
		log.Fatal(err)
	}

	// List all Spaces.
	spaces, err := client.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}
	for _, space := range spaces {
		fmt.Println(space.Name)
	}
}
