package firestore

import(
	"context"
	"fmt"

	"github.com/donkeyroll-shouting/valentinescards/pkg/firestore/docs"
	"cloud.google.com/go/firestore"
)

const (
	userCollection = "users"
	cardCollection = "cards"
)

type CardManager struct {
	client *firestore.Client
}

// CreateCard gets the target user and inserts a new card in user's card collection.
// Returns the ID of the new card if the creation is successful.
func(m *CardManager) CreateCard(ctx context.Context, userID string, card docs.Card) (string, error) {
	user := m.client.Collection(userCollection).Doc(userID)
	cardRef, _, err := user.Collection(cardCollection).Add(ctx, card)
	if err != nil {
		return "", fmt.Errorf("failed to create a new card")
	}
	return cardRef.ID, nil
}