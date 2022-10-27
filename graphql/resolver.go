//go:generate go run github.com/99designs/gqlgen --verbose

package graphql

import (
	"github.com/blueai2022/appsubmission/config"
	db "github.com/blueai2022/appsubmission/db/sqlc"
)

type Resolver struct {
	Config *config.Config
	Store  db.Store
}
