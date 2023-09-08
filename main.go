package main

import (
	"context"
	"log"

	"github.com/khanakia/entautoslow/ent"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=entgodb3 password=root sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resource dds: %v", err)
	}

	// client.Item.Create().Mutation().Set
}
