package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kkomissarov/beggar/db"
	"github.com/kkomissarov/beggar/models"
	"github.com/kkomissarov/beggar/schema/request_schema"
	"github.com/kkomissarov/beggar/utils"
	"log"
	"net/http"
)

func CreateAccount(ctx *gin.Context) {
	body := request_schema.CreateAccountBody{}
	if utils.BindRequestBody(ctx, &body) != nil {
		return
	}

	account := models.Account{
		Name:       body.Name,
		Balance:    body.Balance,
		CurrencyID: body.CurrencyID,
		UserID:     5,
	}

	result := db.DB.Create(&account).Preload("User").Preload("Currency").Find(&account, account.ID)
	if account.ID == 0 || result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Unable to create an account"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": account})
}

func GetAccountList(ctx *gin.Context) {
	//Logic
}

func GetAccount(ctx *gin.Context) {
	//Logic
}

func UpdateAccount(ctx *gin.Context) {
	//Logic
}

func DeleteAccount(ctx *gin.Context) {
	//Logic
}
