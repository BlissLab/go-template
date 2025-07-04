package router

import (
	"context"
	"fmt"
	"go-template/config"
	"go-template/internal/dto"

	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) error {
	g := gin.Default()
	Router(g)
	err := g.Run(config.C.API.ListenAddr())
	return err
}

func Router(g *gin.Engine) {
	app := g.Group(fmt.Sprintf("/%s", config.C.APP))
	app.GET("/ping", func(ctx *gin.Context) {
		dto.Http.Ok(ctx, nil)
	})
	app.GET("/healthz", func(ctx *gin.Context) {
		dto.Http.Ok(ctx, nil)
	})
}
