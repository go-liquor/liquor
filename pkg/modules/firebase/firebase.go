package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/go-liquor/liquor/v3/config"
	"google.golang.org/api/option"
)

// NewApp creates a new Firebase App instance using the provided configuration.
//
// Parameters:
//   - cfg: Configuration object containing Firebase settings
//
// Returns:
//   - *firebase.App: Firebase application instance
//   - error: nil if successful, error otherwise
//
// Example:
//
//	app, err := firebase.NewApp(config)
func NewApp(cfg *config.Config) (*firebase.App, error) {
	if cfg.GetString("firebase.configFile") != "" {
		opt := option.WithCredentialsFile(cfg.GetString("firebase.configFile"))
		return firebase.NewApp(context.Background(), nil, opt)
	}
	return nil, nil
}

// NewAuth creates a new Firebase Auth client from the Firebase App instance.
//
// Parameters:
//   - app: Firebase application instance
//
// Returns:
//   - *auth.Client: Firebase authentication client
//   - error: nil if successful, error otherwise
//
// Example:
//
//	auth, err := firebase.NewAuth(app)
func NewAuth(app *firebase.App) (*auth.Client, error) {
	return app.Auth(context.Background())
}

// NewFirestore creates a new Firestore client from the Firebase App instance.
//
// Parameters:
//   - app: Firebase application instance
//
// Returns:
//   - *firestore.Client: Firestore database client
//   - error: nil if successful, error otherwise
//
// Example:
//
//	firestore, err := firebase.NewFirestore(app)
func NewFirestore(app *firebase.App) (*firestore.Client, error) {
	return app.Firestore(context.Background())
}
