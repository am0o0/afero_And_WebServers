package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// https://github.com/advisories/GHSA-85cf-gj29-f555
	// https://github.com/1Panel-dev/1Panel/commit/f6b84d384e41c1c708d41160b2af9761d984f558#diff-3daefeffef5edafe17bc47b7e75c1332b684aa728ee67efffc6fe795d95d3627L556-L568
	router := gin.Default()

	//router.GET("/get", func(c *gin.Context) {
	//	var embeddedFiles embed.FS
	//	sub, _ := fs.Sub(embeddedFiles, "assets/favicon.png")
	//	c.FileFromFS(c.Request.URL.Path, http.FS(sub))
	//})
	router.POST("/FormUploads", func(c *gin.Context) {
		filepath := c.Query("filepath")
		c.File(filepath)
		http.ServeFile(c.Writer, c.Request, filepath)
		c.FileAttachment(filepath, "file name in response")
		file, _ := c.FormFile("afile")
		_ = c.SaveUploadedFile(file, filepath)
	})
	err := router.Run()
	if err != nil {
		return
	}
}
