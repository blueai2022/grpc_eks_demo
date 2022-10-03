package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	db "github.com/blueai2022/appsubmission/db/sqlc"
	"github.com/blueai2022/appsubmission/token"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey  = "authorization"
	authTypeBearer = "bearer"
	authPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker, store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authHeaderKey)
		if len(authHeader) == 0 {
			err := fmt.Errorf("%s header is not provided", authHeaderKey)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authHeader)

		if len(fields) < 2 {
			err := fmt.Errorf("invalid %s header format", authHeaderKey)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != authTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		getAccountParam := db.GetActiveApiAccountParams{
			Username:    payload.Username,
			ServiceType: "ICD",
		}

		_, err = store.GetActiveApiAccount(ctx, getAccountParam)
		if err != nil {
			if err == sql.ErrNoRows {
				verr := fmt.Errorf("api account not found for user %s", payload.Username)
				ctx.AbortWithStatusJSON(http.StatusNotFound, errorResponse(verr))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.Set(authPayloadKey, payload)
		ctx.Next()
	}
}
