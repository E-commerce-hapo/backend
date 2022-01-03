package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/E-commerce-hapo/backend/pkg/auth"

	"github.com/E-commerce-hapo/backend/pkg/idx"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	router.POST("/login", Login)
	router.POST("/todo", CreateTodo)
	router.POST("/logout", Logout)
	router.POST("/refresh", Refresh)
	log.Fatal(router.Run(":8080"))
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	ts, err := auth.GenerateToken(idx.ID(user.ID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

type Todo struct {
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
}

func CreateTodo(c *gin.Context) {
	//var td Todo
	//if err := c.ShouldBindJSON(&td); err != nil {
	//	c.JSON(http.StatusUnprocessableEntity, "invalid json")
	//	return
	//}
	//Extract the access token metadata
	customClaim, err := auth.GetCustomClaimsFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userID := customClaim.UserID
	//you can proceed to save the Todo to a database
	//but we will just return it to the caller:

	c.JSON(http.StatusCreated, userID)
}

func Logout(c *gin.Context) {
	_, err := auth.GetCustomClaimsFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(auth.JWTCustomClaims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["UserID"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := auth.GenerateToken(idx.ID(userId))
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}
