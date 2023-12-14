package frameworks

import (
	"context"
	"fmt"

	"github.com/Misoten-B/airship-backend/internal/drivers/config"
	"gorm.io/gorm"
)

const (
	ContextKeyConfig = "config"
	ContextKeyUID    = "uid"
	ContextKeyDB     = "db"
)

func GetConfig(ctx context.Context) (*config.Config, error) {
	config, ok := ctx.Value(ContextKeyConfig).(*config.Config)
	if !ok {
		return nil, fmt.Errorf("config is not set")
	}
	return config, nil
}

func GetUID(ctx context.Context) (string, error) {
	uid, ok := ctx.Value(ContextKeyUID).(string)
	if !ok {
		return "", fmt.Errorf("uid is not set")
	}
	return uid, nil
}

func GetDB(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value(ContextKeyDB).(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("db is not set")
	}
	return db, nil
}
