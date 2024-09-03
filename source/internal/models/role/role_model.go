package role_model

import "database/sql"

type RoleModel struct {
	storeContext *sql.DB
}

type Role struct {
	Id          int
	Name        string
	Description *string
}

func (model RoleModel) LinkAccountWithRole(account_id, role_id int) error {
	transaction, err := model.storeContext.Begin()
	if err != nil {
		return err
	}

	query := `
	DECLARE @account_id AS INT = @p1
	DECLARE @role_id AS INT = @p2
	INSERT INTO [rel_accounts_roles](
		account_id, role_id 	
	)
	VALUES (@account_id, @role_id);
	`
	if _, execErr := transaction.Exec(query, account_id, role_id); execErr != nil {
		transaction.Rollback()
		return execErr
	}

	if commitErr := transaction.Commit(); commitErr != nil {
		transaction.Rollback()
		return commitErr
	} else {
		return nil
	}
}

func New(store *sql.DB) RoleModel {
	return RoleModel{
		storeContext: store,
	}
}
