package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"site_about_me/controller"
)

func main() {
	route := gin.Default()

	route.Static("/static", "./static")

	route.LoadHTMLGlob("temp/*")
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main Page",
		})
	})

	achive := route.Group("/achive")
	{
		achive.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "achive.html", gin.H{
				"title": "Achive",
			})
		})
	}

	iam := route.Group("/iam")
	{

		iam.GET("/", func(c *gin.Context) {
			var arr []controller.Me
			arr = controller.GetAllData()

			c.HTML(http.StatusOK, "iam.html", gin.H{
				"title":   "About Me",
				"mySkill": arr,
			})
		})

		iam.GET("/github", func(c *gin.Context) {
			c.HTML(http.StatusOK, "git.html", gin.H{
				"title": "My GIT",
			})
		}) // переход на мой гит
	}
}
