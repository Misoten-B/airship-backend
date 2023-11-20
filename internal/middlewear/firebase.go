package middlewear

import (
	"context"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/drivers"
)

func Guard(ctx context.Context, idToken string, next http.HandlerFunc) {
	app, err := drivers.InitFirebase()
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
}
