package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func BindRequestBody(ctx *gin.Context, body interface{}) error {
	if err := ctx.Bind(&body); err != nil {
		log.Println("Error while binding", err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Wrong request parameters"},
		)
		return err
	}
	return nil
}
