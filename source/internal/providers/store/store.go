package providers_store

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

func NewStore() (*sql.DB, error) {
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=%s", os.Getenv("SQL_USER"), os.Getenv("SQL_PASSWORD"), os.Getenv("SQL_HOST"), 1433, os.Getenv("SQL_SCHEMA"), "disable")
	sqlDb, connectionErr := sql.Open("sqlserver", connectionString)
	if connectionErr != nil {
		return nil, connectionErr
	}

	return sqlDb, nil
}
