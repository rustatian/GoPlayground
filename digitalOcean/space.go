package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

func main() {
	accessKey := ""
	secKey := ""
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
