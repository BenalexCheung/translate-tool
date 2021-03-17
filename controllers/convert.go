package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"translateTool/models"
	"translateTool/utils"

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
		filePath := "static/upload/Localizable.strings"
		err := convertDocs(fi, filePath)
		if err != nil {
			c.Data["json"] = map[string]string{"msg": "转换失败", "code": "10", "file_path": "", "file_name": ""}
			c.ServeJSON()
		}
		c.Data["json"] = map[string]string{"msg": "转换成功", "code": "0", "file_path": filePath, "file_name": "Localizable.strings"}
		c.ServeJSON()
	} else if fi.oper == "loading" {
		loadDocs(fi)
		c.Data["json"] = map[string]string{"msg": "导入成功", "code": "0", "file_path": fi.filePath, "file_name": fi.fileName}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]string{"msg": "Hello world", "code": "10"}
		c.ServeJSON()
	}

	return

}

func transDocs(fi fileInfo) {
	var lang string
	lang = strings.TrimSuffix(fi.fileName, ".xlsx") //获取文件名
	fmt.Println("lang =", lang)

	var newCol []string

	x := utils.NewXLS(fi.filePath)

	err := x.ParsingXLS()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, rowCell := range x.Cols[0] {

		entry, err := models.GetEntry("en", lang, rowCell)
		if err != nil {
			fmt.Println(err)
		}
		newCol = append(newCol, entry)
	}

	err = x.SaveXLS("B", newCol)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func convertDocs(fi fileInfo, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败: ", err)
		return err
	}
	defer f.Close()

	x := utils.NewXLS(fi.filePath)
	err = x.ParsingXLS()
	if err != nil {
		fmt.Println(err)
		return err
	}

	w := bufio.NewWriter(f)
	for i, rowCell := range x.Cols[0] {
		sl := strconv.Quote(rowCell)
		tl := strconv.Quote(x.Cols[1][i])
		s := fmt.Sprintf("%s = %s;", sl, tl)
		fmt.Println(s)
		fmt.Fprintln(w, s)
	}

	return w.Flush()
}

func loadDocs(fi fileInfo) {
	var lang string
	lang = strings.TrimSuffix(fi.fileName, ".xlsx") //获取文件名
	fmt.Println("lang =", lang)

	x := utils.NewXLS(fi.filePath)

	err := x.ParsingXLS()
	if err != nil {
		return
	}

	for i, rowCell := range x.Cols[0] {
		sl := strconv.Quote(rowCell)
		tl := strconv.Quote(x.Cols[1][i])
		data := fmt.Sprintf("{\"en\":%s,\"%s\":%s}", sl, lang, tl)

		str := []byte(data)
		entry := models.Entries{}

		err := json.Unmarshal(str, &entry)
		if err != nil {
			fmt.Println(err)
			break
		}

		e, err := models.GetEntryByKey(entry.En) // 不存在则插入，存在则更新或忽略
		if err != nil {
			entry.Created = time.Now()
			models.SaveEntry(entry)
		} else {
			entry.Id = e.Id
			entry.Updated = time.Now()
			fmt.Println("entry.Updated :", entry.Updated)
			models.UpdateEntry(entry, "update_time", lang)
		}
	}
}
