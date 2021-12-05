package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func AcquireTryLock(ctx context.Context, tx *sqlx.Tx, lockID int32, entityID int32) (bool, error) {
	var isAcquired bool
	err := tx.GetContext(ctx, &isAcquired, fmt.Sprintf("select pg_try_advisory_xact_lock(%d, %d)", lockID, entityID))
	return isAcquired, err
}
