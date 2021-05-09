package main

import(
	"fmt"
	"context"
	"io/ioutil"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

const(
	projectID = "cards-313120"
	adminServiceAccountKeyPath = "/Users/shoutinglyu/go/src/github.com/valentinescards/serviceaccountkeys/firestore-admin.json"
)

func main() {
	firestoreSAKey, err := ioutil.ReadFile(adminServiceAccountKeyPath)
	if err != nil {
		log.Fatalf("Failed to read firestore admin key file: %v", err)
	}
	o := option.WithCredentialsJSON(firestoreSAKey)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, o)
	if err != nil {
		log.Fatalf("Failed to initiate a new firestore client: %v", err)
	}
	defer client.Close()
	// users := client.Collection("users")
	fmt.Println("Successfully initiate firestore client!!!")
}