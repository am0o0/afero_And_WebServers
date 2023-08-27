package main

import (
	"github.com/kataras/iris/v12"
	"mime/multipart"
)

func main() {
	app := iris.New()
	app.UseRouter(iris.Compression)
	app.Get("/", func(ctx iris.Context) {
		filepath := ctx.URLParam("filepath")
		_ = ctx.SendFile(filepath, "file")
		_ = ctx.SendFileWithRate(filepath, "file", 0, 0)
		_ = ctx.ServeFile(filepath)
		_ = ctx.ServeFileWithRate(filepath, 0, 0)
		_, _, _ = ctx.UploadFormFiles(filepath, beforeSave)
		_, fileHeader, _ := ctx.FormFile("file")
		_, _ = ctx.SaveFormFile(fileHeader, filepath)

	})
	app.Listen(":8080")
}
func beforeSave(ctx iris.Context, file *multipart.FileHeader) bool {
	return true
}
