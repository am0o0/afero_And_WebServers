package main

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	filepath := c.Ctx.Request.URL.Query()["filepath"][0]
	c.Ctx.Output.Download(filepath, "license.txt")
	buffer := make([]byte, 10)
	_ = c.SaveToFileWithBuffer("filenameExistsInForm", filepath, buffer)

}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
