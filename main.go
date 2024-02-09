package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "monorail.proxy.rlwy.net:12119",
		Username: "default",
		Password: "i5NmCaCf4CFO42Hb3hDg6jnMOeMCIIPd",
		DB:       0,
	})

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*")
	r.GET("/pie", func(c *gin.Context) {
		pieInfo, _ := rdb.Get(ctx, "pie").Result()
		x, _ := strconv.Atoi(pieInfo)
		pieAmount := x + 1
		rdb.Set(ctx, "pie", fmt.Sprint(pieAmount), 0)
		c.HTML(http.StatusOK, "pie.go.html", gin.H{
			"title": fmt.Sprintf("This is pie number %d!!!", pieAmount),
		})
	})

	r.Run(":9090")
}
