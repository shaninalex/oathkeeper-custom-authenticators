package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Context  context.Context
	Router   *gin.Engine
	Accounts []Account
}

func InitApplication() (*App, error) {
	gin.SetMode(gin.DebugMode)

	app := &App{
		Context: context.TODO(),
		Router:  gin.Default(),
		Accounts: []Account{
			{
				ID:               1,
				Email:            "test@test.com",
				Username:         "test_1",
				SubscriptionPlan: "basic",
			},
			{
				ID:               2,
				Email:            "test2@test2.com",
				Username:         "test_2",
				SubscriptionPlan: "enterprise",
			},
		},
	}
	app.SetupRouter()
	return app, nil
}

func (app *App) SetupRouter() {
	app.Router.GET("/api/v1/account/me", app.AccountDetails)
}

func (app *App) Run(port int) {
	fmt.Printf("App is running on %d\n port", port)
	app.Router.Run(fmt.Sprintf(":%d", port))
}

func (app *App) AccountDetails(c *gin.Context) {
	// user_id_s := c.Request.Header.Get("X-User")
	// if user_id_s == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "user_id required"})
	// 	return
	// }

	// user_id, err := strconv.Atoi(user_id_s)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
	// 	return
	// }

	// for _, account := range app.Accounts {
	// 	if account.ID == int64(user_id) {
	// 		c.JSON(http.StatusBadRequest, account)
	// 		return
	// 	}
	// }

	// c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "account with given id is not found"})
	headers := c.Request.Header
	c.JSON(http.StatusOK, headers)
}
