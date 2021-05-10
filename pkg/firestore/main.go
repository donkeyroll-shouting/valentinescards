package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/donkeyroll-shouting/valentinescards/pkg/constants"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func NewFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	firestoreKeyPath := os.Getenv(constants.FIRESTORE_KEY)
	firestoreSAKey, err := ioutil.ReadFile(firestoreKeyPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read firestore admin key file: %v", err)
	}
	o := option.WithCredentialsJSON(firestoreSAKey)
	client, err := firestore.NewClient(ctx, constants.ProjectID, o)
	if err != nil {
		return nil, fmt.Errorf("Failed to initiate a new firestore client: %v", err)
	}
	return client, nil
}

func main() {
	ctx := context.TODO()
	c, err := NewFirestoreClient(ctx)
	if err != nil {
		log.Fatalf("failed to create a firestore client: %v", err)
	}
	newCard := &docs.Card{
		Name: "testing card",
	}
	m := &CardManager{
		client: c,
	}
	m.CreateCard(ctx, newCard)
}