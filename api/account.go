package api

import (
	"database/sql"
	"net/http"
	"strconv"

	db "github.com/OmarMuhammedAli/FinGo/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
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

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount(ctx *gin.Context) {
	var r getAccountRequest
	if err := ctx.ShouldBindUri(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := s.store.GetAccount(ctx, r.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	PageNo   int32 `form:"page_no" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=50"`
}

func (s *Server) listAccounts(ctx *gin.Context) {
	var r listAccountsRequest
	if err := ctx.ShouldBindQuery(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  r.PageSize,
		Offset: (r.PageNo - 1) * r.PageSize,
	}
	accounts, err := s.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)

}

type updateAccountRequest struct {
	Balance int64 `json:"balance" binding:"required"`
}

func (s *Server) updateAccount(ctx *gin.Context) {
	var r updateAccountRequest
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAccountParams{
		ID:      id,
		Balance: r.Balance,
	}
	account, err := s.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

type deleteAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) deleteAccount(ctx *gin.Context) {
	var r deleteAccountRequest
	if err := ctx.ShouldBindUri(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := s.store.DeleteAccount(ctx, r.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}
