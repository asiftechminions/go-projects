package handlers

import (
	"07072025/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var va *validator.Validate

func RegisterUser(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errot": "invalid input"})

	}

	if err := va.Struct(input); err != nil {
		ve := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": ve.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}
