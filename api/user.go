package api

import (
	"database/sql"
	"net/http"

	db "github.com/DhawalYerekar/Users-App/DB/sqlc"
	"github.com/DhawalYerekar/Users-App/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Name     string         `json:"name" binding:"required"`
	Username sql.NullString `json:"username"`
	Age      sql.NullInt32  `json:"age"`
	Email    string         `json:"email" binding:"required" `
	Password string         `json:"password" binding:"required"`
	Address  string         `json:"address"`
	Contact  string         `json:"contact"`
	Gender   sql.NullString `json:"gender"`
}

type userResponse struct {
	Name     string         `json:"name"`
	Username sql.NullString `json:"username"`
	Age      sql.NullInt32  `json:"age"`
	Email    string         `json:"email"`
	Address  string         `json:"address"`
	Contact  string         `json:"contact"`
	Gender   sql.NullString `json:"gender"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	// Validate incoming JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	hashedPassword, err := util.Hashpassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	arg := db.CreateUserParams{
		Name:     req.Name,
		Username: req.Username,
		Age:      req.Age,
		Email:    req.Email,
		Password: hashedPassword,
		Address:  req.Address,
		Contact:  req.Contact,
		Gender:   req.Gender,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return

	}

	resp := userResponse{
		Name:     user.Name,
		Username: user.Username,
		Age:      user.Age,
		Email:    user.Email,
		Address:  user.Address,
		Contact:  user.Contact,
		Gender:   user.Gender,
	}

	ctx.JSON(http.StatusCreated, resp)

}
