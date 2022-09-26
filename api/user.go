package api

import (
	"database/sql"
	"net/http"

	db "github.com/blueai2022/appsubmission/db/sqlc"
	"github.com/blueai2022/appsubmission/secure"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username        string         `json:"username"`
	Password        string         `json:"password"`
	FullName        string         `json:"full_name"`
	Email           string         `json:"email"`
	Agency          sql.NullString `json:"agency"`
	AppContact      sql.NullString `json:"app_contact"`
	AppContactEmail sql.NullString `json:"app_contact_email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := secure.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:        req.FullName,
		HashedPassword:  hashedPassword,
		FullName:        req.FullName,
		Email:           req.Email,
		Agency:          req.Agency,
		AppContact:      req.AppContact,
		AppContactEmail: req.AppContactEmail,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)

}
