package task

import "time"

type Task struct {
	ID        uint64     `db:"id"`
	StartedAt *time.Time `db:"started_at"`
}
