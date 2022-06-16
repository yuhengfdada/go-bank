package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhengfdada/go-bank/db"
)

type CreateTransferReq struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,validCurrency"`
}

func (server *Server) validateTransfer(c *gin.Context, req *CreateTransferReq) error {
	// check validity of ID
	fromacc, err := server.store.GetAccount(c, req.FromAccountID)
	if err != nil {
		return err
	}
	if fromacc.Balance < req.Amount {
		return fmt.Errorf("transfer: no sufficient balance")
	}
	_, err = server.store.GetAccount(c, req.ToAccountID)
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) createTransfer(c *gin.Context) {
	var req CreateTransferReq
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = server.validateTransfer(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := server.store.TransferTx(c, db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
