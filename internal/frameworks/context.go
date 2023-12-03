package frameworks

import (
	"context"
	"fmt"

	"github.com/Misoten-B/airship-backend/config"
)

const (
	ContextKeyConfig = "config"
	ContextKeyUID    = "uid"
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
