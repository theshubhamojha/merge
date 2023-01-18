package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/megre/app"
	"github.com/megre/dto"
	"github.com/megre/merrors"
	"golang.org/x/exp/slices"
)

const (
	getConfigurationQuery = `SELECT configuration FROM configurations WHERE role = $1`
)

type ResourceType struct {
	Action     string   `json:"action,omitempty"`
	Exceptions []string `json:"exceptions,omitempty"`
}
type Configuration struct {
	Resources map[string]ResourceType `json:"resources,omitempty"`
	Default   string                  `json:"default"`
}

func CheckAllowedRole(resource dto.ResourceIdentifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var db *sqlx.DB = app.GetDB()

			ctx := r.Context()
			role := ctx.Value(role)

			var data json.RawMessage
			err := db.GetContext(ctx, &data, getConfigurationQuery, role)
			if err != nil {
				fmt.Println(err.Error())
				dto.SendAPIResponse(w,
					dto.APIResponse{
						Message:   "something went wrong",
						ErrorCode: merrors.InternalServerError,
					},
					http.StatusInternalServerError,
				)
				return
			}

			var roleConfiguration map[string]Configuration
			err = json.Unmarshal(data, &roleConfiguration)
			if err != nil {
				fmt.Println(err.Error())
				dto.SendAPIResponse(w,
					dto.APIResponse{
						Message:   "something went wrong",
						ErrorCode: merrors.InternalServerError,
					},
					http.StatusInternalServerError,
				)
				return
			}

			isRoleAllowed := checkIsRoleAllowed(ctx, roleConfiguration, resource)
			if !isRoleAllowed {
				dto.SendAPIResponse(w,
					dto.APIResponse{
						Message:   "role not allowed to perform this action",
						ErrorCode: merrors.NotAllowed,
					},
					http.StatusInternalServerError,
				)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func checkIsRoleAllowed(ctx context.Context, configuration map[string]Configuration, resourceIdentifier dto.ResourceIdentifier) (isValid bool) {
	module := resourceIdentifier.Module
	resource := resourceIdentifier.Resource

	resourceConfigs := configuration[module]

	resourceType, ok := resourceConfigs.Resources[resource]
	if !ok {
		return resourceConfigs.Default == "allow"
	}

	accountId, _ := ctx.Value(accountID).(string)
	return resourceType.Action == "allow" && !slices.Contains(resourceType.Exceptions, accountId)
}
