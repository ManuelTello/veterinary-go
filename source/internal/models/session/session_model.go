package session_model

import (
	"database/sql"
	"time"
)

type SessionModel struct {
	storeContext *sql.DB
}

type User struct {
	Id                int
	Email             string
	Password          string
	CreatedOn         time.Time
	FirstName         string
	LastName          string
	PhoneNumber       *int
	AlternativeNumber *int
}

func (model SessionModel) InsertNewUser(password, first_name, last_name, email string, created_on time.Time, phone_number, alternative_number *int) (int, error) {
	var query string
	transaction, transErr := model.storeContext.Begin()
	if transErr != nil {
		return -1, transErr
	}

	var detailsId int
	query = `
	DECLARE @email AS VARCHAR(250) = @p1
	DECLARE @password AS VARCHAR(250) = @p2
	DECLARE @created_on AS DATETIME = @p3
	DECLARE @first_name AS VARCHAR(250) = @p4
	DECLARE @last_name AS VARCHAR(250) = @p5
	DECLARE @phone_number AS INT = @p6
	DECLARE @alternative_number AS INT = @p7
	INSERT INTO [accounts] (
		email, password, created_on, first_name, last_name, phone_number, alternative_number
	) 
	OUTPUT Inserted.id 
	VALUES (@email, @password, @created_on, @first_name, @last_name, @phone_number, @alternative_number);
	`
	err := transaction.QueryRow(query, email, password, created_on, first_name, last_name, phone_number, alternative_number).Scan(&detailsId)
	if err != nil {
		transaction.Rollback()
		return -1, err
	}

	if commitErr := transaction.Commit(); commitErr != nil {
		return -1, commitErr
	} else {
		return detailsId, nil
	}
}

func (model SessionModel) DoesEmailExists(email string) (int, error) {
	transaction, err := model.storeContext.Begin()
	if err != nil {
		return -1, err
	}

	query := `
	DECLARE @email AS VARCHAR(250) = @p1
	SELECT COUNT(*) FROM [accounts] WHERE [accounts].email = @email
	`
	var emailExists int
	scanErr := transaction.QueryRow(query, email).Scan(&emailExists)

	if scanErr != nil {
		transaction.Rollback()
		return -1, scanErr
	}

	if transErr := transaction.Commit(); transErr != nil {
		transaction.Rollback()
		return -1, transErr
	}

	return emailExists, nil
}

func New(store *sql.DB) SessionModel {
	model := SessionModel{
		storeContext: store,
	}

	return model
}
