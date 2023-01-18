package store

const (
	createAccountQuery = `INSERT INTO accounts 
		(id, name, email, password, role, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, true, now(), now())
		RETURNING id, name, email, password, role, is_active, created_at, updated_at
	`

	suspendAccountQuery = `UPDATE accounts SET is_active = false, updated_at = now() WHERE id = $1`

	getAccountDetailsByEmailQuery = `SELECT id, name, email, password, role, is_active, created_at, updated_at FROM accounts WHERE email = $1`
)
