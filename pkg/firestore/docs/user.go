package docs

type User struct {
	ID string `firestore: "id"`
	Username string `firestore: "username"`
	Password string `firestore: "password"`
	Email string `firestore: "email"`
	ConnectedUserIDs []string `firestore: "connectedUserIDs"`
	Cards []Card `firestore: "cards"`
}

type Connection struct {
	userID1 string `firestore: "userID1"`
	userID2 string `firestore: "userID2"`
	relationship string `firestore: "relationship"`
}