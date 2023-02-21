package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kkomissarov/beggar/db"
	"github.com/kkomissarov/beggar/managers/authManager"
	"github.com/kkomissarov/beggar/models"
	"github.com/kkomissarov/beggar/schema/request_schema"
	"github.com/kkomissarov/beggar/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func SignUp(ctx *gin.Context) {
	// Bind request body
	body := request_schema.SignUpBody{}
	if utils.BindRequestBody(ctx, &body) != nil {
		return
	}

	// Hash user password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		log.Println("Unable to hash password", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a user"})
		return
	}

	// Check if user does not exist
	existingUser := models.User{}
	db.DB.First(&existingUser, "email = ?", body.Email)
	if existingUser.ID != 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User with the specified email already exists"})
		return
	}

	// Create user in database
	user := models.User{Email: body.Email, Password: string(passwordHash)}
	db.DB.Create(&user)
	if user.ID == 0 {
		log.Println("Unable to write user to database")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a user"})
		return
	}

	// Return user
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func Login(ctx *gin.Context) {
	// Bind request body
	body := request_schema.LoginBody{}
	if utils.BindRequestBody(ctx, &body) != nil {
		return
	}

	// Find user with requested email
	user := models.User{}
	db.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		log.Println("Unable to find a user with requested email")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong login or password"})
		return
	}

	// Check Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		log.Println("User password does not match the saved hash: ", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong login or password"})
		return
	}

	// Generate JWT token
	var jwtToken *jwt.Token
	jwtToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	jwtTokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("Unable to create JWT-token:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to login"})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{"token": jwtTokenString})
}

func Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	err := authManager.RevokeToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to logout"})
	}
	ctx.AbortWithStatus(http.StatusNoContent)
}
