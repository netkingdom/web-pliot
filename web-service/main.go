package main

import (
	"context"
	"ginTutorial/component"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"time"
)

func main()  {
	app := fx.New(
		fx.Provide(
			//service
			component.NewControlService,
			//handler
			component.NewControlHandler,
		),
		fx.Invoke(
			SetRouter,
		),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil{
		log.Fatal(err)
	}
	//service := component.NewControlService()

	//SetRouter(component.NewControlHandler(service))
}

func SetRouter(c *component.ControlHandler){
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/control/:id" , c.GetControlInfo)
	r.Run()
}