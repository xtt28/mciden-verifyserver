package db

import "time"

type Student struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Academy   string
}

type Registration struct {
	StudentID  int64
	PlayerUUID string
	CreatedAt  *time.Time
}

type VerifyIntent struct {
	ID string
	StudentID int64
	PlayerUUID string
	CreatedAt *time.Time
	ExpiresAt *time.Time
}

func (i VerifyIntent) IsExpired() bool {
	return i.ExpiresAt.UTC().Before(time.Now().UTC())
}

func (i VerifyIntent) IsEmpty() bool {
	return i == VerifyIntent{}
}