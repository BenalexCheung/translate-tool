package controllers

import (
	"log"

	beego "github.com/beego/beego/v2/server/web"
)

type ConvertController struct {
	beego.Controller
}

type fileInfo struct {
	fileName string `form:"file_name"`
	filePath string `form:"file_path"`
	oper     string `form:"op"`
}

func (c *ConvertController) Get() {

	fi := fileInfo{}
	c.Ctx.Input.Bind(&fi.fileName, "file_name")
	c.Ctx.Input.Bind(&fi.filePath, "file_path")
	c.Ctx.Input.Bind(&fi.oper, "op")

	log.Println(fi)
	c.Data["json"] = map[string]string{"msg": "Hello world", "code": "10"}
	c.ServeJSON()
}

func (c *ConvertController) Post() {

	fi := fileInfo{}
	c.Ctx.Input.Bind(&fi.fileName, "file_name")
	c.Ctx.Input.Bind(&fi.filePath, "file_path")
	c.Ctx.Input.Bind(&fi.oper, "op")
	log.Println(fi)

	if fi.oper == "translate" {
		transDocs(fi)
		c.Data["json"] = map[string]string{"msg": "翻译成功", "code": "0", "file_path": fi.filePath, "file_name": fi.fileName}
		c.ServeJSON()
	} else if fi.oper == "convert" {
		convertDocs(fi)
		c.Data["json"] = map[string]string{"msg": "转换成功", "code": "0", "file_path": fi.filePath, "file_name": fi.fileName}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]string{"msg": "Hello world", "code": "10"}
		c.ServeJSON()
	}

	return

}

func transDocs(fi fileInfo) {

}

func convertDocs(fi fileInfo) {

}
