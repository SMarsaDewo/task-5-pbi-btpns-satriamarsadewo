package main

import (
	"net/http"

	"github.com/SMarsaDewo/go-jwt/controllers"
	"github.com/SMarsaDewo/go-jwt/initializers"
	"github.com/SMarsaDewo/go-jwt/middleware"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	var currentImage *imageupload.image
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/image", func(c *gin.Context) {
		if currentImage == nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		currentImage.Write(c.Writer)
	})
	r.GET("/thumbnail", func(c *gin.Context) {
		if currentImage == nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		t, err := imageupload.ThumbnailJPEG(currentImage, 300, 300, 80)

		if err != nil {
			panic(err)
		}

		t.Write(c.Writer)
	})

	r.POST("/uploud", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")
		if err != nil {
			panic(err)
		}

		currentImage = img

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":8080")

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()

}
