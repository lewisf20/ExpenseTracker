package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lewisf20/ExpenseTracker/internal/store"
	"github.com/lewisf20/ExpenseTracker/services/tokenservice"
)

func register(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	store.Users = append(store.Users, user)

	token, err := tokenservice.CreateToken(123)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "Registered successfully.",
		"token": token.Token,
	})
}

func signIn(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for _, u := range store.Users {
		if u.Email == user.Email && u.Password == user.Password {
			token, err := tokenservice.CreateToken(123)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"msg":   "Signed in successfully.",
				"token": token.Token,
			})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "Failed to sign in."})
}
