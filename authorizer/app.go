package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type App struct {
	Context context.Context
	Router  *gin.Engine
}

var (
	SECRET_KEY = []byte("SecretYouShouldHide")
)

func InitApplication() (*App, error) {
	gin.SetMode(gin.DebugMode)

	app := &App{
		Context: context.TODO(),
		Router:  gin.Default(),
	}
	app.SetupRouter()
	return app, nil
}

func (app *App) SetupRouter() {
	app.Router.POST("/api/v1/authorizer/login", app.Login)
	app.Router.POST("/api/v1/authorizer/register", app.Register)
	app.Router.GET("/api/v1/authorizer/verify", app.VerifyToken)
}

func (app *App) Run(port int) {
	fmt.Printf("App is running on %d\n port", port)
	app.Router.Run(fmt.Sprintf(":%d", port))
}

func (app *App) Login(c *gin.Context) {
	token, err := createToken()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (app *App) Register(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (app *App) VerifyToken(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if token[0] == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	log.Println(token)
	ns := strings.Replace(token[0], "Bearer ", "", 1)
	claims, err := verifyToken(ns)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	response_obj := &AuthenticationSession{
		Subject: claims.Id,
		Extra: map[string]interface{}{
			"email": claims.Email,
		},
	}
	c.JSON(http.StatusOK, response_obj)
}

func createToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    "1",
		"email": "test@te",
		"exp":   fmt.Sprintf("%d", time.Now().Add(time.Hour*24).Unix()),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (*MyCustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		log.Println(claims)
		return claims, nil
	} else {
		fmt.Println(err)
		return nil, errors.New("token is not valid")
	}
}
