package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lewisf20/ExpenseTracker/internal/store"
)

func register(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	store.Users = append(store.Users, user)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "Registered successfully.",
		"token": "023a246d-274d-4a84-8a9d-dd12b406d6c7",
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
			ctx.JSON(http.StatusOK, gin.H{
				"msg":   "Signed in successfully.",
				"token": "023a246d-274d-4a84-8a9d-dd12b406d6c7",
			})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "Failed to sign in."})
}
