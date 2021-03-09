package controllers

import (
	"fmt"
	"log"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type params struct {
	sl   string `query:"sl"`
	tl   string `query:"sl"`
	text string `query:"text"`
	op   string `query:"op"`
}

func (c *MainController) Get() {

	p := params{}
	c.Ctx.Input.Bind(&p.sl, "sl")
	c.Ctx.Input.Bind(&p.tl, "tl")
	c.Ctx.Input.Bind(&p.text, "text")
	c.Ctx.Input.Bind(&p.op, "op")
	log.Println(p)

	var tplName string
	tplName = "text-trans.tpl"
	if p.op == "docs" {
		tplName = "file-trans.tpl"
	}

	c.Data["transMap"] = transMap
	c.Layout = "index.tpl"
	c.TplName = tplName

}

func (c *MainController) Post() {

	p := params{}
	c.Ctx.Input.Bind(&p.sl, "sl")
	c.Ctx.Input.Bind(&p.tl, "tl")
	c.Ctx.Input.Bind(&p.text, "text")
	c.Ctx.Input.Bind(&p.op, "op")
	log.Println(p)

	w := new(LocalTranslate)
	trans, _ := w.Text(p.sl, p.tl, p.text)
	fmt.Printf("trans: %v\n", trans)
	c.Data["json"] = trans // json对象
	c.ServeJSON()
	return

}
