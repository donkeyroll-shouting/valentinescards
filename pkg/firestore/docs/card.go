package docs

import (
	"time"
)

type Card struct {
	Name string `firestore: "name"`
	Description string `firestore: "description"`
	ImageURL string `firestore: "imageURL"`
	Created time.Duration `firestore: "created"`
	Used time.Duration`firestore: "used"`
}
