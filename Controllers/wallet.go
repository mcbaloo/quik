package Controllers

import (
	"fmt"
	"net/http"
	"quik-assessment/Config"
	"quik-assessment/Helpers"
	"quik-assessment/Models"
	"quik-assessment/RequestModels"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shopspring/decimal"
)

var message Models.Message

func GetWalletBalanceByID(c *gin.Context) {
	wallet_id := c.Params.ByName("wallet_id")
	var wallet Models.Wallet
	//check redis
	redisClient := Config.RedisConnection()
	data, rederr := Config.GetRedisData(&redisClient, wallet_id, &wallet)
	if rederr == nil {
		c.JSON(http.StatusOK, data)
	} else {
		err := Models.GetWalletBalanceByID(&wallet, wallet_id)
		if err != nil {
			message.Message = "Wallet Not found with specified wallet id"
			c.JSON(http.StatusNotFound, &message)
		} else {
			c.JSON(http.StatusOK, wallet)
		}
	}
}

func CreditWallet(c *gin.Context) {
	wallet_id := c.Params.ByName("wallet_id")
	var wallet Models.Wallet
	var creditObj RequestModels.Transaction

	err := Models.GetWalletBalanceByID(&wallet, wallet_id)
	if err != nil {
		message.Message = "Wallet Not found"
		c.JSON(http.StatusNotFound, &message)
		c.Abort()
		return
	}
	err = c.ShouldBindBodyWith(&creditObj, binding.JSON)
	if err != nil {
		fmt.Println(err.Error())
		message.Message = "An Error Occured"
		c.JSON(http.StatusBadGateway, &message)
		c.Abort()
		return
	}
	wallet.Balance = wallet.Balance + creditObj.Amount.Value
	err = Models.UpdateWallet(&wallet)
	if err != nil {
		message.Message = "Error Occured while processing request"
		c.JSON(http.StatusBadGateway, &message)
		c.Abort()
		return
	}
	//update Redis value
	redisClient := Config.RedisConnection()
	err = Config.RemoveRedisData(&redisClient, wallet_id)
	if err != nil {
		fmt.Println("Error occured while trying to remove redis key")
	}
	err = Config.SetRedisData(&redisClient, wallet_id, wallet)
	if err != nil {
		fmt.Print("Error occured while trying to save request to redis")
	}
	c.JSON(http.StatusOK, wallet)
	c.Abort()
}
func DebitWallet(c *gin.Context) {
	wallet_id := c.Params.ByName("wallet_id")
	var wallet Models.Wallet
	var creditObj RequestModels.Transaction
	err := Models.GetWalletBalanceByID(&wallet, wallet_id)
	if err != nil {
		message.Message = "Wallet Not found"
		c.JSON(http.StatusNotFound, &message)
		c.Abort()
		return
	}
	err = c.ShouldBindBodyWith(&creditObj, binding.JSON)
	if err != nil {
		message.Message = "An Error Occured"
		c.JSON(http.StatusBadGateway, &message)
		c.Abort()
		return
	}
	//check if withdrawal is going to be less than zero
	if Helpers.ComputeBalanceAfterWithdrawal(decimal.NewFromFloat(wallet.Balance), decimal.NewFromFloat(creditObj.Amount.Value)) != nil {
		message.Message = "Insufficient balance"
		c.JSON(http.StatusOK, &message)
		c.Abort()
		return
	}
	wallet.Balance = wallet.Balance - creditObj.Amount.Value
	//update redis value

	redisClient := Config.RedisConnection()
	err = Config.RemoveRedisData(&redisClient, wallet_id)
	if err != nil {
		fmt.Println("Error occured while trying to remove redis key")
	}
	err = Config.SetRedisData(&redisClient, wallet_id, wallet)
	if err != nil {
		fmt.Print("Error occured while trying to save request to redis")
	}
	err = Models.UpdateWallet(&wallet)
	if err != nil {
		message.Message = "Error Occured while processing request"
		c.JSON(http.StatusBadGateway, &message)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &wallet)
	c.Abort()
}
