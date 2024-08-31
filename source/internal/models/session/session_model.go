package session_model

import (
	"database/sql"
	"time"
)

type SessionModel struct {
	storeContext *sql.DB
}

type User struct {
	Id          int
	Username    string
	Password    string
	FirstName   string
	LastName    string
	DateCreated time.Time
	Email       string
	PhoneNumber int
	RoleId      int
}

func (model SessionModel) InsertNewUser(username, password, firstname, lastname, email string, datecreated time.Time, phonenumber *int) error {
	var query string
	transaction, transErr := model.storeContext.Begin()
	if transErr != nil {
		return transErr
	}

	var detailsId int
	query = `
	DECLARE @first_name AS VARCHAR(250) = @p1
	DECLARE @last_name AS VARCHAR(250) = @p2
	DECLARE @email AS VARCHAR(250) = @p3
	DECLARE @date_created AS DATETIME = @p5
	DECLARE @phone_number AS INT = @p6
	INSERT INTO [accounts_details] (
		first_name, last_name, email, date_created, phone_number
	) 
	OUTPUT Inserted.id 
	VALUES (@first_name, @last_name, @email, ,@date_created ,@phone_number);
	`
	err := transaction.QueryRow(query, firstname, lastname, email, datecreated, phonenumber).Scan(&detailsId)
	if err != nil {
		transaction.Rollback()
		return err
	}

	query = `
	DECLARE @username AS VARCHAR(250) = @p1
	DECLARE @password AS VARCHAR(250) = @p2
	DECLARE @details_id AS INT = @p3
	INSERT INTO [accounts](
		username, password, details_id
	)
	VALUES (@username, @password, @details_id)
	`
	_, err = transaction.Exec(query, username, password, detailsId)

	if err != nil {
		transaction.Rollback()
		return err
	}

	if commitErr := transaction.Commit(); commitErr != nil {
		return commitErr
	} else {
		return nil
	}
}

func (model SessionModel) SearchAccountByUsername(username string) (*User, error) {
	transaction, err := model.storeContext.Begin()
	var user *User
	if err != nil {
		return nil, err
	}

	query := `
	DECLARE @username VARCHAR(250) = @p1
	SELECT [accounts].id, [accounts].username, [accounts].password, 
	[accounts_details].first_name, [accounts_details].last_name, [accounts_details].date_created,
	[accounts_details].email, [accounts_details].phone_number,
	FROM [accounts] INNER JOIN [accounts_details] 
	ON [accounts].details_id = [accounts_details].id
	WHERE accounts.username = @username
	`
	rows, queryErr := transaction.Query(query, username)
	if queryErr != nil {
		return nil, queryErr
	}

	for rows.Next() {
		user = new(User)
		scanErr := rows.Scan(&user.Id, &username, &user.Password, &user.FirstName, &user.LastName, &user.DateCreated, &user.Email, &user.PhoneNumber)
		if scanErr != nil {
			return nil, scanErr
		}
	}

	transErr := transaction.Commit()
	if transErr != nil {
		transaction.Rollback()
		return nil, transErr
	}

	return user, nil
}

func (model SessionModel) DoesUsernameExists(username string) (int, error) {
	transaction, err := model.storeContext.Begin()
	if err != nil {
		return -1, err
	}

	query := `
	DECLARE @username AS VARCHAR(250) = @p1
	SELECT COUNT(*) FROM [accounts] WHERE [accounts].username = @username
	`
	var userExists int
	scanErr := transaction.QueryRow(query, username).Scan(&userExists)

	if scanErr != nil {
		transaction.Rollback()
		return -1, scanErr
	}

	if transErr := transaction.Commit(); transErr != nil {
		transaction.Rollback()
		return -1, transErr
	}

	return userExists, nil
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
