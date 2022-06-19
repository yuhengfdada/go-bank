package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhengfdada/go-bank/db"
)

// swagger:parameters createAccount
type CreateAccountReq struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,validCurrency"`
}

// swagger:route POST /accounts accounts createAccount
// Create a new account
func (server *Server) createAccount(c *gin.Context) {
	req := CreateAccountReq{}
	err := c.BindJSON(&req) // 相当于decode JSON
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	acc, err := server.store.CreateAccount(c, db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acc)
}

type GetAccountReq struct {
	ID int64 `uri:"id" binding:"required,gt=0"`
}

func (server *Server) getAccount(c *gin.Context) {
	var getAccReq GetAccountReq
	err := c.ShouldBindUri(&getAccReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	acc, err := server.store.GetAccount(c, getAccReq.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acc)
}

// swagger:parameters listAccount
type ListAccountReq struct {
	// in: path
	PageID int32 `form:"page_id" binding:"required,min=1"`
	// in: path
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// swagger:route GET /accounts accounts listAccount
// List all accounts
func (server *Server) listAccount(c *gin.Context) {
	var req ListAccountReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accs, err := server.store.ListAccounts(c, db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accs)
}
