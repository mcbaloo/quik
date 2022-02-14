package Routes

import (
	"quik-assessment/Controllers"
	"quik-assessment/Handlers"
	"quik-assessment/Middlewares"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api/v1")
	{
		grp1.POST("wallets/:wallet_id/debit", Handlers.WalletHandler, Controllers.DebitWallet)
		grp1.POST("wallets/:wallet_id/credit", Handlers.WalletHandler, Controllers.CreditWallet)
		grp1.GET("wallets/:wallet_id/balance", Controllers.GetWalletBalanceByID)
	}
	r.Use(Middlewares.LoggerToFile())
	return r
}
