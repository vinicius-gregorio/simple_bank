package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/vinicius-gregorio/simple_bank/db/sqlc"
	"github.com/vinicius-gregorio/simple_bank/util"
)

type createUserRequest struct {
	//BODY PARAMS
	Username string `json:"username" binding:"required,alphanum"` //alpla num -> alphanumeric, no special characters
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username          string
	FullName          string
	Email             string
	PasswordChangedAt time.Time
	CreatedAt         time.Time
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { //BODY PARAMS
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
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
	response := createUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
	ctx.JSON(http.StatusOK, response)
}
