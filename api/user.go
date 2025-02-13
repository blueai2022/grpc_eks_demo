package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/blueai2022/appsubmission/crypt"
	db "github.com/blueai2022/appsubmission/db/sqlc"
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

type userResponse struct {
	Username          string         `json:"username"`
	FullName          string         `json:"full_name"`
	Email             string         `json:"email"`
	Agency            sql.NullString `json:"agency"`
	AppContact        sql.NullString `json:"app_contact"`
	AppContactEmail   sql.NullString `json:"app_contact_email"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt         time.Time      `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		Agency:            user.Agency,
		AppContact:        user.AppContact,
		AppContactEmail:   user.AppContactEmail,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := crypt.HashPassword(req.Password)
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

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)

}

type getUserRequest struct {
	Username string `uri:"username" binding:"required,min=6"`
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

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	User        userResponse `json:"user"`
	SessionID   string       `json:"session_id"`
	AccessToken string       `json:"access_token"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = crypt.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, payload, err := server.tokenMaker.CreateToken(
		req.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		User:        newUserResponse(user),
		SessionID:   payload.ID.String(),
		AccessToken: accessToken,
	}
	ctx.JSON(http.StatusOK, rsp)
}
