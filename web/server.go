package web

import (
	"Wildberries_L0/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Server(orderCache *cache.Cache) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("A:/go_workspace/Wildberries_L0/web/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})
	r.GET("/order/", func(c *gin.Context) {
		//id := c.DefaultQuery("order_id", "Guest")
		id := orderCache.GetOrderByUID(c.Query("order_id"))
		//c.PureJSON()
		c.JSON(http.StatusOK, id)
	})
	r.Run(":8080")
}
