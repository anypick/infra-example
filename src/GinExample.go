package src

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-gin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	infra.RegisterApi(&CouponController{})
}

type CouponController struct {}

func (c *CouponController) Init() {
	app := basegin.Gin().Group("/coupon")
	app.GET("/get", c.Get)
}

func (c *CouponController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, `{"name":"张三"}`)
}

