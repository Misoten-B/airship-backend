package drivers

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InitFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %w", err)
	}
	return app, nil
}

func GetFirebaseAuthClient() (*auth.Client, error) {
	app, err := InitFirebase()
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %w", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %w", err)
	}
	return client, nil
}
