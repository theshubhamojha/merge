package middleware

type ContextKey string

const (
	role      ContextKey = "role"
	email     ContextKey = "email"
	accountID ContextKey = "account_id"
)
