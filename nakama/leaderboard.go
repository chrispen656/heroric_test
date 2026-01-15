package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func createLeaderboard(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
	id := "snake_main_leaderboard"
	authoritative := false
	sort := "desc"
	operator := "best"
	reset := ""
	metadata := map[string]interface{}{}
	ranks := false

	nk.LeaderboardDelete(ctx, "4ec4f126-3f9d-11e7-84ef-b7c182b36521")

	if err := nk.LeaderboardCreate(ctx, id, authoritative, sort, operator, reset, metadata, ranks); err != nil {
		logger.Error("Error creating leaderboard: %v", err)
		return runtime.NewError("Cannot create leaderboard", 13)
	}

	return nil
}
