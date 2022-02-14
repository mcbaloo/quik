package Handlers

import (
	"net/http"
	"quik-assessment/Helpers"
	"quik-assessment/RequestModels"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shopspring/decimal"
	"github.com/thedevsaddam/govalidator"
)

func WalletHandler(c *gin.Context) {
	var creditObj RequestModels.Transaction
	err := c.ShouldBindBodyWith(&creditObj, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		c.Abort()
	}
	rules := govalidator.MapData{
		"amount": []string{"required"},
	}
	messages := govalidator.MapData{
		"amount": []string{"required:You must provide amount"},
	}
	opts := govalidator.Options{
		Data:     &creditObj,
		Rules:    rules,
		Messages: messages,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e,
		})
		c.Abort()
	}
	err = Helpers.CheckTransactionAmount(decimal.NewFromFloat(creditObj.Amount.Value))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}
}
