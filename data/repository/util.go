package repository

import (
	"context"
	"fmt"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/zap"
)

func isDevelopmentEnv() bool {
	return config.GetString("APP_ENV", "dev") == "dev"
}

func WithTxResult[T any](ctx context.Context, client *ent.Client, logger *logger.Logger, fn func(tx *ent.Tx) (T, error)) (T, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		var zero T
		logger.Error("Failed to begin transaction", zap.Error(err))
		return zero, fmt.Errorf("starting transaction failed: %w", err)
	}

	if isDevelopmentEnv() {
		logger.Debug("Transaction started in development mode...")
	}

	result, err := fn(tx)
	if err != nil {
		logger.Error("Transaction operation failed, rolling back...", zap.Error(err))
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Error("Rollback failed", zap.Error(rollbackErr))
			return result, fmt.Errorf("rollback failed: %v (original error: %w)", rollbackErr, err)
		}
		return result, err
	}

	if err = tx.Commit(); err != nil {
		logger.Error("Commit failed", zap.Error(err))
		return result, fmt.Errorf("commit failed: %w", err)
	}

	if isDevelopmentEnv() {
		logger.Debug("Transaction committed successfully in development mode")
	}

	return result, nil
}

func WithTx(ctx context.Context, client *ent.Client, logger *logger.Logger, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		logger.Error("Failed to start transaction", zap.Error(err))
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	if isDevelopmentEnv() {
		logger.Info("Transaction started in development mode...")
	}

	if err = fn(tx); err != nil {
		logger.Error("Transaction operation failed, rolling back...", zap.Error(err))
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Error("Rollback failed", zap.Error(rollbackErr))
			return fmt.Errorf("rollback failed: %v (original error: %w)", rollbackErr, err)
		}
		logger.Info("Transaction rolled back successfully")
		return err
	}

	if err = tx.Commit(); err != nil {
		logger.Error("Commit failed", zap.Error(err))
		return fmt.Errorf("commit failed: %w", err)
	}

	if isDevelopmentEnv() {
		logger.Info("Transaction committed successfully in development mode")
	}

	return nil
}
