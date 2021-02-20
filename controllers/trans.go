package controllers

import (
	"time"

	"translateTool/models"

	beego "github.com/beego/beego/v2/server/web"
)

type NewController struct {
	beego.Controller
}

func (this *NewController) Get() {
	this.Layout = "layout.tpl"
	this.TplName = "new.tpl"
}

func (this *NewController) Post() {
	inputs, _ := this.Input()
	var blog models.Blog
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now()
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
