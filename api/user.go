package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/blueai2022/appsubmission/db/sqlc"
	"github.com/blueai2022/appsubmission/secure"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
		Username:        req.Username,
		HashedPassword:  hashedPassword,
		FullName:        req.FullName,
		Email:           req.Email,
		Agency:          req.Agency,
		AppContact:      req.AppContact,
		AppContactEmail: req.AppContactEmail,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//TODO remove hashedPassword from response
	ctx.JSON(http.StatusOK, user)

}

type getUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			notFoundErr := fmt.Errorf("user not found with username %s", req.Username)
			ctx.JSON(http.StatusNotFound, errorResponse(notFoundErr))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//TODO remove hashedPassword from response
	ctx.JSON(http.StatusOK, user)
}
