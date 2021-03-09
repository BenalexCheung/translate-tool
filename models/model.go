package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

// Entries Struct
type Entries struct {
	Id      int       `json:"id,omitempty"  orm:"column(id)"`
	Created time.Time `json:"create_time,omitempty"  orm:"column(create_time)"`
	Updated time.Time `json:"update_time,omitempty"  orm:"column(update_time)"`
	En      string    `json:"en,omitempty"  orm:"column(en);type(text);unique"`
	De      string    `json:"de,omitempty"  orm:"column(de);type(text);null"`
	Fr      string    `json:"fr,omitempty"  orm:"column(fr);type(text);null"`
	Zh_CN   string    `json:"zh_cn,omitempty"  orm:"column(zh_cn);type(text);null"`
	Zh_TW   string    `json:"zh_tw,omitempty"  orm:"column(zh_tw);type(text);null"`
}

func init() {
	// 注册驱动
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	// 设置默认数据库
	orm.RegisterDataBase("default", "sqlite3", "db/stardict.db")
	// 注册定义的 model
	orm.RegisterModel(new(Entries))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	orm.RunSyncdb("default", false, true)
}

func GetLink() orm.Ormer {
	orm := orm.NewOrm()
	return orm
}

func GetAll() (entries []*Entries) {
	o := GetLink()
	num, err := o.QueryTable("entries").All(&entries)
	fmt.Printf("num: %v, ERR: %v\n", num, err)
	return
}

func GetEntry(sl string, tl string, text string) (entry string) {
	o := GetLink()
	rs := o.Raw("SELECT `"+tl+"` FROM `entries` "+
		"WHERE `"+sl+"`=?", text)
	err := rs.QueryRow(&entry)
	fmt.Printf("ERR: %v\n", err)
	return
}

func GetEntryById(id int) (entry *Entries) {
	o := GetLink()
	entry = &Entries{Id: id}
	err := o.Read(entry)
	fmt.Printf("ERR: %v\n", err)
	return
}

func SaveEntry(entry Entries) (id int64) {
	o := GetLink()
	id, err := o.Insert(&entry)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	return
}

func UpdateEntry(entry Entries) (num int64) {
	o := GetLink()
	num, err := o.Update(&entry)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	return
}

func DelEntry(entry Entries) {
	o := GetLink()
	num, err := o.Delete(&entry)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	return
}
