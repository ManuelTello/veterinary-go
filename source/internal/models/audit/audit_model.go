package audit_model

import (
	"database/sql"
	"time"
)

type AuditModel struct {
	storeContext *sql.DB
}

type Audit struct {
	Id           int
	Token        string
	AccountId    int
	DateLoggedIn time.Time
	DateDiesOn   time.Time
}

func New(store *sql.DB) AuditModel {
	model := AuditModel{
		storeContext: store,
	}

	return model
}
