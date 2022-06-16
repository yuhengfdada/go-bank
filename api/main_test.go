package api

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuhengfdada/go-bank/util"
)

func TestMain(m *testing.M) {
	gin.SetMode("test")

	m.Run()
}

func getTestAPIConfig() *util.Config {
	return &util.Config{
		TokenSymmetricKey:   util.GetRandomString(32),
		AccessTokenDuration: time.Minute,
	}
}
