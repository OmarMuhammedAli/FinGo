package api

import (
	"net/http"

	db "github.com/OmarMuhammedAli/FinGo/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding="required"`
	Currency string `json:"currency" binding:"required,oneof=USD EGP EUR"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var r createAccountRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    r.Owner,
		Currency: r.Currency,
		Balance:  0,
	}
	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
