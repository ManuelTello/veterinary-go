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

func (model RoleModel) LinkAccountRole(account_id, role_id int) error {
	transaction, err := model.storeContext.Begin()
	if err != nil {
		return err
	}

	query := `
	`
	if _, execErr := transaction.Exec(query, account_id, role_id); execErr != nil {
		return execErr
	}


	return nil
}

func New(store *sql.DB) RoleModel {
	return RoleModel{
		storeContext: store,
	}
}
