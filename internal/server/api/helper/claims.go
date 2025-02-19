package helper

import (
	"context"
	"fmt"
	"github.com/zasuchilas/gophkeeper/internal/server/jwtmanager"
	"log/slog"
)

func GetCtxUserID(ctx context.Context) (int64, error) {
	claims, err := jwtmanager.GetClaims(ctx)
	if err != nil {
		return 0, err
	}

	userID := claims.ID
	slog.Debug("REQUEST (api level logging)", slog.Int64("UserID", claims.ID))
	if userID == 0 {
		return 0, fmt.Errorf("nil userID received")
	}

	return userID, nil
}
