package middleware

// import (
// 	"context"
// 	"encoding/json"
// 	"testing"

// 	"github.com/megre/dto"
// 	sqlxmock "github.com/zhashkevych/go-sqlxmock"
// )

// func TestCheckAllowedRole(t *testing.T) {
// 	mockDB, mock, _ := sqlxmock.Newx()
// 	defer mockDB.Close()

// 	adminConfig := json.RawMessage(`{"items": {"default": "disallow", "resources": {"add": {"action": "allow"}}}, "accounts": {"default": "disallow", "resources": {"suspend": {"action": "allow"}}}}`)
// 	t.Run("when a role is allowed to perform action", func(t *testing.T) {
// 		rows := sqlxmock.NewRows([]string{"id", "configuration", "role", "created_at", "updated_at"}).AddRow(1)
// 		rows.AddRow("1", adminConfig, "admin", "2023-01-18 17:49:14.591418", "2023-01-18 17:49:14.591418")
// 		mock.ExpectQuery("SELECT configuration FROM configurations WHERE role = 'admin'").WillReturnRows(rows)

// 		ctx := context.Background()
// 		ctx = context.WithValue(ctx, dto.Role, "admin")

// 		response := CheckAllowedRole()
// 	})
// }
