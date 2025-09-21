// internal/dto/account.dto.go
package dto

import db "github.com/chinsiang99/simple-bank-go-project/internal/database/sqlc"

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type GetAccountsRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=30"`
	Offset int32 `form:"offset" binding:"min=0"`
}

type GetAccountsResponse struct {
	Accounts   []db.Account `json:"accounts"`
	TotalCount int64        `json:"total_count"`
}

// AccountListResponse defines response for GetAccounts
type AccountListResponse struct {
	Data      []db.Account `json:"data"`
	TotalData int64        `json:"total_data"`
}
