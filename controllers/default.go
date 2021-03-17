package controllers

import (
	"log"
	"path"

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

	isDocs := false
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
		isDocs = true
	}

	c.Data["IsDocs"] = isDocs
	c.Data["TransMap"] = transMap
	c.Layout = "index.tpl"
	c.TplName = tplName

}

func (c *MainController) Post() {

	p := params{}
	c.Ctx.Input.Bind(&p.op, "op")
	log.Println(p)

	if p.op == "docs" {
		uploadDocs(c, p)
	} else if p.op == "translate" {
		c.Ctx.Input.Bind(&p.sl, "sl")
		c.Ctx.Input.Bind(&p.tl, "tl")
		c.Ctx.Input.Bind(&p.text, "text")
		transText(c, p)
	} else {
		c.Data["json"] = map[string]string{"msg": "Hello world", "code": "10"}
		c.ServeJSON()
	}
	return

}

func transText(c *MainController, p params) {
	w := new(LocalTranslate)
	trans, err := w.Text(p.sl, p.tl, p.text)
	if err != nil {
		log.Println(err)
	}

	c.Data["json"] = trans // json对象
	c.ServeJSON()
}

func uploadDocs(c *MainController, p params) {
	f, h, err := c.GetFile("file")
	if err != nil {
		c.Data["json"] = map[string]string{"msg": "上传文件失败", "code": "10"}
		c.ServeJSON()
	}
	defer f.Close()

	//验证后缀名是否符合要求
	ext := path.Ext(h.Filename)
	var AllowExtMap map[string]bool = map[string]bool{
		// ".xls":  true,
		".xlsx": true,
		// ".csv":  true,
	}

	if _, ok := AllowExtMap[ext]; !ok {
		c.Data["json"] = map[string]string{"msg": "后缀名不符合上传要求", "code": "10"}
		c.ServeJSON()
		return
	}
	filePath := "static/upload/" + h.Filename
	c.SaveToFile("file", filePath) // 保存位置在 static/upload, 没有文件夹要先创建
	c.Data["json"] = map[string]string{"msg": "上传成功", "code": "0", "file_path": filePath, "file_name": h.Filename}
	c.ServeJSON()
}
