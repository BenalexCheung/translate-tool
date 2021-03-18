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
	Created time.Time `json:"create_time,omitempty"  orm:"column(create_time);type(datetime);null"`
	Updated time.Time `json:"update_time,omitempty"  orm:"column(update_time);type(datetime);null"`
	En      string    `json:"en,omitempty"  orm:"column(en);type(text);unique"`
	Fr      string    `json:"fr,omitempty"  orm:"column(fr);type(text);null"`
	De      string    `json:"de,omitempty"  orm:"column(de);type(text);null"`
	It      string    `json:"it,omitempty"  orm:"column(it);type(text);null"`
	Ja      string    `json:"ja,omitempty"  orm:"column(ja);type(text);null"`
	Ko      string    `json:"ko,omitempty"  orm:"column(ko);type(text);null"`
	Pt_PT   string    `json:"pt-PT,omitempty"  orm:"column(pt-PT);type(text);null"`
	Ru      string    `json:"ru,omitempty"  orm:"column(ru);type(text);null"`
	Es      string    `json:"es,omitempty"  orm:"column(es);type(text);null"`
	Ar      string    `json:"ar,omitempty"  orm:"column(ar);type(text);null"`
	He      string    `json:"he,omitempty"  orm:"column(he);type(text);null"`
	Tr      string    `json:"tr,omitempty"  orm:"column(tr);type(text);null"`
	Pl      string    `json:"pl,omitempty"  orm:"column(pl);type(text);null"`
	Ro_RO   string    `json:"ro-RO,omitempty"  orm:"column(ro-RO);type(text);null"`
	Bg_BG   string    `json:"bg-BG,omitempty"  orm:"column(bg-BG);type(text);null"`
	Zh_CN   string    `json:"zh-CN,omitempty"  orm:"column(zh-CN);type(text);null"`
	Zh_TW   string    `json:"zh-TW,omitempty"  orm:"column(zh-TW);type(text);null"`
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

func GetEntry(sl string, tl string, text string) (entry string, err error) {
	o := GetLink()
	rs := o.Raw("SELECT `"+tl+"` FROM `entries` "+
		"WHERE `"+sl+"`=?", text)
	err = rs.QueryRow(&entry)
	fmt.Printf("ERR: %v\n", err)
	return
}

func GetEntryByKey(key string) (*Entries, error) {
	o := GetLink()
	entry := &Entries{En: key}
	err := o.Read(entry, "en")
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return nil, err
	}
	return entry, err
}

func SaveEntry(entry Entries) (id int64) {
	o := GetLink()
	id, err := o.Insert(&entry)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	return
}

func UpdateEntry(entry Entries, col ...string) (num int64) {
	o := GetLink()
	num, err := o.Update(&entry, col...)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	return
}

func DelEntry(entry Entries) {
	o := GetLink()
	num, err := o.Delete(&entry)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	return
}
