package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	initStart := time.Now()

	if err := initializer.RegisterRpc("healthcheck", RpcHealthcheck); err != nil {
		return err
	}

	if err := createLeaderboard(ctx, logger, db, nk); err != nil {
		return err
	}

	logger.Info("Module loading in %dms", time.Since(initStart).Milliseconds())
	return nil
}
