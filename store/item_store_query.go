package store

const (
	addItemQuery = `INSERT INTO items (id, account_id, name, image_url, quantity, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, false, now(), now())
		RETURNING id, account_id, name, image_url, quantity, created_at, updated_at
	`

	getItemByIDQuery = `SELECT id, account_id, quantity, name, image_url FROM items WHERE id = $1`
)
