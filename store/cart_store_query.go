package store

const (
	getCartDetailsQuery = `SELECT 
		id, quantity, account_id, item_id FROM cart 
		WHERE is_deleted = FALSE AND account_id = $1 AND item_id = $2
	`

	insertCartQuery = `INSERT INTO cart
		(id, account_id, item_id, quantity, is_deleted, created_at, updated_at)
		VALUES ($1, $2, $3, $4, false, now(), now())
		RETURNING id, account_id, item_id, quantity, is_deleted, created_at, updated_at
	`

	updateCartQuantityQuery = `UPDATE cart
		SET quantity = $2, updated_at = now()
		WHERE id = $1 AND is_deleted = FALSE AND account_id = $3
		RETURNING id, account_id, quantity, item_id, updated_at, created_at
	`

	listCartItemsQuery = `SELECT c.id, c.account_id, c.quantity, i.name, i.image_url FROM cart AS c 
		INNER JOIN items i ON i.id = c.item_id
		WHERE c.account_id = $1 
		OFFSET $2 LIMIT $3
	`
)
