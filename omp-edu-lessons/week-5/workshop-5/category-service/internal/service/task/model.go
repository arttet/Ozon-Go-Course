package task

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Task struct {
	ID           uint64     `db:"id"`
	ExecDuration Duration   `db:"exec_duration"`
	StartedAt    *time.Time `db:"started_at"`
}

type Duration time.Duration

// Value converts Duration to a primitive value ready to written to a database.
func (d Duration) Value() (driver.Value, error) {
	return driver.Value(int64(d)), nil
}

// Scan reads a Duration value from database driver type.
func (d *Duration) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case int64:
		*d = Duration(v)
	case nil:
		*d = Duration(0)
	default:
		return fmt.Errorf("cannot sql.Scan() strfmt.Duration from: %#v", v)
	}
	return nil
}
